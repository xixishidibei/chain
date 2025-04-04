package keeper

import (
	"time"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/bandprotocol/chain/v3/x/oracle/types"
)

// valWithPower is an internal type to track validator with voting power inside of AllocateTokens.
type valWithPower struct {
	val   stakingtypes.ValidatorI
	power int64
}

// AllocateTokens allocates a portion of fee collected in the previous blocks to validators that
// that are actively performing oracle tasks. Note that this reward is also subjected to comm tax.
func (k Keeper) AllocateTokens(ctx sdk.Context, previousVotes []abci.VoteInfo) error {
	toReward := []valWithPower{}
	totalPower := int64(0)
	for _, vote := range previousVotes {
		val, err := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)
		if err != nil {
			continue
		}
		operator, err := sdk.ValAddressFromBech32(val.GetOperator())
		if err != nil {
			continue
		}
		if k.GetValidatorStatus(ctx, operator).IsActive {
			toReward = append(toReward, valWithPower{val: val, power: vote.Validator.Power})
			totalPower += vote.Validator.Power
		}
	}
	if totalPower == 0 {
		// No active validators performing oracle tasks, nothing needs to be done here.
		return nil
	}
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	totalFee := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())...)
	// Compute the fee allocated for oracle module to distribute to active validators.
	oracleRewardRatio := math.LegacyNewDecWithPrec(int64(k.GetParams(ctx).OracleRewardPercentage), 2)
	oracleRewardInt, _ := totalFee.MulDecTruncate(oracleRewardRatio).TruncateDecimal()

	// Transfer the oracle reward portion from fee collector to distr module.
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, distrtypes.ModuleName, oracleRewardInt)
	if err != nil {
		return err
	}
	// Convert the transferred tokens back to DecCoins for internal distr allocations.
	oracleReward := sdk.NewDecCoinsFromCoins(oracleRewardInt...)
	communityTax, err := k.distrKeeper.GetCommunityTax(ctx)
	if err != nil {
		return err
	}

	// Fund community pool with a portion of the oracle reward.
	communityFund, _ := oracleReward.MulDecTruncate(communityTax).TruncateDecimal()
	err = k.distrKeeper.FundCommunityPool(
		ctx,
		communityFund,
		k.authKeeper.GetModuleAccount(ctx, distrtypes.ModuleName).GetAddress(),
	)
	if err != nil {
		panic(err)
	}
	oracleReward = oracleReward.Sub(sdk.NewDecCoinsFromCoins(communityFund...))
	remaining := oracleReward

	// Allocate non-community pool tokens to active validators weighted by voting power.
	for _, each := range toReward {
		powerFraction := math.LegacyNewDec(each.power).QuoTruncate(math.LegacyNewDec(totalPower))
		reward := oracleReward.MulDecTruncate(powerFraction)
		err := k.distrKeeper.AllocateTokensToValidator(ctx, each.val, reward)
		if err != nil {
			// Should never hit
			return err
		}
		remaining = remaining.Sub(reward)
	}

	// Remaining tokens are sent to proposer.
	proposer, err := k.stakingKeeper.ValidatorByConsAddr(ctx, ctx.BlockHeader().ProposerAddress)
	if err != nil {
		// Should never hit
		return err
	}
	return k.distrKeeper.AllocateTokensToValidator(ctx, proposer, remaining)
}

// GetValidatorStatus returns the validator status for the given validator. Note that validator
// status is default to [inactive, 0], so new validators start with inactive state.
func (k Keeper) GetValidatorStatus(ctx sdk.Context, val sdk.ValAddress) types.ValidatorStatus {
	bz := ctx.KVStore(k.storeKey).Get(types.ValidatorStatusStoreKey(val))
	if bz == nil {
		return types.NewValidatorStatus(false, time.Time{})
	}
	var status types.ValidatorStatus
	k.cdc.MustUnmarshal(bz, &status)
	return status
}

// SetValidatorStatus sets the validator status for the given validator.
func (k Keeper) SetValidatorStatus(ctx sdk.Context, val sdk.ValAddress, status types.ValidatorStatus) {
	ctx.KVStore(k.storeKey).Set(types.ValidatorStatusStoreKey(val), k.cdc.MustMarshal(&status))
}

// Activate changes the given validator's status to active. Returns error if the validator is
// already active or was deactivated recently, as specified by InactivePenaltyDuration parameter.
func (k Keeper) Activate(ctx sdk.Context, val sdk.ValAddress) error {
	status := k.GetValidatorStatus(ctx, val)
	if status.IsActive {
		return types.ErrValidatorAlreadyActive
	}
	penaltyDuration := time.Duration(k.GetParams(ctx).InactivePenaltyDuration)
	if !status.Since.IsZero() && status.Since.Add(penaltyDuration).After(ctx.BlockHeader().Time) {
		return types.ErrTooSoonToActivate
	}
	k.SetValidatorStatus(ctx, val, types.NewValidatorStatus(true, ctx.BlockHeader().Time))
	return nil
}

// MissReport changes the given validator's status to inactive. No-op if already inactive or
// if the validator was active after the time the request happened.
func (k Keeper) MissReport(ctx sdk.Context, val sdk.ValAddress, requestTime time.Time) {
	status := k.GetValidatorStatus(ctx, val)
	if status.IsActive && status.Since.Before(requestTime) {
		k.SetValidatorStatus(ctx, val, types.NewValidatorStatus(false, ctx.BlockHeader().Time))
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			types.EventTypeDeactivate,
			sdk.NewAttribute(types.AttributeKeyValidator, val.String()),
		))
	}
}

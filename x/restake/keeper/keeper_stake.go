package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/restake/types"
)

func (k Keeper) GetStakesIterator(ctx sdk.Context, address sdk.AccAddress) sdk.Iterator {
	return sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.StakesStoreKey(address))
}

func (k Keeper) GetActiveStakes(ctx sdk.Context, address sdk.AccAddress) (stakes []types.Stake) {
	iterator := k.GetStakesIterator(ctx, address)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var stake types.Stake
		k.cdc.MustUnmarshal(iterator.Value(), &stake)

		key, err := k.GetKey(ctx, stake.Key)
		if err != nil || !key.IsActive {
			continue
		}

		stakes = append(stakes, stake)
	}

	return stakes
}

func (k Keeper) GetStakes(ctx sdk.Context, address sdk.AccAddress) (stakes []types.Stake) {
	iterator := k.GetStakesIterator(ctx, address)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var stake types.Stake
		k.cdc.MustUnmarshal(iterator.Value(), &stake)
		stakes = append(stakes, stake)
	}

	return stakes
}

func (k Keeper) HasStake(ctx sdk.Context, address sdk.AccAddress, keyName string) bool {
	return ctx.KVStore(k.storeKey).Has(types.StakeStoreKey(address, keyName))
}

func (k Keeper) GetStake(ctx sdk.Context, address sdk.AccAddress, keyName string) (types.Stake, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.StakeStoreKey(address, keyName))
	if bz == nil {
		return types.Stake{}, types.ErrStakeNotFound.Wrapf(
			"failed to get stake of %s with key name: %s",
			address.String(),
			keyName,
		)
	}

	var stake types.Stake
	k.cdc.MustUnmarshal(bz, &stake)

	return stake, nil
}

func (k Keeper) SetStake(ctx sdk.Context, stake types.Stake) {
	address := sdk.MustAccAddressFromBech32(stake.Address)
	ctx.KVStore(k.storeKey).Set(types.StakeStoreKey(address, stake.Key), k.cdc.MustMarshal(&stake))
}

func (k Keeper) DeleteStake(ctx sdk.Context, address sdk.AccAddress, keyName string) {
	ctx.KVStore(k.storeKey).Delete(types.StakeStoreKey(address, keyName))
}

func (k Keeper) updateRewardLefts(ctx sdk.Context, key types.Key, stake types.Stake) types.Stake {
	diff := key.RewardPerShares.Sub(stake.RewardDebts)
	stake.RewardLefts = stake.RewardLefts.Add(diff.MulDecTruncate(sdk.NewDecFromInt(stake.Amount))...)
	stake.RewardDebts = key.RewardPerShares
	k.SetStake(ctx, stake)

	return stake
}

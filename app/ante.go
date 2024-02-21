package band

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	ibcante "github.com/cosmos/ibc-go/v7/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"

	"github.com/bandprotocol/chain/v2/x/globalfee/feechecker"
	globalfeekeeper "github.com/bandprotocol/chain/v2/x/globalfee/keeper"
	oraclekeeper "github.com/bandprotocol/chain/v2/x/oracle/keeper"
	tsskeeper "github.com/bandprotocol/chain/v2/x/tss/keeper"
	tssmemberkeeper "github.com/bandprotocol/chain/v2/x/tssmember/keeper"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper.
type HandlerOptions struct {
	ante.HandlerOptions
	AuthzKeeper     *authzkeeper.Keeper
	OracleKeeper    *oraclekeeper.Keeper
	IBCKeeper       *ibckeeper.Keeper
	GlobalfeeKeeper *globalfeekeeper.Keeper
	StakingKeeper   *stakingkeeper.Keeper
	TSSKeeper       *tsskeeper.Keeper
	TSSMemberKeeper *tssmemberkeeper.Keeper
}

func NewAnteHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.ErrLogic.Wrap("sign mode handler is required for ante builder")
	}
	if options.AuthzKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("authz keeper is required for AnteHandler")
	}
	if options.OracleKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("oracle keeper is required for AnteHandler")
	}
	if options.TSSKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("tss keeper is required for AnteHandler")
	}
	if options.TSSMemberKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("tssmember keeper is required for AnteHandler")
	}
	if options.IBCKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("IBC keeper is required for AnteHandler")
	}
	if options.StakingKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("Staking keeper is required for AnteHandler")
	}
	if options.GlobalfeeKeeper == nil {
		return nil, sdkerrors.ErrLogic.Wrap("Globalfee keeper is required for AnteHandler")
	}

	sigGasConsumer := options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}

	if options.TxFeeChecker == nil {
		feeChecker := feechecker.NewFeeChecker(
			options.AuthzKeeper,
			options.OracleKeeper,
			options.GlobalfeeKeeper,
			options.StakingKeeper,
			options.TSSKeeper,
			options.TSSMemberKeeper,
		)
		options.TxFeeChecker = feeChecker.CheckTxFeeWithMinGasPrices
	}

	anteDecorators := []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(
			options.AccountKeeper,
			options.BankKeeper,
			options.FeegrantKeeper,
			options.TxFeeChecker,
		),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
	}

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

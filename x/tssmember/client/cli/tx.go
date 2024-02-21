package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	tsstypes "github.com/bandprotocol/chain/v2/x/tss/types"
	"github.com/bandprotocol/chain/v2/x/tssmember/types"
)

const (
	flagGroupID  = "group-id"
	flagFeeLimit = "fee-limit"
)

// GetTxCmd returns a root CLI command handler for all x/tssmember transaction commands.
func GetTxCmd(requestSignatureCmds []*cobra.Command) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "TSS transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// Create the command for requesting a signature.
	cmdRequestSignature := GetTxCmdRequestSignature()

	// Create the command for requesting a signature using text input.
	cmdTextRequestSignature := GetTxCmdTextRequestSignature()

	// Add the text signature command as a subcommand.
	flags.AddTxFlagsToCmd(cmdTextRequestSignature)
	cmdRequestSignature.AddCommand(cmdTextRequestSignature)

	// Loop through and add the provided request signature commands as subcommands.
	for _, cmd := range requestSignatureCmds {
		flags.AddTxFlagsToCmd(cmd)
		cmdRequestSignature.AddCommand(cmd)
	}

	txCmd.AddCommand(
		GetTxCmdActivate(),
		GetTxCmdHealthCheck(),

		cmdRequestSignature,
	)

	return txCmd
}

// GetTxCmdRequestSignature creates a CLI command for CLI command for Msg/RequestSignature.
func GetTxCmdRequestSignature() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-signature",
		Short: "request signature from the group",
	}

	cmd.PersistentFlags().String(flagFeeLimit, "", "The maximum tokens that will be paid for this request")
	cmd.PersistentFlags().Uint64(flagGroupID, 0, "The group that is requested to sign the result")

	_ = cmd.MarkPersistentFlagRequired(flagFeeLimit)
	_ = cmd.MarkPersistentFlagRequired(flagGroupID)

	return cmd
}

// GetTxCmdTextRequestSignature creates a CLI command for CLI command for Msg/TextRequestSignature.
func GetTxCmdTextRequestSignature() *cobra.Command {
	return &cobra.Command{
		Use:   "text [message]",
		Args:  cobra.ExactArgs(1),
		Short: "request signature of the message from the group",
		Example: fmt.Sprintf(
			`%s tx tss request-signature text [message] --group-id 1 --fee-limit 10uband`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gid, err := cmd.Flags().GetUint64(flagGroupID)
			if err != nil {
				return err
			}

			data, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			content := tsstypes.NewTextSignatureOrder(data)

			coinStr, err := cmd.Flags().GetString(flagFeeLimit)
			if err != nil {
				return err
			}

			feeLimit, err := sdk.ParseCoinsNormalized(coinStr)
			if err != nil {
				return err
			}

			msg, err := types.NewMsgRequestSignature(
				tss.GroupID(gid),
				content,
				feeLimit,
				clientCtx.GetFromAddress(),
			)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetTxCmdActivate creates a CLI command for CLI command for Msg/Activate.
func GetTxCmdActivate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate",
		Args:  cobra.NoArgs,
		Short: "activate the status of the address",
		Example: fmt.Sprintf(
			`%s tx tss activate`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgActivate{
				Address: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetTxCmdHealthCheck creates a CLI command for CLI command for Msg/HealthCheck.
func GetTxCmdHealthCheck() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health-check",
		Args:  cobra.NoArgs,
		Short: "update the active status of the address to ensure that the TSS process is still running",
		Example: fmt.Sprintf(
			`%s tx tss health-check`,
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgHealthCheck{
				Address: clientCtx.GetFromAddress().String(),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

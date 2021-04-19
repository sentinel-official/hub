package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func txRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var provider hubtypes.ProvAddress
			if len(s) > 0 {
				provider, err = hubtypes.ProvAddressFromBech32(s)
				if err != nil {
					return err
				}
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			var price sdk.Coins
			if len(s) > 0 {
				price, err = sdk.ParseCoinsNormalized(s)
				if err != nil {
					return err
				}
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterRequest(ctx.FromAddress, provider, price, remoteURL)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagProvider, "", "node provider address")
	cmd.Flags().String(flagPrice, "", "node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "node remote URL")

	_ = cmd.MarkFlagRequired(flagRemoteURL)

	return cmd
}

func txUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var provider hubtypes.ProvAddress
			if len(s) > 0 {
				provider, err = hubtypes.ProvAddressFromBech32(s)
				if err != nil {
					return err
				}
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			var price sdk.Coins
			if len(s) > 0 {
				price, err = sdk.ParseCoinsNormalized(s)
				if err != nil {
					return err
				}
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRequest(ctx.FromAddress.Bytes(), provider, price, remoteURL)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagProvider, "", "node provider address")
	cmd.Flags().String(flagPrice, "", "node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "node remote URL")

	return cmd
}

func txSetStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set [Active | Inactive]",
		Short: "Set a node status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetStatusRequest(ctx.FromAddress.Bytes(), hubtypes.StatusFromString(args[0]))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

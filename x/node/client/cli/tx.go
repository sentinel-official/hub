package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func txRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [remote-url]",
		Short: "Register a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			gigabytePrices, err := GetGigabytePrices(cmd.Flags())
			if err != nil {
				return err
			}

			hourlyPrice, err := GetHourlyPrices(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterRequest(
				ctx.FromAddress,
				gigabytePrices,
				hourlyPrice,
				args[0],
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagGigabytePrices, "", "prices per one gigabyte of bandwidth provision")
	cmd.Flags().String(flagHourlyPrices, "", "prices per one hour of bandwidth provision")

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

			gigabytePrices, err := GetGigabytePrices(cmd.Flags())
			if err != nil {
				return err
			}

			hourlyPrice, err := GetHourlyPrices(cmd.Flags())
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateRequest(
				ctx.FromAddress.Bytes(),
				gigabytePrices,
				hourlyPrice,
				remoteURL,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(flagGigabytePrices, "", "prices per one gigabyte of bandwidth provision")
	cmd.Flags().String(flagHourlyPrices, "", "prices per one hour of bandwidth provision")
	cmd.Flags().String(flagRemoteURL, "", "remote URL address of the node")

	return cmd
}

func txSetStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set [status]",
		Short: "Set status for a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetStatusRequest(
				ctx.FromAddress.Bytes(),
				hubtypes.StatusFromString(args[0]),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

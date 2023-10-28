// DO NOT COVER

package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/v1/types"
	"github.com/sentinel-official/hub/v1/x/node/types"
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
			if err = msg.ValidateBasic(); err != nil {
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

func txUpdateDetails() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-details",
		Short: "Update the details of a node",
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

			msg := types.NewMsgUpdateDetailsRequest(
				ctx.FromAddress.Bytes(),
				gigabytePrices,
				hourlyPrice,
				remoteURL,
			)
			if err = msg.ValidateBasic(); err != nil {
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

func txUpdateStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-status [status]",
		Short: "Update the status for a node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateStatusRequest(
				ctx.FromAddress.Bytes(),
				hubtypes.StatusFromString(args[0]),
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func txSubscribe() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscribe [node-addr] [denom]",
		Short: "Subscribe to a node",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr, err := hubtypes.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			gigabytes, err := cmd.Flags().GetInt64(flagGigabytes)
			if err != nil {
				return err
			}

			hours, err := cmd.Flags().GetInt64(flagHours)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubscribeRequest(
				ctx.FromAddress,
				addr,
				gigabytes,
				hours,
				args[1],
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().Int64(flagGigabytes, 0, "gigabytes")
	cmd.Flags().Int64(flagHours, 0, "hours")

	return cmd
}

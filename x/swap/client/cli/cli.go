package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap",
		Short: "Querying commands for the swap module",
	}

	cmd.AddCommand(
		flags.GetCommands(
			querySwap(cdc),
			querySwaps(cdc),
		)...,
	)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap",
		Short: "Swap module sub-commands",
	}

	cmd.AddCommand(
		flags.PostCommands(
			txSwap(cdc),
		)...,
	)

	return cmd
}

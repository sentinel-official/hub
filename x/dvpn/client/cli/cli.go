package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dvpn",
		Short: "Querying commands for the dVPN module",
	}

	cmd.AddCommand(provider.GetQueryCommands(cdc)...)
	cmd.AddCommand(node.GetQueryCommands(cdc)...)
	cmd.AddCommand(subscription.GetQueryCommands(cdc)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dvpn",
		Short: "dVPN transactions subcommands",
	}

	cmd.AddCommand(provider.GetTxCommands(cdc)...)
	cmd.AddCommand(node.GetTxCommands(cdc)...)
	cmd.AddCommand(subscription.GetTxCommands(cdc)...)

	return cmd
}

package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(provider.GetQueryCommands(cdc)...)
	cmd.AddCommand(node.GetQueryCommands(cdc)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(provider.GetTxCommands(cdc)...)
	cmd.AddCommand(node.GetTxCommands(cdc)...)

	return cmd
}

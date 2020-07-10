package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return client.GetCommands(
		queryNodeCmd(cdc),
		queryNodesCmd(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		txRegisterNodeCmd(cdc),
		txUpdateNodeCmd(cdc),
		txSetNodeStatusCmd(cdc),
	)...)

	return []*cobra.Command{cmd}
}

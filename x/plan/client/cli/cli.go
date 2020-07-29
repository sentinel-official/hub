package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return client.GetCommands(
		queryPlanCmd(cdc),
		queryPlansCmd(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "plan module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		txAddPlanCmd(cdc),
		txSetPlanStatusCmd(cdc),
		txAddNodeForPlanCmd(cdc),
		txRemoveNodeForPlanCmd(cdc),
	)...)

	return []*cobra.Command{cmd}
}

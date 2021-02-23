package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return flags.GetCommands(
		querySession(cdc),
		querySessions(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txUpsert(cdc),
	)...)

	return []*cobra.Command{cmd}
}

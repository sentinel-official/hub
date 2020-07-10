package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return nil
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		txUpdateSession(cdc),
	)...)

	return []*cobra.Command{cmd}
}

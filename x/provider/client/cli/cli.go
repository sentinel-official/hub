package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return flags.GetCommands(
		queryProvider(cdc),
		queryProviders(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "Provider module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txRegister(cdc),
		txUpdate(cdc),
	)...)

	return []*cobra.Command{cmd}
}

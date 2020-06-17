package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return client.GetCommands(
		getQueryProviderCmd(cdc),
		getQueryProvidersCmd(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "Provider module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		getTxRegisterProviderCmd(cdc),
	)...)

	return []*cobra.Command{cmd}
}

package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	provider "github.com/sentinel-official/hub/x/dvpn/provider/client/cli"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(client.GetCommands(
		provider.GetQueryProviderCmd(cdc),
		provider.GetQueryProvidersCmd(cdc),
	)...)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(client.PostCommands(
		provider.GetTxRegisterProviderCmd(cdc),
	)...)

	return cmd
}

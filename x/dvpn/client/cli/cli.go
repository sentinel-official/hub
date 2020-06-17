package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	provider "github.com/sentinel-official/hub/x/dvpn/provider/client/cli"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(
		provider.GetQueryCommands(cdc)...,
	)

	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dvpn",
	}

	cmd.AddCommand(
		provider.GetTxCommands(cdc)...,
	)

	return cmd
}

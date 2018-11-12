package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	app "github.com/ironman0x7b2/sentinel-sdk/apps/sentinel-hub"
	ibcCli "github.com/ironman0x7b2/sentinel-sdk/x/ibc/client/cli"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-hub-cli",
	Short: "Sentinel Hub light-client",
}

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()

	rpc.AddCommands(rootCmd)
	rootCmd.AddCommand(client.LineBreak)
	tx.AddCommands(rootCmd, cdc)
	rootCmd.AddCommand(client.LineBreak)

	rootCmd.AddCommand(
		client.GetCommands(
			authCli.GetAccountCmd("acc", cdc, types.GetAccountDecoder(cdc)),
		)...)

	rootCmd.AddCommand(
		client.PostCommands(
			bankCli.SendTxCmd(cdc),
			ibcCli.IBCRelayCmd(cdc),
		)...)

	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "SH", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

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
	ibcCli "github.com/cosmos/cosmos-sdk/x/ibc/client/cli"
	vpnCli "github.com/ironman0x7b2/sentinel-hub/x/vpn/client/cli"
	"github.com/ironman0x7b2/sentinel-hub/app"
	"github.com/ironman0x7b2/sentinel-hub/types"
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
			ibcCli.IBCTransferCmd(cdc),
			ibcCli.IBCRelayCmd(cdc),
			vpnCli.RegisterVpnCmd(cdc),
		)...)

	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "BC", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

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
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/ironman0x7b2/sentinel-sdk/apps/sentinel-vpn"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	ibcCli "github.com/ironman0x7b2/sentinel-sdk/x/ibc/client/cli"
	vpnCli "github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/cli"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-vpn-cli",
	Short: "Sentinel VPN light-client",
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
			vpnCli.PayVPNServiceCommand(cdc),
			vpnCli.ChangeNodeStatusCommand(cdc),
			vpnCli.RegisterVPNCmd(cdc),
			vpnCli.ChangeSessionStatusCommand(cdc),
			vpnCli.DeregisterVPNCmd(cdc),
		)...)

	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "SV", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

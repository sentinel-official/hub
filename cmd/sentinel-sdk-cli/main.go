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
	slashingCli "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	stakeCli "github.com/cosmos/cosmos-sdk/x/stake/client/cli"
	"github.com/ironman0x7b2/sentinel-sdk/app"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
)

var rootCmd = &cobra.Command{
	Use:   "sentinel-sdk-cli",
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
			stakeCli.GetCmdQueryValidator("stake", cdc),
			stakeCli.GetCmdQueryValidators("stake", cdc),
			stakeCli.GetCmdQueryDelegation("stake", cdc),
			stakeCli.GetCmdQueryDelegations("stake", cdc),
			stakeCli.GetCmdQueryPool("stake", cdc),
			stakeCli.GetCmdQueryParams("stake", cdc),
			stakeCli.GetCmdQueryUnbondingDelegation("stake", cdc),
			stakeCli.GetCmdQueryUnbondingDelegations("stake", cdc),
			stakeCli.GetCmdQueryRedelegation("stake", cdc),
			stakeCli.GetCmdQueryRedelegations("stake", cdc),
			slashingCli.GetCmdQuerySigningInfo("slashing", cdc),
			authCli.GetAccountCmd("acc", cdc, types.GetAccountDecoder(cdc)),
		)...)

	rootCmd.AddCommand(
		client.PostCommands(
			bankCli.SendTxCmd(cdc),
			ibcCli.IBCTransferCmd(cdc),
			ibcCli.IBCRelayCmd(cdc),
			stakeCli.GetCmdCreateValidator(cdc),
			stakeCli.GetCmdEditValidator(cdc),
			stakeCli.GetCmdDelegate(cdc),
			stakeCli.GetCmdUnbond("stake", cdc),
			stakeCli.GetCmdRedelegate("stake", cdc),
			slashingCli.GetCmdUnjail(cdc),
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

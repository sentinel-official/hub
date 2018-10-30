package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	ibcCli "github.com/cosmos/cosmos-sdk/x/ibc/client/cli"
	stakeCli "github.com/cosmos/cosmos-sdk/x/stake/client/cli"
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
			stakeCli.GetCmdQueryValidator("stake", cdc),
			stakeCli.GetCmdQueryValidators("stake", cdc),
			stakeCli.GetCmdQueryDelegation("stake", cdc),
			stakeCli.GetCmdQueryDelegations("stake", cdc),
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
		)...)

	rootCmd.AddCommand(
		client.LineBreak,
		lcd.ServeCommand(cdc),
		keys.Commands(),
		client.LineBreak,
		version.VersionCmd,
	)

	executor := cli.PrepareMainCmd(rootCmd, "BC", os.ExpandEnv("$HOME/.sentinel-hub-cli"))
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

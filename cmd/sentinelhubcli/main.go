package main

import (
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankcli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/types"
)

func main() {
	types.GetConfig().Seal()
	cobra.EnableCommandSorting = false

	var (
		cdc = hub.MakeCodec()
		cmd = &cobra.Command{
			Use:   "sentinelhubcli",
			Short: "Sentinel Hub Command-line Interface (light-client)",
		}
	)

	cmd.PersistentFlags().String(flags.FlagChainID, "", "Chain ID of Tendermint node")
	cmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(cmd)
	}

	cmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(hub.DefaultCLIHome),
		flags.LineBreak,
		keys.Commands(),
		queryCmd(cdc),
		txCmd(cdc),
		lcd.ServeCommand(cdc, registerRoutes),
		flags.LineBreak,
		version.Cmd,
		flags.NewCompletionCmd(cmd, true),
	)

	executor := cli.PrepareMainCmd(cmd, "SENTINELHUB", hub.DefaultCLIHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func queryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query sub-commands",
	}

	cmd.AddCommand(
		rpc.BlockCommand(),
		rpc.ValidatorCommand(cdc),
		flags.LineBreak,
		authcli.GetAccountCmd(cdc),
		authcli.QueryTxCmd(cdc),
		authcli.QueryTxsByEventsCmd(cdc),
		flags.LineBreak,
	)

	hub.ModuleBasics.AddQueryCommands(cmd, cdc)
	return cmd
}

func txCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Transaction sub-commands",
	}

	cmd.AddCommand(
		authcli.GetSignCommand(cdc),
		authcli.GetMultiSignCommand(cdc),
		authcli.GetEncodeCommand(cdc),
		authcli.GetBroadcastCommand(cdc),
		flags.LineBreak,
		bankcli.SendTxCmd(cdc),
		flags.LineBreak,
	)

	hub.ModuleBasics.AddTxCommands(cmd, cdc)
	return cmd
}

func registerRoutes(rs *lcd.RestServer) {
	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authrest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
	hub.ModuleBasics.RegisterRESTRoutes(rs.CliCtx, rs.Mux)
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	file := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(file); err == nil {
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.BindPFlag(flags.FlagChainID, cmd.PersistentFlags().Lookup(flags.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag)); err != nil {
		return err
	}

	return nil
}

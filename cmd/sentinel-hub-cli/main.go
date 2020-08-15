package main

import (
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/version"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/types"
)

func main() {
	cdc := hub.MakeCodec()
	types.GetConfig().Seal()

	cobra.EnableCommandSorting = false
	cmd := &cobra.Command{
		Use:   "sentinel-hub-cli",
		Short: "Sentinel Hub Command-line Interface (light-client)",
	}

	cmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of Tendermint node")
	cmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(cmd)
	}

	cmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(hub.DefaultCLIHome),
		client.LineBreak,
		keys.Commands(),
		queryCmd(cdc),
		txCmd(cdc),
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		version.Cmd,
		client.NewCompletionCmd(cmd, true),
	)

	executor := cli.PrepareMainCmd(cmd, "SENTINEL_HUB", hub.DefaultCLIHome)
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
		client.LineBreak,
		authCli.GetAccountCmd(cdc),
		authCli.QueryTxCmd(cdc),
		authCli.QueryTxsByEventsCmd(cdc),
		client.LineBreak,
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
		authCli.GetSignCommand(cdc),
		authCli.GetMultiSignCommand(cdc),
		authCli.GetEncodeCommand(cdc),
		authCli.GetBroadcastCommand(cdc),
		client.LineBreak,
		bankCli.SendTxCmd(cdc),
		client.LineBreak,
	)

	hub.ModuleBasics.AddTxCommands(cmd, cdc)
	return cmd
}

func registerRoutes(rs *lcd.RestServer) {
	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authRest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
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

	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
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

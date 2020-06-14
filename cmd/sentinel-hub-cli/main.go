package main

import (
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/sentinel-official/hub/types"

	hub "github.com/sentinel-official/hub/app"
)

func main() {
	cdc := hub.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(types.Bech32PrefixAccAddr, types.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr, types.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr, types.Bech32PrefixConsPub)
	config.Seal()

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
		queryCmd(cdc),
		txCmd(cdc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.Cmd,
		client.NewCompletionCmd(cmd, true),
	)

	executor := cli.PrepareMainCmd(cmd, "SENTINEL_HUB", hub.DefaultCLIHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func queryCmd(cdc *amino.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Query sub-commands",
	}

	cmd.AddCommand(
		authCli.GetAccountCmd(cdc),
		client.LineBreak,
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		authCli.QueryTxsByEventsCmd(cdc),
		authCli.QueryTxCmd(cdc),
		client.LineBreak,
	)

	hub.ModuleBasics.AddQueryCommands(cmd, cdc)
	return cmd
}

func txCmd(cdc *amino.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Transaction sub-commands",
	}

	cmd.AddCommand(
		bankCli.SendTxCmd(cdc),
		client.LineBreak,
		authCli.GetSignCommand(cdc),
		authCli.GetMultiSignCommand(cdc),
		client.LineBreak,
		authCli.GetBroadcastCommand(cdc),
		authCli.GetEncodeCommand(cdc),
		client.LineBreak,
	)

	hub.ModuleBasics.AddTxCommands(cmd, cdc)

	for _, cmd := range cmd.Commands() {
		if cmd.Use == auth.ModuleName || cmd.Use == bank.ModuleName {
			cmd.RemoveCommand(cmd)
		}
	}

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

package main

import (
	"os"
	"path"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"
	
	"github.com/sentinel-official/hub/app"
	"github.com/sentinel-official/hub/simapp"
	"github.com/sentinel-official/hub/version"
)

func main() {
	cdc := app.MakeCodec()
	
	config := sdk.GetConfig()
	simapp.SetBech32AddressPrefixes(config)
	config.Seal()
	
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:   "sentinel-hubcli",
		Short: "Sentinel Hub light-client",
	}
	
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}
	
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(app.DefaultCLIHome),
		queryCmd(cdc),
		txCmd(cdc),
		client.LineBreak,
		lcd.ServeCommand(cdc, registerRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
		version.Cmd,
		client.NewCompletionCmd(rootCmd, true),
	)
	
	executor := cli.PrepareMainCmd(rootCmd, "HUB", app.DefaultCLIHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func queryCmd(cdc *_amino.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
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
	
	app.ModuleBasics.AddQueryCommands(cmd, cdc)
	return cmd
}

func txCmd(cdc *_amino.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
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
	
	app.ModuleBasics.AddTxCommands(cmd, cdc)
	
	var cmdsToRemove []*cobra.Command
	for _, cmd := range cmd.Commands() {
		if cmd.Use == auth.ModuleName || cmd.Use == bank.ModuleName {
			cmdsToRemove = append(cmdsToRemove, cmd)
		}
	}
	
	cmd.RemoveCommand(cmdsToRemove...)
	return cmd
}

func registerRoutes(rs *lcd.RestServer) {
	client.RegisterRoutes(rs.CliCtx, rs.Mux)
	authRest.RegisterTxRoutes(rs.CliCtx, rs.Mux)
	app.ModuleBasics.RegisterRESTRoutes(rs.CliCtx, rs.Mux)
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}
	
	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)
		
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
	
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}

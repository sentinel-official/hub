package main

import (
	"net/http"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bankRest "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	crisisClient "github.com/cosmos/cosmos-sdk/x/crisis/client"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distRest "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govClient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govRest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintClient "github.com/cosmos/cosmos-sdk/x/mint/client"
	mintRest "github.com/cosmos/cosmos-sdk/x/mint/client/rest"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingClient "github.com/cosmos/cosmos-sdk/x/slashing/client"
	slashingRest "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingClient "github.com/cosmos/cosmos-sdk/x/staking/client"
	stakingRest "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_amino "github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/sentinel-official/sentinel-hub/app"
	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/version"
	depositClient "github.com/sentinel-official/sentinel-hub/x/deposit/client"
	depositRest "github.com/sentinel-official/sentinel-hub/x/deposit/client/rest"
	vpnClient "github.com/sentinel-official/sentinel-hub/x/vpn/client"
	vpnRest "github.com/sentinel-official/sentinel-hub/x/vpn/client/rest"
)

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hub.Bech32PrefixAccAddr, hub.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hub.Bech32PrefixValAddr, hub.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hub.Bech32PrefixConsAddr, hub.Bech32PrefixConsPub)
	config.Seal()

	mc := []sdk.ModuleClients{
		govClient.NewModuleClient(gov.StoreKey, cdc),
		distClient.NewModuleClient(distribution.StoreKey, cdc),
		stakingClient.NewModuleClient(staking.StoreKey, cdc),
		slashingClient.NewModuleClient(slashing.StoreKey, cdc),
		mintClient.NewModuleClient(mint.StoreKey, cdc),
		crisisClient.NewModuleClient(slashing.StoreKey, cdc),
		depositClient.NewModuleClient(cdc),
		vpnClient.NewModuleClient(cdc),
	}

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
		queryCmd(cdc, mc),
		txCmd(cdc, mc),
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

func queryCmd(cdc *_amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authCli.GetAccountCmd(auth.StoreKey, cdc),
	)

	for _, m := range mc {
		mQueryCmd := m.GetQueryCmd()
		if mQueryCmd != nil {
			queryCmd.AddCommand(mQueryCmd)
		}
	}

	return queryCmd
}

func txCmd(cdc *_amino.Codec, mc []sdk.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankCli.SendTxCmd(cdc),
		client.LineBreak,
		authCli.GetSignCommand(cdc),
		authCli.GetMultiSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
		client.LineBreak,
	)

	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}

	return txCmd
}

func registerRoutes(rs *lcd.RestServer) {
	// registerSwaggerUI(rs)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	authRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, auth.StoreKey)
	bankRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	stakingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	distRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, distribution.StoreKey)
	govRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	mintRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	depositRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	vpnRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
}

// nolint:deadcode,unused
func registerSwaggerUI(rs *lcd.RestServer) {
	staticFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(staticFS)
	rs.Mux.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
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

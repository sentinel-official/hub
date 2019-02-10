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
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authCli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankCli "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bankRest "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distRest "github.com/cosmos/cosmos-sdk/x/distribution/client/rest"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govClient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govRest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
	ibcCli "github.com/cosmos/cosmos-sdk/x/ibc/client/cli"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingClient "github.com/cosmos/cosmos-sdk/x/slashing/client"
	slashingRest "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingClient "github.com/cosmos/cosmos-sdk/x/staking/client"
	stakingRest "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/ironman0x7b2/sentinel-sdk/apps/hub"
	"github.com/ironman0x7b2/sentinel-sdk/version"
)

func main() {
	cdc := app.MakeCodec()

	config := csdkTypes.GetConfig()
	config.SetBech32PrefixForAccount(csdkTypes.Bech32PrefixAccAddr, csdkTypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(csdkTypes.Bech32PrefixValAddr, csdkTypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(csdkTypes.Bech32PrefixConsAddr, csdkTypes.Bech32PrefixConsPub)
	config.Seal()

	mc := []csdkTypes.ModuleClients{
		govClient.NewModuleClient(gov.StoreKey, cdc),
		distClient.NewModuleClient(distribution.StoreKey, cdc),
		stakingClient.NewModuleClient(staking.StoreKey, cdc),
		slashingClient.NewModuleClient(slashing.StoreKey, cdc),
	}

	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:   "hubcli",
		Short: "Hub light-client",
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
		version.VersionCmd,
		client.NewCompletionCmd(rootCmd, true),
	)

	executor := cli.PrepareMainCmd(rootCmd, "HUB", app.DefaultCLIHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func queryCmd(cdc *amino.Codec, mc []csdkTypes.ModuleClients) *cobra.Command {
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
		queryCmd.AddCommand(m.GetQueryCmd())
	}

	return queryCmd
}

func txCmd(cdc *amino.Codec, mc []csdkTypes.ModuleClients) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankCli.SendTxCmd(cdc),
		client.LineBreak,
		authCli.GetSignCommand(cdc),
		authCli.GetMultiSignCommand(cdc),
		authCli.GetBroadcastCommand(cdc),
		authCli.GetEncodeCommand(cdc),
		ibcCli.IBCTransferCmd(cdc),
		ibcCli.IBCRelayCmd(cdc),
		client.LineBreak,
	)

	for _, m := range mc {
		txCmd.AddCommand(m.GetTxCmd())
	}

	return txCmd
}

func registerRoutes(rs *lcd.RestServer) {
	// registerSwaggerUI(rs)
	keys.RegisterRoutes(rs.Mux, rs.CliCtx.Indent)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	authRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, auth.StoreKey)
	bankRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	stakingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashingRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	distRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, distribution.StoreKey)
	govRest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
}

func registerSwaggerUI(rs *lcd.RestServer) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
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

package main

import (
	"os"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/client"
	clientconfig "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/sentinel-official/hub/app"
	hubtypes "github.com/sentinel-official/hub/types"
)

func initAppConfig() (string, interface{}) {
	type Config struct {
		*serverconfig.Config
	}

	cfg := Config{Config: serverconfig.DefaultConfig()}
	cfg.BaseConfig.MinGasPrices = "0.1udvpn"
	cfg.StateSync.SnapshotInterval = 1000

	cfgTemplate := serverconfig.DefaultConfigTemplate

	return cfgTemplate, cfg
}

func moduleInitFlags(cmd *cobra.Command) {
	crisis.AddModuleInitFlags(cmd)
	wasm.AddModuleInitFlags(cmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcli.QueryTxsByEventsCmd(),
		authcli.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcli.GetSignCommand(),
		authcli.GetSignBatchCommand(),
		authcli.GetMultiSignCommand(),
		authcli.GetMultiSignBatchCmd(),
		authcli.GetValidateSignaturesCommand(),
		authcli.GetBroadcastCommand(),
		authcli.GetEncodeCommand(),
		authcli.GetDecodeCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func NewRootCmd(homeDir string) *cobra.Command {
	encCfg := app.DefaultEncodingConfig()
	cmd := &cobra.Command{
		Use:   "sentinelhub",
		Short: "Sentinel Hub application",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) (err error) {
			clientCtx := client.Context{}.
				WithCodec(encCfg.Codec).
				WithInterfaceRegistry(encCfg.InterfaceRegistry).
				WithTxConfig(encCfg.TxConfig).
				WithLegacyAmino(encCfg.Amino).
				WithInput(os.Stdin).
				WithAccountRetriever(authtypes.AccountRetriever{}).
				WithHomeDir(homeDir).
				WithViper("SENTINELHUB")

			clientCtx, err = client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err = clientconfig.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			cfgTemplate, cfg := initAppConfig()
			return server.InterceptConfigsPreRunHandler(cmd, cfgTemplate, cfg)
		},
	}

	cfg := hubtypes.GetConfig()
	cfg.Seal()

	cmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, homeDir),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, homeDir),
		genutilcli.GenTxCmd(app.ModuleBasics, encCfg.TxConfig, banktypes.GenesisBalancesIterator{}, homeDir),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics),
		AddGenesisAccountCmd(homeDir),
		AddGenesisWasmMsgCmd(homeDir),
		tmcli.NewCompletionCmd(cmd, true),
		debug.Cmd(),
		clientconfig.Cmd(),
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(homeDir),
	)

	creator := AppCreator{encCfg: encCfg, homeDir: homeDir}
	server.AddCommands(cmd, homeDir, creator.NewApp, creator.AppExport, moduleInitFlags)

	return cmd
}

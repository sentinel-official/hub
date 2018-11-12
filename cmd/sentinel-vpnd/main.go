package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/tendermint/tendermint/p2p"
	tmTypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	gaiaInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	app "github.com/ironman0x7b2/sentinel-sdk/apps/sentinel-vpn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/common"
	tmDb "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

const flagClientHome = "home-client"

func main() {
	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "sentinel-vpnd",
		Short:             "Sentinel VPN Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	appInit := server.DefaultAppInit
	rootCmd.AddCommand(InitCmd(ctx, cdc, appInit))
	rootCmd.AddCommand(gaiaInit.TestnetFilesCmd(ctx, cdc, appInit))

	server.AddCommands(ctx, cdc, rootCmd, appInit, newApp, exportAppStateAndTMValidators)

	rootDir := app.DefaultNodeHome
	executor := cli.PrepareBaseCmd(rootCmd, "SV", rootDir)

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func InitCmd(ctx *server.Context, cdc *codec.Codec, appInit server.AppInit) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize genesis config, priv-validator file, and p2p-node file",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {

			config := ctx.Config
			config.SetRoot(viper.GetString(cli.HomeFlag))
			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", common.RandStr(6))
			}

			nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
			if err != nil {
				return err
			}
			nodeID := string(nodeKey.ID())

			pk := gaiaInit.ReadOrCreatePrivValidator(config.PrivValidatorFile())
			genTx, appMessage, validator, err := server.SimpleAppGenTx(cdc, pk)
			if err != nil {
				return err
			}

			appState, err := appInit.AppGenState(cdc, []json.RawMessage{genTx})
			if err != nil {
				return err
			}
			appStateJSON, err := cdc.MarshalJSON(appState)
			if err != nil {
				return err
			}

			toPrint := struct {
				ChainID    string          `json:"chain_id"`
				NodeID     string          `json:"node_id"`
				AppMessage json.RawMessage `json:"app_message"`
			}{
				chainID,
				nodeID,
				appMessage,
			}
			out, err := codec.MarshalJSONIndent(cdc, toPrint)
			if err != nil {
				return err
			}
			fmt.Fprintf(os.Stderr, "%s\n", string(out))
			return gaiaInit.WriteGenesisFile(config.GenesisFile(), chainID, []tmTypes.GenesisValidator{validator}, appStateJSON)
		},
	}

	cmd.Flags().String(cli.HomeFlag, app.DefaultNodeHome, "node's home directory")
	cmd.Flags().String(flagClientHome, app.DefaultCLIHome, "client's home directory")
	cmd.Flags().String(client.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().String(client.FlagName, "", "validator's moniker")
	cmd.MarkFlagRequired(client.FlagName)
	return cmd
}

func newApp(logger log.Logger, db tmDb.DB, storeTracer io.Writer) abciTypes.Application {
	return app.NewSentinelVPN(logger, db, baseapp.SetPruning(viper.GetString("pruning")))
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDb.DB, storeTracer io.Writer) (json.RawMessage, []tmTypes.GenesisValidator, error) {
	bapp := app.NewSentinelVPN(logger, db)
	return bapp.ExportAppStateAndValidators()
}

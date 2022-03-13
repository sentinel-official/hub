package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	"github.com/cosmos/ibc-go/v2/modules/core/exported"
	ibctypes "github.com/cosmos/ibc-go/v2/modules/core/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"

	swaptypes "github.com/sentinel-official/hub/x/swap/types"
	v05swaptypes "github.com/sentinel-official/hub/x/swap/types/legacy/v0.5"
	v06swaptypes "github.com/sentinel-official/hub/x/swap/types/legacy/v0.6"
	vpntypes "github.com/sentinel-official/hub/x/vpn/types"
	v05vpntypes "github.com/sentinel-official/hub/x/vpn/types/legacy/v0.5"
	v06vpntypes "github.com/sentinel-official/hub/x/vpn/types/legacy/v0.6"
)

const (
	flagGenesisTime   = "genesis-time"
	flagInitialHeight = "initial-height"
)

func migrateCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "migrate [genesis-file]",
		Short: "Migrate Genesis file from v0.5 to v0.6",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var ctx = client.GetClientContextFromCmd(cmd)

			blob, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			chainID, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}

			genesisTime, err := cmd.Flags().GetString(flagGenesisTime)
			if err != nil {
				return err
			}

			initialHeight, err := cmd.Flags().GetInt64(flagInitialHeight)
			if err != nil {
				return err
			}

			genesis, err := tmtypes.GenesisDocFromJSON(blob)
			if err != nil {
				return err
			}

			var state genutiltypes.AppMap
			if err := json.Unmarshal(genesis.AppState, &state); err != nil {
				return err
			}

			state, err = migrateFunc(state, ctx)
			if err != nil {
				return err
			}

			var (
				bankGenesis    banktypes.GenesisState
				stakingGenesis stakingtypes.GenesisState
				vpnGenesis     vpntypes.GenesisState
			)

			ctx.JSONMarshaler.MustUnmarshalJSON(state[banktypes.ModuleName], &bankGenesis)
			ctx.JSONMarshaler.MustUnmarshalJSON(state[stakingtypes.ModuleName], &stakingGenesis)
			ctx.JSONMarshaler.MustUnmarshalJSON(state[vpntypes.ModuleName], &vpnGenesis)

			bankGenesis.DenomMetadata = []banktypes.Metadata{
				{
					Description: "The native staking token of the Sentinel Hub.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: "dvpn", Exponent: uint32(6), Aliases: []string{}},
						{Denom: "mdvpn", Exponent: uint32(3), Aliases: []string{"millidvpn"}},
						{Denom: "udvpn", Exponent: uint32(0), Aliases: []string{"microdvpn"}},
					},
					Base:    "udvpn",
					Display: "dvpn",
				},
			}

			stakingGenesis.Params.HistoricalEntries = 10000

			vpnGenesis.Nodes.Params.InactiveDuration = 1 * time.Hour
			vpnGenesis.Sessions.Params.InactiveDuration = 2 * time.Hour
			vpnGenesis.Subscriptions.Params.InactiveDuration = 4 * time.Hour

			var (
				ibcTransferGenesis = ibctransfertypes.DefaultGenesisState()
				ibcGenesis         = ibctypes.DefaultGenesisState()
				capabilityGenesis  = capabilitytypes.DefaultGenesis()
				evidenceGenesis    = evidencetypes.DefaultGenesisState()
			)

			ibcTransferGenesis.Params.ReceiveEnabled = true
			ibcTransferGenesis.Params.SendEnabled = true
			ibcGenesis.ClientGenesis.Params.AllowedClients = []string{exported.Tendermint}

			state[banktypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(&bankGenesis)
			state[ibctransfertypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(ibcTransferGenesis)
			state[ibchost.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(ibcGenesis)
			state[capabilitytypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(capabilityGenesis)
			state[evidencetypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(evidenceGenesis)
			state[stakingtypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(&stakingGenesis)
			state[vpntypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(&vpnGenesis)

			genesis.AppState, err = json.Marshal(state)
			if err != nil {
				return err
			}

			if genesisTime != "" {
				var t time.Time
				if err := t.UnmarshalText([]byte(genesisTime)); err != nil {
					return err
				}

				genesis.GenesisTime = t
			}
			if chainID != "" {
				genesis.ChainID = chainID
			}

			genesis.InitialHeight = initialHeight

			blob, err = tmjson.Marshal(genesis)
			if err != nil {
				return err
			}

			sortedBlob, err := sdk.SortJSON(blob)
			if err != nil {
				return err
			}

			fmt.Println(string(sortedBlob))
			return nil
		},
	}

	cmd.Flags().String(flags.FlagChainID, "", "set chain id")
	cmd.Flags().String(flagGenesisTime, "", "set genesis time")
	cmd.Flags().Int64(flagInitialHeight, 0, "set the initial height")

	return &cmd
}

func migrateFunc(state genutiltypes.AppMap, ctx client.Context) (genutiltypes.AppMap, error) {
	migrateFunc := cli.GetMigrationCallback("v0.40")
	if migrateFunc == nil {
		return nil, fmt.Errorf("sdk migration function is not available")
	}

	state = migrateFunc(state, ctx)

	var (
		swapGenesis v05swaptypes.GenesisState
		vpnGenesis  v05vpntypes.GenesisState
		amino       = codec.NewLegacyAmino()
	)

	amino.MustUnmarshalJSON(state[v05swaptypes.ModuleName], &swapGenesis)
	amino.MustUnmarshalJSON(state[v05vpntypes.ModuleName], &vpnGenesis)

	state[swaptypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(v06swaptypes.MigrateGenesisState(&swapGenesis))
	state[vpntypes.ModuleName] = ctx.JSONMarshaler.MustMarshalJSON(v06vpntypes.MigrateGenesisState(&vpnGenesis))

	return state, nil
}

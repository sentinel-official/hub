package upgrades

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibckeeper "github.com/cosmos/ibc-go/v7/modules/core/keeper"
	ibctmmigrations "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint/migrations"
)

func Handler(
	cdc codec.Codec,
	mm *module.Manager,
	configurator module.Configurator,
	consensusKeeper consensuskeeper.Keeper,
	govKeeper *govkeeper.Keeper,
	paramsKeeper paramskeeper.Keeper,
	stakingKeeper *stakingkeeper.Keeper,
	ibcKeeper *ibckeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		keyTables := map[string]paramstypes.KeyTable{
			// Cosmos SDK subspaces
			authtypes.ModuleName:         authtypes.ParamKeyTable(),
			banktypes.ModuleName:         banktypes.ParamKeyTable(),
			crisistypes.ModuleName:       crisistypes.ParamKeyTable(),
			distributiontypes.ModuleName: distributiontypes.ParamKeyTable(),
			govtypes.ModuleName:          govv1types.ParamKeyTable(),
			minttypes.ModuleName:         minttypes.ParamKeyTable(),
			slashingtypes.ModuleName:     slashingtypes.ParamKeyTable(),
			stakingtypes.ModuleName:      stakingtypes.ParamKeyTable(),

			// Other subspaces
			wasmtypes.ModuleName: wasmtypes.ParamKeyTable(),
		}

		for name, table := range keyTables {
			subspace, ok := paramsKeeper.GetSubspace(name)
			if !ok {
				return nil, fmt.Errorf("params subspace does not exist for module: %s", name)
			}
			if subspace.HasKeyTable() {
				continue
			}

			subspace.WithKeyTable(table)
		}

		legacyParamStore := paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, legacyParamStore, &consensusKeeper)

		_, err := ibctmmigrations.PruneExpiredConsensusStates(ctx, cdc, ibcKeeper.ClientKeeper)
		if err != nil {
			return nil, err
		}

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		govParams := govKeeper.GetParams(ctx)
		govParams.MinInitialDepositRatio = sdkmath.LegacyNewDecWithPrec(2, 1).String()
		if err := govKeeper.SetParams(ctx, govParams); err != nil {
			return nil, err
		}

		stakingParams := stakingKeeper.GetParams(ctx)
		stakingParams.MinCommissionRate = sdkmath.LegacyNewDecWithPrec(5, 2)
		if err := stakingKeeper.SetParams(ctx, stakingParams); err != nil {
			return nil, err
		}

		ibcClientParams := ibcKeeper.ClientKeeper.GetParams(ctx)
		ibcClientParams.AllowedClients = append(ibcClientParams.AllowedClients, exported.Localhost)
		ibcKeeper.ClientKeeper.SetParams(ctx, ibcClientParams)

		return newVM, nil
	}
}

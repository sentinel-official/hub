package upgrade3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcconnectionkeeper "github.com/cosmos/ibc-go/v2/modules/core/03-connection/keeper"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v2/modules/core/03-connection/types"
)

func Handler(
	mm *module.Manager,
	configurator module.Configurator,
	ibcConnectionKeeper *ibcconnectionkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
		ibcConnectionKeeper.SetParams(ctx, ibcconnectiontypes.DefaultParams())

		fromVM := make(map[string]uint64)
		for moduleName := range mm.Modules {
			fromVM[moduleName] = 1
		}
		fromVM[authtypes.ModuleName] = 2

		delete(fromVM, authz.ModuleName)
		delete(fromVM, feegrant.ModuleName)

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}
		newVM[authtypes.ModuleName] = 1

		return mm.RunMigrations(ctx, configurator, newVM)
	}
}

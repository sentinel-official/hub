package upgrades

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcicacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	ibcicahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	ibcicahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	ibcicatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
)

func Handler(
	mm *module.Manager,
	configurator module.Configurator,
	ibcICAControllerKeeper ibcicacontrollerkeeper.Keeper,
	ibcICAHostKeeper ibcicahostkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		fromVM[ibcicatypes.ModuleName] = mm.Modules[ibcicatypes.ModuleName].ConsensusVersion()
		fromVM[ibcfeetypes.ModuleName] = mm.Modules[ibcfeetypes.ModuleName].ConsensusVersion()

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		var (
			controllerParams = ibcicacontrollertypes.Params{
				ControllerEnabled: true,
			}
			hostParams = ibcicahosttypes.Params{
				HostEnabled:   true,
				AllowMessages: []string{"*"},
			}
		)

		ibcICAControllerKeeper.SetParams(ctx, controllerParams)
		ibcICAHostKeeper.SetParams(ctx, hostParams)

		return newVM, nil
	}
}

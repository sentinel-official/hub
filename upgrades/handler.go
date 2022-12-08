package upgrades

import (
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	ibcicahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	ibcicatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	nodetypes "github.com/sentinel-official/hub/x/node/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	vpnkeeper "github.com/sentinel-official/hub/x/vpn/keeper"
)

func Handler(mm *module.Manager, configurator module.Configurator, paramsStoreKey sdk.StoreKey, vpnKeeper vpnkeeper.Keeper, wasmKeeper wasmkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		var (
			controllerParams = ibcicacontrollertypes.Params{}
			hostParams       = ibcicahosttypes.Params{
				HostEnabled: true,
				AllowMessages: []string{
					sdk.MsgTypeURL(&authz.MsgExec{}),
					sdk.MsgTypeURL(&authz.MsgGrant{}),
					sdk.MsgTypeURL(&authz.MsgRevoke{}),
					sdk.MsgTypeURL(&banktypes.MsgSend{}),
					sdk.MsgTypeURL(&distributiontypes.MsgFundCommunityPool{}),
					sdk.MsgTypeURL(&distributiontypes.MsgSetWithdrawAddress{}),
					sdk.MsgTypeURL(&distributiontypes.MsgWithdrawDelegatorReward{}),
					sdk.MsgTypeURL(&distributiontypes.MsgWithdrawValidatorCommission{}),
					sdk.MsgTypeURL(&govtypes.MsgVote{}),
					sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
					sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
					sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
					sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
				},
			}
		)

		icaModule, ok := mm.Modules[ibcicatypes.ModuleName].(ibcica.AppModule)
		if !ok {
			panic("mm.Modules[ibcicatypes.ModuleName] is not of type ibcica.AppModule")
		}

		icaModule.InitModule(ctx, controllerParams, hostParams)

		fromVM[ibcicatypes.ModuleName] = mm.Modules[ibcicatypes.ModuleName].ConsensusVersion()

		newVM, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return newVM, err
		}

		var (
			store = ctx.KVStore(paramsStoreKey)
			iter  = sdk.KVStorePrefixIterator(store, append([]byte("provider"), '/'))
		)

		for ; iter.Valid(); iter.Next() {
			ctx.Logger().Info("deleting the parameter", "key", iter.Key(), "value", iter.Value())
			store.Delete(iter.Key())
		}

		iter.Close()

		vpnKeeper.Provider.SetParams(
			ctx,
			providertypes.Params{
				Deposit:      sdk.NewInt64Coin("udvpn", 25000*1e6),
				StakingShare: sdk.NewDecWithPrec(2, 1),
			},
		)
		vpnKeeper.Node.SetParams(
			ctx,
			nodetypes.Params{
				Deposit:          sdk.NewInt64Coin("udvpn", 0),
				InactiveDuration: 60 * time.Minute,
				MaxPrice:         nil,
				MinPrice:         sdk.NewCoins(sdk.NewInt64Coin("udvpn", 100*1e6)),
				StakingShare:     sdk.NewDecWithPrec(2, 1),
			},
		)

		wasmParams := wasmKeeper.GetParams(ctx)
		wasmParams.CodeUploadAccess = wasmtypes.AllowNobody
		wasmKeeper.SetParams(ctx, wasmParams)

		return newVM, nil
	}
}

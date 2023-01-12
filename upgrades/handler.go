package upgrades

import (
	"fmt"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcica "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/types"
	ibcicahosttypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/host/types"
	ibcicatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"

	hubutils "github.com/sentinel-official/hub/utils"
	custommintkeeper "github.com/sentinel-official/hub/x/mint/keeper"
	customminttypes "github.com/sentinel-official/hub/x/mint/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	vpnkeeper "github.com/sentinel-official/hub/x/vpn/keeper"
)

func Handler(
	mm *module.Manager,
	configurator module.Configurator,
	paramsStoreKey sdk.StoreKey,
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	cmk custommintkeeper.Keeper,
	mk mintkeeper.Keeper,
	sk stakingkeeper.Keeper,
	vk vpnkeeper.Keeper,
	wk wasmkeeper.Keeper,
) upgradetypes.UpgradeHandler {
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
					sdk.MsgTypeURL(&govtypes.MsgVoteWeighted{}),
					sdk.MsgTypeURL(&ibctransfertypes.MsgTransfer{}),
					sdk.MsgTypeURL(&stakingtypes.MsgBeginRedelegate{}),
					sdk.MsgTypeURL(&stakingtypes.MsgCreateValidator{}),
					sdk.MsgTypeURL(&stakingtypes.MsgDelegate{}),
					sdk.MsgTypeURL(&stakingtypes.MsgEditValidator{}),
					sdk.MsgTypeURL(&stakingtypes.MsgUndelegate{}),
					sdk.MsgTypeURL(&wasmtypes.MsgExecuteContract{}),
					sdk.MsgTypeURL(&wasmtypes.MsgInstantiateContract{}),
					sdk.MsgTypeURL(&wasmtypes.MsgInstantiateContract2{}),
					sdk.MsgTypeURL(&wasmtypes.MsgStoreCode{}),
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

		denom := sk.BondDenom(ctx)

		vk.Provider.SetParams(
			ctx,
			providertypes.Params{
				Deposit:      sdk.NewInt64Coin(denom, 25000*1e6),
				StakingShare: sdk.NewDecWithPrec(2, 1),
			},
		)
		vk.Node.SetParams(
			ctx,
			nodetypes.Params{
				Deposit:          sdk.NewInt64Coin(denom, 0),
				InactiveDuration: 60 * time.Minute,
				MaxPrice:         nil,
				MinPrice:         sdk.NewCoins(sdk.NewInt64Coin(denom, 100*1e6)),
				StakingShare:     sdk.NewDecWithPrec(2, 1),
			},
		)

		wasmParams := wk.GetParams(ctx)
		wasmParams.CodeUploadAccess = wasmtypes.AllowNobody
		wk.SetParams(ctx, wasmParams)

		for _, voter := range proposal14Voters {
			var (
				startTime = ctx.BlockTime().Add(proposal14LockPeriod)
				endTime   = startTime.Add(proposal14VestingPeriod)
			)

			if err := createContinuousVestingAccount(ctx, ak, bk, mk, sk, denom, proposal14BonusRate, startTime, endTime, voter); err != nil {
				return nil, err
			}
		}

		for _, voter := range proposal15Voters {
			var (
				startTime = ctx.BlockTime().Add(proposal15LockPeriod)
				endTime   = startTime.Add(proposal15VestingPeriod)
			)

			if err := createContinuousVestingAccount(ctx, ak, bk, mk, sk, denom, proposal15BonusRate, startTime, endTime, voter); err != nil {
				return nil, err
			}
		}

		for _, voter := range proposal16Voters {
			var (
				startTime = ctx.BlockTime().Add(proposal16LockPeriod)
				endTime   = startTime.Add(proposal16VestingPeriod)
			)

			if err := createContinuousVestingAccount(ctx, ak, bk, mk, sk, denom, proposal16BonusRate, startTime, endTime, voter); err != nil {
				return nil, err
			}
		}

		cmk.IterateInflations(ctx, func(_ int, item customminttypes.Inflation) (stop bool) {
			cmk.DeleteInflation(ctx, item.Timestamp)
			return false
		})

		inflations := []customminttypes.Inflation{
			{
				Max:        sdk.NewDecWithPrec(37, 2),
				Min:        sdk.NewDecWithPrec(25, 2),
				RateChange: sdk.NewDecWithPrec(12, 2),
				Timestamp:  time.Date(2022, 9, 27, 12, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(25, 2),
				Min:        sdk.NewDecWithPrec(13, 2),
				RateChange: sdk.NewDecWithPrec(12, 2),
				Timestamp:  time.Date(2023, 9, 27, 12, 0, 0, 0, time.UTC),
			},
		}

		for _, inflation := range inflations {
			if err := inflation.Validate(); err != nil {
				return nil, err
			}
		}
		for _, inflation := range inflations {
			cmk.SetInflation(ctx, inflation)
		}

		return newVM, nil
	}
}

func createContinuousVestingAccount(
	ctx sdk.Context,
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	mk mintkeeper.Keeper,
	sk stakingkeeper.Keeper,
	denom string,
	bonusRate sdk.Dec,
	startTime time.Time,
	endTime time.Time,
	address string,
) error {
	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	account := ak.GetAccount(ctx, accAddr)
	if account == nil {
		return nil
	}

	switch account := account.(type) {
	case *authtypes.BaseAccount:
		return createContinuousVestingAccountFromBaseAccount(ctx, ak, bk, mk, sk, denom, bonusRate, startTime, endTime, account)
	case *vestingtypes.ContinuousVestingAccount:
		return createContinuousVestingAccountFromBaseAccount(ctx, ak, bk, mk, sk, denom, bonusRate, startTime, endTime, account.BaseAccount)
	case *vestingtypes.DelayedVestingAccount:
		return createContinuousVestingAccountFromBaseAccount(ctx, ak, bk, mk, sk, denom, bonusRate, startTime, endTime, account.BaseAccount)
	case *vestingtypes.PeriodicVestingAccount:
		return createContinuousVestingAccountFromBaseAccount(ctx, ak, bk, mk, sk, denom, bonusRate, startTime, endTime, account.BaseAccount)
	default:
		return fmt.Errorf("unknown account type for address %s", accAddr)
	}
}

func createContinuousVestingAccountFromBaseAccount(
	ctx sdk.Context,
	ak authkeeper.AccountKeeper,
	bk bankkeeper.Keeper,
	mk mintkeeper.Keeper,
	sk stakingkeeper.Keeper,
	denom string,
	bonusRate sdk.Dec,
	startTime time.Time,
	endTime time.Time,
	baseAccount *authtypes.BaseAccount,
) error {
	var (
		address             = baseAccount.GetAddress()
		balances            = bk.GetAllBalances(ctx, address)
		bondedDelegation    = sk.GetDelegatorBonded(ctx, address)
		unbondingDelegation = sk.GetDelegatorUnbonding(ctx, address)
		totalDelegation     = sdk.NewCoin(denom, bondedDelegation.Add(unbondingDelegation))
		totalBalances       = balances.Add(totalDelegation)
		bonus               = hubutils.GetProportionOfCoin(
			sdk.NewCoin(denom, totalBalances.AmountOf(denom)),
			bonusRate,
		)
	)

	ctx.Logger().Info("creating a continuous vesting account", "address", address, "total_balances", totalBalances, "bonus", bonus)

	if bonus.IsPositive() {
		if err := mk.MintCoins(ctx, sdk.NewCoins(bonus)); err != nil {
			return err
		}

		if err := bk.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, address, sdk.NewCoins(bonus)); err != nil {
			return err
		}

		totalBalances = totalBalances.Add(bonus)
	}

	continuousVestingAccount := vestingtypes.NewContinuousVestingAccount(
		baseAccount,
		sdk.NewCoins(
			sdk.NewCoin(denom, totalBalances.AmountOf(denom)),
		),
		startTime.Unix(),
		endTime.Unix(),
	)

	continuousVestingAccount.TrackDelegation(ctx.BlockTime(), totalBalances, sdk.NewCoins(totalDelegation))
	ak.SetAccount(ctx, continuousVestingAccount)

	return nil
}

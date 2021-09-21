package upgrade2

import (
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	custommintkeeper "github.com/sentinel-official/hub/x/mint/keeper"
	customminttypes "github.com/sentinel-official/hub/x/mint/types"
)

func Handler(
	setStoreLoader func(baseapp.StoreLoader),
	upgradeKeeper upgradekeeper.Keeper,
	customMintKeeper custommintkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	info, err := upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	if info.Name == Name && !upgradeKeeper.IsSkipHeight(info.Height) {
		upgrades := &storetypes.StoreUpgrades{
			Added: []string{customminttypes.ModuleName},
		}

		setStoreLoader(
			upgradetypes.UpgradeStoreLoader(
				info.Height,
				upgrades,
			),
		)
	}

	return func(ctx sdk.Context, _ upgradetypes.Plan) {
		if err := updateInflations(ctx, customMintKeeper); err != nil {
			panic(err)
		}
	}
}

func updateInflations(ctx sdk.Context, k custommintkeeper.Keeper) error {
	var (
		inflations = []customminttypes.Inflation{
			{
				Max:        sdk.NewDecWithPrec(49, 2),
				Min:        sdk.NewDecWithPrec(43, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2021, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(43, 2),
				Min:        sdk.NewDecWithPrec(37, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2022, 3, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(37, 2),
				Min:        sdk.NewDecWithPrec(31, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2022, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(31, 2),
				Min:        sdk.NewDecWithPrec(25, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2023, 3, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(25, 2),
				Min:        sdk.NewDecWithPrec(19, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2023, 9, 27, 0, 0, 0, 0, time.UTC),
			},
			{
				Max:        sdk.NewDecWithPrec(19, 2),
				Min:        sdk.NewDecWithPrec(13, 2),
				RateChange: sdk.NewDecWithPrec(6, 2),
				Timestamp:  time.Date(2024, 3, 27, 0, 0, 0, 0, time.UTC),
			},
		}
	)

	for _, inflation := range inflations {
		if err := inflation.Validate(); err != nil {
			return err
		}
	}

	for _, inflation := range inflations {
		k.SetInflation(ctx, inflation)
	}

	return nil
}

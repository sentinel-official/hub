package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	v1types "github.com/sentinel-official/hub/x/node/legacy/v1/types"
	"github.com/sentinel-official/hub/x/node/types"
)

type Migrator struct {
	Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{k}
}

func (k Migrator) Migrate2to3(ctx sdk.Context) error {
	if err := k.deleteActiveNodeKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteActiveNodeForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeForProviderKeys(ctx); err != nil {
		return err
	}
	if err := k.deleteInactiveNodeAtKeys(ctx); err != nil {
		return err
	}

	if err := k.setParams(ctx); err != nil {
		return err
	}
	if err := k.migrateNodes(ctx); err != nil {
		return err
	}

	return nil
}

func (k Migrator) deleteKeys(ctx sdk.Context, keyPrefix []byte) error {
	store := prefix.NewStore(k.Store(ctx), keyPrefix)

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}

	return nil
}

func (k Migrator) deleteActiveNodeKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x20})
}

func (k Migrator) deleteInactiveNodeKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x21})
}

func (k Migrator) deleteActiveNodeForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x30})
}

func (k Migrator) deleteInactiveNodeForProviderKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x31})
}

func (k Migrator) deleteInactiveNodeAtKeys(ctx sdk.Context) error {
	return k.deleteKeys(ctx, []byte{0x41})
}

func (k Migrator) migrateNodes(ctx sdk.Context) error {
	activeDuration := k.ActiveDuration(ctx)

	store := prefix.NewStore(k.Store(ctx), []byte{0x10})

	iter := store.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var value v1types.Node
		k.cdc.MustUnmarshal(iter.Value(), &value)
		store.Delete(iter.Key())

		node := types.Node{
			Address:        value.Address,
			GigabytePrices: value.Price,
			HourlyPrices:   value.Price,
			RemoteURL:      value.RemoteURL,
			InactiveAt:     time.Time{},
			Status:         value.Status,
			StatusAt:       value.StatusAt,
		}

		if node.Status.Equal(hubtypes.StatusActive) {
			node.InactiveAt = node.StatusAt.Add(activeDuration)
			k.SetNodeForInactiveAt(ctx, node.InactiveAt, node.GetAddress())
		}

		k.SetNode(ctx, node)
	}

	return nil
}

func (k Migrator) setParams(ctx sdk.Context) error {
	k.SetParams(
		ctx,
		types.Params{
			Deposit:        sdk.NewInt64Coin("udvpn", 0),
			ActiveDuration: 60 * time.Minute,
			MaxGigabytePrices: sdk.NewCoins(
				sdk.NewInt64Coin("ibc/31FEE1A2A9F9C01113F90BD0BBCCE8FD6BBB8585FAF109A2101827DD1D5B95B8", 7000000),
				sdk.NewInt64Coin("ibc/A8C2D23A1E6F95DA4E48BA349667E322BD7A6C996D8A4AAE8BA72E190F3D1477", 390000),
				sdk.NewInt64Coin("ibc/B1C0DDB14F25279A2026BC8794E12B259F8BDA546A3C5132CCAEE4431CE36783", 525000000),
				sdk.NewInt64Coin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", 5250000),
				sdk.NewInt64Coin("udvpn", 7000000000),
			),
			MinGigabytePrices: sdk.NewCoins(
				sdk.NewInt64Coin("ibc/31FEE1A2A9F9C01113F90BD0BBCCE8FD6BBB8585FAF109A2101827DD1D5B95B8", 105000),
				sdk.NewInt64Coin("ibc/A8C2D23A1E6F95DA4E48BA349667E322BD7A6C996D8A4AAE8BA72E190F3D1477", 6000),
				sdk.NewInt64Coin("ibc/B1C0DDB14F25279A2026BC8794E12B259F8BDA546A3C5132CCAEE4431CE36783", 800000),
				sdk.NewInt64Coin("ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518", 80000),
				sdk.NewInt64Coin("udvpn", 11000000),
			),
			MaxHourlyPrices:          nil,
			MinHourlyPrices:          nil,
			MaxSubscriptionGigabytes: 1e6,
			MinSubscriptionGigabytes: 1,
			MaxSubscriptionHours:     720,
			MinSubscriptionHours:     1,
			StakingShare:             sdk.NewDecWithPrec(2, 1),
		},
	)

	return nil
}

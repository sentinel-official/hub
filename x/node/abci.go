package node

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abcitypes.ValidatorUpdate {
	var (
		maxGigabytePricesModified = k.IsMaxGigabytePricesModified(ctx)
		minGigabytePricesModified = k.IsMinGigabytePricesModified(ctx)
		maxHourlyPricesModified   = k.IsMaxHourlyPricesModified(ctx)
		minHourlyPricesModified   = k.IsMinHourlyPricesModified(ctx)
	)

	if maxGigabytePricesModified || minGigabytePricesModified || maxHourlyPricesModified || minHourlyPricesModified {
		maxGigabytePrices := sdk.NewCoins()
		if maxGigabytePricesModified {
			maxGigabytePrices = k.MaxGigabytePrices(ctx)
		}

		minGigabytePrices := sdk.NewCoins()
		if minGigabytePricesModified {
			minGigabytePrices = k.MinGigabytePrices(ctx)
		}

		maxHourlyPrices := sdk.NewCoins()
		if maxHourlyPricesModified {
			maxHourlyPrices = k.MaxHourlyPrices(ctx)
		}

		minHourlyPrices := sdk.NewCoins()
		if minHourlyPricesModified {
			minHourlyPrices = k.MinHourlyPrices(ctx)
		}

		k.IterateNodes(ctx, func(_ int, item types.Node) bool {
			if item.GigabytePrices != nil {
				if maxGigabytePricesModified {
					for _, coin := range maxGigabytePrices {
						amount := item.GigabytePrices.AmountOf(coin.Denom)
						if amount.GT(coin.Amount) {
							item.GigabytePrices = item.GigabytePrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}

				if minGigabytePricesModified {
					for _, coin := range minGigabytePrices {
						amount := item.GigabytePrices.AmountOf(coin.Denom)
						if amount.LT(coin.Amount) {
							item.GigabytePrices = item.GigabytePrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}
			}

			if item.HourlyPrices != nil {
				if maxHourlyPricesModified {
					for _, coin := range maxHourlyPrices {
						amount := item.HourlyPrices.AmountOf(coin.Denom)
						if amount.GT(coin.Amount) {
							item.HourlyPrices = item.HourlyPrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}

				if minHourlyPricesModified {
					for _, coin := range minHourlyPrices {
						amount := item.HourlyPrices.AmountOf(coin.Denom)
						if amount.LT(coin.Amount) {
							item.HourlyPrices = item.HourlyPrices.Sub(
								sdk.NewCoins(
									sdk.NewCoin(coin.Denom, amount),
								),
							).Add(coin)
						}
					}
				}
			}

			k.SetNode(ctx, item)
			return false
		})
	}

	var (
		log              = k.Logger(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
	)

	k.IterateInactiveNodesAt(ctx, ctx.BlockTime(), func(_ int, item types.Node) bool {
		log.Info("found an inactive node", "address", item.Address)

		var (
			nodeAddr   = item.GetAddress()
			inactiveAt = item.StatusAt.Add(inactiveDuration)
		)

		k.DeleteActiveNode(ctx, nodeAddr)
		k.DeleteInactiveNodeAt(ctx, inactiveAt, nodeAddr)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetInactiveNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				Address: item.Address,
				Status:  item.Status,
			},
		)

		return false
	})

	return nil
}

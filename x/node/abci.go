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

	k.IterateNodesForExpiryAt(ctx, ctx.BlockTime(), func(_ int, item types.Node) bool {
		k.Logger(ctx).Info("found an expired node", "address", item.Address)

		nodeAddr := item.GetAddress()
		k.DeleteActiveNode(ctx, nodeAddr)
		k.DeleteNodeForExpiryAt(ctx, item.ExpiryAt, nodeAddr)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetNode(ctx, item)
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

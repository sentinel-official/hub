package keeper

import (
	"time"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
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
			k.Logger(ctx).Info("Updating prices for node", "address", item.Address)

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

			k.SetNode(ctx, item)
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateDetails{
					Address: item.Address,
				},
			)

			return false
		})
	}

	k.IterateNodesForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item types.Node) bool {
		k.Logger(ctx).Info("Found an inactive node", "address", item.Address)

		nodeAddr := item.GetAddress()
		k.DeleteActiveNode(ctx, nodeAddr)
		k.DeleteNodeForInactiveAt(ctx, item.InactiveAt, nodeAddr)

		item.InactiveAt = time.Time{}
		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()

		k.SetNode(ctx, item)
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				Status:  hubtypes.StatusInactive,
				Address: item.Address,
			},
		)

		return false
	})

	return nil
}

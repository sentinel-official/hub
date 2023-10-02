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
		log              = k.Logger(ctx)
		inactiveDuration = k.InactiveDuration(ctx)
		maxPriceModified = k.IsMaxPriceModified(ctx)
		minPriceModified = k.IsMinPriceModified(ctx)
	)

	if maxPriceModified || minPriceModified {
		var (
			maxPrice = sdk.NewCoins()
			minPrice = sdk.NewCoins()
		)

		if maxPriceModified {
			maxPrice = k.MaxPrice(ctx)
		}
		if minPriceModified {
			minPrice = k.MinPrice(ctx)
		}

		k.IterateNodes(ctx, func(_ int, item types.Node) (stop bool) {
			if item.Price == nil {
				return false
			}

			if maxPriceModified {
				for _, coin := range maxPrice {
					amount := item.Price.AmountOf(coin.Denom)
					if amount.GT(coin.Amount) {
						item.Price = item.Price.Sub(
							sdk.NewCoins(sdk.NewCoin(coin.Denom, amount)),
						).Add(coin)
					}
				}
			}

			if minPriceModified {
				for _, coin := range minPrice {
					amount := item.Price.AmountOf(coin.Denom)
					if amount.LT(coin.Amount) {
						item.Price = item.Price.Sub(
							sdk.NewCoins(sdk.NewCoin(coin.Denom, amount)),
						).Add(coin)
					}
				}
			}

			k.SetNode(ctx, item)
			return false
		})
	}

	k.IterateInactiveNodesAt(ctx, ctx.BlockTime(), func(_ int, item types.Node) bool {
		log.Info("found an inactive node", "address", item.Address)

		nodeAddr := item.GetAddress()
		k.DeleteActiveNode(ctx, nodeAddr)
		k.SetInactiveNode(ctx, nodeAddr)

		if item.Provider != "" {
			provAddr := item.GetProvider()
			k.DeleteActiveNodeForProvider(ctx, provAddr, nodeAddr)
			k.SetInactiveNodeForProvider(ctx, provAddr, nodeAddr)
		}

		k.DeleteInactiveNodeAt(ctx, item.StatusAt.Add(inactiveDuration), nodeAddr)

		item.Status = hubtypes.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetNode(ctx, item)

		ctx.EventManager().EmitTypedEvent(
			&types.EventSetNodeStatus{
				From:    "",
				Address: item.Address,
				Status:  item.Status,
			},
		)

		return false
	})

	return nil
}

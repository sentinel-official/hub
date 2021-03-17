package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgSwap:
			return handleSwap(ctx, k, msg)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownMsgType, "%s", msg.Type())
		}
	}
}

func handleSwap(ctx sdk.Context, k keeper.Keeper, msg types.MsgSwap) (*sdk.Result, error) {
	if !k.SwapEnabled(ctx) {
		return nil, types.ErrorSwapIsDisabled
	}
	if !k.ApproveBy(ctx).Equals(msg.From) {
		return nil, types.ErrorUnauthorized
	}
	if k.HasSwap(ctx, msg.TxHash) {
		return nil, types.ErrorDuplicateSwap
	}

	var (
		coin = sdk.NewCoin(k.SwapDenom(ctx), msg.Amount.QuoRaw(100))
		swap = types.Swap{
			TxHash:   msg.TxHash,
			Receiver: msg.Receiver,
			Amount:   coin,
		}
	)

	if err := k.MintCoin(ctx, coin); err != nil {
		return nil, err
	}
	if err := k.SendCoinFromModuleToAccount(ctx, msg.Receiver, coin); err != nil {
		return nil, err
	}

	k.SetSwap(ctx, swap)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyTxHash, swap.TxHash.String()),
		sdk.NewAttribute(types.AttributeKeyAddress, swap.Receiver.String()),
		sdk.NewAttribute(types.AttributeKeyAmount, swap.Amount.String()),
	))

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

package swap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	server := keeper.NewMsgServiceServer(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSwapRequest:
			res, err := server.MsgSwap(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(types.ErrorUnknownMsgType, "%s", msg)
		}
	}
}

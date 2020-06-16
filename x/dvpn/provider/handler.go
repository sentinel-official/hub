package provider

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/provider/keeper"
	"github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func HandleRegisterProvider(ctx sdk.Context, k keeper.Keeper, msg types.MsgRegisterProvider) sdk.Result {
	return sdk.Result{}
}

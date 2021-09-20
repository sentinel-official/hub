package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

type MintKeeper interface {
	GetParams(ctx sdk.Context) minttypes.Params
	SetParams(ctx sdk.Context, params minttypes.Params)
}

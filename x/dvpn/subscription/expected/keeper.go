package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

type BankKeeper interface {
	SendCoins(ctx sdk.Context, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Coins) sdk.Error
}

type DepositKeeper interface {
	Add(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error
	Subtract(ctx sdk.Context, address sdk.AccAddress, coins sdk.Coins) sdk.Error
}

type ProviderKeeper interface {
	GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool)
}

type NodeKeeper interface {
	GetNode(ctx sdk.Context, address hub.NodeAddress) (node.Node, bool)
}

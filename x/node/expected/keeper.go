package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, address sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, address sdk.AccAddress) sdk.Coins
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type ProviderKeeper interface {
	HasProvider(ctx sdk.Context, address hubtypes.ProvAddress) bool
}

type PlanKeeper interface {
	GetCountForNodeByProvider(ctx sdk.Context, p hubtypes.ProvAddress, n hubtypes.NodeAddress) uint64
}

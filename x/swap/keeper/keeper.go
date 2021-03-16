package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/swap/expected"
	"github.com/sentinel-official/hub/x/swap/types"
)

type Keeper struct {
	cdc    *codec.Codec
	key    sdk.StoreKey
	params params.Subspace
	supply expected.SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, params params.Subspace, supply expected.SupplyKeeper) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
		supply: supply,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}

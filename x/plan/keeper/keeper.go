package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/types"
)

type Keeper struct {
	cdc      codec.BinaryMarshaler
	key      sdk.StoreKey
	provider expected.ProviderKeeper
	node     expected.NodeKeeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}

func (k *Keeper) WithProviderKeeper(keeper expected.ProviderKeeper) {
	k.provider = keeper
}

func (k *Keeper) WithNodeKeeper(keeper expected.NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}

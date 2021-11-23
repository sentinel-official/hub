package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/mint/expected"
	"github.com/sentinel-official/hub/x/mint/types"
)

type Keeper struct {
	appCodec codec.BinaryCodec
	key      sdk.StoreKey
	mint     expected.MintKeeper
}

func NewKeeper(appCodec codec.BinaryCodec, key sdk.StoreKey, mint expected.MintKeeper) Keeper {
	return Keeper{
		appCodec: appCodec,
		key:      key,
		mint:     mint,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}

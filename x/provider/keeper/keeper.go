package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	key          sdk.StoreKey
	params       paramstypes.Subspace
	distribution expected.DistributionKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, params paramstypes.Subspace) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
	}
}

func (k *Keeper) WithDistributionKeeper(keeper expected.DistributionKeeper) {
	k.distribution = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}

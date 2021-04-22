package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/types"
)

type Keeper struct {
	cdc          codec.BinaryMarshaler
	key          sdk.StoreKey
	params       params.Subspace
	distribution expected.DistributionKeeper
	provider     expected.ProviderKeeper
	plan         expected.PlanKeeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, params params.Subspace) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
	}
}

func (k *Keeper) WithDistributionKeeper(keeper expected.DistributionKeeper) {
	k.distribution = keeper
}

func (k *Keeper) WithProviderKeeper(keeper expected.ProviderKeeper) {
	k.provider = keeper
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}

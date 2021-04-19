package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/swap/expected"
	"github.com/sentinel-official/hub/x/swap/types"
)

type Keeper struct {
	cdc    codec.BinaryMarshaler
	key    sdk.StoreKey
	params params.Subspace
	bank   expected.BankKeeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, params params.Subspace, bank expected.BankKeeper) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
		bank:   bank,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}

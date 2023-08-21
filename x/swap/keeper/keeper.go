package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/hub/x/swap/expected"
	"github.com/sentinel-official/hub/x/swap/types"
)

type Keeper struct {
	cdc     codec.BinaryCodec
	key     sdk.StoreKey
	params  paramstypes.Subspace
	account expected.AccountKeeper
	bank    expected.BankKeeper
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, params paramstypes.Subspace, account expected.AccountKeeper, bank expected.BankKeeper) Keeper {
	return Keeper{
		cdc:     cdc,
		key:     key,
		params:  params.WithKeyTable(types.ParamsKeyTable()),
		account: account,
		bank:    bank,
	}
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.key)
}

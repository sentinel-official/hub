package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/sentinel-official/hub/v12/x/session/expected"
	"github.com/sentinel-official/hub/v12/x/session/types"
)

type Keeper struct {
	cdc              codec.BinaryCodec
	key              storetypes.StoreKey
	params           paramstypes.Subspace
	account          expected.AccountKeeper
	bank             expected.BankKeeper
	deposit          expected.DepositKeeper
	node             expected.NodeKeeper
	plan             expected.PlanKeeper
	subscription     expected.SubscriptionKeeper
	feeCollectorName string
}

func NewKeeper(cdc codec.BinaryCodec, key storetypes.StoreKey, params paramstypes.Subspace, feeCollectorName string) Keeper {
	return Keeper{
		cdc:              cdc,
		key:              key,
		params:           params.WithKeyTable(types.ParamsKeyTable()),
		feeCollectorName: feeCollectorName,
	}
}

func (k *Keeper) WithAccountKeeper(keeper expected.AccountKeeper) {
	k.account = keeper
}

func (k *Keeper) WithBankKeeper(keeper expected.BankKeeper) {
	k.bank = keeper
}

func (k *Keeper) WithDepositKeeper(keeper expected.DepositKeeper) {
	k.deposit = keeper
}

func (k *Keeper) WithNodeKeeper(keeper expected.NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k *Keeper) WithSubscriptionKeeper(keeper expected.SubscriptionKeeper) {
	k.subscription = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}

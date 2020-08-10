package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func SimulateMsgRegister(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from        = simulation.RandomAcc(r, accounts).Address
			name        = simulation.RandStringOfLength(r, r.Intn(64)+1)
			identity    = simulation.RandStringOfLength(r, r.Intn(64)+1)
			website     = simulation.RandStringOfLength(r, r.Intn(64)+1)
			description = simulation.RandStringOfLength(r, r.Intn(256)+1)
		)

		msg := types.NewMsgRegister(from, name, identity, website, description)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := provider.HandleRegister(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdate(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from        = RandomProvider(r, k.GetProviders(ctx)).Address
			name        = simulation.RandStringOfLength(r, r.Intn(64))
			identity    = simulation.RandStringOfLength(r, r.Intn(64))
			website     = simulation.RandStringOfLength(r, r.Intn(64))
			description = simulation.RandStringOfLength(r, r.Intn(256))
		)

		msg := types.NewMsgUpdate(from, name, identity, website, description)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := provider.HandleUpdate(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/simulation"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
)

func SimulateMsgAddPlan(pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from      = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			price     = sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(r.Int63n(100)+1)))
			validity  = time.Duration(r.Intn(24)+1) * time.Hour
			bandwidth = hub.NewBandwidthFromInt64(r.Int63n(1e12)+1, r.Int63n(1e12)+1)
		)

		msg := types.NewMsgAddPlan(from, price, validity, bandwidth)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := plan.HandleAddPlan(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgSetPlanStatus(pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from   = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			id     = RandomPlan(r, k.GetPlansForProvider(ctx, from)).ID
			status hub.Status
		)

		switch r.Intn(2) {
		case 0:
			status = hub.StatusActive
		case 1:
			status = hub.StatusInactive
		}

		msg := types.NewMsgSetPlanStatus(from, id, status)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := plan.HandleSetPlanStatus(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgAddNodeForPlan(pk expected.ProviderKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from    = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			id      = RandomPlan(r, k.GetPlansForProvider(ctx, from)).ID
			address = node.RandomNode(r, nk.GetNodes(ctx)).Address
		)

		msg := types.NewMsgAddNodeForPlan(from, id, address)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := plan.HandleAddNodeForPlan(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgRemoveNodeForPlan(pk expected.ProviderKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from    = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			id      = RandomPlan(r, k.GetPlansForProvider(ctx, from)).ID
			address = node.RandomNode(r, nk.GetNodes(ctx)).Address
		)

		msg := types.NewMsgRemoveNodeForPlan(from, id, address)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := plan.HandleRemoveNodeForPlan(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

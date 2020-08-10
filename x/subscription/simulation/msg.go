package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/simulation"
	plan "github.com/sentinel-official/hub/x/plan/simulation"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func SimulateMsgSubscribeToPlan(pk expected.PlanKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from  = simulation.RandomAcc(r, accounts).Address
			id    = plan.RandomPlan(r, pk.GetPlans(ctx)).ID
			denom = "stake"
		)

		msg := types.NewMsgSubscribeToPlan(from, id, denom)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := subscription.HandleSubscribeToPlan(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgSubscribeToNode(nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from    = simulation.RandomAcc(r, accounts).Address
			address = node.RandomNode(r, nk.GetNodes(ctx)).Address
			deposit = sdk.NewCoin("stake", simulation.RandomAmount(r, sdk.NewInt(1e3)))
		)

		msg := types.NewMsgSubscribeToNode(from, address, deposit)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := subscription.HandleSubscribeToNode(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgEnd(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from = simulation.RandomAcc(r, accounts).Address
			id   = RandomSubscription(r, k.GetSubscriptions(ctx)).ID
		)

		msg := types.NewMsgEnd(from, id)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := subscription.HandleEnd(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgAddQuota(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from      = simulation.RandomAcc(r, accounts).Address
			id        = RandomSubscription(r, k.GetSubscriptions(ctx)).ID
			address   = simulation.RandomAcc(r, accounts).Address
			bandwidth = hub.NewBandwidthFromInt64(r.Int63n(1e9)+1, r.Int63n(1e9)+1)
		)

		msg := types.NewMsgAddQuota(from, id, address, bandwidth)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := subscription.HandleAddQuota(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdateQuota(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from      = simulation.RandomAcc(r, accounts).Address
			id        = RandomSubscription(r, k.GetSubscriptions(ctx)).ID
			address   = simulation.RandomAcc(r, accounts).Address
			bandwidth = hub.NewBandwidthFromInt64(r.Int63n(1e9)+1, r.Int63n(1e9)+1)
		)

		msg := types.NewMsgUpdateQuota(from, id, address, bandwidth)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := subscription.HandleUpdateQuota(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

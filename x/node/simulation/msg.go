package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
)

func SimulateMsgRegister(pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from     = simulation.RandomAcc(r, accounts).Address
			prov     = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			speed    = hub.NewBandwidthFromInt64(r.Int63n(1e9)+1, r.Int63n(1e9)+1)
			moniker  = simulation.RandStringOfLength(r, 64)
			remote   = simulation.RandStringOfLength(r, 64)
			version  = simulation.RandStringOfLength(r, 64)
			category types.Category
		)

		switch r.Intn(2) {
		case 0:
			category = types.CategoryOpenVPN
		case 1:
			category = types.CategoryWireGuard
		}

		msg := types.NewMsgRegister(from, moniker, prov, nil, speed, remote, version, category)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := node.HandleRegister(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdate(pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from     = RandomNode(r, k.GetNodes(ctx)).Address
			prov     = provider.RandomProvider(r, pk.GetProviders(ctx)).Address
			speed    = hub.NewBandwidthFromInt64(r.Int63n(1e9)+1, r.Int63n(1e9)+1)
			moniker  = simulation.RandStringOfLength(r, 64)
			remote   = simulation.RandStringOfLength(r, 64)
			version  = simulation.RandStringOfLength(r, 64)
			category types.Category
		)

		switch r.Intn(2) {
		case 0:
			category = types.CategoryOpenVPN
		case 1:
			category = types.CategoryWireGuard
		}

		msg := types.NewMsgUpdate(from, moniker, prov, nil, speed, remote, version, category)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := node.HandleUpdate(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgSetStatus(k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from   = RandomNode(r, k.GetNodes(ctx)).Address
			status hub.Status
		)

		switch r.Intn(2) {
		case 0:
			status = hub.StatusActive
		case 1:
			status = hub.StatusInactive
		}

		msg := types.NewMsgSetStatus(from, status)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := node.HandleSetStatus(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

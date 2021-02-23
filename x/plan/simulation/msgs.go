package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/simulation"
	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
)

func SimulateMsgAdd(ak expected.AccountKeeper, pk expected.ProviderKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		providers := pk.GetProviders(ctx, 0, 0)
		if len(providers) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rProvider := provider.RandomProvider(r, providers)

		rAccount, found := simulation.FindAccount(accounts, rProvider.Address)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account  = ak.GetAccount(ctx, rAccount.Address)
			price    = sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(r.Int63n(100)+1)))
			validity = time.Duration(r.Intn(24)+1) * time.Hour
			bytes    = sdk.NewInt(r.Int63n(1e12) + 1)
		)

		msg := types.NewMsgAdd(rProvider.Address, price, validity, bytes)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			nil,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgSetStatus(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		plans := k.GetPlans(ctx, 0, 0)
		if len(plans) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rPlan := RandomPlan(r, plans)

		rAccount, found := simulation.FindAccount(accounts, rPlan.Provider)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		account := ak.GetAccount(ctx, rAccount.Address)

		var status hub.Status
		switch r.Intn(2) {
		case 0:
			status = hub.StatusActive
		case 1:
			status = hub.StatusInactive
		}

		msg := types.NewMsgSetStatus(rPlan.Provider, rPlan.ID, status)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			nil,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgAddNode(ak expected.AccountKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		plans := k.GetPlans(ctx, 0, 0)
		if len(plans) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rPlan := RandomPlan(r, plans)

		rAccount, found := simulation.FindAccount(accounts, rPlan.Provider)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		account := ak.GetAccount(ctx, rAccount.Address)

		nodes := nk.GetNodesForProvider(ctx, rPlan.Provider, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := node.RandomNode(r, nodes)

		msg := types.NewMsgAddNode(rPlan.Provider, rPlan.ID, rNode.Address)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			nil,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgRemoveNode(ak expected.AccountKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		plans := k.GetPlans(ctx, 0, 0)
		if len(plans) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rPlan := RandomPlan(r, plans)

		rAccount, found := simulation.FindAccount(accounts, rPlan.Provider)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		account := ak.GetAccount(ctx, rAccount.Address)

		nodes := nk.GetNodesForProvider(ctx, rPlan.Provider, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := node.RandomNode(r, nodes)

		msg := types.NewMsgRemoveNode(rPlan.Provider, rPlan.ID, rNode.Address)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			nil,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

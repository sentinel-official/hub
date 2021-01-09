package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	node "github.com/sentinel-official/hub/x/node/simulation"
	plan "github.com/sentinel-official/hub/x/plan/simulation"
	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func SimulateMsgSubscribeToNode(ak expected.AccountKeeper, nk expected.NodeKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		nodes := nk.GetActiveNodes(ctx, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := node.RandomNode(r, nodes)
		if rNode.Provider != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			rAccount, _ = simulation.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
		)

		amount := simulation.RandomAmount(r, account.SpendableCoins(ctx.BlockTime()).AmountOf("stake"))
		if !amount.IsPositive() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		deposit := sdk.NewCoin("stake", amount)

		msg := types.NewMsgSubscribeToNode(rAccount.Address, rNode.Address, deposit)
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

func SimulateMsgSubscribeToPlan(ak expected.AccountKeeper, pk expected.PlanKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		plans := pk.GetActivePlans(ctx, 0, 0)
		if len(plans) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			rPlan       = plan.RandomPlan(r, plans)
			rAccount, _ = simulation.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
			denom       = "stake"
		)

		if account.SpendableCoins(ctx.BlockTime()).AmountOf(denom).LT(rPlan.Price.AmountOf(denom)) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgSubscribeToPlan(rAccount.Address, rPlan.ID, denom)
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

func SimulateMsgCancel(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		rAccount, _ := simulation.RandomAcc(r, accounts)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, rAccount.Address, 0, 0)
		if len(subscriptions) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account       = ak.GetAccount(ctx, rAccount.Address)
			rSubscription = RandomSubscription(r, subscriptions)
		)

		msg := types.NewMsgCancel(rAccount.Address, rSubscription.ID)
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

func SimulateMsgAddQuota(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		rAccount, _ := simulation.RandomAcc(r, accounts)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, rAccount.Address, 0, 0)
		if len(subscriptions) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rSubscription := RandomSubscription(r, subscriptions)
		if rSubscription.Plan == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		if rSubscription.Free.IsZero() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		toAccount, _ := simulation.RandomAcc(r, accounts)

		if k.HasQuota(ctx, rSubscription.ID, toAccount.Address) {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account = ak.GetAccount(ctx, rAccount.Address)
			bytes   = sdk.NewInt(r.Int63n(rSubscription.Free.Int64()) + 1)
		)

		msg := types.NewMsgAddQuota(rAccount.Address, rSubscription.ID, toAccount.Address, bytes)
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

func SimulateMsgUpdateQuota(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		rAccount, _ := simulation.RandomAcc(r, accounts)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, rAccount.Address, 0, 0)
		if len(subscriptions) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rSubscription := RandomSubscription(r, subscriptions)
		if rSubscription.Plan == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		quotas := k.GetQuotas(ctx, rSubscription.ID, 0, 0)
		if len(quotas) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account = ak.GetAccount(ctx, rAccount.Address)
			rQuota  = RandomQuota(r, quotas)
			bytes   = sdk.NewInt(r.Int63n(rSubscription.Free.
				Add(rQuota.Allocated).Int64()) + rQuota.Consumed.Int64())
		)

		msg := types.NewMsgUpdateQuota(rAccount.Address, rSubscription.ID, rQuota.Address, bytes)
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

package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
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
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			address = node.RandomNode(r, nk.GetNodes(ctx, 0, 0)).Address
			deposit = sdk.NewCoin("stake", simulation.RandomAmount(r, sdk.NewInt(1e3)))
		)

		msg := types.NewMsgSubscribeToNode(from.Address, address, deposit)
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
			from.PrivKey,
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
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id    = plan.RandomPlan(r, pk.GetPlans(ctx, 0, 0)).ID
			denom = "stake"
		)

		msg := types.NewMsgSubscribeToPlan(from.Address, id, denom)
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
			from.PrivKey,
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
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id = RandomSubscription(r, k.GetSubscriptions(ctx, 0, 0)).ID
		)

		msg := types.NewMsgCancel(from.Address, id)
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
			from.PrivKey,
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
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id    = RandomSubscription(r, k.GetSubscriptions(ctx, 0, 0)).ID
			to, _ = simulation.RandomAcc(r, accounts)
			bytes = sdk.NewInt(r.Int63n(1e9) + 1)
		)

		msg := types.NewMsgAddQuota(from.Address, id, to.Address, bytes)
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
			from.PrivKey,
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
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id    = RandomSubscription(r, k.GetSubscriptions(ctx, 0, 0)).ID
			to, _ = simulation.RandomAcc(r, accounts)
			bytes = sdk.NewInt(r.Int63n(1e9) + 1)
		)

		msg := types.NewMsgUpdateQuota(from.Address, id, to.Address, bytes)
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
			from.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

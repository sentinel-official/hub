package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
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
		var (
			rProvider     = provider.RandomProvider(r, pk.GetProviders(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rProvider.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

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
			from.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgSetStatus(ak expected.AccountKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			rProvider     = provider.RandomProvider(r, pk.GetProviders(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rProvider.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id     = RandomPlan(r, k.GetPlansForProvider(ctx, rProvider.Address, 0, 0)).ID
			status hub.Status
		)

		switch r.Intn(2) {
		case 0:
			status = hub.StatusActive
		case 1:
			status = hub.StatusInactive
		}

		msg := types.NewMsgSetStatus(rProvider.Address, id, status)
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

func SimulateMsgAddNode(ak expected.AccountKeeper, pk expected.ProviderKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			rProvider     = provider.RandomProvider(r, pk.GetProviders(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rProvider.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id      = RandomPlan(r, k.GetPlansForProvider(ctx, rProvider.Address, 0, 0)).ID
			address = node.RandomNode(r, nk.GetNodes(ctx, 0, 0)).Address
		)

		msg := types.NewMsgAddNode(rProvider.Address, id, address)
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

func SimulateMsgRemoveNode(ak expected.AccountKeeper, pk expected.ProviderKeeper, nk expected.NodeKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			rProvider     = provider.RandomProvider(r, pk.GetProviders(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rProvider.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id      = RandomPlan(r, k.GetPlansForProvider(ctx, rProvider.Address, 0, 0)).ID
			address = node.RandomNode(r, nk.GetNodes(ctx, 0, 0)).Address
		)

		msg := types.NewMsgRemoveNode(rProvider.Address, id, address)
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

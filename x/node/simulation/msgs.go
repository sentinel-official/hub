package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
)

func SimulateMsgRegister(ak expected.AccountKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		rAccount, _ := simulation.RandomAcc(r, accounts)
		if _, found := k.GetNode(ctx, rAccount.Address.Bytes()); found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		account := ak.GetAccount(ctx, rAccount.Address)

		providers := pk.GetProviders(ctx, 0, 0)
		if len(providers) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			rProvider = provider.RandomProvider(r, providers)
			remote    = simulation.RandStringOfLength(r, 64)
		)

		msg := types.NewMsgRegister(rAccount.Address, rProvider.Address, nil, remote)
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

func SimulateMsgUpdate(ak expected.AccountKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		nodes := k.GetNodes(ctx, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := RandomNode(r, nodes)

		rAccount, found := simulation.FindAccount(accounts, rNode.Address)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		account := ak.GetAccount(ctx, rAccount.Address)

		providers := pk.GetProviders(ctx, 0, 0)
		if len(providers) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			rProvider = provider.RandomProvider(r, providers)
			remoteURL = simulation.RandStringOfLength(r, 64)
		)

		msg := types.NewMsgUpdate(rNode.Address, rProvider.Address, nil, remoteURL)
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
		nodes := k.GetNodes(ctx, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := RandomNode(r, nodes)

		rAccount, found := simulation.FindAccount(accounts, rNode.Address)
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

		msg := types.NewMsgSetStatus(rNode.Address, status)
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

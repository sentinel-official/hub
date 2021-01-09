package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func SimulateMsgRegister(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		rAccount, _ := simulation.RandomAcc(r, accounts)
		if _, found := k.GetProvider(ctx, rAccount.Address.Bytes()); found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account     = ak.GetAccount(ctx, rAccount.Address)
			name        = simulation.RandStringOfLength(r, r.Intn(64)+1)
			identity    = simulation.RandStringOfLength(r, r.Intn(64)+1)
			website     = simulation.RandStringOfLength(r, r.Intn(64)+1)
			description = simulation.RandStringOfLength(r, r.Intn(256)+1)
		)

		msg := types.NewMsgRegister(rAccount.Address, name, identity, website, description)
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

func SimulateMsgUpdate(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		providers := k.GetProviders(ctx, 0, 0)
		if len(providers) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rProvider := RandomProvider(r, providers)

		rAccount, found := simulation.FindAccount(accounts, rProvider.Address)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account     = ak.GetAccount(ctx, rAccount.Address)
			name        = simulation.RandStringOfLength(r, r.Intn(64))
			identity    = simulation.RandStringOfLength(r, r.Intn(64))
			website     = simulation.RandStringOfLength(r, r.Intn(64))
			description = simulation.RandStringOfLength(r, r.Intn(256))
		)

		msg := types.NewMsgUpdate(rProvider.Address, name, identity, website, description)
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

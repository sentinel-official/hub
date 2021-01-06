package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func SimulateMsgRegister(ak expected.AccountKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.RandomAcc(r, accounts)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			name        = simulation.RandStringOfLength(r, r.Intn(64)+1)
			identity    = simulation.RandStringOfLength(r, r.Intn(64)+1)
			website     = simulation.RandStringOfLength(r, r.Intn(64)+1)
			description = simulation.RandStringOfLength(r, r.Intn(256)+1)
		)

		msg := types.NewMsgRegister(from.Address, name, identity, website, description)
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

func SimulateMsgUpdate(ak expected.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			rProvider     = RandomProvider(r, k.GetProviders(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rProvider.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

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
			from.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

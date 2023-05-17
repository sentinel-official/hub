// DO NOT COVER

package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

var (
	typeMsgRegister = sdk.MsgTypeURL((*types.MsgRegisterRequest)(nil))
	typeMsgUpdate   = sdk.MsgTypeURL((*types.MsgUpdateRequest)(nil))
)

func WeightedOperations(
	cdc codec.Codec,
	txConfig client.TxConfig,
	params simtypes.AppParams,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgRegister int
		weightMsgUpdate   int
	)

	params.GetOrGenerate(
		cdc,
		typeMsgRegister,
		&weightMsgRegister,
		nil,
		func(_ *rand.Rand) {
			weightMsgRegister = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgUpdate,
		&weightMsgUpdate,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdate = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgRegister, SimulateMsgRegister(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUpdate, SimulateMsgUpdate(txConfig, ak, bk, k)),
	}
}

func SimulateMsgRegister(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			rAccount, _ = simtypes.RandomAcc(r, accounts)
			fromAccount = ak.GetAccount(ctx, rAccount.Address)
		)

		found := k.HasProvider(ctx, hubtypes.ProvAddress(fromAccount.GetAddress()))
		if found {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, ""), nil, nil
		}

		sCoins := bk.SpendableCoins(ctx, fromAccount.GetAddress())
		if !sCoins.IsAnyNegative() {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, ""), nil, nil
		}

		rFees, err := simtypes.RandomFees(r, ctx, sCoins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, err.Error()), nil, err
		}

		deposit := k.Deposit(ctx)
		if sCoins.Sub(rFees).AmountOf(deposit.Denom).LT(deposit.Amount) {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, ""), nil, nil
		}

		var (
			name        = simtypes.RandStringOfLength(r, r.Intn(MaxNameLength)+8)
			identity    = simtypes.RandStringOfLength(r, r.Intn(MaxIdentityLength))
			website     = simtypes.RandStringOfLength(r, r.Intn(MaxWebsiteLength))
			description = simtypes.RandStringOfLength(r, r.Intn(MaxDescriptionLength))
		)

		msg := types.NewMsgRegisterRequest(
			fromAccount.GetAddress(),
			name,
			identity,
			website,
			description,
		)

		tx, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			rFees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{fromAccount.GetAccountNumber()},
			[]uint64{fromAccount.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, err.Error()), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func SimulateMsgUpdate(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			rAccount, _ = simtypes.RandomAcc(r, accounts)
			fromAccount = ak.GetAccount(ctx, rAccount.Address)
		)

		found := k.HasProvider(ctx, hubtypes.ProvAddress(fromAccount.GetAddress()))
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdate, ""), nil, nil
		}

		sCoins := bk.SpendableCoins(ctx, fromAccount.GetAddress())
		if !sCoins.IsAnyNegative() {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdate, ""), nil, nil
		}

		rFees, err := simtypes.RandomFees(r, ctx, sCoins)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdate, err.Error()), nil, err
		}

		var (
			name        = simtypes.RandStringOfLength(r, r.Intn(MaxNameLength+8))
			identity    = simtypes.RandStringOfLength(r, r.Intn(MaxIdentityLength))
			website     = simtypes.RandStringOfLength(r, r.Intn(MaxWebsiteLength))
			description = simtypes.RandStringOfLength(r, r.Intn(MaxDescriptionLength))
		)

		status := hubtypes.StatusUnspecified
		if rand.Intn(2) == 0 {
			status = hubtypes.StatusActive
		} else {
			status = hubtypes.StatusInactive
		}

		msg := types.NewMsgUpdateRequest(
			hubtypes.ProvAddress(fromAccount.GetAddress()),
			name,
			identity,
			website,
			description,
			status,
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			rFees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{fromAccount.GetAccountNumber()},
			[]uint64{fromAccount.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdate, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdate, err.Error()), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

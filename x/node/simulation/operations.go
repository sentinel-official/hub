package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	simulationhubtypes "github.com/sentinel-official/hub/types/simulation"
	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

var (
	OperationWeightMsgRegisterRequest  = "op_weight_" + types.TypeMsgRegisterRequest
	OperationWeightMsgUpdateRequest    = "op_weight_" + types.TypeMsgUpdateRequest
	OperationWeightMsgSetStatusRequest = "op_weight_" + types.TypeMsgSetStatusRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONMarshaler,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgRegisterRequest  int
		weightMsgSetStatusRequest int
		weightMsgUpdateRequest    int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgRegisterRequest,
		&weightMsgRegisterRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgRegisterRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgSetStatusRequest,
		&weightMsgSetStatusRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgSetStatusRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUpdateRequest,
		&weightMsgUpdateRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgRegisterRequest,
			SimulateMsgRegisterRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateRequest,
			SimulateMsgUpdateRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgSetStatusRequest,
			SimulateMsgSetStatusRequest(ak, bk, k),
		),
	}
}

func SimulateMsgRegisterRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rAccount, _ = simulationtypes.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
			deposit     = k.Deposit(ctx)
		)

		_, found := k.GetNode(ctx, hubtypes.NodeAddress(account.GetAddress()))
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, "node already exists"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, "balance cannot be negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, err.Error()), nil, err
		}

		balance = balance.Sub(fees)
		if balance.AmountOf(deposit.Denom).LT(deposit.Amount) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, "balance cannot be less than deposit"), nil, err
		}

		var (
			price     = simulationhubtypes.RandomCoins(r, balance)
			remoteURL = fmt.Sprintf(
				"https://%s:8080",
				simulationtypes.RandStringOfLength(r, r.Intn(MaxRemoteURLLength)),
			)
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgRegisterRequest(
				account.GetAddress(),
				nil,
				price,
				remoteURL,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

func SimulateMsgUpdateRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rAccount, _ = simulationtypes.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
		)

		_, found := k.GetNode(ctx, hubtypes.NodeAddress(account.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "node does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "balance cannot be negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		var (
			price     = simulationhubtypes.RandomCoins(r, balance)
			remoteURL = fmt.Sprintf(
				"https://%s:8080",
				simulationtypes.RandStringOfLength(r, r.Intn(MaxRemoteURLLength)),
			)
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgUpdateRequest(
				hubtypes.NodeAddress(account.GetAddress()),
				nil,
				price,
				remoteURL,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

func SimulateMsgSetStatusRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rAccount, _ = simulationtypes.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
		)

		_, found := k.GetNode(ctx, hubtypes.NodeAddress(account.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "node does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "balance cannot be negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		var (
			status = hubtypes.StatusActive
		)

		if rand.Intn(2) == 0 {
			status = hubtypes.StatusInactive
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgSetStatusRequest(
				hubtypes.NodeAddress(account.GetAddress()),
				status,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

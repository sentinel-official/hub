package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

const (
	OpWeightMsgStartRequest  = "op_weight_msg_start_request"
	OpWeightMsgUpdateRequest = "op_weight_msg_update_session_request"
	OpWeightMsgEndRequest    = "op_weight_msg_end_request"
)

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var (
		weightMsgStartRequest  int
		weightMsgEndRequst     int
		weightMsgUpdateRequest int
	)

	randMsgStartRequest := func(_ *rand.Rand) {
		weightMsgStartRequest = 100
	}

	randMsgUpdateRequest := func(_ *rand.Rand) {
		weightMsgUpdateRequest = 100
	}

	randMsgEndRequest := func(_ *rand.Rand) {
		weightMsgEndRequst = 100
	}

	ap.GetOrGenerate(cdc, OpWeightMsgStartRequest, &weightMsgStartRequest, nil, randMsgStartRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgUpdateRequest, &weightMsgUpdateRequest, nil, randMsgUpdateRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgEndRequest, &weightMsgEndRequst, nil, randMsgEndRequest)

	startSessionOperation := simulation.NewWeightedOperation(weightMsgStartRequest, SimulateMsgStartRequest(k))
	updateSessionOperation := simulation.NewWeightedOperation(weightMsgUpdateRequest, SimulateMsgUpdateRequest(k))
	endSesionOperation := simulation.NewWeightedOperation(weightMsgEndRequst, SimulateMsgMsgEndRequest(k))

	return simulation.WeightedOperations{
		startSessionOperation,
		updateSessionOperation,
		endSesionOperation,
	}
}

func SimulateMsgStartRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		nodeAddress := hubtypes.NodeAddress(acc.Address.Bytes())

		sessionID := uint64(r.Uint64())

		_, found := k.GetSession(ctx, sessionID)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "start_session", "session is already started"), nil, nil
		}

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "start_session", err.Error()), nil, err
		}

		msg := types.NewMsgStartRequest(from.GetAddress(), sessionID, nodeAddress)
		txConfig := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "start_session", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "start_session", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgMsgEndRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)

		session := getRandomSession(r, k.GetActiveSessionsForAddress(ctx, from.GetAddress(), 0, 0))

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "end_session", err.Error()), nil, err
		}

		msg := types.NewMsgEndRequest(from.GetAddress(), session.Id, uint64(r.Intn(5)+1))
		txConfig := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "end_session", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "end_session", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgUpdateRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		nodeAddress := hubtypes.NodeAddress(acc.Address.Bytes())

		_, found := k.GetNode(ctx, nodeAddress)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_session_request", "session not found"), nil, nil
		}

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_session_request", err.Error()), nil, err
		}

		proof := types.Proof{
			Id:        r.Uint64(),
			Duration:  time.Duration(r.Int63n(1800) + 600),
			Bandwidth: hubtypes.NewBandwidth(hubtypes.Gigabyte, hubtypes.Gigabyte),
		}

		signature := make([]byte, 64)
		if _, err := r.Read(signature); err != nil {
			panic(err)
		}

		msg := types.NewMsgUpdateRequest(nodeAddress, proof, signature)
		txConfig := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_session_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_session_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func getRandomSession(r *rand.Rand, sessions types.Sessions) types.Session {
	if len(sessions) == 0 {
		return types.Session{
			Id: uint64(1),
		}
	}

	return sessions[r.Intn(len(sessions))]
}

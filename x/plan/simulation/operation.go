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
	"github.com/sentinel-official/hub/x/plan/keeper"
	types "github.com/sentinel-official/hub/x/plan/types"
)

const (
	OpWeightMsgAddRequest        = "op_weight_msg_add_request"
	OpWeightMsgSetStatusRequest  = "op_weight_msg_set_status_request"
	OpWeightMsgAddNodeRequest    = "op_weight_msg_add_node_request"
	OpWeightMsgRemoveNodeRequest = "op_weight_msg_remove_node_request"
)

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var (
		weightMsgAddRequest        int
		weightMsgSetStatus         int
		weightMsgAddNodeRequest    int
		weightMsgRemoveNodeRequest int
	)

	randAddRequest := func(_ *rand.Rand) {
		weightMsgAddRequest = 100
	}

	randSetStatusRequest := func(_ *rand.Rand) {
		weightMsgSetStatus = 100
	}

	randAddNodeRequest := func(_ *rand.Rand) {
		weightMsgAddNodeRequest = 100
	}

	randRemoveNodeRequest := func(_ *rand.Rand) {
		weightMsgRemoveNodeRequest = 100
	}

	ap.GetOrGenerate(cdc, OpWeightMsgAddRequest, &weightMsgAddRequest, nil, randAddRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgSetStatusRequest, &weightMsgSetStatus, nil, randSetStatusRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgAddNodeRequest, &weightMsgAddNodeRequest, nil, randAddNodeRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgRemoveNodeRequest, &weightMsgRemoveNodeRequest, nil, randRemoveNodeRequest)

	addOperation := simulation.NewWeightedOperation(weightMsgAddRequest, SimulateMsgAddRequest(k, cdc))
	setStatusOperation := simulation.NewWeightedOperation(weightMsgSetStatus, SimulateMsgSetStatusRequest(k, cdc))
	addNodeOperation := simulation.NewWeightedOperation(weightMsgAddNodeRequest, SimulateMsgAddNodeRequest(k, cdc))
	removeNodeOperation := simulation.NewWeightedOperation(weightMsgRemoveNodeRequest, SimulateMsgRemoveNodeRequest(k, cdc))

	return simulation.WeightedOperations{
		addOperation,
		setStatusOperation,
		addNodeOperation,
		removeNodeOperation,
	}
}

func SimulateMsgAddRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		fromProvAddress := hubtypes.ProvAddress(from.GetAddress().Bytes())
		id := r.Uint64()

		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: "tsent", Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_request", err.Error()), nil, err
		}

		_, found := k.GetPlan(ctx, id)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_request", "plan already exists"), nil, nil
		}

		validity := time.Duration(sdksimulation.RandIntBetween(r, weekDurationInSeconds, monthDurationInSeconds))
		bytes := sdk.NewInt(int64(sdksimulation.RandIntBetween(r, gigabytes, terabytes)))

		msg := types.NewMsgAddRequest(fromProvAddress, price, validity, bytes)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgSetStatusRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		fromProvAddress := hubtypes.ProvAddress(from.GetAddress().Bytes())
		plan := getRandomPlan(r, k.GetPlans(ctx, 0, 0))

		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: "tsent", Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		_, found := k.GetPlan(ctx, plan.Id)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", "plan does not exist"), nil, nil
		}

		msg := types.NewMsgSetStatusRequest(fromProvAddress, plan.Id, hubtypes.Status(r.Intn(3)))
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgAddNodeRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		nodeAddress := hubtypes.NodeAddress(from.GetAddress().Bytes())
		fromProvAddress := hubtypes.ProvAddress(from.GetAddress().Bytes())
		plan := getRandomPlan(r, k.GetPlans(ctx, 0, 0))
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: "tsent", Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_node_request", err.Error()), nil, err
		}

		_, found := k.GetPlan(ctx, plan.Id)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_node_request", "plan does not exist"), nil, nil
		}

		msg := types.NewMsgAddNodeRequest(fromProvAddress, plan.Id, nodeAddress)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_node_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_node_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgRemoveNodeRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		nodeAddress := hubtypes.NodeAddress(from.GetAddress().Bytes())
		plan := getRandomPlan(r, k.GetPlans(ctx, 0, 0))
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: "tsent", Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "remove_node_request", err.Error()), nil, err
		}

		_, found := k.GetNode(ctx, nodeAddress)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "remove_node_request", "node does not exist"), nil, nil
		}

		msg := types.NewMsgRemoveNodeRequest(from.GetAddress(), plan.Id, nodeAddress)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "remove_node_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "remove_node_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

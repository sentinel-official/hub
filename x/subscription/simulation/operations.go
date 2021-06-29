package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	OpWeightMsgSubscribeToNodeRequest = "op_weight_msg_subscribe_to_node"
	OpWeightMsgSubscribeToPlanRequest = "op_weight_msg_subscribe_to_plan"
	OpWeightMsgAddQuotaRequest        = "op_weight_msg_add_quota_request"
	OpWeightMsgUpdateQuotaRequest     = "op_weight_msg_update_quota_request"
	OpWeightMsgCancelRequest          = "op_weight_msg_cancel_request"
)

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var (
		weightMsgSubscribeToNode int
		weightMsgSubscribeToPlan int
		weightMsgAddQuota        int
		weightMsgUpdateQuota     int
		weightMsgCancel          int
	)

	randMsgSubscribeToNode := func(_ *rand.Rand) {
		weightMsgSubscribeToNode = 100
	}

	randMsgSubscribeToPlan := func(_ *rand.Rand) {
		weightMsgSubscribeToPlan = 100
	}

	randMsgAddQuota := func(_ *rand.Rand) {
		weightMsgAddQuota = 100
	}

	randMsgUpdateQuota := func(_ *rand.Rand) {
		weightMsgUpdateQuota = 100
	}

	randMsgCancel := func(_ *rand.Rand) {
		weightMsgCancel = 100
	}

	ap.GetOrGenerate(cdc, OpWeightMsgSubscribeToNodeRequest, &weightMsgSubscribeToNode, nil, randMsgSubscribeToNode)
	ap.GetOrGenerate(cdc, OpWeightMsgSubscribeToPlanRequest, &weightMsgSubscribeToPlan, nil, randMsgSubscribeToPlan)
	ap.GetOrGenerate(cdc, OpWeightMsgAddQuotaRequest, &weightMsgAddQuota, nil, randMsgAddQuota)
	ap.GetOrGenerate(cdc, OpWeightMsgUpdateQuotaRequest, &weightMsgUpdateQuota, nil, randMsgUpdateQuota)
	ap.GetOrGenerate(cdc, OpWeightMsgCancelRequest, &weightMsgCancel, nil, randMsgCancel)

	subscribeToNodeOperation := simulation.NewWeightedOperation(weightMsgSubscribeToNode, SimulateMsgSubscribeToNodeRequest(k))
	subscribeToPlanOperation := simulation.NewWeightedOperation(weightMsgSubscribeToPlan, SimulateMsgSubscribeToPlanRequest(k))
	addQuotaOperation := simulation.NewWeightedOperation(weightMsgAddQuota, SimulateMsgAddQuotaRequest(k))
	updaateQuotaOperation := simulation.NewWeightedOperation(weightMsgUpdateQuota, SimulateMsgUpdateQuotaRequest(k))
	cancelOperation := simulation.NewWeightedOperation(weightMsgCancel, SimulateMsgCancelRequest(k))

	return simulation.WeightedOperations{
		subscribeToNodeOperation,
		subscribeToPlanOperation,
		addQuotaOperation,
		updaateQuotaOperation,
		cancelOperation,
	}
}

func SimulateMsgSubscribeToNodeRequest(k keeper.Keeper) sdksimulation.Operation {
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

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		subscription := getRandomSubscription(r, subscriptions, "node")

		_, found := k.GetSubscription(ctx, subscription.Id)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_node_request", "already subscribed to node"), nil, nil
		}

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		deposit := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, deposit)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_node_request", err.Error()), nil, err
		}

		msg := types.NewMsgSubscribeToNodeRequest(from.GetAddress(), nodeAddress, deposit[0])
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
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_node_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_node_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgSubscribeToPlanRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		plan := getRandomSubscription(r, subscriptions, "plan")

		_, found := k.GetPlan(ctx, plan.Id)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_plan_request", "node is not registered"), nil, nil
		}

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_plan_request", err.Error()), nil, err
		}

		msg := types.NewMsgSubscribeToPlanRequest(from.GetAddress(), plan.Plan, denom)
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
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_plan_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "subscribe_to_plan_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgAddQuotaRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc1, _ := sdksimulation.RandomAcc(r, accounts)
		acc2, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc1.Address)
		address := k.GetAccount(ctx, acc2.Address)

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(1000))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_quota_request", err.Error()), nil, err
		}

		plan := getRandomSubscription(r, k.GetSubscriptions(ctx, 0, 0), "node")

		_, found := k.GetQuota(ctx, plan.Id, address.GetAddress())
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_quota_request", "quota already exists"), nil, nil
		}

		msg := types.NewMsgAddQuotaRequest(from.GetAddress(), plan.Id, address.GetAddress(), getRandomBytes(r))
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
			return sdksimulation.NoOpMsg(types.ModuleName, "add_quota_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "add_quota_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgUpdateQuotaRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc1, _ := sdksimulation.RandomAcc(r, accounts)
		acc2, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc1.Address)
		address := k.GetAccount(ctx, acc2.Address)

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(1000))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_quota_request", err.Error()), nil, err
		}

		plan := getRandomSubscription(r, k.GetSubscriptions(ctx, 0, 0), "node")

		_, found := k.GetQuota(ctx, plan.Id, address.GetAddress())
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_quota_request", "quota does not exists"), nil, nil
		}

		msg := types.NewMsgUpdateQuotaRequest(from.GetAddress(), plan.Id, address.GetAddress(), getRandomBytes(r))
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
			return sdksimulation.NoOpMsg(types.ModuleName, "update_quota_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_quota_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgCancelRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc1, _ := sdksimulation.RandomAcc(r, accounts)
		acc2, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc1.Address)
		address := k.GetAccount(ctx, acc2.Address)

		denom := "tsent"
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(1000))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "cancel_request", err.Error()), nil, err
		}

		plan := getRandomSubscription(r, k.GetSubscriptions(ctx, 0, 0), "node")

		_, found := k.GetQuota(ctx, plan.Id, address.GetAddress())
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "cancel_request", "quota does not exists"), nil, nil
		}

		msg := types.NewMsgCancelRequest(from.GetAddress(), plan.Id)
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
			return sdksimulation.NoOpMsg(types.ModuleName, "cancel_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "cancel_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

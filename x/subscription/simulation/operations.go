package simulation

import (
	"math"
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
	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	OperationWeightMsgSubscribeToNodeRequest = "op_weight_" + types.TypeMsgSubscribeToNodeRequest
	OperationWeightMsgSubscribeToPlanRequest = "op_weight_" + types.TypeMsgSubscribeToPlanRequest
	OperationWeightMsgCancelRequest          = "op_weight_" + types.TypeMsgCancelRequest
	OperationWeightMsgShareRequest           = "op_weight_" + types.TypeMsgShareRequest
	OperationWeightMsgUpdateQuotaRequest     = "op_weight_" + types.TypeMsgUpdateQuotaRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgSubscribeToNodeRequest int
		weightMsgSubscribeToPlanRequest int
		weightMsgCancelRequest          int
		weightMsgShareRequest           int
		weightMsgUpdateQuotaRequest     int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgSubscribeToNodeRequest,
		&weightMsgSubscribeToNodeRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgSubscribeToNodeRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgSubscribeToPlanRequest,
		&weightMsgSubscribeToPlanRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgSubscribeToPlanRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgCancelRequest,
		&weightMsgCancelRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgCancelRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgShareRequest,
		&weightMsgShareRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgShareRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUpdateQuotaRequest,
		&weightMsgUpdateQuotaRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateQuotaRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgSubscribeToNodeRequest,
			SimulateMsgSubscribeToNodeRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgSubscribeToPlanRequest,
			SimulateMsgSubscribeToPlanRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgCancelRequest,
			SimulateMsgCancelRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgShareRequest,
			SimulateMsgShareRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateQuotaRequest,
			SimulateMsgUpdateQuotaRequest(ak, bk, k),
		),
	}
}

func SimulateMsgSubscribeToNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rFrom, _    = simulationtypes.RandomAcc(r, accounts)
			from        = ak.GetAccount(ctx, rFrom.Address)
			rAddress, _ = simulationtypes.RandomAcc(r, accounts)
			address     = ak.GetAccount(ctx, rAddress.Address)
		)

		node, found := k.GetNode(ctx, hubtypes.NodeAddress(address.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, "node does not exist"), nil, nil
		}
		if node.Provider != "" {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, "provider of the node is not empty"), nil, nil
		}
		if !node.Status.Equal(hubtypes.StatusActive) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, "node status is not active"), nil, nil
		}

		var (
			deposit = simulationhubtypes.RandomCoin(
				r,
				sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					MaxSubscriptionDepositAmount,
				),
			)
		)

		_, found = node.PriceForDenom(deposit.Denom)
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, "price for denom does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgSubscribeToNodeRequest(
				from.GetAddress(),
				hubtypes.NodeAddress(address.GetAddress()),
				deposit,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgSubscribeToPlanRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rFrom, _ = simulationtypes.RandomAcc(r, accounts)
			from     = ak.GetAccount(ctx, rFrom.Address)
			rId      = uint64(r.Int63n(1 << 18))
		)

		plan, found := k.GetPlan(ctx, rId)
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, "plan does not exist"), nil, nil
		}
		if !plan.Status.Equal(hubtypes.StatusActive) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, "plan status is not active"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgSubscribeToPlanRequest(
				from.GetAddress(),
				rId,
				sdk.DefaultBondDenom,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSubscribeToPlanRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgCancelRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rFrom, _ = simulationtypes.RandomAcc(r, accounts)
			from     = ak.GetAccount(ctx, rFrom.Address)
		)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		if len(subscriptions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelRequest, "active subscriptions for address does not exist"), nil, nil
		}

		var (
			rSubscription = subscriptions[r.Intn(len(subscriptions))]
		)

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgCancelRequest(
				from.GetAddress(),
				rSubscription.Id,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCancelRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgShareRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rFrom, _    = simulationtypes.RandomAcc(r, accounts)
			from        = ak.GetAccount(ctx, rFrom.Address)
			rAddress, _ = simulationtypes.RandomAcc(r, accounts)
			address     = ak.GetAccount(ctx, rAddress.Address)
		)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		if len(subscriptions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "active subscriptions for address does not exist"), nil, nil
		}

		rSubscription := subscriptions[r.Intn(len(subscriptions))]
		if rSubscription.Plan == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "plan of the subscription is zero"), nil, nil
		}

		found := k.HasQuota(ctx, rSubscription.Id, address.GetAddress())
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "quota already exists"), nil, nil
		}

		bytes := sdk.NewInt(r.Int63n(math.MaxInt32))
		if bytes.GT(rSubscription.Free) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "no enough quota"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgShareRequest(
				from.GetAddress(),
				rSubscription.Id,
				address.GetAddress(),
				bytes,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgUpdateQuotaRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rFrom, _    = simulationtypes.RandomAcc(r, accounts)
			from        = ak.GetAccount(ctx, rFrom.Address)
			rAddress, _ = simulationtypes.RandomAcc(r, accounts)
			address     = ak.GetAccount(ctx, rAddress.Address)
		)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		if len(subscriptions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "active subscriptions for address does not exist"), nil, nil
		}

		rSubscription := subscriptions[r.Intn(len(subscriptions))]
		if rSubscription.Plan == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "plan of the subscription is zero"), nil, nil
		}

		quota, found := k.GetQuota(ctx, rSubscription.Id, address.GetAddress())
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "quota does not exist"), nil, nil
		}

		bytes := sdk.NewInt(r.Int63n(math.MaxInt32))
		if bytes.LT(quota.Consumed) || bytes.GT(rSubscription.Free.Add(quota.Allocated)) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "no enough quota"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateQuotaRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateQuotaRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgUpdateQuotaRequest(
				from.GetAddress(),
				rSubscription.Id,
				address.GetAddress(),
				bytes,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateQuotaRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateQuotaRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

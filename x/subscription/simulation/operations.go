package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	OperationWeightMsgCancelRequest = "op_weight_" + types.TypeMsgCancelRequest
	OperationWeightMsgShareRequest  = "op_weight_" + types.TypeMsgShareRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgCancelRequest int
		weightMsgShareRequest  int
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

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCancelRequest,
			SimulateMsgCancelRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgShareRequest,
			SimulateMsgShareRequest(ak, bk, k),
		),
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

		subscriptions := k.GetSubscriptionsForAccount(ctx, from.GetAddress())
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
				rSubscription.GetID(),
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

		subscriptions := k.GetSubscriptionsForAccount(ctx, from.GetAddress())
		if len(subscriptions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "active subscriptions for address does not exist"), nil, nil
		}

		rSubscription := subscriptions[r.Intn(len(subscriptions))]
		if rSubscription.Type() == types.TypeNode {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "plan of the subscription is zero"), nil, nil
		}

		found := k.HasQuota(ctx, rSubscription.GetID(), address.GetAddress())
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgShareRequest, "quota already exists"), nil, nil
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
				rSubscription.GetID(),
				address.GetAddress(),
				sdk.ZeroInt(),
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

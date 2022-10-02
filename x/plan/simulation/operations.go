package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	simulationhubtypes "github.com/sentinel-official/hub/types/simulation"
	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

var (
	OperationWeightMsgAddRequest        = "op_weight_" + types.TypeMsgAddRequest
	OperationWeightMsgSetStatusRequest  = "op_weight_" + types.TypeMsgSetStatusRequest
	OperationWeightMsgAddNodeRequest    = "op_weight_" + types.TypeMsgAddNodeRequest
	OperationWeightMsgRemoveNodeRequest = "op_weight_" + types.TypeMsgRemoveNodeRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgAddRequest        int
		weightMsgSetStatusRequest  int
		weightMsgAddNodeRequest    int
		weightMsgRemoveNodeRequest int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgAddRequest,
		&weightMsgAddRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgAddRequest = 100
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
		OperationWeightMsgAddNodeRequest,
		&weightMsgAddNodeRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgAddNodeRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgRemoveNodeRequest,
		&weightMsgRemoveNodeRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgRemoveNodeRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgAddRequest,
			SimulateMsgAddRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgSetStatusRequest,
			SimulateMsgSetStatusRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgAddNodeRequest,
			SimulateMsgAddNodeRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRemoveNodeRequest,
			SimulateMsgRemoveNodeRequest(ak, bk, k),
		),
	}
}

func SimulateMsgAddRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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

		found := k.HasProvider(ctx, hubtypes.ProvAddress(from.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, "provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, err.Error()), nil, err
		}

		var (
			price = simulationhubtypes.RandomCoins(
				r,
				sdk.NewCoins(
					sdk.NewInt64Coin(
						sdk.DefaultBondDenom,
						MaxPlanPriceAmount,
					),
				),
			)
			validity = time.Duration(r.Int63n(MaxPlanValidity)) * time.Minute
			bytes    = sdk.NewInt(r.Int63n(MaxPlanBytes))
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgAddRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				price,
				validity,
				bytes,
			)
		)

		txn, err := helpers.GenTx(
			r,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
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
			rFrom, _ = simulationtypes.RandomAcc(r, accounts)
			from     = ak.GetAccount(ctx, rFrom.Address)
		)

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(from.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "plans for provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		var (
			id     = plans[r.Intn(len(plans))].Id
			status = hubtypes.StatusActive
		)

		if r.Int63n(2) == 0 {
			status = hubtypes.StatusInactive
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgSetStatusRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				id,
				status,
			)
		)

		txn, err := helpers.GenTx(
			r,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgAddNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "node does not exist"), nil, nil
		}
		if node.Provider != hubtypes.ProvAddress(from.GetAddress()).String() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "node has different provider"), nil, nil
		}

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(from.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "plans for provider does not exist"), nil, nil
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found = k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(address.GetAddress()))
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "node for plan already exists"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgAddNodeRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				id,
				hubtypes.NodeAddress(address.GetAddress()),
			)
		)

		txn, err := helpers.GenTx(
			r,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgRemoveNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(from.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "plans for provider does not exist"), nil, nil
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found := k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(address.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "node for plan does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgRemoveNodeRequest(
				from.GetAddress(),
				id,
				hubtypes.NodeAddress(address.GetAddress()),
			)
		)

		txn, err := helpers.GenTx(
			r,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

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
	OperationWeightMsgCreateRequest       = "op_weight_" + types.TypeMsgCreateRequest
	OperationWeightMsgUpdateStatusRequest = "op_weight_" + types.TypeMsgUpdateStatusRequest
	OperationWeightMsgLinkNodeRequest     = "op_weight_" + types.TypeMsgLinkNodeRequest
	OperationWeightMsgUnlinkNodeRequest   = "op_weight_" + types.TypeMsgUnlinkNodeRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgCreateRequest       int
		weightMsgUpdateStatusRequest int
		weightMsgLinkNodeRequest     int
		weightMsgUnlinkNodeRequest   int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgCreateRequest,
		&weightMsgCreateRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgCreateRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUpdateStatusRequest,
		&weightMsgUpdateStatusRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStatusRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgLinkNodeRequest,
		&weightMsgLinkNodeRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgLinkNodeRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUnlinkNodeRequest,
		&weightMsgUnlinkNodeRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUnlinkNodeRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateRequest,
			SimulateMsgCreateRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateStatusRequest,
			SimulateMsgUpdateStatusRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgLinkNodeRequest,
			SimulateMsgLinkNodeRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUnlinkNodeRequest,
			SimulateMsgUnlinkNodeRequest(ak, bk, k),
		),
	}
}

func SimulateMsgCreateRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRequest, "provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRequest, err.Error()), nil, err
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
			message  = types.NewMsgCreateRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				price,
				validity,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgCreateRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgUpdateStatusRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateStatusRequest, "plans for provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateStatusRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateStatusRequest, err.Error()), nil, err
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
			message  = types.NewMsgUpdateStatusRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				id,
				status,
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateStatusRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateStatusRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgLinkNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, "node does not exist"), nil, nil
		}
		if node.Provider != hubtypes.ProvAddress(from.GetAddress()).String() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, "node has different provider"), nil, nil
		}

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(from.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, "plans for provider does not exist"), nil, nil
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found = k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(address.GetAddress()))
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, "node for plan already exists"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgLinkNodeRequest(
				hubtypes.ProvAddress(from.GetAddress()),
				id,
				hubtypes.NodeAddress(address.GetAddress()),
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgLinkNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

func SimulateMsgUnlinkNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, "plans for provider does not exist"), nil, nil
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found := k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(address.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, "node for plan does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgUnlinkNodeRequest(
				from.GetAddress(),
				id,
				hubtypes.NodeAddress(address.GetAddress()),
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
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUnlinkNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

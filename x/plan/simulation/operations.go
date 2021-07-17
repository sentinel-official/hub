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
	cdc codec.JSONMarshaler,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	pk expected.ProviderKeeper,
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
			SimulateMsgAddRequest(ak, bk, pk),
		),
		simulation.NewWeightedOperation(
			weightMsgSetStatusRequest,
			SimulateMsgSetStatusRequest(ak, bk, pk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgAddNodeRequest,
			SimulateMsgAddNodeRequest(ak, bk, pk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgRemoveNodeRequest,
			SimulateMsgRemoveNodeRequest(ak, bk, pk, k),
		),
	}
}

func SimulateMsgAddRequest(ak expected.AccountKeeper, bk expected.BankKeeper, pk expected.ProviderKeeper) simulationtypes.Operation {
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

		found := pk.HasProvider(ctx, hubtypes.ProvAddress(account.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, "provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, "balance cannot be negative"), nil, nil
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
				hubtypes.ProvAddress(account.GetAddress()),
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
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

func SimulateMsgSetStatusRequest(ak expected.AccountKeeper, bk expected.BankKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulationtypes.Operation {
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

		found := pk.HasProvider(ctx, hubtypes.ProvAddress(account.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "provider does not exist"), nil, nil
		}

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(account.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgSetStatusRequest, "plans does not exist for provider"), nil, nil
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
			id     = plans[r.Intn(len(plans))].Id
			status = hubtypes.StatusActive
		)

		if r.Int63n(2) == 0 {
			status = hubtypes.StatusInactive
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgSetStatusRequest(
				hubtypes.ProvAddress(account.GetAddress()),
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

func SimulateMsgAddNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rProvider, _ = simulationtypes.RandomAcc(r, accounts)
			provider     = ak.GetAccount(ctx, rProvider.Address)
			rNode, _     = simulationtypes.RandomAcc(r, accounts)
			node         = ak.GetAccount(ctx, rNode.Address)
		)

		found := pk.HasProvider(ctx, hubtypes.ProvAddress(provider.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "provider does not exist"), nil, nil
		}

		_, found = k.GetNode(ctx, hubtypes.NodeAddress(node.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "node does not exist"), nil, nil
		}

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(provider.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "plans does not exist for provider"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, provider.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "balance cannot be negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found = k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(node.GetAddress()))
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, "node for plan already exists"), nil, nil
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgAddNodeRequest(
				hubtypes.ProvAddress(provider.GetAddress()),
				id,
				hubtypes.NodeAddress(node.GetAddress()),
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{provider.GetAccountNumber()},
			[]uint64{provider.GetSequence()},
			rProvider.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgAddNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

func SimulateMsgRemoveNodeRequest(ak expected.AccountKeeper, bk expected.BankKeeper, pk expected.ProviderKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rProvider, _ = simulationtypes.RandomAcc(r, accounts)
			provider     = ak.GetAccount(ctx, rProvider.Address)
			rNode, _     = simulationtypes.RandomAcc(r, accounts)
			node         = ak.GetAccount(ctx, rNode.Address)
		)

		found := pk.HasProvider(ctx, hubtypes.ProvAddress(provider.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "provider does not exist"), nil, nil
		}

		_, found = k.GetNode(ctx, hubtypes.NodeAddress(node.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "node does not exist"), nil, nil
		}

		plans := k.GetPlansForProvider(ctx, hubtypes.ProvAddress(provider.GetAddress()), 0, 0)
		if len(plans) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "plans does not exist for provider"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, provider.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "balance cannot be negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		var (
			id = plans[r.Intn(len(plans))].Id
		)

		found = k.HasNodeForPlan(ctx, id, hubtypes.NodeAddress(node.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, "node for plan does not exist"), nil, nil
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgRemoveNodeRequest(
				provider.GetAddress(),
				id,
				hubtypes.NodeAddress(node.GetAddress()),
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{provider.GetAccountNumber()},
			[]uint64{provider.GetSequence()},
			rProvider.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRemoveNodeRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, ""), nil, nil
	}
}

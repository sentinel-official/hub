package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

const (
	OpWeightMsgRegisterRequest  = "op_weight_msg_register_request"
	OpWeightMsgSetStatusRequest = "op_weight_msg_set_status_request"
	OpWeightMsgUpdateRequest    = "op_weight_msg_update_request"
)

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var (
		weightMsgRegisterRequest  int
		weightMsgSetStatusRequest int
		weightMsgUpdateRequest    int
	)

	randMsgRegisterRequest := func(_ *rand.Rand) {
		weightMsgRegisterRequest = 100
	}

	randMsgSetStatus := func(_ *rand.Rand) {
		weightMsgSetStatusRequest = 100
	}

	randMsgUpdateRequest := func(_ *rand.Rand) {
		weightMsgUpdateRequest = 100
	}

	ap.GetOrGenerate(cdc, OpWeightMsgRegisterRequest, &weightMsgRegisterRequest, nil, randMsgRegisterRequest)
	ap.GetOrGenerate(cdc, OpWeightMsgSetStatusRequest, &weightMsgSetStatusRequest, nil, randMsgSetStatus)
	ap.GetOrGenerate(cdc, OpWeightMsgUpdateRequest, &weightMsgUpdateRequest, nil, randMsgUpdateRequest)

	registerOperation := simulation.NewWeightedOperation(weightMsgRegisterRequest, SimulateMsgRegisterRequest(k))
	setStatusOperation := simulation.NewWeightedOperation(weightMsgSetStatusRequest, SimulateMsgSetStatus(k))
	updateOperation := simulation.NewWeightedOperation(weightMsgUpdateRequest, SimulateMsgUpdateRequest(k))

	return simulation.WeightedOperations{
		registerOperation,
		setStatusOperation,
		updateOperation,
	}
}

func SimulateMsgRegisterRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		acc, _ := sdksimulation.RandomAcc(r, accounts)
		from := k.GetAccount(ctx, acc.Address)
		nodeAddress := getNodeAddress()
		provider := hubtypes.ProvAddress(acc.Address.Bytes())

		_, found := k.GetNode(ctx, nodeAddress)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", "node is already registered"), nil, nil
		}

		denom := k.GetParams(ctx).Deposit.Denom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
		}

		remoteURL := sdksimulation.RandStringOfLength(r, r.Intn(28)+4)

		msg := types.NewMsgRegisterRequest(from.GetAddress(), provider, price, remoteURL)
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
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
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
		nodeAddress, nodeAccount := getNodeAccountI(ctx, k)
		provider := hubtypes.ProvAddress(acc.Address.Bytes())

		_, found := k.GetNode(ctx, nodeAddress)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", "node is not registered"), nil, nil
		}

		denom := k.GetParams(ctx).Deposit.Denom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		remoteURL := sdksimulation.RandStringOfLength(r, r.Intn(28)+4)

		msg := types.NewMsgUpdateRequest(nodeAddress, provider, price, remoteURL)
		txConfig := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{nodeAccount.GetAccountNumber()},
			[]uint64{nodeAccount.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgSetStatus(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		_ []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		nodeAddress, nodeAccount := getNodeAccountI(ctx, k)

		_, found := k.GetNode(ctx, nodeAddress)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", "node is not registered"), nil, nil
		}

		denom := k.GetParams(ctx).Deposit.Denom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		price := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, price)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		msg := types.NewMsgSetStatusRequest(nodeAddress, hubtypes.Active)
		txConfig := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{nodeAccount.GetAccountNumber()},
			[]uint64{nodeAccount.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "set_status_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func getNodeAccountI(ctx sdk.Context, k keeper.Keeper) (hubtypes.NodeAddress, authtypes.AccountI) {
	nodeAddress := getNodeAddress()
	nodeAccount := k.GetAccount(ctx, sdk.AccAddress(nodeAddress.Bytes()))

	return nodeAddress, nodeAccount
}

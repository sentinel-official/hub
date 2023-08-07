// DO NOT COVER

package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/node/expected"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

var (
	typeMsgRegister      = sdk.MsgTypeURL((*types.MsgRegisterRequest)(nil))
	typeMsgUpdateDetails = sdk.MsgTypeURL((*types.MsgUpdateDetailsRequest)(nil))
	typeMsgUpdateStatus  = sdk.MsgTypeURL((*types.MsgUpdateStatusRequest)(nil))
	typeMsgSubscribe     = sdk.MsgTypeURL((*types.MsgSubscribeRequest)(nil))
)

func WeightedOperations(
	cdc codec.Codec,
	txConfig client.TxConfig,
	params simtypes.AppParams,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgRegister      int
		weightMsgUpdateDetails int
		weightMsgUpdateStatus  int
		weightMsgSubscribe     int
	)

	params.GetOrGenerate(
		cdc,
		typeMsgRegister,
		&weightMsgRegister,
		nil,
		func(_ *rand.Rand) {
			weightMsgRegister = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgUpdateDetails,
		&weightMsgUpdateDetails,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDetails = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgUpdateStatus,
		&weightMsgUpdateStatus,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStatus = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgSubscribe,
		&weightMsgSubscribe,
		nil,
		func(_ *rand.Rand) {
			weightMsgSubscribe = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgRegister, SimulateMsgRegister(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUpdateDetails, SimulateMsgUpdateDetails(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUpdateStatus, SimulateMsgUpdateStatus(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgSubscribe, SimulateMsgSubscribe(txConfig, ak, bk, k)),
	}
}

func SimulateMsgRegister(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgRegister, ""), nil, nil
	}
}

func SimulateMsgUpdateDetails(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdateDetails, ""), nil, nil
	}
}

func SimulateMsgUpdateStatus(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdateStatus, ""), nil, nil
	}
}

func SimulateMsgSubscribe(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgSubscribe, ""), nil, nil
	}
}

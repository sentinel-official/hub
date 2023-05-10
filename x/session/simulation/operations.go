// DO NOT COVER

package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

var (
	OperationWeightMsgStartRequest         = "op_weight_" + types.TypeMsgStartRequest
	OperationWeightMsgUpdateDetailsRequest = "op_weight_" + types.TypeMsgUpdateDetailsRequest
	OperationWeightMsgEndRequest           = "op_weight_" + types.TypeMsgEndRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgStartRequest         int
		weightMsgUpdateDetailsRequest int
		weightMsgEndRequest           int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgStartRequest,
		&weightMsgStartRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgStartRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUpdateDetailsRequest,
		&weightMsgUpdateDetailsRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDetailsRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgEndRequest,
		&weightMsgEndRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgEndRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgStartRequest,
			SimulateMsgStartRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateDetailsRequest,
			SimulateMsgUpdateDetailsRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgEndRequest,
			SimulateMsgEndRequest(ak, bk, k),
		),
	}
}

func SimulateMsgStartRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, ""), nil, nil
	}
}

func SimulateMsgUpdateDetailsRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateDetailsRequest, ""), nil, nil
	}
}

func SimulateMsgEndRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, ""), nil, nil
	}
}

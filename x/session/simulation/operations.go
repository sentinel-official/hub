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

	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

var (
	typeMsgStart         = sdk.MsgTypeURL((*types.MsgStartRequest)(nil))
	typeMsgUpdateDetails = sdk.MsgTypeURL((*types.MsgUpdateDetailsRequest)(nil))
	typeMsgEnd           = sdk.MsgTypeURL((*types.MsgEndRequest)(nil))
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
		weightMsgStart         int
		weightMsgUpdateDetails int
		weightMsgEnd           int
	)

	params.GetOrGenerate(
		cdc,
		typeMsgStart,
		&weightMsgStart,
		nil,
		func(_ *rand.Rand) {
			weightMsgStart = 100
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
		typeMsgEnd,
		&weightMsgEnd,
		nil,
		func(_ *rand.Rand) {
			weightMsgEnd = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgStart, SimulateMsgStart(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUpdateDetails, SimulateMsgUpdateDetails(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgEnd, SimulateMsgEnd(txConfig, ak, bk, k)),
	}
}

func SimulateMsgStart(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgStart, ""), nil, nil
	}
}

func SimulateMsgUpdateDetails(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdateDetails, ""), nil, nil
	}
}

func SimulateMsgEnd(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgEnd, ""), nil, nil
	}
}

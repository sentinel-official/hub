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

	"github.com/sentinel-official/hub/x/plan/expected"
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/types"
)

var (
	typeMsgCreate       = sdk.MsgTypeURL((*types.MsgCreateRequest)(nil))
	typeMsgUpdateStatus = sdk.MsgTypeURL((*types.MsgUpdateStatusRequest)(nil))
	typeMsgLinkNode     = sdk.MsgTypeURL((*types.MsgLinkNodeRequest)(nil))
	typeMsgUnlinkNode   = sdk.MsgTypeURL((*types.MsgUnlinkNodeRequest)(nil))
	typeMsgSubscribe    = sdk.MsgTypeURL((*types.MsgSubscribeRequest)(nil))
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
		weightMsgCreate       int
		weightMsgUpdateStatus int
		weightMsgLinkNode     int
		weightMsgUnlinkNode   int
		weightMsgSubscribe    int
	)

	params.GetOrGenerate(
		cdc,
		typeMsgCreate,
		&weightMsgCreate,
		nil,
		func(_ *rand.Rand) {
			weightMsgCreate = 100
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
		typeMsgLinkNode,
		&weightMsgLinkNode,
		nil,
		func(_ *rand.Rand) {
			weightMsgLinkNode = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgUnlinkNode,
		&weightMsgUnlinkNode,
		nil,
		func(_ *rand.Rand) {
			weightMsgUnlinkNode = 100
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
		simulation.NewWeightedOperation(weightMsgCreate, SimulateMsgCreate(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUpdateStatus, SimulateMsgUpdateStatus(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgLinkNode, SimulateMsgLinkNode(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgUnlinkNode, SimulateMsgUnlinkNode(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgSubscribe, SimulateMsgSubscribe(txConfig, ak, bk, k)),
	}
}

func SimulateMsgCreate(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgCreate, ""), nil, nil
	}
}

func SimulateMsgUpdateStatus(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgUpdateStatus, ""), nil, nil
	}
}

func SimulateMsgLinkNode(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgLinkNode, ""), nil, nil
	}
}

func SimulateMsgUnlinkNode(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgUnlinkNode, ""), nil, nil
	}
}

func SimulateMsgSubscribe(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgSubscribe, ""), nil, nil
	}
}

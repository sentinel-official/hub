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

	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

var (
	typeMsgCancel   = sdk.MsgTypeURL((*types.MsgCancelRequest)(nil))
	typeMsgAllocate = sdk.MsgTypeURL((*types.MsgAllocateRequest)(nil))
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
		weightMsgCancel   int
		weightMsgAllocate int
	)

	params.GetOrGenerate(
		cdc,
		typeMsgCancel,
		&weightMsgCancel,
		nil,
		func(_ *rand.Rand) {
			weightMsgCancel = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		typeMsgAllocate,
		&weightMsgAllocate,
		nil,
		func(_ *rand.Rand) {
			weightMsgAllocate = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightMsgCancel, SimulateMsgCancel(txConfig, ak, bk, k)),
		simulation.NewWeightedOperation(weightMsgAllocate, SimulateMsgAllocate(txConfig, ak, bk, k)),
	}
}

func SimulateMsgCancel(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgCancel, ""), nil, nil
	}
}

func SimulateMsgAllocate(txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simtypes.Account, chainID string) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		return simtypes.NoOpMsg(types.ModuleName, typeMsgAllocate, ""), nil, nil
	}
}

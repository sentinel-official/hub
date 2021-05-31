package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/sentinel-official/hub/x/swap/keeper"
	types "github.com/sentinel-official/hub/x/swap/types"
)

const OpWeightMsgSwapRequest = "op_weight_msg_swap_request"

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var weightMsgSwapRequest int

	randSwapFn := func(_ *rand.Rand) {
		weightMsgSwapRequest = 100
	}
	ap.GetOrGenerate(cdc, OpWeightMsgSwapRequest, &weightMsgSwapRequest, nil, randSwapFn)

	operation := simulation.NewWeightedOperation(weightMsgSwapRequest, SimulateMsgSwapRequest(k, cdc))
	return simulation.WeightedOperations{operation}
}

func SimulateMsgSwapRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc1, _ := sdksimulation.RandomAcc(r, accounts)
		acc2, _ := sdksimulation.RandomAcc(r, accounts)
		sender := k.GetAccount(ctx, acc1.Address)
		receiver := k.GetAccount(ctx, acc2.Address)
		hash := types.EthereumHash{}

		denom := k.GetParams(ctx).SwapDenom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		coins := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, coins)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "swap_request", err.Error()), nil, err
		}

		_, found := k.GetSwap(ctx, hash)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "swap_request", "swap already exists for this txn hash"), nil, nil
		}

		msg := types.NewMsgSwapRequest(sender.GetAddress(), hash, receiver.GetAddress(), amount)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{sender.GetAccountNumber()},
			[]uint64{sender.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "swap_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "swap_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

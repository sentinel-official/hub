package simulation

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	swapkeeper "github.com/sentinel-official/hub/x/swap/expected"
	"github.com/sentinel-official/hub/x/swap/keeper"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
)

const (
	OpWeightMsgSwapRequest = "op_weight_msg_swap_request"
)

func SimulateMsgSwapRequest(bk swapkeeper.BankKeeper, ak swapkeeper.AccountKeeper, k keeper.Keeper, cdc codec.JSONMarshaler) simtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accs []simtypes.Account,
		chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		acc1, _ := simtypes.RandomAcc(r, accs)
		acc2, _ := simtypes.RandomAcc(r, accs)
		sender := ak.GetAccount(ctx, acc1.Address)
		receiver := ak.GetAccount(ctx, acc2.Address)
		hash := swaptypes.EthereumHash{}

		denom := k.GetParams(ctx).SwapDenom
		amount := simtypes.RandomAmount(r, sdk.NewInt(60<<13))

		coins := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := simtypes.RandomFees(r, ctx, coins)
		if err != nil {
			return simtypes.NoOpMsg(swaptypes.ModuleName, "swap_request", err.Error()), nil, err
		}

		_, found := k.GetSwap(ctx, hash)
		if found {
			return simtypes.NoOpMsg(swaptypes.ModuleName, "swap_request", "swap already exists for this txn hash"), nil, nil
		}

		msg := swaptypes.NewMsgSwapRequest(sender.GetAddress(), hash, receiver.GetAddress(), amount)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig

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
			return simtypes.NoOpMsg(swaptypes.ModuleName, "swap_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return simtypes.NoOpMsg(swaptypes.ModuleName, "swap_request", err.Error()), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

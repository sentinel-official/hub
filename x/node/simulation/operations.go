package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	hubbtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/keeper"
	"github.com/sentinel-official/hub/x/node/types"
)

func SimulateMsgRegisterRequest(k keeper.Keeper) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) ( sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		bz := make([]byte, 20)
		_, err := r.Read(bz)
		if err != nil {
			panic(err)
		}

		from := sdk.AccAddress(bz)
		nodeAddress := getNodeAddress()
		provider := hubbtypes.ProvAddress(bz)

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

		msg := types.NewMsgRegisterRequest(from, provider, price, remoteURL)
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
	}
}

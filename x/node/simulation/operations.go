package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	hubtypes "github.com/sentinel-official/hub/types"
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
	) ( sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

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
	) ( sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {

		nodeAddress, nodeAccount := getNodeAccountI(ctx, k)

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
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func getNodeAccountI(ctx sdk.Context, k keeper.Keeper) (hubtypes.NodeAddress, authtypes.AccountI) {
	nodeAddress := getNodeAddress()
	nodeAccount := k.GetAccount(ctx, sdk.AccAddress(nodeAddress.Bytes()))

	return nodeAddress, nodeAccount
}

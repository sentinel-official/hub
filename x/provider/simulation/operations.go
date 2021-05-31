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
	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/keeper"
	types "github.com/sentinel-official/hub/x/provider/types"
)

const (
	OpWeightMsgRegisterRequest = "op_weight_msg_register_request"
	OpWeightMsgUpdateRequest   = "op_weight_msg_update_request"
)

func WeightedOperations(ap sdksimulation.AppParams, cdc codec.JSONMarshaler, k keeper.Keeper) simulation.WeightedOperations {
	var (
		weightMsgRegisterRequest int
		weightMsgUpdateRequest   int
	)

	randRegisterFn := func(_ *rand.Rand) {
		weightMsgRegisterRequest = 100
	}

	randUpdateFn := func(_ *rand.Rand) {
		weightMsgUpdateRequest = 100
	}
	ap.GetOrGenerate(cdc, OpWeightMsgRegisterRequest, &weightMsgRegisterRequest, nil, randRegisterFn)
	ap.GetOrGenerate(cdc, OpWeightMsgUpdateRequest, &weightMsgUpdateRequest, nil, randUpdateFn)

	registerOperation := simulation.NewWeightedOperation(weightMsgRegisterRequest, SimulateMsgRegisterRequest(k, cdc))
	updateOperation := simulation.NewWeightedOperation(weightMsgUpdateRequest, SimulateMsgUpdateRequest(k, cdc))

	return simulation.WeightedOperations{
		registerOperation,
		updateOperation,
	}
}

func SimulateMsgRegisterRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		providerAccount := k.GetAccount(ctx, acc.Address)
		providerAddress := hubtypes.ProvAddress(providerAccount.GetAddress().Bytes())

		denom := k.GetParams(ctx).Deposit.Denom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		coins := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, coins)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
		}

		_, found := k.GetProvider(ctx, providerAddress)
		if found {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", "provider already exists"), nil, nil
		}

		name := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		identity := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		website := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		description := sdksimulation.RandStringOfLength(r, r.Intn(250)+6)

		msg := types.NewMsgRegisterRequest(providerAccount.GetAddress(), name, identity, website, description)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{providerAccount.GetAccountNumber()},
			[]uint64{providerAccount.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "register_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func SimulateMsgUpdateRequest(k keeper.Keeper, cdc codec.JSONMarshaler) sdksimulation.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []sdksimulation.Account,
		chainID string,
	) (sdksimulation.OperationMsg, []sdksimulation.FutureOperation, error) {
		acc, _ := sdksimulation.RandomAcc(r, accounts)
		providerAccount := k.GetAccount(ctx, acc.Address)
		providerAddress := hubtypes.ProvAddress(providerAccount.GetAddress().Bytes())

		denom := k.GetParams(ctx).Deposit.Denom
		amount := sdksimulation.RandomAmount(r, sdk.NewInt(60<<13))

		coins := sdk.Coins{
			{Denom: denom, Amount: amount},
		}

		fees, err := sdksimulation.RandomFees(r, ctx, coins)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		_, found := k.GetProvider(ctx, providerAddress)
		if !found {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", "provider does not exist"), nil, nil
		}

		name := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		identity := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		website := sdksimulation.RandStringOfLength(r, r.Intn(60)+4)
		description := sdksimulation.RandStringOfLength(r, r.Intn(250)+6)

		msg := types.NewMsgUpdateRequest(providerAddress, name, identity, website, description)
		txGen := params.MakeTestEncodingConfig().TxConfig

		txn, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{providerAccount.GetAccountNumber()},
			[]uint64{providerAccount.GetSequence()},
		)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), txn)
		if err != nil {
			return sdksimulation.NoOpMsg(types.ModuleName, "update_request", err.Error()), nil, err
		}

		return sdksimulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

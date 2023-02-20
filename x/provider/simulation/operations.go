package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/expected"
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

var (
	OperationWeightMsgRegisterRequest = "op_weight_" + types.TypeMsgRegisterRequest
	OperationWeightMsgUpdateRequest   = "op_weight_" + types.TypeMsgUpdateRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgRegisterRequest int
		weightMsgUpdateRequest   int
	)

	params.GetOrGenerate(
		cdc,
		OperationWeightMsgRegisterRequest,
		&weightMsgRegisterRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgRegisterRequest = 100
		},
	)
	params.GetOrGenerate(
		cdc,
		OperationWeightMsgUpdateRequest,
		&weightMsgUpdateRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequest = 100
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgRegisterRequest,
			SimulateMsgRegisterRequest(ak, bk, k),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateRequest,
			SimulateMsgUpdateRequest(ak, bk, k),
		),
	}
}

func SimulateMsgRegisterRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rAccount, _ = simulationtypes.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
		)

		found := k.HasProvider(ctx, hubtypes.ProvAddress(account.GetAddress()))
		if found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "provider already exists"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, err.Error()), nil, err
		}

		deposit := k.Deposit(ctx)
		if balance.Sub(fees...).AmountOf(deposit.Denom).LT(deposit.Amount) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgRegisterRequest, "balance is less than deposit"), nil, nil
		}

		var (
			name        = simulationtypes.RandStringOfLength(r, r.Intn(MaxNameLength)+8)
			identity    = simulationtypes.RandStringOfLength(r, r.Intn(MaxIdentityLength))
			website     = simulationtypes.RandStringOfLength(r, r.Intn(MaxWebsiteLength))
			description = simulationtypes.RandStringOfLength(r, r.Intn(MaxDescriptionLength))
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgRegisterRequest(
				account.GetAddress(),
				name,
				identity,
				website,
				description,
			)
		)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         txConfig,
			Cdc:           nil,
			Msg:           message,
			MsgType:       message.Type(),
			Context:       ctx,
			SimAccount:    rAccount,
			AccountKeeper: ak,
			ModuleName:    types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}

func SimulateMsgUpdateRequest(ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) simulationtypes.Operation {
	return func(
		r *rand.Rand,
		app *baseapp.BaseApp,
		ctx sdk.Context,
		accounts []simulationtypes.Account,
		chainID string,
	) (simulationtypes.OperationMsg, []simulationtypes.FutureOperation, error) {
		var (
			rAccount, _ = simulationtypes.RandomAcc(r, accounts)
			account     = ak.GetAccount(ctx, rAccount.Address)
		)

		found := k.HasProvider(ctx, hubtypes.ProvAddress(account.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "provider does not exist"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, account.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		var (
			name        = simulationtypes.RandStringOfLength(r, r.Intn(MaxNameLength+8))
			identity    = simulationtypes.RandStringOfLength(r, r.Intn(MaxIdentityLength))
			website     = simulationtypes.RandStringOfLength(r, r.Intn(MaxWebsiteLength))
			description = simulationtypes.RandStringOfLength(r, r.Intn(MaxDescriptionLength))
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgUpdateRequest(
				hubtypes.ProvAddress(account.GetAddress()),
				name,
				identity,
				website,
				description,
			)
		)

		txCtx := simulation.OperationInput{
			R:             r,
			App:           app,
			TxGen:         txConfig,
			Cdc:           nil,
			Msg:           message,
			MsgType:       message.Type(),
			Context:       ctx,
			SimAccount:    rAccount,
			AccountKeeper: ak,
			ModuleName:    types.ModuleName,
		}

		return simulation.GenAndDeliverTx(txCtx, fees)
	}
}

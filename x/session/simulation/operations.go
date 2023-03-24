package simulation

import (
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
)

var (
	OperationWeightMsgStartRequest  = "op_weight_" + types.TypeMsgStartRequest
	OperationWeightMsgUpdateRequest = "op_weight_" + types.TypeMsgUpdateRequest
	OperationWeightMsgEndRequest    = "op_weight_" + types.TypeMsgEndRequest
)

func WeightedOperations(
	params simulationtypes.AppParams,
	cdc codec.JSONCodec,
	ak expected.AccountKeeper,
	bk expected.BankKeeper,
	k keeper.Keeper,
) simulation.WeightedOperations {
	var (
		weightMsgStartRequest  int
		weightMsgUpdateRequest int
		weightMsgEndRequest    int
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
		OperationWeightMsgUpdateRequest,
		&weightMsgUpdateRequest,
		nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRequest = 100
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
			weightMsgUpdateRequest,
			SimulateMsgUpdateRequest(ak, bk, k),
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
		var (
			rFrom, _    = simulationtypes.RandomAcc(r, accounts)
			from        = ak.GetAccount(ctx, rFrom.Address)
			rAddress, _ = simulationtypes.RandomAcc(r, accounts)
			address     = ak.GetAccount(ctx, rAddress.Address)
		)

		subscriptions := k.GetActiveSubscriptionsForAddress(ctx, from.GetAddress(), 0, 0)
		if len(subscriptions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, "active subscriptions for address does not exist"), nil, nil
		}

		var (
			rSubscription = subscriptions[r.Intn(len(subscriptions))]
		)

		node, found := k.GetNode(ctx, hubtypes.NodeAddress(address.GetAddress()))
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, "node does not exist"), nil, nil
		}
		if !node.Status.Equal(hubtypes.StatusActive) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, "node status is not active"), nil, nil
		}
		if rSubscription.Plan != 0 {
			if k.HasNodeForPlan(ctx, rSubscription.Plan, hubtypes.NodeAddress(address.String())) {
				return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, "node for plan does not exist"), nil, nil
			}
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgStartRequest(
				from.GetAddress(),
				rSubscription.Id,
				hubtypes.NodeAddress(address.GetAddress()),
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgStartRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
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
			rFrom, _ = simulationtypes.RandomAcc(r, accounts)
			from     = ak.GetAccount(ctx, rFrom.Address)
			rId      = uint64(r.Int63n(1 << 18))
		)

		rSession, found := k.GetSession(ctx, rId)
		if !found {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "session does not exist"), nil, nil
		}
		if rSession.Status.Equal(hubtypes.StatusInactive) {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "session status is inactive"), nil, nil
		}
		if rSession.Node == hubtypes.NodeAddress(from.GetAddress()).String() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "session does not belong to node"), nil, nil
		}

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		var (
			duration  = time.Duration(r.Int63n(MaxSessionDuration)) * time.Minute
			bandwidth = hubtypes.Bandwidth{
				Upload:   sdk.NewInt(r.Int63n(MaxSessionBandwidthUpload)),
				Download: sdk.NewInt(r.Int63n(MaxSessionBandwidthDownload)),
			}
		)

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgUpdateRequest(
				hubtypes.NodeAddress(from.GetAddress()),
				types.Proof{
					Id:        rSession.Id,
					Duration:  duration,
					Bandwidth: bandwidth,
				},
				nil,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgUpdateRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
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
		var (
			rFrom, _ = simulationtypes.RandomAcc(r, accounts)
			from     = ak.GetAccount(ctx, rFrom.Address)
		)

		sessions := k.GetActiveSessionsForAddress(ctx, from.GetAddress(), 0, 0)
		if len(sessions) == 0 {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, "sessions for address does not exist"), nil, nil
		}

		var (
			rSession = sessions[r.Intn(len(sessions))]
		)

		balance := bk.SpendableCoins(ctx, from.GetAddress())
		if !balance.IsAnyNegative() {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, "balance is negative"), nil, nil
		}

		fees, err := simulationtypes.RandomFees(r, ctx, balance)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, err.Error()), nil, err
		}

		var (
			txConfig = params.MakeTestEncodingConfig().TxConfig
			message  = types.NewMsgEndRequest(
				from.GetAddress(),
				rSession.Id,
				0,
			)
		)

		txn, err := helpers.GenTx(
			txConfig,
			[]sdk.Msg{message},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{from.GetAccountNumber()},
			[]uint64{from.GetSequence()},
			rFrom.PrivKey,
		)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, err.Error()), nil, err
		}

		_, _, err = app.Deliver(txConfig.TxEncoder(), txn)
		if err != nil {
			return simulationtypes.NoOpMsg(types.ModuleName, types.TypeMsgEndRequest, err.Error()), nil, err
		}

		return simulationtypes.NewOperationMsg(message, true, "", nil), nil, nil
	}
}

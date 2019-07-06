package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/vpn"
)

func SimulateMsgRegisterNode(vpnKeeper vpn.Keeper, accountKeeper auth.AccountKeeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		comment, msg, ok := createMsgRegisterNode(r, ctx, accounts, accountKeeper)
		operationMsg = simulation.NewOperationMsg(msg, ok, comment)
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgRegisterNode(r *rand.Rand, ctx sdk.Context, accounts []simulation.Account,
	accountKeeper auth.AccountKeeper) (comment string, msg *vpn.MsgRegisterNode, ok bool) {

	randomAcc := simulation.RandomAcc(r, accounts)
	coins := accountKeeper.GetAccount(ctx, randomAcc.Address).SpendableCoins(ctx.BlockHeader().Time)
	if len(coins) == 0 {
		return "skipping register_node, no coins in account", &vpn.MsgRegisterNode{}, false
	}

	msg = vpn.NewMsgRegisterNode(randomAcc.Address, getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
		getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

	return "", msg, true
}

func SimulateMsgUpdateNodeInfo(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		msg, ok := createMsgUpdateNodeInfo(r, accounts)
		operationMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgUpdateNodeInfo(r *rand.Rand, accounts []simulation.Account) (msg *vpn.MsgUpdateNodeInfo, ok bool) {
	randomAcc := simulation.RandomAcc(r, accounts)
	msg = vpn.NewMsgUpdateNodeInfo(randomAcc.Address, getRandomID(r), getRandomType(r), getRandomVersion(r),
		getRandomMoniker(r), getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

	return msg, true
}

func SimulateMsgUpdateNodeStatus(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		msg, ok := createMsgUpdateNodeStatus(r, accounts)
		operationMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgUpdateNodeStatus(r *rand.Rand, accounts []simulation.Account) (msg *vpn.MsgUpdateNodeStatus, ok bool) {
	randomAcc := simulation.RandomAcc(r, accounts)
	msg = vpn.NewMsgUpdateNodeStatus(randomAcc.Address, getRandomID(r), getRandomStatus(r))

	return msg, true
}

func SimulateMsgStartSubscription(vpnKeeper vpn.Keeper, accountKeeper auth.AccountKeeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		comment, msg, ok := createMsgStartSubscription(r, ctx, accounts, accountKeeper)
		operationMsg = simulation.NewOperationMsg(msg, ok, comment)
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgStartSubscription(r *rand.Rand, ctx sdk.Context, accounts []simulation.Account,
	accountKeeper auth.AccountKeeper) (comment string, msg *vpn.MsgStartSubscription, ok bool) {

	randomAcc := simulation.RandomAcc(r, accounts)
	coins := accountKeeper.GetAccount(ctx, randomAcc.Address).SpendableCoins(ctx.BlockHeader().Time)
	if len(coins) == 0 {
		return "skipping start_subscription, no coins in account", &vpn.MsgStartSubscription{}, false
	}

	msg = vpn.NewMsgStartSubscription(randomAcc.Address, getRandomID(r), getRandomCoin(r))

	return "", msg, true
}

func SimulateMsgEndSubscription(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		msg, ok := createMsgEndSubscription(r, accounts)
		operationMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgEndSubscription(r *rand.Rand, accounts []simulation.Account) (msg *vpn.MsgEndSubscription, ok bool) {
	randomAcc := simulation.RandomAcc(r, accounts)
	msg = vpn.NewMsgEndSubscription(randomAcc.Address, getRandomID(r))

	return msg, true
}

func SimulateMsgUpdateSessionInfo(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		operationMsg simulation.OperationMsg, futureOps []simulation.FutureOperation, err error) {

		msg, ok := createMsgUpdateSessionInfo(r, accounts)
		operationMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return operationMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return operationMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return operationMsg, nil, err
			}
		}

		return operationMsg, nil, nil
	}
}

func createMsgUpdateSessionInfo(r *rand.Rand, accounts []simulation.Account) (msg *vpn.MsgUpdateSessionInfo, ok bool) {
	randomAcc := simulation.RandomAcc(r, accounts)
	msg = vpn.NewMsgUpdateSessionInfo(randomAcc.Address, getRandomID(r), getRandomBandwidth(r),
		getRandomSignData(r, accounts), getRandomSignData(r, accounts))

	return msg, true
}

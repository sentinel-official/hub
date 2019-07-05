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
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {

		comment, msg, ok := createMsgRegisterNode(r, ctx, accs, accountKeeper)
		opMsg = simulation.NewOperationMsg(msg, ok, comment)
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgRegisterNode(r *rand.Rand, ctx sdk.Context, accs []simulation.Account,
	accountKeeper auth.AccountKeeper) (comment string, msg *vpn.MsgRegisterNode, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)
	coins := accountKeeper.GetAccount(ctx, randAcc.Address).SpendableCoins(ctx.BlockHeader().Time)
	if len(coins) == 0 {
		return "skipping register_node, no coins in account", &vpn.MsgRegisterNode{}, false
	}

	msg = vpn.NewMsgRegisterNode(randAcc.Address, getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
		getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

	return "", msg, true
}

func SimulateMsgUpdateNodeInfo(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		msg, ok := createMsgUpdateNodeInfo(r, accs)
		opMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgUpdateNodeInfo(r *rand.Rand, accs []simulation.Account) (msg *vpn.MsgUpdateNodeInfo, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)
	msg = vpn.NewMsgUpdateNodeInfo(randAcc.Address, getRandomID(r), getRandomType(r), getRandomVersion(r),
		getRandomMoniker(r), getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

	return msg, true
}

func SimulateMsgUpdateNodeStatus(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		msg, ok := createMsgUpdateNodeStatus(r, accs)
		opMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgUpdateNodeStatus(r *rand.Rand, accs []simulation.Account) (msg *vpn.MsgUpdateNodeStatus, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)
	msg = vpn.NewMsgUpdateNodeStatus(randAcc.Address, getRandomID(r), getRandomStatus(r))

	return msg, true
}

func SimulateMsgStartSubscription(vpnKeeper vpn.Keeper, accountKeeper auth.AccountKeeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		comment, msg, ok := createMsgStartSubscription(r, ctx, accs, accountKeeper)
		opMsg = simulation.NewOperationMsg(msg, ok, comment)
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgStartSubscription(r *rand.Rand, ctx sdk.Context, accs []simulation.Account,
	accountKeeper auth.AccountKeeper) (comment string, msg *vpn.MsgStartSubscription, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)

	coins := accountKeeper.GetAccount(ctx, randAcc.Address).SpendableCoins(ctx.BlockHeader().Time)
	if len(coins) == 0 {
		return "skipping start_subscription, no coins in account", &vpn.MsgStartSubscription{}, false
	}

	msg = vpn.NewMsgStartSubscription(randAcc.Address, getRandomID(r), getRandomCoin(r))

	return "", msg, true
}

func SimulateMsgEndSubscription(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		msg, ok := createMsgEndSubscription(r, accs)
		opMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgEndSubscription(r *rand.Rand, accs []simulation.Account) (msg *vpn.MsgEndSubscription, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)
	msg = vpn.NewMsgEndSubscription(randAcc.Address, getRandomID(r))

	return msg, true
}

func SimulateMsgUpdateSessionInfo(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account) (
		opMsg simulation.OperationMsg, fOps []simulation.FutureOperation, err error) {
		msg, ok := createMsgUpdateSessionInfo(r, accs)
		opMsg = simulation.NewOperationMsg(msg, ok, "")
		if !ok {
			return opMsg, nil, nil
		}

		if msg.ValidateBasic() != nil {
			return opMsg, nil, nil
		}

		if handler != nil {
			res := handler(ctx, *msg)
			if !res.IsOK() {
				return opMsg, nil, err
			}
		}

		return opMsg, nil, nil
	}
}

func createMsgUpdateSessionInfo(r *rand.Rand, accs []simulation.Account) (msg *vpn.MsgUpdateSessionInfo, ok bool) {
	randAcc := simulation.RandomAcc(r, accs)
	msg = vpn.NewMsgUpdateSessionInfo(randAcc.Address, getRandomID(r), getRandomBandwidth(r), getRandomSignData(r, accs), getRandomSignData(r, accs))

	return msg, true
}

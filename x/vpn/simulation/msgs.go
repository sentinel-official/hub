package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/vpn"
)

func SimulateMsgRegisterNode(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgRegisterNode(randomAcc.Address,
			getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
			getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdateNodeInfo(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgUpdateNodeInfo(randomAcc.Address, getRandomID(r),
			getRandomType(r), getRandomVersion(r), getRandomMoniker(r),
			getRandomCoins(r), getRandomBandwidth(r), getRandomEncryption(r))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdateNodeStatus(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgUpdateNodeStatus(randomAcc.Address, getRandomID(r), getRandomStatus(r))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgStartSubscription(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgStartSubscription(randomAcc.Address, getRandomID(r), getRandomCoin(r))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgEndSubscription(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgEndSubscription(randomAcc.Address, getRandomID(r))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdateSessionInfo(vpnKeeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(vpnKeeper)

	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {

		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgUpdateSessionInfo(randomAcc.Address, getRandomID(r),
			getRandomBandwidth(r), getRandomBandwidthSignature(r, accounts), getRandomBandwidthSignature(r, accounts))

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

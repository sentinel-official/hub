package simulation

import (
	"fmt"
	"math/rand"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	
	hub "github.com/sentinel-official/hub/types"
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
		if len(keeper.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		node := vpn.RandomNode(r, ctx, keeper)
		msg := vpn.NewMsgUpdateNodeInfo(node.Owner, node.ID,
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

func SimulateMsgDeregisterNode(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)
	
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		if len(keeper.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		node := vpn.RandomNode(r, ctx, keeper)
		msg := vpn.NewMsgDeregisterNode(node.Owner, node.ID)
		
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
		if len(keeper.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		node := vpn.RandomNode(r, ctx, keeper)
		resolver := vpn.RandomResolver(r, ctx, keeper)
		node.Status = vpn.StatusRegistered
		keeper.SetNode(ctx, node)
		
		randomAcc := simulation.RandomAcc(r, accounts)
		msg := vpn.NewMsgStartSubscription(randomAcc.Address, resolver.ID, node.ID, getRandomCoin(r))
		
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
		if len(keeper.GetAllSubscriptions(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		subscription := vpn.RandomSubscription(r, ctx, keeper)
		msg := vpn.NewMsgEndSubscription(subscription.Client, subscription.ID)
		
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}
		
		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateMsgUpdateSessionInfo(keeper vpn.Keeper) simulation.Operation {
	handler := vpn.NewHandler(keeper)
	
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		if len(keeper.GetAllSessions(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		clientAccount := simulation.RandomAcc(r, accounts)
		nodeOwnerAccount := simulation.RandomAcc(r, accounts)
		
		session := vpn.RandomSession(r, ctx, keeper)
		subscription, _ := keeper.GetSubscription(ctx, session.SubscriptionID)
		
		subscription.Status = vpn.StatusActive
		subscription.Client = clientAccount.Address
		keeper.SetSubscription(ctx, subscription)
		
		scs := keeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
		node, _ := keeper.GetNode(ctx, subscription.NodeID)
		node.Owner = nodeOwnerAccount.Address
		keeper.SetNode(ctx, node)
		
		bandwidth := getRandomBandwidth(r)
		
		bandWidthSignData := hub.NewBandwidthSignatureData(subscription.ID, scs, bandwidth)
		clientAccountSignedData, _ := clientAccount.PrivKey.Sign(bandWidthSignData.Bytes())
		nodeOwnerAccountSignedData, _ := nodeOwnerAccount.PrivKey.Sign(bandWidthSignData.Bytes())
		
		clienStdSig := auth.StdSignature{
			PubKey:    clientAccount.PubKey,
			Signature: clientAccountSignedData,
		}
		nodeOwnerStdSig := auth.StdSignature{
			PubKey:    nodeOwnerAccount.PubKey,
			Signature: nodeOwnerAccountSignedData,
		}
		
		msg := vpn.NewMsgUpdateSessionInfo(clientAccount.Address, session.SubscriptionID,
			bandwidth, nodeOwnerStdSig, clienStdSig)
		
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(vpn.ModuleName), nil,
				fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}
		
		ok := handler(ctx, *msg).IsOK()
		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}

func SimulateEndBlock(keeper vpn.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		if len(keeper.GetAllNodes(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		if len(keeper.GetAllSessions(ctx)) == 0 {
			return simulation.NoOpMsg(vpn.ModuleName), nil, nil
		}
		
		vpn.EndBlock(ctx, keeper)
		return simulation.NewOperationMsgBasic(vpn.ModuleName, "end_block", "", true, nil), nil, nil
	}
}

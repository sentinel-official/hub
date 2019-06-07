package vpn

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func Test_handleRegisterNode(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := types.TestNodeValid

	msg := NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDZero, node.ID)
	require.Equal(t, types.TestMonikerValid, node.Moniker)
	count := vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	vpnKeeper.SetNodesCountOfAddress(ctx, types.TestAddress1, DefaultFreeNodesCount)
	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, coins)

	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsNil, coins)

	node, found = vpnKeeper.GetNode(ctx, types.TestIDPos)
	require.Equal(t, true, found)
	require.Equal(t, types.TestIDPos, node.ID)

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(2), count)
}

func Test_handleUpdateNodeInfo(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := types.TestNodeValid
	node.Status = StatusDeRegistered

	vpnKeeper.SetNode(ctx, node)
	msg := NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestNewNodeType, node.Type)
	require.Equal(t, types.TestNewVersion, node.Version)
	require.Equal(t, types.TestNewMonikerValid, node.Moniker)
	require.Equal(t, types.TestCoinsPos, node.PricesPerGB)
	require.Equal(t, types.TestNewEncryption, node.Encryption)
}

func Test_handleUpdateNodeStatus(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := types.TestNodeValid
	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)

	msg := NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusInactive)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusInactive)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, node.Status)
}

func Test_handleDeregisterNode(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := types.TestNodeValid
	node.Deposit = types.TestCoinZero

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg := NewMsgDeregisterNode(node.Owner, node.ID)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.NotEqual(t, StatusDeRegistered, node.Status)

	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node = types.TestNodeValid
	node.Deposit = types.TestCoinZero
	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)
}

func Test_handleStartSubscription(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := types.TestNodeValid
	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg := NewMsgStartSubscription(types.TestAddress2, types.TestIDZero, types.TestCoinPos)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, types.TestIDZero, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, types.TestIDZero, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	msg = NewMsgStartSubscription(types.TestAddress2, types.TestIDZero, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsEmpty, coins)

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	subscription, found := vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, types.TestAddress2, subscription.Client)
	require.Equal(t, types.TestIDZero, subscription.NodeID)
}

func Test_handleEndSubscription(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	msg := NewMsgEndSubscription(types.TestAddress2, types.TestIDZero)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	_, err = vpnKeeper.AddDeposit(ctx, types.TestAddress2, types.TestCoinPos)
	require.Nil(t, err)

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsNil, coins)

	subscription := types.TestSubscriptionValid
	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, subscription.Status)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)
}

func Test_handleUpdateSessionInfo(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	node := types.TestNodeValid
	subscription := types.TestSubscriptionValid
	subscription.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	vpnKeeper.SetSubscription(ctx, subscription)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	node, _ = vpnKeeper.GetNode(ctx, node.ID)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)
	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos2, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos2)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos2,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthNeg,
		types.TestNodeOwnerStdSignatureNeg, types.TestClientStdSignatureNeg)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthZero,
		types.TestNodeOwnerStdSignatureZero, types.TestClientStdSignatureZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	session, found := vpnKeeper.GetSession(ctx, types.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, types.TestBandwidthPos1, session.Bandwidth)

	count := vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	session.Status = StatusInactive
	vpnKeeper.SetSession(ctx, session)

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())
}

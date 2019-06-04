package vpn

import (
	"testing"

	"github.com/stretchr/testify/require"

	test "github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func Test_handleRegisterNode(t *testing.T) {
	ctx, _, _, vpnKeeper, accountKeeper, bankKeeper := test.TestCreateInput()

	account := accountKeeper.NewAccountWithAddress(ctx, test.TestAddress1)
	require.Nil(t, account.SetPubKey(test.TestPubkey1))
	accountKeeper.SetAccount(ctx, account)
	account = accountKeeper.GetAccount(ctx, test.TestAddress1)
	require.NotNil(t, account)

	coins, _, err := bankKeeper.AddCoins(ctx, test.TestAddress1, test.TestCoinsPos.Add(test.TestCoinsPos))
	require.Nil(t, err)
	require.Equal(t, test.TestCoinsPos.Add(test.TestCoinsPos), coins)

	params := types.Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		NodeInactiveInterval:    DefaultNodeInactiveInterval,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}

	vpnKeeper.SetParams(ctx, params)
	handler := NewHandler(vpnKeeper)
	node := test.TestNodeValid
	msg := NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	vpnKeeper.SetNodesCountOfAddress(ctx, test.TestAddress1, DefaultFreeNodesCount)
	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	account = accountKeeper.GetAccount(ctx, test.TestAddress1)
	require.NotNil(t, account)
	require.Equal(t, test.TestCoinsPos, account.GetCoins())

	count := vpnKeeper.GetNodesCountOfAddress(ctx, test.TestAddress1)
	require.Equal(t, DefaultFreeNodesCount+1, count)
}

func Test_handleUpdateNodeInfo(t *testing.T) {
	ctx, _, _, vpnKeeper, _, _ := test.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := test.TestNodeValid
	vpnKeeper.SetNode(ctx, node)
	msg := NewMsgUpdateNodeInfo(node.Owner, node.ID, test.TestNewNodeType, test.TestNewVersion, test.TestNewMonikerValid, test.TestCoinsPos, test.TestBandwidthPos1, test.TestNewEncryption)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, test.TestNewNodeType, node.Type)
	require.Equal(t, test.TestNewVersion, node.Version)
	require.Equal(t, test.TestNewMonikerValid, node.Moniker)
	require.Equal(t, test.TestCoinsPos, node.PricesPerGB)
	require.Equal(t, test.TestNewEncryption, node.Encryption)
}

func Test_handleUpdateNodeStatus(t *testing.T) {
	ctx, _, _, vpnKeeper, _, _ := test.TestCreateInput()

	handler := NewHandler(vpnKeeper)
	node := test.TestNodeValid
	vpnKeeper.SetNode(ctx, node)
	msg := NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusActive)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found := vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)
}

func Test_handleDeregisterNode(t *testing.T) {
	ctx, _, _, vpnKeeper, _, _ := test.TestCreateInput()

	params := types.Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		NodeInactiveInterval:    DefaultNodeInactiveInterval,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}

	vpnKeeper.SetParams(ctx, params)
	handler := NewHandler(vpnKeeper)
	node := test.TestNodeValid
	node.Deposit = test.TestCoinZero
	vpnKeeper.SetNode(ctx, node)
	node, found := vpnKeeper.GetNode(ctx, node.ID)
	msg := NewMsgDeregisterNode(node.Owner, node.ID)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)
}

func Test_handleStartSubscription(t *testing.T) {
	ctx, _, depositKeeper, vpnKeeper, accountKeeper, bankKeeper := test.TestCreateInput()

	account := accountKeeper.NewAccountWithAddress(ctx, test.TestAddress2)
	require.Nil(t, account.SetPubKey(test.TestPubkey2))
	accountKeeper.SetAccount(ctx, account)
	account = accountKeeper.GetAccount(ctx, test.TestAddress2)
	require.NotNil(t, account)

	coins, _, err := bankKeeper.AddCoins(ctx, test.TestAddress2, test.TestCoinsPos.Add(test.TestCoinsPos))
	require.Nil(t, err)
	require.Equal(t, test.TestCoinsPos.Add(test.TestCoinsPos), coins)

	params := types.Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		NodeInactiveInterval:    DefaultNodeInactiveInterval,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}

	vpnKeeper.SetParams(ctx, params)
	handler := NewHandler(vpnKeeper)
	node := test.TestNodeValid
	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	node, found := vpnKeeper.GetNode(ctx, node.ID)

	msg := NewMsgStartSubscription(test.TestAddress2, test.TestIDZero, test.TestCoinPos)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	deposit, found := depositKeeper.GetDeposit(ctx, test.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, test.TestCoinsPos, deposit.Coins)
	subscription, found := vpnKeeper.GetSubscription(ctx, test.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, test.TestAddress2, subscription.Client)
	require.Equal(t, test.TestIDZero, subscription.NodeID)
}

func Test_handleEndSubscription(t *testing.T) {
	ctx, _, depositKeeper, vpnKeeper, accountKeeper, bankKeeper := test.TestCreateInput()

	account := accountKeeper.NewAccountWithAddress(ctx, test.TestAddress2)
	require.Nil(t, account.SetPubKey(test.TestPubkey2))
	accountKeeper.SetAccount(ctx, account)
	account = accountKeeper.GetAccount(ctx, test.TestAddress2)
	require.NotNil(t, account)

	coins, _, err := bankKeeper.AddCoins(ctx, test.TestAddress2, test.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, test.TestCoinsPos, coins)

	params := types.Params{
		FreeNodesCount:          DefaultFreeNodesCount,
		Deposit:                 DefaultDeposit,
		NodeInactiveInterval:    DefaultNodeInactiveInterval,
		SessionInactiveInterval: DefaultSessionInactiveInterval,
	}

	vpnKeeper.SetParams(ctx, params)
	_, err = vpnKeeper.AddDeposit(ctx, test.TestAddress2, test.TestCoinPos)
	require.Nil(t, err)

	deposit, found := depositKeeper.GetDeposit(ctx, test.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, test.TestCoinsPos, deposit.Coins)

	account = accountKeeper.GetAccount(ctx, test.TestAddress2)
	require.Equal(t, test.TestCoinsNil, account.GetCoins())

	subscription := test.TestSubscriptionValid
	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgEndSubscription(test.TestAddress2, subscription.ID)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, subscription.Status)

	deposit, found = depositKeeper.GetDeposit(ctx, test.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, test.TestCoinsNil, deposit.Coins)

	account = accountKeeper.GetAccount(ctx, test.TestAddress2)
	require.Equal(t, test.TestCoinsPos, account.GetCoins())
}

func Test_handleUpdateSessionInfo(t *testing.T) {
	ctx, _, _, vpnKeeper, _, _ := test.TestCreateInput()

	node := test.TestNodeValid
	subscription := test.TestSubscriptionValid
	node.Status = StatusActive
	subscription.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	vpnKeeper.SetSubscription(ctx, subscription)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	node, _ = vpnKeeper.GetNode(ctx, subscription.ID)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgUpdateSessionInfo(node.Owner, subscription.ID, test.TestBandwidthPos1, test.TestNodeOwnerstdSignaturePos1, test.TestClientstdSignaturePos1)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	session, found := vpnKeeper.GetSession(ctx, test.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, test.TestBandwidthPos1, session.Bandwidth)
}

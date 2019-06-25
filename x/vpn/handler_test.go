package vpn

import (
	"testing"

	hub "github.com/sentinel-official/hub/types"

	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func Test_handleRegisterNode(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	count := vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	count = vpnKeeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

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

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	vpnKeeper.SetNodesCount(ctx, DefaultFreeNodesCount)
	vpnKeeper.SetNodesCountOfAddress(ctx, types.TestAddress1, DefaultFreeNodesCount)
	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(5), count)

	count = vpnKeeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(5), count)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, coins)

	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(6), count)

	count = vpnKeeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(6), count)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsNil, coins)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(6), count)

	id := hub.NewIDFromUInt64(count - 1)
	node, found = vpnKeeper.GetNode(ctx, id)
	require.Equal(t, true, found)
	require.Equal(t, id, node.ID)

	coins, _, err = bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos.Add(types.TestCoinsPos))
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), coins)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), coins)

	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB,
		node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(7), count)

	count = vpnKeeper.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(7), count)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, coins)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)

	count = vpnKeeper.GetNodesCount(ctx)
	require.Equal(t, uint64(7), count)

	id = hub.NewIDFromUInt64(count - 1)
	node, found = vpnKeeper.GetNode(ctx, id)
	require.Equal(t, true, found)
	require.Equal(t, id, node.ID)
}

func Test_handleUpdateNodeInfo(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	node, found := vpnKeeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestNodeEmpty, node)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node = types.TestNodeValid
	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateNodeInfo(types.TestAddress2, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNewNodeType, types.TestNewVersion,
		types.TestNewMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestNewEncryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestNewNodeType, node.Type)
	require.Equal(t, types.TestNewVersion, node.Version)
	require.Equal(t, types.TestNewMonikerValid, node.Moniker)
	require.Equal(t, types.TestCoinsPos, node.PricesPerGB)
	require.Equal(t, types.TestNewEncryption, node.Encryption)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, types.TestNodeType, types.TestVersion,
		types.TestMonikerValid, types.TestCoinsPos, types.TestBandwidthPos1, types.TestEncryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestNodeType, node.Type)
	require.Equal(t, types.TestVersion, node.Version)
	require.Equal(t, types.TestMonikerValid, node.Moniker)
	require.Equal(t, types.TestCoinsPos, node.PricesPerGB)
	require.Equal(t, types.TestEncryption, node.Encryption)
}

func Test_handleUpdateNodeStatus(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	node, found := vpnKeeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestNodeEmpty, node)

	handler := NewHandler(vpnKeeper)
	node = types.TestNodeValid
	msg := NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusInactive)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusRegistered)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusInactive)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeStatus(types.TestAddress2, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeStatus(types.TestAddress2, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateNodeStatus(node.Owner, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)

	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgUpdateNodeStatus(types.TestAddress2, node.ID, StatusActive)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

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
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	node, found := vpnKeeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestNodeEmpty, node)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgDeregisterNode(node.Owner, node.ID)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node = types.TestNodeValid
	node.Deposit = types.TestCoinZero

	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusRegistered, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins := bankKeeper.GetCoins(ctx, node.Owner)
	require.Equal(t, types.TestCoinsEmpty, coins)

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusActive, node.Status)

	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bankKeeper.GetCoins(ctx, node.Owner)
	require.Equal(t, types.TestCoinsEmpty, coins)

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	coins, _, err := bankKeeper.AddCoins(ctx, node.Owner, types.TestCoinsPos.Add(types.TestCoinsPos))
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), coins)

	_, err = vpnKeeper.AddDeposit(ctx, node.Owner, types.TestCoinPos.Add(types.TestCoinPos))
	require.Nil(t, err)

	node.Status = StatusInactive
	node.Deposit = types.TestCoinPos.Add(types.TestCoinPos).Add(types.TestCoinPos)
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, node.Status)

	deposit, found := depositKeeper.GetDeposit(ctx, node.Owner)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, node.Owner)
	require.Equal(t, types.TestCoinsNil, coins)

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, node.Status)

	node.Status = StatusInactive
	node.Deposit = types.TestCoinPos
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, node.Status)

	deposit, found = depositKeeper.GetDeposit(ctx, node.Owner)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, node.Owner)
	require.Equal(t, types.TestCoinsPos, coins)

	node, found = vpnKeeper.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)
}

func Test_handleStartSubscription(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	node, found := vpnKeeper.GetNode(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestNodeEmpty, node)

	subscription, found := vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgStartSubscription(types.TestAddress2, types.TestIDPos, types.TestCoinPos)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node = types.TestNodeValid
	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinInvalid)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	node.Status = StatusInactive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusDeRegistered
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	coins, _, err = bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsInvalid)
	require.Nil(t, err)

	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinInvalid)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsInvalid, deposit.Coins)

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsNil, coins)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsInvalid.AmountOf(types.TestCoinInvalid.Denom), deposit.Coins.AmountOf(types.TestCoinInvalid.Denom))

	subscription, found = vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, types.TestAddress2, subscription.Client)
	require.Equal(t, types.TestIDZero, subscription.NodeID)

	count := vpnKeeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	id, found := vpnKeeper.GetSubscriptionIDByAddress(ctx, types.TestAddress2, 0)
	require.Equal(t, true, found)
	require.Equal(t, id, subscription.ID)

	id, found = vpnKeeper.GetSubscriptionIDByNodeID(ctx, node.ID, 0)
	require.Equal(t, true, found)
	require.Equal(t, id, subscription.ID)

	count = vpnKeeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(1), count)

	subscriptions := vpnKeeper.GetSubscriptionsOfNode(ctx, node.ID)
	require.Equal(t, types.TestSubscriptionsValid, subscriptions)

	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	coins, _, err = bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos.Add(types.TestCoinsPos))
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), coins)

	node.Status = StatusActive
	vpnKeeper.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)

	deposit, found = depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos).Add(types.TestCoinsInvalid), deposit.Coins)

	count = vpnKeeper.GetSubscriptionsCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)

	subscriptionID := hub.NewIDFromUInt64(count - 1)
	subscription, found = vpnKeeper.GetSubscription(ctx, subscriptionID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestAddress2, subscription.Client)
	require.Equal(t, types.TestIDZero, subscription.NodeID)

	count = vpnKeeper.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)

	id, found = vpnKeeper.GetSubscriptionIDByAddress(ctx, types.TestAddress2, 1)
	require.Equal(t, true, found)
	require.Equal(t, id, subscriptionID)

	id, found = vpnKeeper.GetSubscriptionIDByNodeID(ctx, node.ID, 1)
	require.Equal(t, true, found)
	require.Equal(t, id, subscriptionID)

	subscriptions = vpnKeeper.GetSubscriptionsOfNode(ctx, node.ID)
	require.Len(t, subscriptions, 2)
	require.Equal(t, types.TestSubscriptionValid, subscriptions[0])
	require.Equal(t, subscription, subscriptions[1])
}

func Test_handleEndSubscription(t *testing.T) {
	ctx, depositKeeper, vpnKeeper, bankKeeper := keeper.TestCreateInput()

	subscription, found := vpnKeeper.GetSubscription(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSubscriptionEmpty, subscription)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgEndSubscription(types.TestAddress1, subscription.ID)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription = types.TestSubscriptionValid
	subscription.Status = StatusInactive
	vpnKeeper.SetSubscription(ctx, subscription)

	msg = NewMsgEndSubscription(types.TestAddress1, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)
	msg = NewMsgEndSubscription(types.TestAddress1, types.TestIDZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	coins, _, err := bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	_, err = vpnKeeper.AddDeposit(ctx, types.TestAddress2, types.TestCoinPos)
	require.Nil(t, err)

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsNil, coins)

	deposit, found := depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	subscription.Status = StatusInactive
	vpnKeeper.SetSubscription(ctx, subscription)

	msg = NewMsgEndSubscription(types.TestAddress1, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, types.TestIDZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = vpnKeeper.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)
	msg = NewMsgEndSubscription(types.TestAddress1, types.TestIDZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

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

	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)

	_, err = vpnKeeper.AddDeposit(ctx, types.TestAddress2, types.TestCoinPos)
	require.Nil(t, err)

	coins, _, err = bankKeeper.AddCoins(ctx, types.TestAddress2, types.TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, types.TestCoinsPos, coins)

	vpnKeeper.SetSubscription(ctx, types.TestSubscriptionValid)
	vpnKeeper.SetSession(ctx, types.TestSessionValid)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	vpnKeeper.SetSubscription(ctx, types.TestSubscriptionValid)
	vpnKeeper.SetSession(ctx, types.TestSessionValid)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	vpnKeeper.SetSessionIDBySubscriptionID(ctx, subscription.ID, 1, types.TestIDZero)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())
}

func Test_handleUpdateSessionInfo(t *testing.T) {
	ctx, _, vpnKeeper, _ := keeper.TestCreateInput()

	session, found := vpnKeeper.GetSession(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSessionEmpty, session)

	handler := NewHandler(vpnKeeper)
	msg := NewMsgUpdateSessionInfo(types.TestAddress2, types.TestIDPos, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription := types.TestSubscriptionValid
	subscription.Status = StatusInactive
	vpnKeeper.SetSubscription(ctx, subscription)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 0)

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, types.TestIDZero)
	require.Equal(t, false, found)
	require.Equal(t, types.TestSessionEmpty, session)

	count := vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(0), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node := types.TestNodeValid
	vpnKeeper.SetNode(ctx, node)
	subscription.Status = StatusActive
	vpnKeeper.SetSubscription(ctx, subscription)
	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 0)
	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1,
		types.TestClientStdSignaturePos1, types.TestNodeOwnerStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1,
		types.TestClientStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	vpnKeeper.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos2,
		types.TestNodeOwnerStdSignaturePos2, types.TestClientStdSignaturePos2)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos2, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos2)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos2,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthNeg,
		types.TestNodeOwnerStdSignatureNeg, types.TestClientStdSignatureNeg)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthZero,
		types.TestNodeOwnerStdSignatureZero, types.TestClientStdSignatureZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSessionValid, session)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1,
		types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	session, found = vpnKeeper.GetSession(ctx, types.TestIDZero)
	require.Equal(t, true, found)
	require.Equal(t, types.TestBandwidthPos1, session.Bandwidth)

	count = vpnKeeper.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = vpnKeeper.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)
}

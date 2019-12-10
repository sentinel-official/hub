package vpn

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func Test_handleRegisterNode(t *testing.T) {
	ctx, k, dk, bk := keeper.CreateTestInput(t, false)

	count := k.GetNodesCount(ctx)
	require.Equal(t, uint64(0), count)

	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(0), count)

	handler := NewHandler(k)
	node := types.TestNode

	msg := NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB, node.InternetSpeed, node.Encryption)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found := k.GetNode(ctx, hub.NewNodeID(0))
	require.Equal(t, true, found)
	require.Equal(t, hub.NewNodeID(0), node.ID)
	require.Equal(t, "moniker", node.Moniker)

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(1), count)

	deposit, found := dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	k.SetNodesCount(ctx, DefaultFreeNodesCount)
	k.SetNodesCountOfAddress(ctx, types.TestAddress1, DefaultFreeNodesCount)
	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB, node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(5), count)

	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(5), count)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	coins, err := bk.AddCoins(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB, node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(6), count)

	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(6), count)

	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins(nil), coins)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(6), count)

	id := hub.NewNodeID(count - 1)
	node, found = k.GetNode(ctx, id)
	require.Equal(t, true, found)
	require.Equal(t, id, node.ID)

	coins, err = bk.AddCoins(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}))
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), coins)

	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), coins)

	msg = NewMsgRegisterNode(node.Owner, node.Type, node.Version, node.Moniker, node.PricesPerGB, node.InternetSpeed, node.Encryption)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(7), count)

	count = k.GetNodesCountOfAddress(ctx, types.TestAddress1)
	require.Equal(t, uint64(7), count)

	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), deposit.Coins)

	count = k.GetNodesCount(ctx)
	require.Equal(t, uint64(7), count)

	id = hub.NewNodeID(count - 1)
	node, found = k.GetNode(ctx, id)
	require.Equal(t, true, found)
	require.Equal(t, id, node.ID)
}

func Test_handleUpdateNodeInfo(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)

	node, found := k.GetNode(ctx, hub.NewNodeID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Node{}, node)

	handler := NewHandler(k)

	node = types.TestNode
	node.Status = StatusDeRegistered
	k.SetNode(ctx, node)
	msg := NewMsgUpdateNodeInfo(node.Owner, node.ID, "new_node_type", "new_version", "new_moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, types.TestBandwidthPos1, "new_encryption")
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateNodeInfo(types.TestAddress2, node.ID, "new_node_type", "new_version", "new_moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, types.TestBandwidthPos1, "new_encryption")
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node.Status = StatusInactive
	k.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, "new_node_type", "new_version", "new_moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, types.TestBandwidthPos1, "new_encryption")
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, "new_node_type", node.Type)
	require.Equal(t, "new_version", node.Version)
	require.Equal(t, "new_moniker", node.Moniker)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, node.PricesPerGB)
	require.Equal(t, "new_encryption", node.Encryption)

	node.Status = StatusRegistered
	k.SetNode(ctx, node)
	msg = NewMsgUpdateNodeInfo(node.Owner, node.ID, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, types.TestBandwidthPos1, "encryption")
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, "node_type", node.Type)
	require.Equal(t, "version", node.Version)
	require.Equal(t, "moniker", node.Moniker)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, node.PricesPerGB)
	require.Equal(t, "encryption", node.Encryption)
}

func Test_handleDeregisterNode(t *testing.T) {
	ctx, k, dk, bk := keeper.CreateTestInput(t, false)

	node, found := k.GetNode(ctx, hub.NewNodeID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Node{}, node)

	handler := NewHandler(k)

	node = types.TestNode
	node.Deposit = sdk.NewInt64Coin("stake", 0)

	node.Status = StatusDeRegistered
	k.SetNode(ctx, node)
	msg := NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	k.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusRegistered, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins := bk.GetCoins(ctx, node.Owner)
	require.Equal(t, sdk.Coins{}, coins)

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	k.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusRegistered, node.Status)

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusRegistered, node.Status)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bk.GetCoins(ctx, node.Owner)
	require.Equal(t, sdk.Coins{}, coins)

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	coins, err := bk.AddCoins(ctx, node.Owner, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}))
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), coins)

	err = k.AddDeposit(ctx, node.Owner, sdk.NewInt64Coin("stake", 100).Add(sdk.NewInt64Coin("stake", 100)))
	require.Nil(t, err)

	node.Status = StatusDeRegistered
	node.Deposit = sdk.NewInt64Coin("stake", 100).Add(sdk.NewInt64Coin("stake", 100)).Add(sdk.NewInt64Coin("stake", 100))
	k.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	deposit, found := dk.GetDeposit(ctx, node.Owner)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), deposit.Coins)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), deposit.Coins)

	coins = bk.GetCoins(ctx, node.Owner)
	require.Equal(t, sdk.Coins(nil), coins)

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)

	node.Status = StatusRegistered
	node.Deposit = sdk.NewInt64Coin("stake", 100)
	k.SetNode(ctx, node)
	msg = NewMsgDeregisterNode(types.TestAddress2, node.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusRegistered, node.Status)

	deposit, found = dk.GetDeposit(ctx, node.Owner)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), deposit.Coins)

	msg = NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	coins = bk.GetCoins(ctx, node.Owner)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	node, found = k.GetNode(ctx, node.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusDeRegistered, node.Status)
}

func Test_handleStartSubscription(t *testing.T) {
	ctx, k, dk, bk := keeper.CreateTestInput(t, false)

	node, found := k.GetNode(ctx, hub.NewNodeID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Node{}, node)

	subscription, found := k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Subscription{}, subscription)

	handler := NewHandler(k)
	msg := NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100))
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	node = types.TestNode
	node.Status = StatusDeRegistered
	k.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, node.ID, sdk.NewInt64Coin("stake", 100))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found := dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	subscription, found = k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Subscription{}, subscription)

	node.Status = StatusRegistered
	k.SetNode(ctx, node)
	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, node.ID, sdk.NewInt64Coin("stake", 100))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	subscription, found = k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Subscription{}, subscription)

	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress2, node.ID, sdk.NewInt64Coin("invalid", 100))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	subscription, found = k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Subscription{}, subscription)

	coins, err := bk.AddCoins(ctx, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, node.ID, sdk.NewInt64Coin("stake", 100))
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.NewInt(100), deposit.Coins.AmountOf("stake"))

	subscription, found = k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, true, found)
	require.Equal(t, types.TestAddress2, subscription.Client)
	require.Equal(t, hub.NewNodeID(0), subscription.NodeID)

	count := k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(1), count)

	id, found := k.GetSubscriptionIDByAddress(ctx, types.TestAddress2, 0)
	require.Equal(t, true, found)
	require.Equal(t, id, subscription.ID)

	id, found = k.GetSubscriptionIDByNodeID(ctx, node.ID, 0)
	require.Equal(t, true, found)
	require.Equal(t, id, subscription.ID)

	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(1), count)

	subscriptions := k.GetSubscriptionsOfNode(ctx, node.ID)
	require.Equal(t, []types.Subscription{types.TestSubscription}, subscriptions)

	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, node.ID, sdk.NewInt64Coin("stake", 100))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	coins, err = bk.AddCoins(ctx, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}))
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 100)}), coins)

	msg = NewMsgStartSubscription(types.TestAddress2, types.TestAddress1, node.ID, sdk.NewInt64Coin("stake", 100))
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 200)}, deposit.Coins)

	count = k.GetSubscriptionsCountOfAddress(ctx, types.TestAddress2)
	require.Equal(t, uint64(2), count)

	subscriptionID := hub.NewSubscriptionID(count - 1)
	subscription, found = k.GetSubscription(ctx, subscriptionID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestAddress2, subscription.Client)
	require.Equal(t, hub.NewNodeID(0), subscription.NodeID)

	count = k.GetSubscriptionsCount(ctx)
	require.Equal(t, uint64(2), count)

	id, found = k.GetSubscriptionIDByAddress(ctx, types.TestAddress2, 1)
	require.Equal(t, true, found)
	require.Equal(t, id, subscriptionID)

	id, found = k.GetSubscriptionIDByNodeID(ctx, node.ID, 1)
	require.Equal(t, true, found)
	require.Equal(t, id, subscriptionID)

	subscriptions = k.GetSubscriptionsOfNode(ctx, node.ID)
	require.Len(t, subscriptions, 2)
	require.Equal(t, types.TestSubscription, subscriptions[0])
	require.Equal(t, subscription, subscriptions[1])
}

func Test_handleEndSubscription(t *testing.T) {
	ctx, k, dk, bk := keeper.CreateTestInput(t, false)

	subscription, found := k.GetSubscription(ctx, hub.NewSubscriptionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Subscription{}, subscription)

	handler := NewHandler(k)

	subscription = types.TestSubscription
	subscription.Status = StatusInactive
	k.SetSubscription(ctx, subscription)

	msg := NewMsgEndSubscription(types.TestAddress1, subscription.ID)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	subscription.Status = StatusActive
	k.SetSubscription(ctx, subscription)
	msg = NewMsgEndSubscription(types.TestAddress1, hub.NewSubscriptionID(0))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	coins, err := bk.AddCoins(ctx, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	err = k.AddDeposit(ctx, types.TestAddress2, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)

	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Nil(t, err)
	require.Equal(t, sdk.Coins(nil), coins)

	deposit, found := dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	subscription.Status = StatusInactive
	k.SetSubscription(ctx, subscription)

	msg = NewMsgEndSubscription(types.TestAddress1, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, hub.NewSubscriptionID(0))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	subscription.Status = StatusActive
	k.SetSubscription(ctx, subscription)
	msg = NewMsgEndSubscription(types.TestAddress1, hub.NewSubscriptionID(0))
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	subscription, found = k.GetSubscription(ctx, subscription.ID)
	require.Equal(t, true, found)
	require.Equal(t, StatusInactive, subscription.Status)

	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	err = k.AddDeposit(ctx, types.TestAddress2, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)

	coins, err = bk.AddCoins(ctx, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSession(ctx, types.TestSession)
	k.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	k.SetSubscription(ctx, types.TestSubscription)
	k.SetSession(ctx, types.TestSession)
	k.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	k.SetSessionIDBySubscriptionID(ctx, subscription.ID, 1, hub.NewSessionID(0))

	msg = NewMsgEndSubscription(types.TestAddress2, subscription.ID)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())
}

func Test_handleUpdateSessionInfo(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)

	session, found := k.GetSession(ctx, hub.NewSessionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Session{}, session)

	handler := NewHandler(k)
	msg := NewMsgUpdateSessionInfo(types.TestAddress2, hub.NewSubscriptionID(1), types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res := handler(ctx, *msg)
	require.False(t, res.IsOK())

	subscription := types.TestSubscription
	subscription.Status = StatusInactive
	k.SetSubscription(ctx, subscription)
	k.SetSessionsCountOfSubscription(ctx, subscription.ID, 0)

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, hub.NewSessionID(0))
	require.Equal(t, false, found)
	require.Equal(t, types.Session{}, session)

	count := k.GetSessionsCount(ctx)
	require.Equal(t, uint64(0), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(0), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	node := types.TestNode
	k.SetNode(ctx, node)
	subscription.Status = StatusActive
	k.SetSubscription(ctx, subscription)
	k.SetSessionsCountOfSubscription(ctx, subscription.ID, 0)
	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1, types.TestClientStdSignaturePos1, types.TestNodeOwnerStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1, types.TestClientStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	k.SetSessionsCountOfSubscription(ctx, subscription.ID, 1)
	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos2, types.TestNodeOwnerStdSignaturePos2, types.TestClientStdSignaturePos2)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	msg = NewMsgUpdateSessionInfo(types.TestAddress2, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	id, _ := k.GetSessionIDBySubscriptionID(ctx, subscription.ID, k.GetSessionsCountOfSubscription(ctx, subscription.ID))
	session, found = k.GetSession(ctx, id)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)
	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos2, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos2)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthPos2, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(subscription.Client, subscription.ID, types.TestBandwidthNeg, types.TestNodeOwnerStdSignatureNeg, types.TestClientStdSignatureNeg)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthZero, types.TestNodeOwnerStdSignatureZero, types.TestClientStdSignatureZero)
	res = handler(ctx, *msg)
	require.False(t, res.IsOK())

	session, found = k.GetSession(ctx, session.ID)
	require.Equal(t, true, found)
	require.Equal(t, types.TestSession, session)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)

	msg = NewMsgUpdateSessionInfo(node.Owner, subscription.ID, types.TestBandwidthPos1, types.TestNodeOwnerStdSignaturePos1, types.TestClientStdSignaturePos1)
	res = handler(ctx, *msg)
	require.True(t, res.IsOK())

	session, found = k.GetSession(ctx, hub.NewSessionID(0))
	require.Equal(t, true, found)
	require.Equal(t, types.TestBandwidthPos1, session.Bandwidth)

	count = k.GetSessionsCount(ctx)
	require.Equal(t, uint64(1), count)

	count = k.GetSessionsCountOfSubscription(ctx, subscription.ID)
	require.Equal(t, uint64(1), count)
}

func Test_handleResolverNode(t *testing.T) {
	ctx, k, _, _ := keeper.CreateTestInput(t, false)
	handler := NewHandler(k)

	resolver := types.TestResolver

	data, found := k.GetResolver(ctx, resolver.Owner)
	require.False(t, found)

	msg := NewMsgRegisterResolver(resolver.Owner, resolver.Commission)
	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	data, found = k.GetResolver(ctx, resolver.Owner)
	require.True(t, found)
	require.Equal(t, data, resolver)

	msg = NewMsgRegisterResolver(resolver.Owner, resolver.Commission)
	res = handler(ctx, msg)
	require.False(t, res.IsOK())
	require.Equal(t, res.Log, types.ErrorResolverAlreadyExist().ABCILog())

	updateResolverInfoMsg := NewMsgUpdateResolverInfo(types.TestAddress2, sdk.NewDecWithPrec(2, 1))
	res = handler(ctx, updateResolverInfoMsg)
	require.False(t, res.IsOK())

	resolver.Status = StatusDeRegistered
	k.SetResolver(ctx, resolver)
	updateResolverInfoMsg = NewMsgUpdateResolverInfo(types.TestAddress1, sdk.NewDecWithPrec(2, 1))
	res = handler(ctx, updateResolverInfoMsg)
	require.False(t, res.IsOK())

	resolver.Status = StatusRegistered
	k.SetResolver(ctx, resolver)
	updateResolverInfoMsg = NewMsgUpdateResolverInfo(types.TestAddress1, sdk.NewDecWithPrec(2, 1))
	res = handler(ctx, updateResolverInfoMsg)
	require.True(t, res.IsOK())

	resolver, found = k.GetResolver(ctx, types.TestAddress1)
	require.True(t, found)
	require.Equal(t, sdk.NewDecWithPrec(2, 1), resolver.Commission)

	deRegisterResolverMsg := NewMsgDeregisterResolver(types.TestAddress2)
	res = handler(ctx, deRegisterResolverMsg)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorResolverDoesNotExist().ABCILog(), res.Log)

	resolver.Status = StatusDeRegistered
	k.SetResolver(ctx, resolver)
	deRegisterResolverMsg = NewMsgDeregisterResolver(types.TestAddress1)
	res = handler(ctx, deRegisterResolverMsg)
	require.False(t, res.IsOK())

	resolver.Status = StatusRegistered
	k.SetResolver(ctx, resolver)
	deRegisterResolverMsg = NewMsgDeregisterResolver(types.TestAddress1)
	res = handler(ctx, deRegisterResolverMsg)
	require.True(t, res.IsOK())
}

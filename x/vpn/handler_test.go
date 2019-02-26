package vpn

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func Test_handleRegisterNode(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, node.Owner)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, node.ID)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)
}

func Test_handleUpdateNodeDetails(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)

	msgUpdateNodeDetails := types.NewMsgUpdateNodeDetails(node.Owner, types.TestNodeIDInvalid,
		csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_1", 1)},
		csdkTypes.NewInt(100), csdkTypes.NewInt(100), uint16(8080), "", "")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorNodeNotExists().Code(), res.Code)

	msgUpdateNodeDetails = types.NewMsgUpdateNodeDetails(csdkTypes.AccAddress([]byte("invalid_address")), types.TestNodeIDValid,
		csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_1", 1)},
		csdkTypes.NewInt(100), csdkTypes.NewInt(100), uint16(8080), "", "")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorUnauthorized().Code(), res.Code)

	msgUpdateNodeDetails = types.NewMsgUpdateNodeDetails(node.Owner, types.TestNodeIDValid,
		csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_1", 1)},
		csdkTypes.NewInt(100), csdkTypes.NewInt(100), uint16(8080), "", "")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[0].Value)))

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, csdkTypes.Coins{csdkTypes.NewInt64Coin("coin_1", 1)}, nodeRes.PricesPerGB)
	require.Equal(t, sdkTypes.NewBandwidthFromInt64(100, 100), nodeRes.NetSpeed)
	require.Equal(t, uint16(8080), nodeRes.APIPort)
	require.Equal(t, types.TestEncryption, nodeRes.Encryption)
	require.Equal(t, types.TestVersion, nodeRes.Version)
}

func Test_handleUpdateNodeStatus(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)

	msgUpdateNodeStatus := types.NewMsgUpdateNodeStatus(node.Owner, types.TestNodeIDInvalid, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorNodeNotExists().Code(), res.Code)

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(csdkTypes.AccAddress([]byte("invalid_address")), node.ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorUnauthorized().Code(), res.Code)

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(node.Owner, node.ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[0].Value)))

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, types.StatusActive, nodeRes.Status)
	require.Equal(t, int64(0), nodeRes.StatusAtHeight)

	nodes, err := vpnKeeper.GetActiveNodeIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{node.ID}, nodes)
}

func Test_handleDeregisterNode(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)

	msgDeregister := types.NewMsgDeregisterNode(node.Owner, types.TestNodeIDInvalid)
	res = handler(ctx, *msgDeregister)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorNodeNotExists().Code(), res.Code)

	msgDeregister = types.NewMsgDeregisterNode(csdkTypes.AccAddress([]byte("invalid_address")), node.ID)
	res = handler(ctx, *msgDeregister)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorUnauthorized().Code(), res.Code)

	msgDeregister = types.NewMsgDeregisterNode(node.Owner, node.ID)
	res = handler(ctx, *msgDeregister)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[0].Value)))

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, types.StatusDeregistered, nodeRes.Status)
	require.Equal(t, int64(0), nodeRes.StatusAtHeight)

	account = accountKeeper.GetAccount(ctx, node.Owner)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())
}

func Test_handleInitSession(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)

	account = accountKeeper.NewAccountWithAddress(ctx, types.TestAddress2)
	require.Nil(t, account.SetPubKey(types.TestPubkey2))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Equal(t, types.TestPubkey2, account.GetPubKey())

	msgInitSession := types.NewMsgInitSession(types.TestAddress2, types.TestNodeIDInvalid, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorNodeNotExists().Code(), res.Code)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorInvalidNodeStatus().Code(), res.Code)

	msgUpdateNodeStatus := types.NewMsgUpdateNodeStatus(types.TestAddress1, node.ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinNil)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorInvalidPriceDenom().Code(), res.Code)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, csdkTypes.CodeType(10), res.Code)

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Nil(t, account.SetPubKey(types.TestPubkey2))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestSessionIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{types.TestCoinPos}), account.GetCoins())

	session, err := vpnKeeper.GetSessionDetails(ctx, types.TestSessionIDValid)
	require.Nil(t, err)
	require.Equal(t, types.TestSessionIDValid, session.ID)
}

func Test_handleUpdateSessionBandwidth(t *testing.T) {
	ctx, _, vpnKeeper, accountKeeper, bankKeeper := keeper.TestCreateInput()
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	node := keeper.TestNodeValid

	msg := types.NewMsgRegisterNode(node.Owner, node.LockedAmount, node.PricesPerGB, node.NetSpeed.Upload,
		node.NetSpeed.Download, node.APIPort, node.Encryption, node.NodeType, node.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestNodeIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{node.LockedAmount}), account.GetCoins())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, types.TestNodeIDValid)
	require.Nil(t, err)
	require.Equal(t, &node, nodeRes)

	account = accountKeeper.NewAccountWithAddress(ctx, types.TestAddress2)
	require.Nil(t, account.SetPubKey(types.TestPubkey2))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Equal(t, types.TestPubkey2, account.GetPubKey())

	msgInitSession := types.NewMsgInitSession(types.TestAddress2, types.TestNodeIDInvalid, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorNodeNotExists().Code(), res.Code)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorInvalidNodeStatus().Code(), res.Code)

	msgUpdateNodeStatus := types.NewMsgUpdateNodeStatus(types.TestAddress1, node.ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinNil)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorInvalidPriceDenom().Code(), res.Code)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.False(t, res.IsOK())
	require.Equal(t, csdkTypes.CodeType(10), res.Code)

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Nil(t, account.SetPubKey(types.TestPubkey2))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	msgInitSession = types.NewMsgInitSession(types.TestAddress2, node.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitSession)
	require.True(t, res.IsOK())
	require.Equal(t, types.TestSessionIDValid, sdkTypes.NewID(string(res.Tags[1].Value)))

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos.Minus(csdkTypes.Coins{types.TestCoinPos}), account.GetCoins())

	session, err := vpnKeeper.GetSessionDetails(ctx, types.TestSessionIDValid)
	require.Nil(t, err)
	require.Equal(t, types.TestSessionIDValid, session.ID)

	signDataBytes := sdkTypes.NewBandwidthSignData(session.ID, types.TestBandwidthPos, node.Owner, session.Client).GetBytes()
	nodeOwnerSign, err1 := types.TestPrivKey1.Sign(signDataBytes)
	require.Nil(t, err1)
	clientSign, err1 := types.TestPrivKey2.Sign(signDataBytes)
	require.Nil(t, err1)

	msgUpdateSessionBandwidth := types.NewMsgUpdateSessionBandwidth(session.Client, types.TestSessionIDInvalid,
		types.TestUploadPos, types.TestDownloadPos, clientSign, nodeOwnerSign)
	res = handler(ctx, *msgUpdateSessionBandwidth)
	require.False(t, res.IsOK())
	require.Equal(t, types.ErrorSessionNotExists().Code(), res.Code)

	sessionRes, err := vpnKeeper.GetSessionDetails(ctx, session.ID)
	require.Nil(t, err)
	require.Equal(t, types.TestBandwidthZero, sessionRes.Bandwidth.Consumed)

	msgUpdateSessionBandwidth = types.NewMsgUpdateSessionBandwidth(session.Client, session.ID,
		types.TestUploadPos, types.TestDownloadPos, clientSign, nodeOwnerSign)
	res = handler(ctx, *msgUpdateSessionBandwidth)
	require.True(t, res.IsOK())

	sessionRes, err = vpnKeeper.GetSessionDetails(ctx, session.ID)
	require.Nil(t, err)
	require.Equal(t, types.TestBandwidthPos, sessionRes.Bandwidth.Consumed)
	require.Equal(t, clientSign, sessionRes.Bandwidth.ClientSign)
	require.Equal(t, nodeOwnerSign, sessionRes.Bandwidth.NodeOwnerSign)

	sessions, err := vpnKeeper.GetActiveSessionIDsAtHeight(ctx, 0)
	require.Nil(t, err)
	require.Equal(t, sdkTypes.IDs{types.TestSessionIDValid}, sessions)
}

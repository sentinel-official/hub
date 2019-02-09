package vpn

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestHandler_RegisterNode(t *testing.T) {
	ctx, vpnKeeper, accountKeeper := keeper.TestCreateInput()
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	cdc := keeper.TestMakeCodec()
	RegisterCodec(cdc)

	nodeDetails := keeper.TestNodeValid

	account := accountKeeper.NewAccountWithAddress(ctx, nodeDetails.Owner)
	err := account.SetCoins(types.TestCoinsPos)
	require.Nil(t, err)
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, nodeDetails.Owner)
	coins := account.GetCoins()
	require.Equal(t, csdkTypes.Coins{csdkTypes.NewInt64Coin("sent", 10), csdkTypes.Coin{"sut", csdkTypes.NewInt(100)}}, coins)

	msg := types.NewMsgRegisterNode(nodeDetails.Owner, nodeDetails.LockedAmount, nodeDetails.PricesPerGB, nodeDetails.NetSpeed.Upload,
		nodeDetails.NetSpeed.Download, nodeDetails.APIPort, nodeDetails.EncMethod, nodeDetails.NodeType, nodeDetails.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, NewNodeID(string(res.Tags.ToKVPairs()[1].Value)))
	require.Nil(t, err)
	require.Equal(t, cdc.MustMarshalJSON(nodeDetails.NetSpeed), cdc.MustMarshalJSON(nodeRes.NetSpeed))

	msg1 := types.NewMsgRegisterNode(nodeDetails.Owner, types.TestCoinNeg, nodeDetails.PricesPerGB, nodeDetails.NetSpeed.Upload,
		nodeDetails.NetSpeed.Download, nodeDetails.APIPort, nodeDetails.EncMethod, nodeDetails.NodeType, nodeDetails.Version)
	res = handler(ctx, *msg1)
	require.True(t, res.IsOK())
}

func TestHandler_UpdateNodeStatusAndNodeDetails(t *testing.T) {
	ctx, vpnKeeper, accountKeeper := keeper.TestCreateInput()
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	cdc := keeper.TestMakeCodec()
	RegisterCodec(cdc)

	nodeDetails := keeper.TestNodeValid

	account := accountKeeper.NewAccountWithAddress(ctx, nodeDetails.Owner)
	err := account.SetCoins(types.TestCoinsPos)
	require.Nil(t, err)
	accountKeeper.SetAccount(ctx, account)

	msg := types.NewMsgRegisterNode(nodeDetails.Owner, nodeDetails.LockedAmount, nodeDetails.PricesPerGB, nodeDetails.NetSpeed.Upload,
		nodeDetails.NetSpeed.Download, nodeDetails.APIPort, nodeDetails.EncMethod, nodeDetails.NodeType, nodeDetails.Version)
	res := handler(ctx, *msg)
	require.True(t, res.IsOK())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, NewNodeID(string(res.Tags.ToKVPairs()[1].Value)))
	require.Nil(t, err)

	msgUpdateNodeDetails := types.NewMsgUpdateNodeDetails(nodeDetails.Owner, NewNodeID(string(res.Tags.ToKVPairs()[1].Value)),
		types.TestCoinsPos, types.TestUploadPos, types.TestDownloadPos, types.TestAPIPortValid, "", "0.01")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.True(t, res.IsOK())
	nodeID := NewNodeID(string(res.Tags.ToKVPairs()[0].Value))

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, nodeID)
	require.Nil(t, err)
	require.Equal(t, csdkTypes.NewInt(100), nodeRes.PricesPerGB.AmountOf("sut"))
	require.Equal(t, "enc_method", nodeRes.EncMethod)

	msgUpdateNodeDetails = types.NewMsgUpdateNodeDetails(nodeDetails.Owner, types.TestNodeIDInvalid, types.TestCoinsPos,
		types.TestUploadPos, types.TestDownloadPos, types.TestAPIPortValid, "", "0.01")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.False(t, res.IsOK())

	msgUpdateNodeDetails = types.NewMsgUpdateNodeDetails(types.TestAddressEmpty, nodeID, types.TestCoinsPos,
		types.TestUploadPos, types.TestDownloadPos, types.TestAPIPortValid, "", "0.01")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.Equal(t, csdkTypes.CodeType(204), res.Code)

	msgUpdateNodeStatus := types.NewMsgUpdateNodeStatus(nodeDetails.Owner, nodeID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, nodeID)
	require.Nil(t, err)
	require.Equal(t, StatusActive, nodeRes.Status)

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(nodeDetails.Owner, nodeRes.ID, types.StatusRegistered)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	msgDeregister := types.NewMsgDeregisterNode(nodeRes.Owner, nodeRes.ID)
	res = handler(ctx, *msgDeregister)
	require.Nil(t, err)

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(nodeRes.Owner, nodeRes.ID, types.StatusInactive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.False(t, res.IsOK())

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(nodeDetails.Owner, types.TestNodeIDInvalid, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.False(t, res.IsOK())

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(types.TestAddress, nodeID, types.StatusDeregistered)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.False(t, res.IsOK())

	msgDeregister = types.NewMsgDeregisterNode(nodeDetails.Owner, types.TestNodeIDInvalid)
	res = handler(ctx, *msgDeregister)
	require.False(t, res.IsOK())

	msgDeregister = types.NewMsgDeregisterNode(types.TestAddress, nodeID)
	res = handler(ctx, *msgDeregister)
	require.False(t, res.IsOK())

	nodeRes, err = vpnKeeper.GetNodeDetails(ctx, nodeID)
	require.Equal(t, StatusDeregistered, nodeRes.Status)

	msgUpdateNodeDetails = types.NewMsgUpdateNodeDetails(nodeDetails.Owner, nodeID, types.TestCoinsPos,
		types.TestUploadPos, types.TestDownloadPos, types.TestAPIPortValid, "", "0.01")
	res = handler(ctx, *msgUpdateNodeDetails)
	require.False(t, res.IsOK())

	msgDeregister = types.NewMsgDeregisterNode(nodeRes.Owner, nodeID)
	res = handler(ctx, *msgDeregister)
	require.False(t, res.IsOK())
}

func TestHandler_Session(t *testing.T) {
	ctx, vpnKeeper, accountKeeper := keeper.TestCreateInput()
	bankKeeper := bank.NewBaseKeeper(accountKeeper)
	handler := NewHandler(vpnKeeper, accountKeeper, bankKeeper)

	cdc := keeper.TestMakeCodec()
	RegisterCodec(cdc)

	nodeDetails := keeper.TestNodeValid
	sessionDetails := keeper.TestSessionValid

	vpnAccount := accountKeeper.NewAccountWithAddress(ctx, testAddress1)
	err := vpnAccount.SetCoins(types.TestCoinsPos)
	require.Nil(t, err)
	err = vpnAccount.SetPubKey(testPubkey1)
	require.Nil(t, err)
	accountKeeper.SetAccount(ctx, vpnAccount)

	clientAccount := accountKeeper.NewAccountWithAddress(ctx, testAddress2)
	err = clientAccount.SetCoins(types.TestCoinsPos)
	require.Nil(t, err)
	err = clientAccount.SetPubKey(testPubkey2)
	require.Nil(t, err)
	accountKeeper.SetAccount(ctx, clientAccount)

	nodeDetails.Owner = testAddress1
	sessionDetails.Client = testAddress2

	msgVPNRegister := types.NewMsgRegisterNode(nodeDetails.Owner, types.TestCoinPos, nodeDetails.PricesPerGB, nodeDetails.NetSpeed.Upload,
		nodeDetails.NetSpeed.Download, nodeDetails.APIPort, nodeDetails.EncMethod, nodeDetails.NodeType, nodeDetails.Version)
	res := handler(ctx, *msgVPNRegister)
	require.True(t, res.IsOK())

	nodes, err := vpnKeeper.GetNodesOfOwner(ctx, nodeDetails.Owner)
	require.Nil(t, err)

	msgUpdateNodeStatus := types.NewMsgUpdateNodeStatus(nodeDetails.Owner, nodes[0].ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	msgInitClientSession := types.NewMsgInitSession(testAddress2, nodes[0].ID, types.TestCoinPos)
	res = handler(ctx, *msgInitClientSession)
	require.True(t, res.IsOK())

	sessionRes, err := vpnKeeper.GetSessionDetails(ctx, NewSessionID(string(res.Tags[1].Value)))
	require.NotNil(t, sessionRes)

	clientAccount = accountKeeper.GetAccount(ctx, sessionDetails.Client)
	require.Equal(t, csdkTypes.Coins{csdkTypes.Coin{"sut", csdkTypes.NewInt(100)}}, clientAccount.GetCoins())

	msgInitClientSession = types.NewMsgInitSession(sessionDetails.Client, types.TestNodeIDInvalid, types.TestCoinPos)
	res = handler(ctx, *msgInitClientSession)
	require.False(t, res.IsOK())

	msgInitClientSession = types.NewMsgInitSession(sessionDetails.Client, nodes[0].ID, types.TestCoinNeg)
	res = handler(ctx, *msgInitClientSession)
	require.True(t, res.IsOK())

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(nodeDetails.Owner, nodes[0].ID, types.StatusRegistered)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	msgInitClientSession = types.NewMsgInitSession(sessionDetails.Client, nodes[0].ID, types.TestCoinPos)
	res = handler(ctx, *msgInitClientSession)
	require.False(t, res.IsOK())

	msgUpdateNodeStatus = types.NewMsgUpdateNodeStatus(nodeDetails.Owner, nodes[0].ID, types.StatusActive)
	res = handler(ctx, *msgUpdateNodeStatus)
	require.True(t, res.IsOK())

	nodeRes, err := vpnKeeper.GetNodeDetails(ctx, nodes[0].ID)
	require.Equal(t, types.StatusActive, nodeRes.Status)

	msgInitClientSession = types.NewMsgInitSession(sessionDetails.Client, nodeRes.ID, types.TestCoinPos)
	res = handler(ctx, *msgInitClientSession)
	require.True(t, res.IsOK())

	sessionDetails.ID = NewSessionID(string(res.Tags.ToKVPairs()[1].Value))

	bandwidth := sdkTypes.Bandwidth{
		Upload:   types.TestUploadPos,
		Download: types.TestDownloadPos,
	}

	signBytes, err := types.NewBandwidthSign(sessionDetails.ID, bandwidth, nodeRes.Owner, sessionDetails.Client).GetBytes()
	require.Nil(t, err)

	nodeOwnerSign, err := testPrivKey1.Sign(signBytes)
	require.Nil(t, err)

	clientSign, err := testPrivKey2.Sign(signBytes)
	require.Nil(t, err)

	msgUpdateSessionBandwidth := types.NewMsgUpdateSessionBandwidth(sessionDetails.Client, sessionDetails.ID, types.TestUploadPos, types.TestDownloadPos, clientSign, nodeOwnerSign)
	res = handler(ctx, *msgUpdateSessionBandwidth)
	require.True(t, res.IsOK())

	sessionRes1, err := vpnKeeper.GetSessionDetails(ctx, sessionDetails.ID)
	require.Nil(t, err)
	require.Equal(t, bandwidth, sessionRes1.Bandwidth.Consumed)

	msgUpdateSessionBandwidth = types.NewMsgUpdateSessionBandwidth(sessionDetails.Client, types.TestSessionIDInvalid, types.TestUploadPos, types.TestDownloadPos, clientSign, nodeOwnerSign)
	res = handler(ctx, *msgUpdateSessionBandwidth)
	require.False(t, res.IsOK())

	msgUpdateSessionBandwidth = types.NewMsgUpdateSessionBandwidth(sessionDetails.Client, sessionDetails.ID, types.TestUploadPos.AddRaw(1), types.TestDownloadPos.AddRaw(1), clientSign, nodeOwnerSign)
	res = handler(ctx, *msgUpdateSessionBandwidth)
	require.Equal(t, csdkTypes.CodeType(255), res.Code)
}

var (
	testPrivKey1 = ed25519.GenPrivKey()
	testPrivKey2 = ed25519.GenPrivKey()

	testPubkey1 = testPrivKey1.PubKey()
	testPubkey2 = testPrivKey2.PubKey()

	testAddress1 = csdkTypes.AccAddress(testPubkey1.Address())
	testAddress2 = csdkTypes.AccAddress(testPubkey2.Address())
)

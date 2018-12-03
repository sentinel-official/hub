package vpn

import (
	"encoding/json"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterNode
		want csdkTypes.Error
	}{
		{"From nil", NewMsgRegisterNode(
			nil, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyAddress(),
		},
		{"From empty", NewMsgRegisterNode(
			csdkTypes.AccAddress{}, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyAddress(),
		},
		{"API port small", NewMsgRegisterNode(
			accAddress1, 1000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidAPIPort(),
		},
		{"API port large", NewMsgRegisterNode(
			accAddress1, 65600, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidAPIPort(),
		},
		{"City empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidLocation(),
		},
		{"Country empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidLocation(),
		},
		{
			"Download speed negative", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, -1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidNetSpeed(),
		},
		{
			"Upload speed negative", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", -1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidNetSpeed(),
		},
		{
			"Encryption method empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyEncMethod(),
		},
		{
			"Price per GB negative", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", -1, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorInvalidPricePerGB(),
		},
		{
			"Version empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyVersion(),
		},
		{
			"Locker ID empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyLockerID(),
		},
		{
			"Coins nil", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			nil, pubKey1, getMsgLockCoinsSignature("locker_id", nil)),
			errorInvalidCoins(),
		},
		{
			"Coins empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{})),
			errorInvalidCoins(),
		},
		{
			"Coins invalid", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "y"), coin(100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "y"), coin(100, "x")})),
			errorInvalidCoins(),
		},
		{
			"Coins negative", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(-100, "x")}, pubKey1, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(-100, "x")})),
			errorInvalidCoins(),
		},
		{
			"Public key nil", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, nil, getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyPubKey(),
		},
		{
			"Signature empty", NewMsgRegisterNode(
			accAddress1, 3000, 0, 0, "city", "country", 1024, 1024, "enc_method", 0, "version", "locker_id",
			csdkTypes.Coins{coin(100, "x")}, pubKey1, []byte("")),
			errorEmptySignature(),
		},
		{
			"Valid", msgRegisterNode, nil,
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.want, tc.msg.ValidateBasic())
	}
}

func TestMsgRegisterNode_GetSignBytes(t *testing.T) {
	msg := msgRegisterNode
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := msgRegisterNode
	require.Equal(t, []csdkTypes.AccAddress{accAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := msgRegisterNode
	require.Equal(t, sdkTypes.KeyVPN, msg.Route())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := msgRegisterNode
	require.Equal(t, "msg_register_node", msg.Type())
}

func TestMsgPayVPNService_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgPayVPNService
		want csdkTypes.Error
	}{
		{"From nil", NewMsgPayVPNService(
			nil, "vpn_id", "locker_id", csdkTypes.Coins{coin(100, "x")}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyAddress(),
		},
		{"From empty", NewMsgPayVPNService(
			csdkTypes.AccAddress{}, "vpn_id", "locker_id", csdkTypes.Coins{coin(100, "x")}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyAddress(),
		},
		{"VPN ID empty", NewMsgPayVPNService(
			accAddress1, "", "locker_id", csdkTypes.Coins{coin(100, "x")}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyVPNID(),
		},
		{"Locker ID empty", NewMsgPayVPNService(
			accAddress1, "vpn_id", "", csdkTypes.Coins{coin(100, "x")}, pubKey1,
			getMsgLockCoinsSignature("", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyLockerID(),
		},
		{"Coins nil", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", nil, pubKey1,
			getMsgLockCoinsSignature("locker_id", nil)),
			errorInvalidCoins(),
		},
		{"Coins empty", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", csdkTypes.Coins{}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{})),
			errorInvalidCoins(),
		},
		{"Coins invalid", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", csdkTypes.Coins{coin(100, "y"), coin(100, "x")}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "y"), coin(100, "x")})),
			errorInvalidCoins(),
		},
		{"Coins negative", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", csdkTypes.Coins{coin(-100, "x")}, pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(-100, "x")})),
			errorInvalidCoins(),
		},
		{"Public key nil", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", csdkTypes.Coins{coin(100, "x")}, nil,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(100, "x")})),
			errorEmptyPubKey(),
		},
		{"Signature empty", NewMsgPayVPNService(
			accAddress1, "vpn_id", "locker_id", csdkTypes.Coins{coin(100, "x")}, pubKey1,
			[]byte("")),
			errorEmptySignature(),
		},
		{
			"Valid", msgPayVPNService, nil,
		},
	}
	for _, tc := range tests {
		require.Equal(t, tc.want, tc.msg.ValidateBasic())
	}
}

func TestMsgPayVPNService_GetSignBytes(t *testing.T) {
	msg := msgPayVPNService
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgPayVPNService_GetSigners(t *testing.T) {
	msg := msgPayVPNService
	require.Equal(t, []csdkTypes.AccAddress{accAddress1}, msg.GetSigners())
}

func TestMsgPayVPNService_Route(t *testing.T) {
	msg := msgPayVPNService
	require.Equal(t, sdkTypes.KeyVPN, msg.Route())
}

func TestMsgPayVPNService_Type(t *testing.T) {
	msg := msgPayVPNService
	require.Equal(t, "msg_pay_vpn_service", msg.Type())
}

func TestMsgUpdateSessionStatus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateSessionStatus
		want csdkTypes.Error
	}{
		{"From nil", NewMsgUpdateSessionStatus(
			nil, "session_id", sdkTypes.StatusInactive),
			errorEmptyAddress(),
		},
		{"From empty", NewMsgUpdateSessionStatus(
			csdkTypes.AccAddress{}, "session_id", sdkTypes.StatusInactive),
			errorEmptyAddress(),
		},
		{"Session ID empty", NewMsgUpdateSessionStatus(
			accAddress1, "", sdkTypes.StatusInactive),
			errorEmptySessionID(),
		},
		{"Status invalid", NewMsgUpdateSessionStatus(
			accAddress1, "session_id", "INVALID"),
			errorInvalidSessionStatus(),
		},
		{
			"Valid", msgUpdateSessionStatus, nil,
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.want, tc.msg.ValidateBasic())
	}
}

func TestMsgUpdateSessionStatus_GetSignBytes(t *testing.T) {
	msg := msgUpdateSessionStatus
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionStatus_GetSigners(t *testing.T) {
	msg := msgUpdateSessionStatus
	require.Equal(t, []csdkTypes.AccAddress{accAddress1}, msg.GetSigners())
}

func TestMsgUpdateSessionStatus_Route(t *testing.T) {
	msg := msgUpdateSessionStatus
	require.Equal(t, sdkTypes.KeySession, msg.Route())
}

func TestMsgUpdateSessionStatus_Type(t *testing.T) {
	msg := msgUpdateSessionStatus
	require.Equal(t, "msg_update_session_status", msg.Type())
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgDeregisterNode
		want csdkTypes.Error
	}{
		{"From nil", NewMsgDeregisterNode(
			nil, "vpn_id", "locker_id", pubKey1, getMsgReleaseCoinsSignature("locker_id")),
			errorEmptyAddress(),
		},
		{"From empty", NewMsgDeregisterNode(
			csdkTypes.AccAddress{}, "vpn_id", "locker_id", pubKey1, getMsgReleaseCoinsSignature("locker_id")),
			errorEmptyAddress(),
		},
		{"VPN ID empty", NewMsgDeregisterNode(
			accAddress1, "", "locker_id", pubKey1, getMsgReleaseCoinsSignature("locker_id")),
			errorEmptyVPNID(),
		},
		{"Locker ID empty", NewMsgDeregisterNode(
			accAddress1, "vpn_id", "", pubKey1, getMsgReleaseCoinsSignature("")),
			errorEmptyLockerID(),
		},
		{"Public key nil", NewMsgDeregisterNode(
			accAddress1, "vpn_id", "locker_id", nil, getMsgReleaseCoinsSignature("locker_id")),
			errorEmptyPubKey(),
		},
		{"Signature empty", NewMsgDeregisterNode(
			accAddress1, "vpn_id", "locker_id", pubKey1, []byte("")),
			errorEmptySignature(),
		},
		{
			"Valid", msgDeregisterNode, nil,
		},
	}

	for _, tc := range tests {
		require.Equal(t, tc.want, tc.msg.ValidateBasic())
	}
}

func TestMsgDeregisterNode_GetSignBytes(t *testing.T) {
	msg := msgDeregisterNode
	msgBytes, err := json.Marshal(msg)

	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := msgDeregisterNode
	require.Equal(t, []csdkTypes.AccAddress{accAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := msgDeregisterNode
	require.Equal(t, sdkTypes.KeyVPN, msg.Route())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := msgDeregisterNode
	require.Equal(t, "msg_deregister_node", msg.Type())
}

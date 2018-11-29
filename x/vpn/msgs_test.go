package vpn

import (
	"encoding/json"
	"fmt"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	"testing"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	type fields struct {
		From csdkTypes.AccAddress

		APIPort    int64
		Location   sdkTypes.Location
		NetSpeed   sdkTypes.NetSpeed
		EncMethod  string
		PricePerGB int64
		Version    string

		LockerID  string
		Coins     csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{name: "test1", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: true},
		{name: "test2", fields: fields{From: nil, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: false},

		{name: "test3", fields: fields{From: addr3, APIPort: -1, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: false},

		{name: "test4", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  -9000000000,
			Longitude: -1,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test5", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "",
			Country:   "",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test6", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   -1,
			Download: -1,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test7", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 0, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: pubKey1, Signature: sign1}, expectedPass: true},
		{name: "test8", fields: fields{From: addr3, APIPort: 1234, Location: sdkTypes.Location{
			Latitude:  100,
			Longitude: 100,
			City:      "city",
			Country:   "country",
		}, NetSpeed: sdkTypes.NetSpeed{
			Upload:   100,
			Download: 100,
		},
			EncMethod: "AES-256", PricePerGB: 100, Version: "1.0", LockerID: lockerID1,
			Coins: csdkTypes.Coins{coin1}, PubKey: nil, Signature: nil}, expectedPass: false},
	}

	for _, val := range tests {

		msg := MsgRegisterNode{
			val.fields.From, val.fields.APIPort, val.fields.Location,
			val.fields.NetSpeed, val.fields.EncMethod, val.fields.PricePerGB,
			val.fields.Version, val.fields.LockerID,
			val.fields.Coins, val.fields.PubKey, val.fields.Signature,
		}
		if val.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}

}

func TestMsgRegisterNode_GetSignBytes(t *testing.T) {
	msg := TestGetMsgRegisterNode()

	bytes, _ := json.Marshal(msg)

	functionBytes := msg.GetSignBytes()
	require.Equal(t, bytes, functionBytes)

}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	signer := []csdkTypes.AccAddress{addr3}

	msg := TestGetMsgRegisterNode()

	functionSigner := msg.GetSigners()

	require.Equal(t, signer, functionSigner)
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := TestGetMsgRegisterNode()

	functionRoute := msg.Route()
	require.Equal(t, vpnRoute, functionRoute)
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := TestGetMsgRegisterNode()

	functionType := msg.Type()
	require.Equal(t, functionType, vpnRegisterType)
}

func TestMsgPayVPNService_ValidateBasic(t *testing.T) {

	type fields struct {
		From csdkTypes.AccAddress

		VPNID string

		LockerID  string
		Coins     csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}

	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{name: "test1", fields: fields{From: addr3, VPNID: vpnID, LockerID: lockerID1, Coins: csdkTypes.Coins{coin1},
			PubKey: pubKey1, Signature: sign1}, expectedPass: true},
		{name: "test2", fields: fields{From: nil, VPNID: vpnID, LockerID: lockerID1, Coins: csdkTypes.Coins{coin1},
			PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test3", fields: fields{From: addr3, VPNID: "", LockerID: lockerID1, Coins: csdkTypes.Coins{coin1},
			PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test4", fields: fields{From: addr3, VPNID: vpnID, LockerID: "", Coins: csdkTypes.Coins{coin1},
			PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test5", fields: fields{From: addr3, VPNID: vpnID, LockerID: lockerID1, Coins: nil,
			PubKey: pubKey1, Signature: sign1}, expectedPass: false},
		{name: "test6", fields: fields{From: addr3, VPNID: vpnID, LockerID: lockerID1, Coins: csdkTypes.Coins{coin1},
			PubKey: nil, Signature: sign1}, expectedPass: false},
		{name: "test7", fields: fields{From: addr3, VPNID: vpnID, LockerID: lockerID1, Coins: csdkTypes.Coins{coin1},
			PubKey: pubKey1, Signature: nil}, expectedPass: false},
	}

	for _, val := range tests {
		msg := MsgPayVPNService{From: val.fields.From, VPNID: val.fields.VPNID, LockerID: val.fields.LockerID,
			Coins: val.fields.Coins, PubKey: val.fields.PubKey, Signature: val.fields.Signature}
		if val.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}

}

func TestMsgPayVPNService_GetSignBytes(t *testing.T) {
	msg := TestGetMsgPayVpnService()

	bytes, _ := json.Marshal(msg)

	functionBytes := msg.GetSignBytes()

	require.Equal(t, bytes, functionBytes)
}

func TestMsgPayVPNService_GetSigners(t *testing.T) {
	msg := TestGetMsgPayVpnService()

	signer := addr3

	functionSigner := msg.GetSigners()
	require.Equal(t, []csdkTypes.AccAddress{signer}, functionSigner)
}

func TestMsgPayVPNService_Route(t *testing.T) {
	msg := TestGetMsgPayVpnService()
	functionRoute := msg.Route()

	require.Equal(t, vpnRoute, functionRoute)
}

func TestMsgPayVPNService_Type(t *testing.T) {
	msg := TestGetMsgPayVpnService()
	functionType := msg.Type()

	require.Equal(t, functionType, payVpnServiceType)
}

func TestMsgUpdateSessionStatus_ValidateBasic(t *testing.T) {
	type fields struct {
		From      csdkTypes.AccAddress
		SessionId string
		Status    string
	}

	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{name: "test1", fields: fields{From: addr3, SessionId: sessionID1, Status: status}, expectedPass: true},
		{name: "test2", fields: fields{From: nil, SessionId: sessionID1, Status: status}, expectedPass: false},
		{name: "test3", fields: fields{From: addr3, SessionId: "", Status: status}, expectedPass: false},
		{name: "test4", fields: fields{From: addr3, SessionId: sessionID1, Status: ""}, expectedPass: false},
	}

	for _, val := range tests {
		msg := MsgUpdateSessionStatus{From: val.fields.From, SessionID: val.fields.SessionId, Status: val.fields.Status}

		if val.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}
}

func TestMsgUpdateSessionStatus_GetSignBytes(t *testing.T) {
	msg := TestGetMsgUpdateSessionStatus()
	bytes, _ := json.Marshal(msg)
	functionBytes := msg.GetSignBytes()

	require.Equal(t, bytes, functionBytes)

}

func TestMsgUpdateSessionStatus_GetSigners(t *testing.T) {
	msg := TestGetMsgUpdateSessionStatus()
	signer := msg.From
	functionSigners := msg.GetSigners()

	require.Equal(t, functionSigners, []csdkTypes.AccAddress{signer})
}

func TestMsgUpdateSessionStatus_Route(t *testing.T) {
	msg := TestGetMsgUpdateSessionStatus()
	function_route := msg.Route()

	require.Equal(t, function_route, sessionRoute)
}

func TestMsgUpdateSessionStatus_Type(t *testing.T) {
	msg := TestGetMsgUpdateSessionStatus()
	functionType := msg.Type()

	require.Equal(t, functionType, updatedSessionStatusType)
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	type fields struct {
		From      csdkTypes.AccAddress
		VpnID     string
		LockerID  string
		PubKey    crypto.PubKey
		Signature []byte
	}

	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{name: "test1", fields: fields{From: addr3, VpnID: vpnID, LockerID: lockerID1, PubKey: pubKey1, Signature: sign1},
			expectedPass: true},
		{name: "test2", fields: fields{From: nil, VpnID: vpnID, LockerID: lockerID1, PubKey: pubKey1, Signature: sign1},
			expectedPass: false},
		{name: "test3", fields: fields{From: addr3, VpnID: "", LockerID: lockerID1, PubKey: pubKey1, Signature: sign1},
			expectedPass: false},
		{name: "test4", fields: fields{From: addr3, VpnID: vpnID, LockerID: "", PubKey: pubKey1, Signature: sign1},
			expectedPass: false},
		{name: "test5", fields: fields{From: addr3, VpnID: vpnID, LockerID: lockerID1, PubKey: nil, Signature: sign1},
			expectedPass: false},
		{name: "test6", fields: fields{From: addr3, VpnID: vpnID, LockerID: lockerID1, PubKey: pubKey1, Signature: nil},
			expectedPass: false},
	}

	for _, val := range tests {

		msg := MsgDeregisterNode{From: val.fields.From, VPNID: val.fields.VpnID, LockerID: val.fields.LockerID,
			PubKey: val.fields.PubKey, Signature: val.fields.Signature}

		if val.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}
}

func TestMsgDeregisterNode_GetSignBytes(t *testing.T) {
	msg := TestGetMsgDeregisterNode()

	bytes, _ := json.Marshal(msg)
	functionBytes := msg.GetSignBytes()

	require.Equal(t, bytes, functionBytes)
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := TestGetMsgDeregisterNode()

	signers := msg.From
	functionSigners := msg.GetSigners()

	require.Equal(t, []csdkTypes.AccAddress{signers}, functionSigners)
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := TestGetMsgDeregisterNode()
	functionRoute := msg.Route()

	require.Equal(t, functionRoute, vpnRoute)
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := TestGetMsgDeregisterNode()
	functionType := msg.Type()
	fmt.Println(functionType, deregisterNodeIDType)
	require.Equal(t, functionType, deregisterNodeIDType)
}

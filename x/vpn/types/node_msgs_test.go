package types

import (
	"encoding/json"
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterNode
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgRegisterNode(nil, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode(TestAddressEmpty, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is more",
			NewMsgRegisterNode(TestAddress1, TestMonikerLenGT128, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("moniker"),
		}, {
			"amount_to_lock is nil",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinNil, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinNeg, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinNeg, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinZero, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, nil, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsEmpty, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsInvalid, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsNeg, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsZero, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"net_speed is negative",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthNeg, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"net_speed is zero",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthZero, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortInvalid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"encryption id empty",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, "", TestNodeType, TestVersion),
			ErrorInvalidField("encryption_method"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, "", TestVersion),
			ErrorInvalidField("type_"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, ""),
			ErrorInvalidField("version"),
		}, {
			"",
			NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgRegisterNode_GetSignBytes(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, "msg_register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestMonikerValid, TestCoinPos, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeDetails
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeDetails(nil, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeDetails(TestAddressEmpty, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is more",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerLenGT128, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("moniker"),
		}, {
			"id is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDEmpty, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"id in invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDInvalid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, nil, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsEmpty, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsInvalid, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsNeg, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsZero, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"net_speed is zero",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			nil,
		}, {
			"net_speed is negative",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthNeg, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortInvalid, TestEncryptionMethod, TestNodeType, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, "", TestNodeType, TestVersion),
			nil,
		}, {
			"type is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, "", TestVersion),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, ""),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateNode_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, "msg_update_node_details", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestMonikerValid, TestCoinsPos, TestBandwidthPos, TestAPIPortValid, TestEncryptionMethod, TestNodeType, TestVersion)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNodeStatus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeStatus
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeStatus(nil, TestNodeIDValid, StatusActive),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeStatus(TestAddressEmpty, TestNodeIDValid, StatusActive),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDEmpty, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDInvalid, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"status is empty",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, ""),
			ErrorInvalidField("status"),
		}, {
			"status is invalid",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, TestStatusInvalid),
			ErrorInvalidField("status"),
		}, {
			"status is active",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusActive),
			nil,
		}, {
			"status is inactive",
			NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusInactive),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateNodeStatus_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusActive)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusActive)
	require.Equal(t, "msg_update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestNodeIDValid, StatusActive)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgDeregisterNode
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgDeregisterNode(nil, TestNodeIDValid),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode(TestAddressEmpty, TestNodeIDValid),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgDeregisterNode(TestAddress1, TestNodeIDEmpty),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgDeregisterNode(TestAddress1, TestNodeIDInvalid),
			ErrorInvalidField("id"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress1, TestNodeIDValid),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgDeregisterNode_GetSignBytes(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestNodeIDValid)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestNodeIDValid)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestNodeIDValid)
	require.Equal(t, "msg_deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestNodeIDValid)
	require.Equal(t, RouterKey, msg.Route())
}

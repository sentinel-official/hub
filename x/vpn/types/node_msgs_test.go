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
			NewMsgRegisterNode(nil, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode(TestAddressEmpty, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"amount_to_lock is nil",
			NewMsgRegisterNode(TestAddress1, TestCoinNil, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgRegisterNode(TestAddress1, TestCoinNeg, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgRegisterNode(TestAddress1, TestCoinNeg, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgRegisterNode(TestAddress1, TestCoinZero, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, nil, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsEmpty, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsInvalid, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsNeg, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsZero, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is negative",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadNeg, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"upload is zero",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadZero, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is negative",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadNeg, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadZero, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortInvalid, TestEncryption, TestNodeType, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"encryption id empty",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, "", TestNodeType, TestVersion),
			ErrorInvalidField("encryption"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, "", TestVersion),
			ErrorInvalidField("node_type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, ""),
			ErrorInvalidField("version"),
		}, {
			"",
			NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion),
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
	msg := NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion)
	require.Equal(t, "msg_register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestNodeType, TestVersion)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNodeDetails_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeDetails
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeDetails(nil, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeDetails(TestAddressEmpty, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDEmpty, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"id in invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDInvalid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, nil, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsEmpty, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsInvalid, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsNeg, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsZero, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is zero",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadZero, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			nil,
		}, {
			"upload is negative",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadNeg, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadZero, TestAPIPortValid, TestEncryption, TestVersion),
			nil,
		}, {
			"download is negative",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadNeg, TestAPIPortValid, TestEncryption, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortInvalid, TestEncryption, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, "", TestVersion),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, ""),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion),
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

func TestMsgUpdateNodeDetails_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeDetails_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNodeDetails_Type(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion)
	require.Equal(t, "msg_update_node_details", msg.Type())
}

func TestMsgUpdateNodeDetails_Route(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress1, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncryption, TestVersion)
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

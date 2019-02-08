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
			NewMsgRegisterNode(nil, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode(TestAddressEmpty, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"amount_to_lock is nil",
			NewMsgRegisterNode(TestAddress, TestCoinNil, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgRegisterNode(TestAddress, TestCoinNeg, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgRegisterNode(TestAddress, TestCoinNeg, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgRegisterNode(TestAddress, TestCoinZero, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress, TestCoinPos, nil, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsEmpty, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsInvalid, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsNeg, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsZero, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is negative",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadNeg, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"upload is zero",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadZero, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is negative",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadNeg, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadZero, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortInvalid, TestEncMethod, TestNodeType, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"enc_method id empty",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, "", TestNodeType, TestVersion),
			ErrorInvalidField("enc_method"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, "", TestVersion),
			ErrorInvalidField("node_type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, ""),
			ErrorInvalidField("version"),
		}, {
			"",
			NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion),
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
	msg := NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion)
	require.Equal(t, "msg_register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress, TestCoinPos, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestNodeType, TestVersion)
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
			NewMsgUpdateNodeDetails(nil, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeDetails(TestAddressEmpty, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDEmpty, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"id in invalid",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDInvalid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("id"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, nil, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsEmpty, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsInvalid, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsNeg, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsZero, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is zero",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadZero, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			nil,
		}, {
			"upload is negative",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadNeg, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadZero, TestAPIPortValid, TestEncMethod, TestVersion),
			nil,
		}, {
			"download is negative",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadNeg, TestAPIPortValid, TestEncMethod, TestVersion),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortInvalid, TestEncMethod, TestVersion),
			ErrorInvalidField("api_port"),
		}, {
			"enc_method is empty",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, "", TestVersion),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, ""),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion),
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
	msg := NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeDetails_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress}, msg.GetSigners())
}

func TestMsgUpdateNodeDetails_Type(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion)
	require.Equal(t, "msg_update_node_details", msg.Type())
}

func TestMsgUpdateNodeDetails_Route(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(TestAddress, TestNodeIDValid, TestCoinsPos, TestUploadPos, TestDownloadPos, TestAPIPortValid, TestEncMethod, TestVersion)
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
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDEmpty, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDInvalid, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"status is empty",
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, ""),
			ErrorInvalidField("status"),
		}, {
			"status is invalid",
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, TestStatusInvalid),
			ErrorInvalidField("status"),
		}, {
			"status is active",
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusActive),
			nil,
		}, {
			"status is inactive",
			NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusInactive),
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
	msg := NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusActive)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusActive)
	require.Equal(t, "msg_update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress, TestNodeIDValid, StatusActive)
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
			NewMsgDeregisterNode(TestAddress, TestNodeIDEmpty),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgDeregisterNode(TestAddress, TestNodeIDInvalid),
			ErrorInvalidField("id"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress, TestNodeIDValid),
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
	msg := NewMsgDeregisterNode(TestAddress, TestNodeIDValid)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress, TestNodeIDValid)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress, TestNodeIDValid)
	require.Equal(t, "msg_deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress, TestNodeIDValid)
	require.Equal(t, RouterKey, msg.Route())
}

package types

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRegisterNode(nil, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode(TestAddressEmpty, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("from"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, "", TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, TestNodeType, "", TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("version"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerLengthGT128, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, nil,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsEmpty,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsNeg,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsZero,
				TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is negative",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthNeg, TestEncryption),
			ErrorInvalidField("internet_speed"),
		}, {
			"internet_speed is zero",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthZero, TestEncryption),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, ""),
			ErrorInvalidField("encryption"),
		}, {
			"valid",
			NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos,
				TestBandwidthPos1, TestEncryption),
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
	msg := NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, "register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNodeInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeInfo
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeInfo(nil, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeInfo(TestAddressEmpty, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerLengthGT128, TestCoinsPos, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, nil, TestBandwidthPos1, TestEncryption),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsEmpty, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsNeg, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsZero, TestBandwidthPos1, TestEncryption),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is zero",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthZero, TestEncryption),
			nil,
		}, {
			"internet_speed is negative",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthNeg, TestEncryption),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, ""),
			nil,
		}, {
			"type is empty",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, "", TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, "", TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption),
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
	msg := NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, "update_node_info", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, TestIDPos, TestNodeType, TestVersion, TestMonikerValid, TestCoinsPos, TestBandwidthPos1, TestEncryption)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNodeStatus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeStatus
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeStatus(nil, TestIDPos, TestStatusActive),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeStatus(TestAddressEmpty, TestIDPos, TestStatusActive),
			ErrorInvalidField("from"),
		}, {
			"status is empty",
			NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, ""),
			ErrorInvalidField("status"),
		}, {
			"status is invalid",
			NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, TestStatusInValid),
			ErrorInvalidField("status"),
		}, {
			"status is active",
			NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, TestStatusActive),
			nil,
		}, {
			"status is inactive",
			NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, StatusInactive),
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
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, StatusActive)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, StatusActive)
	require.Equal(t, "update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, TestIDPos, StatusActive)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgDeregisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgDeregisterNode(nil, TestIDPos),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode(TestAddressEmpty, TestIDPos),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress1, TestIDPos),
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
	msg := NewMsgDeregisterNode(TestAddress1, TestIDPos)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestIDPos)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestIDPos)
	require.Equal(t, "deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, TestIDPos)
	require.Equal(t, RouterKey, msg.Route())
}

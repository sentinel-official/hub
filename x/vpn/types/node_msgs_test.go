package types

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRegisterNode(nil, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode([]byte(""), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("version"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", nil,
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"internet_speed is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), ""),
			ErrorInvalidField("encryption"),
		}, {
			"valid",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
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
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, "register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
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
			NewMsgUpdateNodeInfo(nil, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeInfo([]byte(""), hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", nil, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), "encryption"),
			nil,
		}, {
			"internet_speed is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), ""),
			nil,
		}, {
			"type is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
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
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, "update_node_info", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
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
			NewMsgUpdateNodeStatus(nil, hub.NewIDFromUInt64(1), StatusActive),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeStatus([]byte(""), hub.NewIDFromUInt64(1), StatusActive),
			ErrorInvalidField("from"),
		}, {
			"status is empty",
			NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), ""),
			ErrorInvalidField("status"),
		}, {
			"status is invalid",
			NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), "status"),
			ErrorInvalidField("status"),
		}, {
			"status is active",
			NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusActive),
			nil,
		}, {
			"status is inactive",
			NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusInactive),
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
	msg := NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusActive)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusActive)
	require.Equal(t, "update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(TestAddress1, hub.NewIDFromUInt64(1), StatusActive)
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
			NewMsgDeregisterNode(nil, hub.NewIDFromUInt64(1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode([]byte(""), hub.NewIDFromUInt64(1)),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress1, hub.NewIDFromUInt64(1)),
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
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewIDFromUInt64(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewIDFromUInt64(1))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewIDFromUInt64(1))
	require.Equal(t, "deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewIDFromUInt64(1))
	require.Equal(t, RouterKey, msg.Route())
}

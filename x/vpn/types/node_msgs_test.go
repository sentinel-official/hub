package types_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgRegisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgRegisterNode(nil, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgRegisterNode([]byte(""), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("from"),
		}, {
			"node_type is empty",
			types.NewMsgRegisterNode(address1, "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("type"),
		}, {
			"version is empty",
			types.NewMsgRegisterNode(address1, "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("version"),
		}, {
			"node_moniker length is greater than 128",
			types.NewMsgRegisterNode(address1, "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", nil,
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is negative",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), "encryption"),
			types.ErrorInvalidField("internet_speed"),
		}, {
			"internet_speed is zero",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), "encryption"),
			types.ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
				hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), ""),
			types.ErrorInvalidField("encryption"),
		}, {
			"valid",
			types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)},
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
	msg := types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, "register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := types.NewMsgRegisterNode(address1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgUpdateNodeInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgUpdateNodeInfo
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgUpdateNodeInfo(nil, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgUpdateNodeInfo([]byte(""), hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("from"),
		}, {
			"node_moniker length is greater than 128",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", nil, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"prices_per_gb is empty",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			types.ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is zero",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), "encryption"),
			nil,
		}, {
			"internet_speed is negative",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), "encryption"),
			types.ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), ""),
			nil,
		}, {
			"type is empty",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"version is empty",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
			nil,
		}, {
			"valid",
			types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption"),
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
	msg := types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, "update_node_info", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := types.NewMsgUpdateNodeInfo(address1, hub.NewIDFromUInt64(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), "encryption")
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgUpdateNodeStatus_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgUpdateNodeStatus
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgUpdateNodeStatus(nil, hub.NewIDFromUInt64(1), types.StatusActive),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgUpdateNodeStatus([]byte(""), hub.NewIDFromUInt64(1), types.StatusActive),
			types.ErrorInvalidField("from"),
		}, {
			"status is empty",
			types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), ""),
			types.ErrorInvalidField("status"),
		}, {
			"status is invalid",
			types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), "status"),
			types.ErrorInvalidField("status"),
		}, {
			"status is active",
			types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusActive),
			nil,
		}, {
			"status is inactive",
			types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusInactive),
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
	msg := types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusActive)
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusActive)
	require.Equal(t, "update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := types.NewMsgUpdateNodeStatus(address1, hub.NewIDFromUInt64(1), types.StatusActive)
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgDeregisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgDeregisterNode(nil, hub.NewIDFromUInt64(1)),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgDeregisterNode([]byte(""), hub.NewIDFromUInt64(1)),
			types.ErrorInvalidField("from"),
		}, {
			"valid",
			types.NewMsgDeregisterNode(address1, hub.NewIDFromUInt64(1)),
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
	msg := types.NewMsgDeregisterNode(address1, hub.NewIDFromUInt64(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := types.NewMsgDeregisterNode(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := types.NewMsgDeregisterNode(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, "deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := types.NewMsgDeregisterNode(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, types.RouterKey, msg.Route())
}

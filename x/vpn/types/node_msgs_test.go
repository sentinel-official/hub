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
			NewMsgRegisterNode(nil, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode([]byte(""), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("version"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", nil, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthNeg, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"internet_speed is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthZero, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, ""),
			ErrorInvalidField("encryption"),
		}, {
			"valid",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
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
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, "register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
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
			NewMsgUpdateNodeInfo(nil, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeInfo([]byte(""), hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", nil, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthZero, "encryption"),
			nil,
		}, {
			"internet_speed is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthNeg, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, ""),
			nil,
		}, {
			"type is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
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
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, "update_node_info", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgAddFreeClient_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgAddFreeClient
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgAddFreeClient(nil, hub.NewNodeID(1), TestAddress1),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgAddFreeClient([]byte(""), hub.NewNodeID(1), TestAddress1),
			ErrorInvalidField("from"),
		}, {
			"node_id is nil",
			NewMsgAddFreeClient(TestAddress1, nil, TestAddress1),
			ErrorInvalidField("node_id"),
		}, {
			"client is nil",
			NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), nil),
			ErrorInvalidField("client"),
		}, {
			"client is empty",
			NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), []byte("")),
			ErrorInvalidField("client"),
		}, {
			"valid",
			NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1),
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

func TestMsgAddFreeClient_GetSignBytes(t *testing.T) {
	msg := NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgAddFreeClient_GetSigners(t *testing.T) {
	msg := NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgAddFreeClient_Type(t *testing.T) {
	msg := NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, "add_free_client", msg.Type())
}

func TestMsgAddFreeClient_Route(t *testing.T) {
	msg := NewMsgAddFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgRemoveFreeClient_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRemoveFreeClient
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRemoveFreeClient(nil, hub.NewNodeID(1), TestAddress1),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRemoveFreeClient([]byte(""), hub.NewNodeID(1), TestAddress1),
			ErrorInvalidField("from"),
		}, {
			"node_id is nil",
			NewMsgRemoveFreeClient(TestAddress1, nil, TestAddress1),
			ErrorInvalidField("node_id"),
		}, {
			"client is nil",
			NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), nil),
			ErrorInvalidField("client"),
		}, {
			"client is empty",
			NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), []byte("")),
			ErrorInvalidField("client"),
		}, {
			"valid",
			NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1),
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

func TestMsgRemoveFreeClient_GetSignBytes(t *testing.T) {
	msg := NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRemoveFreeClient_GetSigners(t *testing.T) {
	msg := NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRemoveFreeClient_Type(t *testing.T) {
	msg := NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, "remove_free_client", msg.Type())
}

func TestMsgRemoveFreeClient_Route(t *testing.T) {
	msg := NewMsgRemoveFreeClient(TestAddress1, hub.NewNodeID(1), TestAddress1)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgAddVPNOnResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterVPNOnResolver
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRegisterVPNOnResolver(nil, hub.NewNodeID(1), hub.NewResolverID(0)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterVPNOnResolver([]byte(""), hub.NewNodeID(1), hub.NewResolverID(0)),
			ErrorInvalidField("from"),
		}, {
			"node_id is nil",
			NewMsgRegisterVPNOnResolver(TestAddress1, nil, hub.NewResolverID(0)),
			ErrorInvalidField("node_id"),
		}, {
			"resolver is nil",
			NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), nil),
			ErrorInvalidField("resolver"),
		}, {
			"valid",
			NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0)),
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

func TestMsgRegisterVPNOnResolver_GetSignBytes(t *testing.T) {
	msg := NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterVPNOnResolver_GetSigners(t *testing.T) {
	msg := NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterVPNOnResolver_Type(t *testing.T) {
	msg := NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	require.Equal(t, "register_vpn_on_resolver", msg.Type())
}

func TestMsgRegisterVPNOnResolver_Route(t *testing.T) {
	msg := NewMsgRegisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterVPNOnResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgDeregisterVPNOnResolver
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgDeregisterVPNOnResolver(nil, hub.NewNodeID(1), hub.NewResolverID(0)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterVPNOnResolver([]byte(""), hub.NewNodeID(1), hub.NewResolverID(0)),
			ErrorInvalidField("from"),
		}, {
			"node_id is nil",
			NewMsgDeregisterVPNOnResolver(TestAddress1, nil, hub.NewResolverID(0)),
			ErrorInvalidField("node_id"),
		}, {
			"resolver is nil",
			NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), nil),
			ErrorInvalidField("resolver"),
		}, {
			"valid",
			NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0)),
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

func TestMsgDeregisterVPNOnResolver_GetSignBytes(t *testing.T) {
	msg := NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterVPNOnResolver_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterVPNOnResolver_Type(t *testing.T) {
	msg := NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
	require.Equal(t, "deregister_vpn_on_resolver", msg.Type())
}

func TestMsgDeregisterVPNOnResolver_Route(t *testing.T) {
	msg := NewMsgDeregisterVPNOnResolver(TestAddress1, hub.NewNodeID(1), hub.NewResolverID(0))
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
			NewMsgDeregisterNode(nil, hub.NewNodeID(1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode([]byte(""), hub.NewNodeID(1)),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1)),
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
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, "deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, RouterKey, msg.Route())
}

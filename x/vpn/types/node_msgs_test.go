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
			NewMsgRegisterNode(nil, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode(AddressEmpty, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("from"),
		}, {
			"amount_to_lock is nil",
			NewMsgRegisterNode(Address1, CoinNil, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgRegisterNode(Address1, CoinEmpty, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgRegisterNode(Address1, CoinNeg, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgRegisterNode(Address1, CoinZero, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(Address1, CoinPos, nil, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(Address1, CoinPos, CoinsEmpty, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgRegisterNode(Address1, CoinPos, CoinsInvalid, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(Address1, CoinPos, CoinsNeg, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(Address1, CoinPos, CoinsZero, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is negative",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadNeg, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"upload is zero",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadZero, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"download is negative",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadNeg, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadZero, APIPortValid, EncMethod, NodeType, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortInvalid, EncMethod, NodeType, Version),
			ErrorInvalidField("api_port"),
		}, {
			"enc_method id empty",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, "", NodeType, Version),
			ErrorInvalidField("enc_method"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, "", Version),
			ErrorInvalidField("node_type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, ""),
			ErrorInvalidField("version"),
		}, {
			"",
			NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version),
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
	msg := NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version)
	require.Equal(t, "msg_register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(Address1, CoinPos, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, NodeType, Version)
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
			NewMsgUpdateNodeDetails(nil, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeDetails(AddressEmpty, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateNodeDetails(Address1, NodeIDEmpty, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("id"),
		}, {
			"id in invalid",
			NewMsgUpdateNodeDetails(Address1, NodeIDInvalid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("id"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, nil, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsEmpty, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is invalid",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsInvalid, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsNeg, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsZero, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"upload is zero",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadZero, DownloadPos, APIPortValid, EncMethod, Version),
			nil,
		}, {
			"upload is negative",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadNeg, DownloadPos, APIPortValid, EncMethod, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"download is zero",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadZero, APIPortValid, EncMethod, Version),
			nil,
		}, {
			"download is negative",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadNeg, APIPortValid, EncMethod, Version),
			ErrorInvalidField("net_speed"),
		}, {
			"api_port is invalid",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortInvalid, EncMethod, Version),
			ErrorInvalidField("api_port"),
		}, {
			"enc_method is empty",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, "", Version),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, ""),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version),
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
	msg := NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeDetails_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgUpdateNodeDetails_Type(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version)
	require.Equal(t, "msg_update_node_details", msg.Type())
}

func TestMsgUpdateNodeDetails_Route(t *testing.T) {
	msg := NewMsgUpdateNodeDetails(Address1, NodeIDValid, CoinsPos, UploadPos, DownloadPos, APIPortValid, EncMethod, Version)
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
			NewMsgUpdateNodeStatus(nil, NodeIDValid, StatusActive),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeStatus(AddressEmpty, NodeIDValid, StatusActive),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateNodeStatus(Address1, NodeIDEmpty, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgUpdateNodeStatus(Address1, NodeIDInvalid, StatusActive),
			ErrorInvalidField("id"),
		}, {
			"status is empty",
			NewMsgUpdateNodeStatus(Address1, NodeIDValid, ""),
			ErrorInvalidField("status"),
		}, {
			"status is invalid",
			NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusInvalid),
			ErrorInvalidField("status"),
		}, {
			"status is active",
			NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusActive),
			nil,
		}, {
			"status is inactive",
			NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusInactive),
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
	msg := NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusActive)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNodeStatus_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusActive)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgUpdateNodeStatus_Type(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusActive)
	require.Equal(t, "msg_update_node_status", msg.Type())
}

func TestMsgUpdateNodeStatus_Route(t *testing.T) {
	msg := NewMsgUpdateNodeStatus(Address1, NodeIDValid, StatusActive)
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
			NewMsgDeregisterNode(nil, NodeIDValid),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode(AddressEmpty, NodeIDValid),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgDeregisterNode(Address1, NodeIDEmpty),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgDeregisterNode(Address1, NodeIDInvalid),
			ErrorInvalidField("id"),
		}, {
			"valid",
			NewMsgDeregisterNode(Address1, NodeIDValid),
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
	msg := NewMsgDeregisterNode(Address1, NodeIDValid)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(Address1, NodeIDValid)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(Address1, NodeIDValid)
	require.Equal(t, "msg_deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(Address1, NodeIDValid)
	require.Equal(t, RouterKey, msg.Route())
}

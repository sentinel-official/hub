package types

import (
	"encoding/json"
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgInitSession_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgInitSession
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgInitSession(nil, TestNodeIDValid, TestCoinPos),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgInitSession(TestAddressEmpty, TestNodeIDValid, TestCoinPos),
			ErrorInvalidField("from"),
		}, {
			"node_id is empty",
			NewMsgInitSession(TestAddress2, TestNodeIDEmpty, TestCoinPos),
			ErrorInvalidField("node_id"),
		}, {
			"node_id is invalid",
			NewMsgInitSession(TestAddress2, TestNodeIDInvalid, TestCoinPos),
			ErrorInvalidField("node_id"),
		}, {
			"amount_to_lock is nil",
			NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinNil),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinEmpty),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinNeg),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinZero),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"valid",
			NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinPos),
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

func TestMsgInitSession_GetSignBytes(t *testing.T) {
	msg := NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinPos)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgInitSession_GetSigners(t *testing.T) {
	msg := NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinPos)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress2}, msg.GetSigners())
}

func TestMsgInitSession_Type(t *testing.T) {
	msg := NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinPos)
	require.Equal(t, "msg_init_session", msg.Type())
}

func TestMsgInitSession_Route(t *testing.T) {
	msg := NewMsgInitSession(TestAddress2, TestNodeIDValid, TestCoinPos)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateSessionBandwidth_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateSessionBandwidth
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgUpdateSessionBandwidth(nil, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateSessionBandwidth(TestAddressEmpty, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDEmpty, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDInvalid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("id"),
		}, {
			"upload is negative",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadNeg, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"upload is zero",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadZero, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"download is negative",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadNeg, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"download is zero",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadNeg, TestClientSign, TestNodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"client_sign is nil",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadPos, nil, TestNodeOwnerSign),
			ErrorInvalidField("client_sign"),
		}, {
			"client_sign is empty",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadPos, []byte{}, TestNodeOwnerSign),
			ErrorInvalidField("client_sign"),
		}, {
			"node_owner_sign is nil",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, nil),
			ErrorInvalidField("node_owner_sign"),
		}, {
			"node_owner_sign is empty",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, []byte{}),
			ErrorInvalidField("node_owner_sign"),
		}, {
			"valid",
			NewMsgUpdateSessionBandwidth(TestAddress1, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign),
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

func TestMsgUpdateSessionBandwidth_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(TestAddress2, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionBandwidth_GetSigners(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(TestAddress2, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress2}, msg.GetSigners())
}

func TestMsgUpdateSessionBandwidth_Type(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(TestAddress2, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign)
	require.Equal(t, "msg_update_session_bandwidth", msg.Type())
}

func TestMsgUpdateSessionBandwidth_Route(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(TestAddress2, TestSessionIDValid, TestUploadPos, TestDownloadPos, TestClientSign, TestNodeOwnerSign)
	require.Equal(t, RouterKey, msg.Route())
}

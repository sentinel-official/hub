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
			NewMsgInitSession(nil, NodeIDValid, CoinPos),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgInitSession(AddressEmpty, NodeIDValid, CoinPos),
			ErrorInvalidField("from"),
		}, {
			"node_id is empty",
			NewMsgInitSession(Address1, NodeIDEmpty, CoinPos),
			ErrorInvalidField("node_id"),
		}, {
			"node_id is invalid",
			NewMsgInitSession(Address1, NodeIDInvalid, CoinPos),
			ErrorInvalidField("node_id"),
		}, {
			"amount_to_lock is nil",
			NewMsgInitSession(Address1, NodeIDValid, CoinNil),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is empty",
			NewMsgInitSession(Address1, NodeIDValid, CoinEmpty),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is negative",
			NewMsgInitSession(Address1, NodeIDValid, CoinNeg),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"amount_to_lock is zero",
			NewMsgInitSession(Address1, NodeIDValid, CoinZero),
			ErrorInvalidField("amount_to_lock"),
		}, {
			"valid",
			NewMsgInitSession(Address1, NodeIDValid, CoinPos),
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
	msg := NewMsgInitSession(Address1, NodeIDValid, CoinPos)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgInitSession_GetSigners(t *testing.T) {
	msg := NewMsgInitSession(Address1, NodeIDValid, CoinPos)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgInitSession_Type(t *testing.T) {
	msg := NewMsgInitSession(Address1, NodeIDValid, CoinPos)
	require.Equal(t, "msg_init_session", msg.Type())
}

func TestMsgInitSession_Route(t *testing.T) {
	msg := NewMsgInitSession(Address1, NodeIDValid, CoinPos)
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
			NewMsgUpdateSessionBandwidth(nil, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateSessionBandwidth(AddressEmpty, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("from"),
		}, {
			"id is empty",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDEmpty, UploadPos, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("id"),
		}, {
			"id is invalid",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDInvalid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("id"),
		}, {
			"upload is negative",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadNeg, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"upload is zero",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadZero, DownloadPos, ClientSign, NodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"download is negative",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadNeg, ClientSign, NodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"download is zero",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadNeg, ClientSign, NodeOwnerSign),
			ErrorInvalidField("bandwidth"),
		}, {
			"client_sign is nil",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, nil, NodeOwnerSign),
			ErrorInvalidField("client_sign"),
		}, {
			"client_sign is empty",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, []byte{}, NodeOwnerSign),
			ErrorInvalidField("client_sign"),
		}, {
			"node_owner_sign is nil",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, nil),
			ErrorInvalidField("node_owner_sign"),
		}, {
			"node_owner_sign is empty",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, []byte{}),
			ErrorInvalidField("node_owner_sign"),
		}, {
			"valid",
			NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign),
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
	msg := NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionBandwidth_GetSigners(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign)
	require.Equal(t, []csdkTypes.AccAddress{Address1}, msg.GetSigners())
}

func TestMsgUpdateSessionBandwidth_Type(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign)
	require.Equal(t, "msg_update_session_bandwidth", msg.Type())
}

func TestMsgUpdateSessionBandwidth_Route(t *testing.T) {
	msg := NewMsgUpdateSessionBandwidth(Address1, SessionIDValid, UploadPos, DownloadPos, ClientSign, NodeOwnerSign)
	require.Equal(t, RouterKey, msg.Route())
}

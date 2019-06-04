package types

import (
	"encoding/json"
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateSessionInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateSessionInfo
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgUpdateSessionInfo(nil, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateSessionInfo(TestAddressEmpty, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
			ErrorInvalidField("from"),
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthZero, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is neg",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthNeg, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthZero, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"node owner sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TeststdSignatureEmpty, TestClientstdSignaturePos1),
			ErrorInvalidField("node_owner_signature"),
		}, {
			"client sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TeststdSignatureEmpty),
			ErrorInvalidField("client_signature"),
		}, {
			"valid ",
			NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1),
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

func TestMsgUpdateSessionInfo_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionInfo_GetSigners(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateSessionInfo_Type(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1)
	require.Equal(t, "MsgUpdateSessionInfo", msg.Type())
}

func TestMsgUpdateSessionInfo_Route(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, TestIDPos, TestBandwidthPos1, TestNodeOwnerstdSignaturePos1, TestClientstdSignaturePos1)
	require.Equal(t, RouterKey, msg.Route())
}

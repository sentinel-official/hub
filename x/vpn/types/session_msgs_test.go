package types

import (
	"encoding/json"
	"reflect"
	"testing"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/stretchr/testify/require"
	
	hub "github.com/sentinel-official/hub/types"
)

func TestMsgUpdateSessionInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateSessionInfo
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgUpdateSessionInfo(nil, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateSessionInfo([]byte(""), hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrorInvalidField("from"),
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthZero, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is neg",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthNeg, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"bandwidth is zero",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthZero, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
			ErrorInvalidField("bandwidth"),
		}, {
			"node owner sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, auth.StdSignature{}, TestClientStdSignaturePos1),
			ErrorInvalidField("node_owner_signature"),
		}, {
			"client sign is empty  ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, auth.StdSignature{}),
			ErrorInvalidField("client_signature"),
		}, {
			"valid ",
			NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1),
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
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	
	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateSessionInfo_GetSigners(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateSessionInfo_Type(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, "update_session_info", msg.Type())
}

func TestMsgUpdateSessionInfo_Route(t *testing.T) {
	msg := NewMsgUpdateSessionInfo(TestAddress1, hub.NewSubscriptionID(1), TestBandwidthPos1, TestNodeOwnerStdSignaturePos1, TestClientStdSignaturePos1)
	require.Equal(t, RouterKey, msg.Route())
}

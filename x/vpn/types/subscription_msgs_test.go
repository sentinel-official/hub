package types

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgStartSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgStartSubscription
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgStartSubscription(nil, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgStartSubscription([]byte(""), TestAddress1, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("from"),
		}, {
			"resolver address is nil",
			NewMsgStartSubscription(TestAddress2, nil, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("resolver"),
		}, {
			"resolver is empty",
			NewMsgStartSubscription(TestAddress1, []byte(""), hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("resolver"),
		}, {
			"node id is nil",
			NewMsgStartSubscription(TestAddress1, TestAddress2, nil, sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("node_id"),
		},
		{
			"node id is empty",
			NewMsgStartSubscription(TestAddress1, TestAddress2, []byte(""), sdk.NewInt64Coin("stake", 100)),
			ErrorInvalidField("node_id"),
		},
		{
			"deposit is empty",
			NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.Coin{}),
			ErrorInvalidField("deposit"),
		}, {
			"deposit is zero",
			NewMsgStartSubscription(TestAddress1, TestAddress1, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 0)),
			ErrorInvalidField("deposit"),
		}, {
			"valid",
			NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100)),
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

func TestMsgStartSubscription_GetSignBytes(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgStartSubscription_GetSigners(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgStartSubscription_Type(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, "start_subscription", msg.Type())
}

func TestMsgStartSubscription_Route(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestAddress2, hub.NewNodeID(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgEndSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgEndSubscription
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgEndSubscription(nil, hub.NewSubscriptionID(1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgEndSubscription([]byte(""), hub.NewSubscriptionID(1)),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgEndSubscription(TestAddress1, hub.NewSubscriptionID(1)),
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

func TestMsgEndSubscription_GetSignBytes(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, hub.NewSubscriptionID(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgEndSubscription_GetSigners(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, hub.NewSubscriptionID(1))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgEndSubscription_Type(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, hub.NewSubscriptionID(1))
	require.Equal(t, "end_subscription", msg.Type())
}

func TestMsgEndSubscription_Route(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, hub.NewSubscriptionID(1))
	require.Equal(t, RouterKey, msg.Route())
}

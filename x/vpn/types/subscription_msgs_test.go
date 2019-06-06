package types

import (
	"encoding/json"
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgStartSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgStartSubscription
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgStartSubscription(nil, TestIDPos, TestCoinPos),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgStartSubscription(TestAddressEmpty, TestIDPos, TestCoinPos),
			ErrorInvalidField("from"),
		}, {
			"deposit is empty",
			NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinEmpty),
			ErrorInvalidField("deposit"),
		}, {
			"deposit is zero",
			NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinZero),
			ErrorInvalidField("deposit"),
		}, {
			"valid",
			NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinPos),
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
	msg := NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinPos)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgStartSubscription_GetSigners(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinPos)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgStartSubscription_Type(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinPos)
	require.Equal(t, "MsgStartSubscription", msg.Type())
}

func TestMsgStartSubscription_Route(t *testing.T) {
	msg := NewMsgStartSubscription(TestAddress1, TestIDPos, TestCoinPos)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgEndSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgEndSubscription
		want csdkTypes.Error
	}{
		{
			"from is nil",
			NewMsgEndSubscription(nil, TestIDPos),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgEndSubscription(TestAddressEmpty, TestIDPos),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgEndSubscription(TestAddress1, TestIDPos),
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
	msg := NewMsgEndSubscription(TestAddress1, TestIDPos)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgEndSubscription_GetSigners(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, TestIDPos)
	require.Equal(t, []csdkTypes.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgEndSubscription_Type(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, TestIDPos)
	require.Equal(t, "MsgEndSubscription", msg.Type())
}

func TestMsgEndSubscription_Route(t *testing.T) {
	msg := NewMsgEndSubscription(TestAddress1, TestIDPos)
	require.Equal(t, RouterKey, msg.Route())
}

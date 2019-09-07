package types_test

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestMsgStartSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgStartSubscription
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgStartSubscription(nil, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100)),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgStartSubscription([]byte(""), hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100)),
			types.ErrorInvalidField("from"),
		}, {
			"deposit is empty",
			types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.Coin{}),
			types.ErrorInvalidField("deposit"),
		}, {
			"deposit is zero",
			types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 0)),
			types.ErrorInvalidField("deposit"),
		}, {
			"valid",
			types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100)),
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
	msg := types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgStartSubscription_GetSigners(t *testing.T) {
	msg := types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgStartSubscription_Type(t *testing.T) {
	msg := types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, "start_subscription", msg.Type())
}

func TestMsgStartSubscription_Route(t *testing.T) {
	msg := types.NewMsgStartSubscription(address1, hub.NewIDFromUInt64(1), sdk.NewInt64Coin("stake", 100))
	require.Equal(t, types.RouterKey, msg.Route())
}

func TestMsgEndSubscription_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *types.MsgEndSubscription
		want sdk.Error
	}{
		{
			"from is nil",
			types.NewMsgEndSubscription(nil, hub.NewIDFromUInt64(1)),
			types.ErrorInvalidField("from"),
		}, {
			"from is empty",
			types.NewMsgEndSubscription([]byte(""), hub.NewIDFromUInt64(1)),
			types.ErrorInvalidField("from"),
		}, {
			"valid",
			types.NewMsgEndSubscription(address1, hub.NewIDFromUInt64(1)),
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
	msg := types.NewMsgEndSubscription(address1, hub.NewIDFromUInt64(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgEndSubscription_GetSigners(t *testing.T) {
	msg := types.NewMsgEndSubscription(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, []sdk.AccAddress{address1}, msg.GetSigners())
}

func TestMsgEndSubscription_Type(t *testing.T) {
	msg := types.NewMsgEndSubscription(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, "end_subscription", msg.Type())
}

func TestMsgEndSubscription_Route(t *testing.T) {
	msg := types.NewMsgEndSubscription(address1, hub.NewIDFromUInt64(1))
	require.Equal(t, types.RouterKey, msg.Route())
}

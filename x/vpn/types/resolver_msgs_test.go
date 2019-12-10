package types

import (
	"encoding/json"
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterResolver_GetSignBytes(t *testing.T) {
	msg := NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterResolver_GetSigners(t *testing.T) {
	msg := NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterResolver_Route(t *testing.T) {
	msg := NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgRegisterResolver_Type(t *testing.T) {
	msg := NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(1, 2))
	require.Equal(t, "register_resolver_node", msg.Type())
}

func TestMsgRegisterResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterResolver
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRegisterResolver(nil, sdk.NewDecWithPrec(2, 1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterResolver([]byte(""), sdk.OneDec()),
			ErrorInvalidField("from"),
		}, {
			"commission is negative",
			NewMsgRegisterResolver(TestAddress1, sdk.NewDecWithPrec(-1, 0)),
			ErrorInvalidField("commission"),
		}, {
			"commission is grater than 1",
			NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(2, 0)),
			ErrorInvalidField("commission"),
		}, {
			"commission with zero",
			NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(0, 0)),
			nil,
		}, {
			"commission with one",
			NewMsgRegisterResolver(TestAddress2, sdk.NewDecWithPrec(1, 0)),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateResolverInfo_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateResolverInfo(TestAddress1, sdk.OneDec())
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateResolverInfo_GetSigners(t *testing.T) {
	msg := NewMsgUpdateResolverInfo(TestAddress1, sdk.OneDec())
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateResolverInfo_Route(t *testing.T) {
	msg := NewMsgUpdateResolverInfo(TestAddress1, sdk.OneDec())
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateResolverInfo_Type(t *testing.T) {
	msg := NewMsgUpdateResolverInfo(TestAddress1, sdk.OneDec())
	require.Equal(t, "update_resolver_info", msg.Type())
}

func TestMsgUpdateResolverInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateResolverInfo
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgUpdateResolverInfo(nil, sdk.NewDecWithPrec(2, 1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateResolverInfo([]byte(""), sdk.OneDec()),
			ErrorInvalidField("from"),
		}, {
			"commission is negative",
			NewMsgUpdateResolverInfo(TestAddress1, sdk.NewDecWithPrec(-1, 0)),
			ErrorInvalidField("commission"),
		}, {
			"commission is grater than 1",
			NewMsgUpdateResolverInfo(TestAddress2, sdk.NewDecWithPrec(2, 0)),
			ErrorInvalidField("commission"),
		}, {
			"commission with zero",
			NewMsgUpdateResolverInfo(TestAddress2, sdk.NewDecWithPrec(0, 0)),
			nil,
		}, {
			"commission with one",
			NewMsgUpdateResolverInfo(TestAddress2, sdk.NewDecWithPrec(1, 0)),
			nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMsgDeregisterResolver_GetSignBytes(t *testing.T) {
	msg := NewMsgDeregisterResolver(TestAddress1)
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterResolver_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterResolver(TestAddress1)
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterResolver_Route(t *testing.T) {
	msg := NewMsgDeregisterResolver(TestAddress1)
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterResolver_Type(t *testing.T) {
	msg := NewMsgDeregisterResolver(TestAddress1)
	require.Equal(t, "deregister_resolver_node", msg.Type())
}

func TestMsgDeregisterResolver_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeregisterResolver
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgDeregisterResolver(nil),
			ErrorInvalidField("from"),
		},
		{
			"from is empty",
			NewMsgDeregisterResolver([]byte("")),
			ErrorInvalidField("from"),
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

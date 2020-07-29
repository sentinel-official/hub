package types

import (
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		m    MsgRegisterProvider
		want sdk.Error
	}{
		{"from nil", NewMsgRegisterProvider(nil, "", "", "", ""), ErrorInvalidField("from")},
		{"from zero", NewMsgRegisterProvider(sdk.AccAddress{}, "", "", "", ""), ErrorInvalidField("from")},
		{"from empty", NewMsgRegisterProvider(sdk.AccAddress(""), "", "", "", ""), ErrorInvalidField("from")},
		{"name length 0", NewMsgRegisterProvider(sdk.AccAddress("address-1"), "", "", "", ""), ErrorInvalidField("name")},
		{"name length greater than 64", NewMsgRegisterProvider(sdk.AccAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), ErrorInvalidField("name")},
		{"identity length greater than 64", NewMsgRegisterProvider(sdk.AccAddress("address-1"), "name", strings.Repeat("-", 64+1), "", ""), ErrorInvalidField("identity")},
		{"website length greater than 64", NewMsgRegisterProvider(sdk.AccAddress("address-1"), "name", "", strings.Repeat("-", 64+1), ""), ErrorInvalidField("website")},
		{"description length greater than 256", NewMsgRegisterProvider(sdk.AccAddress("address-1"), "name", "", "", strings.Repeat("-", 256+1)), ErrorInvalidField("description")},
		{"valid", NewMsgRegisterProvider(sdk.AccAddress("address-1"), "name", "", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUpdateProvider_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		m    MsgUpdateProvider
		want sdk.Error
	}{
		{"from nil", NewMsgUpdateProvider(nil, "", "", "", ""), ErrorInvalidField("from")},
		{"from zero", NewMsgUpdateProvider(hub.ProvAddress{}, "", "", "", ""), ErrorInvalidField("from")},
		{"from empty", NewMsgUpdateProvider(hub.ProvAddress(""), "", "", "", ""), ErrorInvalidField("from")},
		{"name length greater than 64", NewMsgUpdateProvider(hub.ProvAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), ErrorInvalidField("name")},
		{"identity length greater than 64", NewMsgUpdateProvider(hub.ProvAddress("address-1"), "", strings.Repeat("-", 64+1), "", ""), ErrorInvalidField("identity")},
		{"website length greater than 64", NewMsgUpdateProvider(hub.ProvAddress("address-1"), "", "", strings.Repeat("-", 64+1), ""), ErrorInvalidField("website")},
		{"description length greater than 256", NewMsgUpdateProvider(hub.ProvAddress("address-1"), "", "", "", strings.Repeat("-", 256+1)), ErrorInvalidField("description")},
		{"valid", NewMsgUpdateProvider(hub.ProvAddress("address-1"), "", "", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

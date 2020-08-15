package types

import (
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgRegister_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		m    MsgRegister
		want sdk.Error
	}{
		{"from nil", NewMsgRegister(nil, "", "", "", ""), ErrorInvalidField("from")},
		{"from zero", NewMsgRegister(sdk.AccAddress{}, "", "", "", ""), ErrorInvalidField("from")},
		{"from empty", NewMsgRegister(sdk.AccAddress(""), "", "", "", ""), ErrorInvalidField("from")},
		{"name length zero", NewMsgRegister(sdk.AccAddress("address-1"), "", "", "", ""), ErrorInvalidField("name")},
		{"name length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), ErrorInvalidField("name")},
		{"identity length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), "name", strings.Repeat("-", 64+1), "", ""), ErrorInvalidField("identity")},
		{"website length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), "name", "", strings.Repeat("-", 64+1), ""), ErrorInvalidField("website")},
		{"description length greater than 256", NewMsgRegister(sdk.AccAddress("address-1"), "name", "", "", strings.Repeat("-", 256+1)), ErrorInvalidField("description")},
		{"valid", NewMsgRegister(sdk.AccAddress("address-1"), "name", "", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgUpdate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		m    MsgUpdate
		want sdk.Error
	}{
		{"from nil", NewMsgUpdate(nil, "", "", "", ""), ErrorInvalidField("from")},
		{"from zero", NewMsgUpdate(hub.ProvAddress{}, "", "", "", ""), ErrorInvalidField("from")},
		{"from empty", NewMsgUpdate(hub.ProvAddress(""), "", "", "", ""), ErrorInvalidField("from")},
		{"name length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), ErrorInvalidField("name")},
		{"identity length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), "", strings.Repeat("-", 64+1), "", ""), ErrorInvalidField("identity")},
		{"website length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), "", "", strings.Repeat("-", 64+1), ""), ErrorInvalidField("website")},
		{"description length greater than 256", NewMsgUpdate(hub.ProvAddress("address-1"), "", "", "", strings.Repeat("-", 256+1)), ErrorInvalidField("description")},
		{"valid", NewMsgUpdate(hub.ProvAddress("address-1"), "", "", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

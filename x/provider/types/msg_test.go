package types

import (
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgRegister_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		m    MsgRegister
		want error
	}{
		{"from nil", NewMsgRegister(nil, "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"from zero", NewMsgRegister(sdk.AccAddress{}, "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"from empty", NewMsgRegister(sdk.AccAddress(""), "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"name length zero", NewMsgRegister(sdk.AccAddress("address-1"), "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "name")},
		{"name length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "name")},
		{"identity length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), "name", strings.Repeat("-", 64+1), "", ""), errors.Wrapf(ErrorInvalidField, "%s", "identity")},
		{"website length greater than 64", NewMsgRegister(sdk.AccAddress("address-1"), "name", "", strings.Repeat("-", 64+1), ""), errors.Wrapf(ErrorInvalidField, "%s", "website")},
		{"description length greater than 256", NewMsgRegister(sdk.AccAddress("address-1"), "name", "", "", strings.Repeat("-", 256+1)), errors.Wrapf(ErrorInvalidField, "%s", "description")},
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
		want error
	}{
		{"from nil", NewMsgUpdate(nil, "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"from zero", NewMsgUpdate(hub.ProvAddress{}, "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"from empty", NewMsgUpdate(hub.ProvAddress(""), "", "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "from")},
		{"name length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""), errors.Wrapf(ErrorInvalidField, "%s", "name")},
		{"identity length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), "", strings.Repeat("-", 64+1), "", ""), errors.Wrapf(ErrorInvalidField, "%s", "identity")},
		{"website length greater than 64", NewMsgUpdate(hub.ProvAddress("address-1"), "", "", strings.Repeat("-", 64+1), ""), errors.Wrapf(ErrorInvalidField, "%s", "website")},
		{"description length greater than 256", NewMsgUpdate(hub.ProvAddress("address-1"), "", "", "", strings.Repeat("-", 256+1)), errors.Wrapf(ErrorInvalidField, "%s", "description")},
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

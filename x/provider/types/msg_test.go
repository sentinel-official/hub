package types

import (
	"errors"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	GT64  = strings.Repeat("sentinel", 9)
	GT256 = strings.Repeat("sentinel", 33)
)

func TestMsgRegister_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("failed: %s\n", err)
	}

	tests := []struct {
		name string
		m    *MsgRegisterRequest
		want error
	}{
		{"from nil", NewMsgRegisterRequest(nil, "", "", "", ""), ErrorInvalidFieldFrom},
		{"from zero", NewMsgRegisterRequest(sdk.AccAddress{}, "", "", "", ""), ErrorInvalidFieldFrom},
		{"from empty", NewMsgRegisterRequest(sdk.AccAddress(""), "", "", "", ""), ErrorInvalidFieldFrom},
		{"name zero", NewMsgRegisterRequest(correctAddress, "", "", "", ""), ErrorInvalidFieldName},
		{"name GT64", NewMsgRegisterRequest(correctAddress, GT64, "", "", ""), ErrorInvalidFieldName},
		{"identity GT64", NewMsgRegisterRequest(correctAddress, "test-name", GT64, "", ""), ErrorInvalidFieldIdentity},
		{"website GT64", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", GT64, ""), ErrorInvalidFieldWebsite},
		{"description GT256", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "test-website", GT256), ErrorInvalidFieldDescription},
		{"from correct", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"name correct", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"identity empty", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"identity correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"website empty", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"website correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "test-website", ""), nil},
		{"description empty", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"description correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", "test-description"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(tt.want, err) {
				t.Errorf("ValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

func TestMsgUpdate_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("failed: %s\n", err)
	}

	tests := []struct {
		name string
		m    *MsgRegisterRequest
		want error
	}{
		{"from nil", NewMsgRegisterRequest(nil, "", "", "", ""), ErrorInvalidFieldFrom},
		{"from zero", NewMsgRegisterRequest(sdk.AccAddress{}, "", "", "", ""), ErrorInvalidFieldFrom},
		{"from empty", NewMsgRegisterRequest(sdk.AccAddress(""), "", "", "", ""), ErrorInvalidFieldFrom},
		{"name zero", NewMsgRegisterRequest(correctAddress, "", "", "", ""), ErrorInvalidFieldName},
		{"name GT64", NewMsgRegisterRequest(correctAddress, GT64, "", "", ""), ErrorInvalidFieldName},
		{"identity GT64", NewMsgRegisterRequest(correctAddress, "test-name", GT64, "", ""), ErrorInvalidFieldIdentity},
		{"website GT64", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", GT64, ""), ErrorInvalidFieldWebsite},
		{"description GT256", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "test-website", GT256), ErrorInvalidFieldDescription},
		{"from correct", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"name correct", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"identity empty", NewMsgRegisterRequest(correctAddress, "test-name", "", "", ""), nil},
		{"identity correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"website empty", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"website correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "test-website", ""), nil},
		{"description empty", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", ""), nil},
		{"description correct", NewMsgRegisterRequest(correctAddress, "test-name", "test-identity", "", "test-description"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(tt.want, err) {
				t.Errorf("ValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

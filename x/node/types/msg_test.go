package types

import (
	"errors"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	GT64 = strings.Repeat("sentinel", 9)
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	wrongProvider := hubtypes.ProvAddress(correctAddress.String())

	correctCoins := sdk.NewCoins(sdk.Coin{
		Denom:  "sent",
		Amount: sdk.NewInt(10),
	})

	wrongCoins := sdk.Coins{sdk.Coin{Denom: "--!sent--!", Amount: sdk.NewInt(-1)}}

	tests := []struct {
		name string
		m    *MsgRegisterRequest
		want error
	}{
		{"nil from address", NewMsgRegisterRequest(nil, nil, sdk.Coins{}, ""), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgRegisterRequest(sdk.AccAddress{}, nil, sdk.Coins{}, ""), ErrorInvalidFieldFrom},
		{"nil provider and price", NewMsgRegisterRequest(correctAddress, nil, nil, ""), ErrorInvalidFieldProviderOrPrice},
		{"empty provider and nil price", NewMsgRegisterRequest(correctAddress, hubtypes.ProvAddress{}, nil, ""), ErrorInvalidFieldProviderOrPrice},
		{"invalid provider", NewMsgRegisterRequest(correctAddress, wrongProvider, nil, ""), ErrorInvalidFieldProvider},
		{"wrong price", NewMsgRegisterRequest(correctAddress, hubtypes.ProvAddress{}, wrongCoins, ""), ErrorInvalidFieldPrice},
		{"empty remote_url", NewMsgRegisterRequest(correctAddress, nil, correctCoins, ""), ErrorInvalidFieldRemoteURL},
		{"remote_url GT64", NewMsgRegisterRequest(correctAddress, nil, correctCoins, GT64), ErrorInvalidFieldRemoteURL},
		{"valid", NewMsgRegisterRequest(correctAddress, nil, correctCoins, "https://sentinel.co"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgUpdateRequest_ValidateBasic(t *testing.T) {
	nodeAddress, err := hubtypes.NodeAddressFromBech32("sentnode1grdunxx5jxd0ja75wt508sn6v39p70hhczsm43")
	if err != nil {
		t.Errorf("invalid provider address: %s", err)
	}

	provAddress, err := hubtypes.ProvAddressFromBech32("sentprov1grdunxx5jxd0ja75wt508sn6v39p70hhxrdetl")
	if err != nil {
		t.Errorf("invalid provider address: %s", err)
	}

	wrongProvAddress := hubtypes.ProvAddress("sentwrongprov1grdunxx5jxd0ja75wt508sn6v39p70hhxrdetl")

	correctCoins := sdk.NewCoins(sdk.Coin{
		Denom:  "sent",
		Amount: sdk.NewInt(10),
	})

	wrongCoins := sdk.Coins{sdk.Coin{Denom: "--!sent--!", Amount: sdk.NewInt(-1)}}

	tests := []struct {
		name string
		m    *MsgUpdateRequest
		want error
	}{
		{"nil from address", NewMsgUpdateRequest(nil, nil, sdk.Coins{}, ""), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgUpdateRequest(hubtypes.NodeAddress{}, nil, sdk.Coins{}, ""), ErrorInvalidFieldFrom},
		{"non-empty provider and price", NewMsgUpdateRequest(nodeAddress, provAddress, correctCoins, ""), ErrorInvalidFieldProviderOrPrice},
		{"wrong provider", NewMsgUpdateRequest(nodeAddress, wrongProvAddress, nil, ""), ErrorInvalidFieldProvider},
		{"wrong price", NewMsgUpdateRequest(nodeAddress, hubtypes.ProvAddress{}, wrongCoins, ""), ErrorInvalidFieldPrice},
		{"remote_url GT64", NewMsgUpdateRequest(nodeAddress, hubtypes.ProvAddress{}, correctCoins, GT64), ErrorInvalidFieldRemoteURL},
		{"valid", NewMsgUpdateRequest(nodeAddress, hubtypes.ProvAddress{}, correctCoins, "https://sentinel.co"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgSetStatusRequest_ValidateBasic(t *testing.T) {

	nodeAddress, err := hubtypes.NodeAddressFromBech32("sentnode1grdunxx5jxd0ja75wt508sn6v39p70hhczsm43")
	if err != nil {
		t.Errorf("invalid provider address: %s", err)
	}

	status := hubtypes.StatusInactivePending

	tests := []struct {
		name string
		m    *MsgSetStatusRequest
		want error
	}{
		{"nil from address", NewMsgSetStatusRequest(nil, status), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgSetStatusRequest(hubtypes.NodeAddress{}, status), ErrorInvalidFieldFrom},
		{"wrong status", NewMsgSetStatusRequest(nodeAddress, status), ErrorInvalidFieldStatus},
		{"valid", NewMsgSetStatusRequest(nodeAddress, hubtypes.Active), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

package types

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgAddQuotaRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	tests := []struct {
		name string
		m    *MsgAddQuotaRequest
		want error
	}{
		{"nil from address", NewMsgAddQuotaRequest(nil, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgAddQuotaRequest(sdk.AccAddress{}, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"zero id", NewMsgAddQuotaRequest(correctAddress, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldId},
		{"nil address", NewMsgAddQuotaRequest(correctAddress, 10, nil, sdk.NewInt(0)), ErrorInvalidFieldAddress},
		{"empty address", NewMsgAddQuotaRequest(correctAddress, 10, sdk.AccAddress{}, sdk.NewInt(0)), ErrorInvalidFieldAddress},
		{"empty bytes", NewMsgAddQuotaRequest(correctAddress, 10, sdk.AccAddress{}, sdk.Int{}), ErrorInvalidFieldAddress},
		{"zero bytes", NewMsgAddQuotaRequest(correctAddress, 10, correctAddress, sdk.NewInt(0)), ErrorInvalidFieldBytes},
		{"valid", NewMsgAddQuotaRequest(correctAddress, 10, correctAddress, sdk.NewInt(1000)), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgUpdateQuotaRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	tests := []struct {
		name string
		m    *MsgUpdateQuotaRequest
		want error
	}{
		{"nil from address", NewMsgUpdateQuotaRequest(nil, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgUpdateQuotaRequest(sdk.AccAddress{}, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"zero id", NewMsgUpdateQuotaRequest(correctAddress, 0, nil, sdk.NewInt(0)), ErrorInvalidFieldId},
		{"nil address", NewMsgUpdateQuotaRequest(correctAddress, 10, nil, sdk.NewInt(0)), ErrorInvalidFieldAddress},
		{"empty address", NewMsgUpdateQuotaRequest(correctAddress, 10, sdk.AccAddress{}, sdk.NewInt(0)), ErrorInvalidFieldAddress},
		{"empty bytes", NewMsgUpdateQuotaRequest(correctAddress, 10, sdk.AccAddress{}, sdk.Int{}), ErrorInvalidFieldAddress},
		{"zero bytes", NewMsgUpdateQuotaRequest(correctAddress, 10, correctAddress, sdk.NewInt(0)), ErrorInvalidFieldBytes},
		{"valid", NewMsgUpdateQuotaRequest(correctAddress, 10, correctAddress, sdk.NewInt(1000)), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgCancelRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	tests := []struct {
		name string
		m    *MsgCancelRequest
		want error
	}{
		{"nil from address", NewMsgCancelRequest(nil, 0), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgCancelRequest(sdk.AccAddress{}, 0), ErrorInvalidFieldFrom},
		{"zero id", NewMsgCancelRequest(correctAddress, 0), ErrorInvalidFieldId},
		{"valid", NewMsgCancelRequest(correctAddress, 10), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgSubscribeToNodeRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	nodeAddress, err := hubtypes.NodeAddressFromBech32("sentnode1grdunxx5jxd0ja75wt508sn6v39p70hhczsm43")
	if err != nil {
		t.Errorf("invalid node address: %s", err)
	}

	coinWithWrongDenom := sdk.Coin{
		Denom:  "wrongdenom",
		Amount: sdk.NewInt(0),
	}
	coinWithWrongAmount := sdk.Coin{
		Denom:  "sent",
		Amount: sdk.NewInt(-10),
	}
	validCoin := sdk.Coin{
		Denom:  "sent",
		Amount: sdk.NewInt(1000),
	}

	tests := []struct {
		name string
		m    *MsgSubscribeToNodeRequest
		want error
	}{
		{"nil from address", NewMsgSubscribeToNodeRequest(nil, nil, sdk.Coin{}), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgSubscribeToNodeRequest(sdk.AccAddress{}, nil, sdk.Coin{}), ErrorInvalidFieldFrom},
		{"nil address", NewMsgSubscribeToNodeRequest(correctAddress, nil, sdk.Coin{}), ErrorInvalidFieldAddress},
		{"empty address", NewMsgSubscribeToNodeRequest(correctAddress, hubtypes.NodeAddress{}, sdk.Coin{}), ErrorInvalidFieldAddress},
		{"empty coin", NewMsgSubscribeToNodeRequest(correctAddress, nodeAddress, sdk.Coin{}), ErrorInvalidFieldDeposit},
		{"wrong coin denom", NewMsgSubscribeToNodeRequest(correctAddress, nodeAddress, coinWithWrongDenom), ErrorInvalidFieldDeposit},
		{"wrong coin amount", NewMsgSubscribeToNodeRequest(correctAddress, nodeAddress, coinWithWrongAmount), ErrorInvalidFieldDeposit},
		{"valid", NewMsgSubscribeToNodeRequest(correctAddress, nodeAddress, validCoin), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

func TestMsgSubscribeToPlanRequest_ValidateBasic(t *testing.T) {
	correctAddress, err := sdk.AccAddressFromBech32("sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8")
	if err != nil {
		t.Errorf("invalid address: %s", err)
	}

	tests := []struct {
		name string
		m    *MsgSubscribeToPlanRequest
		want error
	}{
		{"nil from address", NewMsgSubscribeToPlanRequest(nil, 0, ""), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgSubscribeToPlanRequest(sdk.AccAddress{}, 0, ""), ErrorInvalidFieldFrom},
		{"zero id", NewMsgSubscribeToPlanRequest(correctAddress, 0, ""), ErrorInvalidFieldId},
		{"wrong denom", NewMsgSubscribeToPlanRequest(correctAddress, 10, "!!wrongdenom"), ErrorInvalidFieldDenom},
		{"valid", NewMsgSubscribeToPlanRequest(correctAddress, 10, "wrongdenom"), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %s, want %s", err, tt.want)
			}
		})
	}
}

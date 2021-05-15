package types

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgSwapRequest_ValidateBasic(t *testing.T) {

	correctAddress, err := sdk.AccAddressFromBech32("cosmos1hlzx7raug6wea9fs955nkzwe2xaa65gnpmddal")
	if err != nil {
		t.Errorf("invalid address provided: %s", err)
	}

	tests := []struct {
		name string
		m    *MsgSwapRequest
		want error
	}{
		{"nil from address", NewMsgSwapRequest(nil, EthereumHash{}, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"empty from address", NewMsgSwapRequest(sdk.AccAddress{}, EthereumHash{}, nil, sdk.NewInt(0)), ErrorInvalidFieldFrom},
		{"nil receiver address", NewMsgSwapRequest(correctAddress, EthereumHash{}, nil, sdk.NewInt(0)), ErrorInvalidFieldReceiver},
		{"empty receiver address", NewMsgSwapRequest(correctAddress, EthereumHash{}, sdk.AccAddress{}, sdk.NewInt(0)), ErrorInvalidFieldReceiver},
		{"invalid amount", NewMsgSwapRequest(correctAddress, EthereumHash{}, correctAddress, sdk.NewInt(10)), ErrorInvalidFieldAmount},
		{"valid", NewMsgSwapRequest(correctAddress, EthereumHash{}, correctAddress, sdk.NewInt(1000)), nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.ValidateBasic(); !errors.Is(err, tt.want) {
				t.Errorf("ValidateBasic() = %v, want %v", err, tt.want)
			}
		})
	}
}

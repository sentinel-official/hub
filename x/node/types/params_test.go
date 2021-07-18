package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit          sdk.Coin
		InactiveDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty denom deposit",
			fields{
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"invalid denom deposit",
			fields{
				Deposit: sdk.Coin{Denom: "0", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"negative amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"negative inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: -1000,
			},
			true,
		},
		{
			"zero inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 0,
			},
			true,
		},
		{
			"positive inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Params{
				Deposit:          tt.fields.Deposit,
				InactiveDuration: tt.fields.InactiveDuration,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

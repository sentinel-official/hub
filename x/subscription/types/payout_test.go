package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestPayout_Validate(t *testing.T) {
	type fields struct {
		ID     uint64
		Hours  int64
		Price  sdk.Coin
		NextAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"id zero",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"id positive",
			fields{
				ID:     1000,
				Hours:  1000,
				NextAt: time.Now(),
			},
			false,
		},
		{
			"hours negative",
			fields{
				ID:    1000,
				Hours: -1000,
			},
			true,
		},
		{
			"hours zero",
			fields{
				ID:    1000,
				Hours: 0,
			},
			true,
		},
		{
			"hours positive",
			fields{
				ID:     1000,
				Hours:  1000,
				NextAt: time.Now(),
			},
			false,
		},
		{
			"price empty",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  sdk.Coin{},
				NextAt: time.Now(),
			},
			false,
		},
		{
			"price empty denom",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  sdk.Coin{Denom: ""},
				NextAt: time.Now(),
			},
			false,
		},
		{
			"price invalid denom",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"price empty amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"price negative amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"price zero amount",
			fields{
				ID:    1000,
				Hours: 1000,
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"price positive amount",
			fields{
				ID:     1000,
				Hours:  1000,
				Price:  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				NextAt: time.Now(),
			},
			false,
		},
		{
			"next_at zero",
			fields{
				ID:     1000,
				Hours:  1000,
				NextAt: time.Time{},
			},
			true,
		},
		{
			"next_at positive",
			fields{
				ID:     1000,
				Hours:  1000,
				NextAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Payout{
				ID:     tt.fields.ID,
				Hours:  tt.fields.Hours,
				Price:  tt.fields.Price,
				NextAt: tt.fields.NextAt,
			}
			if err := l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

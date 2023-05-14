package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestLease_Validate(t *testing.T) {
	type fields struct {
		ID      uint64
		Bytes   sdk.Int
		Minutes int64
		Price   sdk.Coin
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
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			false,
		},
		{
			"bytes empty and minutes empty",
			fields{
				ID:      1000,
				Bytes:   sdk.Int{},
				Minutes: 0,
			},
			true,
		},
		{
			"bytes negative and minutes negative",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(-1000),
				Minutes: -1000,
			},
			true,
		},
		{
			"bytes zero and minutes zero",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(0),
				Minutes: 0,
			},
			true,
		},
		{
			"bytes positive and minutes positive",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 1000,
			},
			true,
		},
		{
			"bytes negative",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"bytes zero",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(0),
			},
			true,
		},
		{
			"bytes positive",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(1000),
			},
			false,
		},
		{
			"minutes negative",
			fields{
				ID:      1000,
				Bytes:   sdk.Int{},
				Minutes: -1000,
			},
			true,
		},
		{
			"minutes zero",
			fields{
				ID:      1000,
				Bytes:   sdk.Int{},
				Minutes: 0,
			},
			true,
		},
		{
			"minutes positive",
			fields{
				ID:      1000,
				Bytes:   sdk.Int{},
				Minutes: 1000,
			},
			false,
		},
		{
			"price empty",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{},
			},
			false,
		},
		{
			"price empty denom",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: ""},
			},
			false,
		},
		{
			"price invalid denom",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"price empty amount",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"price negative amount",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"price zero amount",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"price positive amount",
			fields{
				ID:      1000,
				Bytes:   sdk.NewInt(1000),
				Minutes: 0,
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lease{
				ID:      tt.fields.ID,
				Bytes:   tt.fields.Bytes,
				Minutes: tt.fields.Minutes,
				Price:   tt.fields.Price,
			}
			if err := l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestLease_Validate(t *testing.T) {
	type fields struct {
		ID             uint64
		Bytes          sdk.Int
		Duration       int64
		Price          sdk.Coin
		DistributionAt time.Time
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
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"bytes empty and duration empty",
			fields{
				ID:       1000,
				Bytes:    sdk.Int{},
				Duration: 0,
			},
			true,
		},
		{
			"bytes negative and duration negative",
			fields{
				ID:       1000,
				Bytes:    sdk.NewInt(-1000),
				Duration: -1000,
			},
			true,
		},
		{
			"bytes zero and duration zero",
			fields{
				ID:       1000,
				Bytes:    sdk.NewInt(0),
				Duration: 0,
			},
			true,
		},
		{
			"bytes positive and duration positive",
			fields{
				ID:       1000,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
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
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"duration negative",
			fields{
				ID:       1000,
				Duration: -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				ID:       1000,
				Duration: 0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				ID:             1000,
				Duration:       1000,
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"price empty",
			fields{
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"price empty denom",
			fields{
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				Price:          sdk.Coin{Denom: ""},
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"price invalid denom",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(1000),
				Price: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"price empty amount",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(1000),
				Price: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"price negative amount",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(1000),
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"price zero amount",
			fields{
				ID:    1000,
				Bytes: sdk.NewInt(1000),
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"price positive amount",
			fields{
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				Price:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				DistributionAt: time.Now(),
			},
			false,
		},
		{
			"distribution_at zero",
			fields{
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				DistributionAt: time.Time{},
			},
			true,
		},
		{
			"distribution_at positive",
			fields{
				ID:             1000,
				Bytes:          sdk.NewInt(1000),
				DistributionAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Lease{
				ID:             tt.fields.ID,
				Bytes:          tt.fields.Bytes,
				Duration:       tt.fields.Duration,
				Price:          tt.fields.Price,
				DistributionAt: tt.fields.DistributionAt,
			}
			if err := l.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

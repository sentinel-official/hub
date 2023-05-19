package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestPlan_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.ProvAddress
	}{
		{
			"empty",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			hubtypes.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Address: tt.fields.Address,
			}
			if got := p.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlan_Price(t *testing.T) {
	type fields struct {
		Prices sdk.Coins
	}
	type args struct {
		denom string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   sdk.Coin
		want1  bool
	}{
		{
			"nil prices and empty denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			true,
		},
		{
			"nil prices and one denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"nil prices and two denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"empty prices and empty denom",
			fields{
				Prices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"empty prices and one denom",
			fields{
				Prices: sdk.Coins{},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"empty prices and two denom",
			fields{
				Prices: sdk.Coins{},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"1one prices and empty denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"1one prices and one denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Denom: "one", Amount: sdk.NewInt(1)},
			true,
		},
		{
			"1one prices and two denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"1two prices and empty denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "two", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"1two prices and one denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "two", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"1two prices and two denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "two", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Denom: "two", Amount: sdk.NewInt(1)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Prices: tt.fields.Prices,
			}
			got, got1 := p.Price(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Price() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Price() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPlan_Validate(t *testing.T) {
	type fields struct {
		ID       uint64
		Address  string
		Bytes    sdk.Int
		Duration time.Duration
		Prices   sdk.Coins
		Status   hubtypes.Status
		StatusAt time.Time
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
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Duration: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address empty",
			fields{
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgsutj8xr",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"bytes empty",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:   sdk.Int{},
			},
			true,
		},
		{
			"bytes negative",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:   sdk.NewInt(-1000),
			},
			true,
		},
		{
			"bytes zero",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:   sdk.NewInt(0),
			},
			true,
		},
		{
			"bytes positive",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"duration negative",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"prices nil",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   nil,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"prices empty",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{},
			},
			true,
		},
		{
			"prices empty denom",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"prices empty amount",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"prices invalid denom",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"prices negative amount",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"prices zero amount",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"prices positive amount",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status unspecified",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status inactive pending",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusInactive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status_at zero",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Time{},
			},
			true,
		},
		{
			"status_at positive",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				ID:       tt.fields.ID,
				Address:  tt.fields.Address,
				Bytes:    tt.fields.Bytes,
				Duration: tt.fields.Duration,
				Prices:   tt.fields.Prices,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

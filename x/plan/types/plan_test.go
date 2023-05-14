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
			sdk.Coin{},
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
			sdk.Coin{},
			true,
		},
		{
			"nil prices and two denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "two",
			},
			sdk.Coin{},
			true,
		},
		{
			"empty prices and empty denom",
			fields{
				Prices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
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
			sdk.Coin{},
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
			sdk.Coin{},
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
			sdk.Coin{},
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
			sdk.Coin{},
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
			sdk.Coin{},
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
			sdk.Coin{},
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
		Prices   sdk.Coins
		Validity time.Duration
		Bytes    sdk.Int
		Status   hubtypes.Status
		StatusAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero id",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"positive id",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"empty provider",
			fields{
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid provider",
			fields{
				ID:      1000,
				Address: "sentprov",
			},
			true,
		},
		{
			"invalid prefix provider",
			fields{
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes provider",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgsutj8xr",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"20 bytes provider",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"30 bytes provider",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"nil price",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   nil,
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"empty price",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:  sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:  sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:  sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				ID:      1000,
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"negative validity",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: -1000,
			},
			true,
		},
		{
			"zero validity",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 0,
			},
			true,
		},
		{
			"positive validity",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"negative bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"positive bytes",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"unspecified status",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"zero status_at",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Time{},
			},
			true,
		},
		{
			"positive status_at",
			fields{
				ID:       1000,
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
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
				Prices:   tt.fields.Prices,
				Validity: tt.fields.Validity,
				Bytes:    tt.fields.Bytes,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

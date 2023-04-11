package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestPlan_GetProviderAddr(t *testing.T) {
	type fields struct {
		ID           uint64
		ProviderAddr string
		Prices       sdk.Coins
		Validity     time.Duration
		Bytes        sdk.Int
		Status       hubtypes.Status
		StatusAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.ProvAddress
	}{
		{
			"empty",
			fields{
				ProviderAddr: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			hubtypes.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				ID:           tt.fields.ID,
				ProviderAddr: tt.fields.ProviderAddr,
				Prices:       tt.fields.Prices,
				Validity:     tt.fields.Validity,
				Bytes:        tt.fields.Bytes,
				Status:       tt.fields.Status,
				StatusAt:     tt.fields.StatusAt,
			}
			if got := p.GetProviderAddr(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProviderAddr() = %v, want %v", got, tt.want)
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
			"nil price and empty denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"nil price and one denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "one",
			},
			sdk.Coin{},
			false,
		},
		{
			"nil price and two denom",
			fields{
				Prices: nil,
			},
			args{
				denom: "two",
			},
			sdk.Coin{},
			false,
		},
		{
			"empty price and empty denom",
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
			"empty price and one denom",
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
			"empty price and two denom",
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
			"1one price and empty denom",
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
			"1one price and one denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "one",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one price and two denom",
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
			"1two price and empty denom",
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
			"1two price and one denom",
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
			"1two price and two denom",
			fields{
				Prices: sdk.Coins{sdk.Coin{Denom: "two", Amount: sdk.NewInt(1)}},
			},
			args{
				denom: "two",
			},
			sdk.Coin{},
			false,
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
		ID           uint64
		ProviderAddr string
		Prices       sdk.Coins
		Validity     time.Duration
		Bytes        sdk.Int
		Status       hubtypes.Status
		StatusAt     time.Time
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
			false,
		},
		{
			"positive id",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"empty provider",
			fields{
				ID:           1000,
				ProviderAddr: "",
			},
			false,
		},
		{
			"invalid provider",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov",
			},
			false,
		},
		{
			"invalid prefix provider",
			fields{
				ID:           1000,
				ProviderAddr: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			false,
		},
		{
			"10 bytes provider",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgsutj8xr",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"20 bytes provider",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"30 bytes provider",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"nil price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       nil,
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"empty price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{},
			},
			false,
		},
		{
			"empty denom price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: ""}},
			},
			false,
		},
		{
			"invalid denom price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			false,
		},
		{
			"negative amount price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			false,
		},
		{
			"zero amount price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			false,
		},
		{
			"positive amount price",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"negative validity",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     -1000,
			},
			false,
		},
		{
			"zero validity",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     0,
			},
			false,
		},
		{
			"positive validity",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"negative bytes",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(-1000),
			},
			false,
		},
		{
			"zero bytes",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(0),
			},
			false,
		},
		{
			"positive bytes",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"unspecified status",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusUnspecified,
			},
			false,
		},
		{
			"active status",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusInactivePending,
			},
			false,
		},
		{
			"inactive status",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusInactive,
			},
			false,
		},
		{
			"zero status_at",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Time{},
			},
			false,
		},
		{
			"now status_at",
			fields{
				ID:           1000,
				ProviderAddr: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:       sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity:     1000,
				Bytes:        sdk.NewInt(1000),
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				ID:           tt.fields.ID,
				ProviderAddr: tt.fields.ProviderAddr,
				Prices:       tt.fields.Prices,
				Validity:     tt.fields.Validity,
				Bytes:        tt.fields.Bytes,
				Status:       tt.fields.Status,
				StatusAt:     tt.fields.StatusAt,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

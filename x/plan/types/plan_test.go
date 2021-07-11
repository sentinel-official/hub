package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestPlan_GetProvider(t *testing.T) {
	type fields struct {
		Id       uint64
		Provider string
		Price    sdk.Coins
		Validity time.Duration
		Bytes    sdk.Int
		Status   hubtypes.Status
		StatusAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.ProvAddress
	}{
		{
			"empty",
			fields{
				Provider: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			hubtypes.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Id:       tt.fields.Id,
				Provider: tt.fields.Provider,
				Price:    tt.fields.Price,
				Validity: tt.fields.Validity,
				Bytes:    tt.fields.Bytes,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			if got := p.GetProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlan_PriceForDenom(t *testing.T) {
	type fields struct {
		Id       uint64
		Provider string
		Price    sdk.Coins
		Validity time.Duration
		Bytes    sdk.Int
		Status   hubtypes.Status
		StatusAt time.Time
	}
	type args struct {
		d string
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
				Price: nil,
			},
			args{
				d: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"empty price and empty denom",
			fields{
				Price: sdk.Coins{},
			},
			args{
				d: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one price and empty denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				d: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one price and one denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				d: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one price and two denom",
			fields{
				Price: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				d: "two",
			},
			sdk.Coin{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Id:       tt.fields.Id,
				Provider: tt.fields.Provider,
				Price:    tt.fields.Price,
				Validity: tt.fields.Validity,
				Bytes:    tt.fields.Bytes,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			got, got1 := p.PriceForDenom(tt.args.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriceForDenom() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PriceForDenom() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPlan_Validate(t *testing.T) {
	type fields struct {
		Id       uint64
		Provider string
		Price    sdk.Coins
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
				Id: 0,
			},
			true,
		},
		{
			"empty provider",
			fields{
				Id:       1000,
				Provider: "",
			},
			true,
		},
		{
			"invalid provider",
			fields{
				Id:       1000,
				Provider: "invalid",
			},
			true,
		},
		{
			"invalid prefix provider",
			fields{
				Id:       1000,
				Provider: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes provider",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes provider",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes provider",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"nil price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    nil,
			},
			true,
		},
		{
			"empty price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"negative validity",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: -1000,
			},
			true,
		},
		{
			"zero validity",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 0,
			},
			true,
		},
		{
			"positive validity",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"negative bytes",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"positive bytes",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			true,
		},
		{
			"unknown status",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusUnknown,
			},
			true,
		},
		{
			"active status",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
			},
			true,
		},
		{
			"inactive pending status",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactive,
			},
			true,
		},
		{
			"zero status_at",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactive,
				StatusAt: time.Time{},
			},
			true,
		},
		{
			"now status_at",
			fields{
				Id:       1000,
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
				Status:   hubtypes.StatusInactive,
				StatusAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plan{
				Id:       tt.fields.Id,
				Provider: tt.fields.Provider,
				Price:    tt.fields.Price,
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

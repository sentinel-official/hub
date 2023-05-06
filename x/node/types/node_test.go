package types

import (
	"reflect"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestNode_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
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
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address: tt.fields.Address,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GigabytePrice(t *testing.T) {
	type fields struct {
		GigabytePrices sdk.Coins
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
			"nil gigabyte_prices and empty denom",
			fields{
				GigabytePrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			true,
		},
		{
			"empty gigabyte_prices and empty denom",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one gigabyte_prices and empty denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one gigabyte_prices and one denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one gigabyte_prices and two denom",
			fields{
				GigabytePrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
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
			m := &Node{
				GigabytePrices: tt.fields.GigabytePrices,
			}
			got, got1 := m.GigabytePrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GigabytePrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GigabytePrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_HourlyPrice(t *testing.T) {
	type fields struct {
		HourlyPrices sdk.Coins
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
			"nil hourly_prices and empty denom",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			true,
		},
		{
			"empty hourly_prices and empty denom",
			fields{
				HourlyPrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one hourly_prices and empty denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "",
			},
			sdk.Coin{},
			false,
		},
		{
			"1one hourly_prices and one denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
			},
			args{
				denom: "one",
			},
			sdk.NewInt64Coin("one", 1),
			true,
		},
		{
			"1one hourly_prices and two denom",
			fields{
				HourlyPrices: sdk.Coins{sdk.NewInt64Coin("one", 1)},
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
			m := &Node{
				HourlyPrices: tt.fields.HourlyPrices,
			}
			got, got1 := m.HourlyPrice(tt.args.denom)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HourlyPrice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HourlyPrice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNode_Validate(t *testing.T) {
	type fields struct {
		Address        string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
		Status         hubtypes.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				Address: "sentnode",
			},
			true,
		},
		{
			"invalid address prefix",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgse4wwrm",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"nil gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"empty gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive gigabyte_prices",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"nil hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: nil,
				RemoteURL:    "https://remote.url:443",
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"empty hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive hourly_prices",
			fields{
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:    "https://remote.url:443",
				Status:       hubtypes.StatusActive,
				StatusAt:     time.Now(),
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"unspecified status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"inactive_pending status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusInactive,
				StatusAt:  time.Now(),
			},
			false,
		},
		{
			"zero status_at",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Time{},
			},
			true,
		},
		{
			"positive status_at",
			fields{
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				RemoteURL: "https://remote.url:443",
				Status:    hubtypes.StatusActive,
				StatusAt:  time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Node{
				Address:        tt.fields.Address,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

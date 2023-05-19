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
			"address empty",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"address 20 bytes",
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
			"gigabyte_prices nil and denom empty",
			fields{
				GigabytePrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			true,
		},
		{
			"gigabyte_prices empty and denom empty",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices 1000one and denom empty",
			fields{
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices nil and denom one",
			fields{
				GigabytePrices: nil,
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices empty and denom one",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices 1000one and denom one",
			fields{
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			true,
		},
		{
			"gigabyte_prices nil and denom two",
			fields{
				GigabytePrices: nil,
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices empty and denom two",
			fields{
				GigabytePrices: sdk.Coins{},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"gigabyte_prices 1000one and denom two",
			fields{
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
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
			"hourly_prices nil and denom empty",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			true,
		},
		{
			"hourly_prices empty and denom empty",
			fields{
				HourlyPrices: sdk.Coins{},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices 1000one and denom empty",
			fields{
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices nil and denom one",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices empty and denom one",
			fields{
				HourlyPrices: sdk.Coins{},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices 1000one and denom one",
			fields{
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "one",
			},
			sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			true,
		},
		{
			"hourly_prices nil and denom two",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices empty and denom two",
			fields{
				HourlyPrices: sdk.Coins{},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
			false,
		},
		{
			"hourly_prices 1000one and denom two",
			fields{
				HourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			args{
				denom: "two",
			},
			sdk.Coin{Amount: sdk.NewInt(0)},
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
		ExpiryAt       time.Time
		Status         hubtypes.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"address empty",
			fields{
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgse4wwrm",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"expiry_at empty",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Time{},
			},
			true,
		},
		{
			"expiry_at now",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status unspecified",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Time{},
				Status:         hubtypes.StatusInactive,
				StatusAt:       time.Now(),
			},
			false,
		},
		{
			"status_at empty",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Time{},
			},
			true,
		},
		{
			"status_at now",
			fields{
				Address:        "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
				ExpiryAt:       time.Now(),
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
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
				ExpiryAt:       tt.fields.ExpiryAt,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

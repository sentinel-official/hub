package types

import (
	"reflect"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
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
				Address: hubtypes.TestAddrEmpty,
			},
			nil,
		},
		{
			"address 20 bytes",
			fields{
				Address: hubtypes.TestBech32NodeAddr20Bytes,
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
				GigabytePrices: hubtypes.TestCoinsNil,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices empty and denom empty",
			fields{
				GigabytePrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices 1000one and denom empty",
			fields{
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices nil and denom one",
			fields{
				GigabytePrices: hubtypes.TestCoinsNil,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices empty and denom one",
			fields{
				GigabytePrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices 1000one and denom one",
			fields{
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinPositiveAmount,
			true,
		},
		{
			"gigabyte_prices nil and denom two",
			fields{
				GigabytePrices: hubtypes.TestCoinsNil,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices empty and denom two",
			fields{
				GigabytePrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"gigabyte_prices 1000one and denom two",
			fields{
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
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
				HourlyPrices: hubtypes.TestCoinsNil,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices empty and denom empty",
			fields{
				HourlyPrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices 1000one and denom empty",
			fields{
				HourlyPrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomEmpty,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices nil and denom one",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices empty and denom one",
			fields{
				HourlyPrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices 1000one and denom one",
			fields{
				HourlyPrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomOne,
			},
			hubtypes.TestCoinPositiveAmount,
			true,
		},
		{
			"hourly_prices nil and denom two",
			fields{
				HourlyPrices: nil,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices empty and denom two",
			fields{
				HourlyPrices: hubtypes.TestCoinsEmpty,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
			false,
		},
		{
			"hourly_prices 1000one and denom two",
			fields{
				HourlyPrices: hubtypes.TestCoinsPositiveAmount,
			},
			args{
				denom: hubtypes.TestDenomTwo,
			},
			hubtypes.TestCoinEmpty,
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
		InactiveAt     time.Time
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
				Address: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"address invalid",
			fields{
				Address: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				Address:        hubtypes.TestBech32NodeAddr10Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				Address:        hubtypes.TestBech32NodeAddr30Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsNil,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			true,
		},
		{
			"gigabyte_prices empty",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsEmpty,
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsNil,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			true,
		},
		{
			"hourly_prices empty",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsEmpty,
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"inactive_at empty",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeZero,
			},
			true,
		},
		{
			"inactive_at now",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeZero,
				Status:         hubtypes.StatusInactive,
				StatusAt:       hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status_at empty",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeZero,
			},
			true,
		},
		{
			"status_at now",
			fields{
				Address:        hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: hubtypes.TestCoinsPositiveAmount,
				HourlyPrices:   hubtypes.TestCoinsPositiveAmount,
				RemoteURL:      "https://remote.url:443",
				InactiveAt:     hubtypes.TestTimeNow,
				Status:         hubtypes.StatusActive,
				StatusAt:       hubtypes.TestTimeNow,
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
				InactiveAt:     tt.fields.InactiveAt,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

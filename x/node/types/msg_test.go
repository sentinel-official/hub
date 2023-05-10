package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From           string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: "",
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: "invalid",
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           "sent1qypqxpq9qcrsszgslawd5s",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRegisterRequest{
				From:           tt.fields.From,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateDetailsRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From           string
		GigabytePrices sdk.Coins
		HourlyPrices   sdk.Coins
		RemoteURL      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: "",
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: "invalid",
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgse4wwrm",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			false,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateDetailsRequest{
				From:           tt.fields.From,
				GigabytePrices: tt.fields.GigabytePrices,
				HourlyPrices:   tt.fields.HourlyPrices,
				RemoteURL:      tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateStatusRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From   string
		Status hubtypes.Status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: "",
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: "invalid",
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgse4wwrm",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateStatusRequest{
				From:   tt.fields.From,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSubscribeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Address   string
		Hours     int64
		Gigabytes int64
		Denom     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: "",
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: "invalid",
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgslawd5s",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"address empty",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgse4wwrm",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
		{
			"hours negative gigabytes negative",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     -1000,
				Gigabytes: -1000,
			},
			true,
		},
		{
			"hours zero gigabytes zero",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     0,
				Gigabytes: 0,
			},
			true,
		},
		{
			"hours positive gigabytes positive",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 1000,
			},
			true,
		},
		{
			"hours negative",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:   -1000,
			},
			true,
		},
		{
			"hours positive",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:   1000,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     0,
				Gigabytes: -1000,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     0,
				Gigabytes: 1000,
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "",
			},
			false,
		},
		{
			"denom invalid",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "o",
			},
			true,
		},
		{
			"denom valid",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:     1000,
				Gigabytes: 0,
				Denom:     "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:      tt.fields.From,
				Address:   tt.fields.Address,
				Hours:     tt.fields.Hours,
				Gigabytes: tt.fields.Gigabytes,
				Denom:     tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

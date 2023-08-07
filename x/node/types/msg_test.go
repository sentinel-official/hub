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
				From: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           hubtypes.TestBech32AccAddr10Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           hubtypes.TestBech32AccAddr30Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			true,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           hubtypes.TestBech32AccAddr20Bytes,
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:           hubtypes.TestBech32NodeAddr10Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:           hubtypes.TestBech32NodeAddr30Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices nil",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"gigabyte_prices empty",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"gigabyte_prices empty denom",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"gigabyte_prices invalid denom",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"gigabyte_prices empty amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"gigabyte_prices negative amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"gigabyte_prices zero amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"gigabyte_prices positive amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices nil",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"hourly_prices empty",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"hourly_prices empty denom",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"hourly_prices invalid denom",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"hourly_prices empty amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"hourly_prices negative amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"hourly_prices zero amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"hourly_prices positive amount",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"remote_url empty",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			false,
		},
		{
			"remote_url 72 chars",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"remote_url invalid",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"remote_url invalid scheme",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"remote_url without port",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:           hubtypes.TestBech32NodeAddr20Bytes,
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:   hubtypes.TestBech32NodeAddr10Bytes,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   hubtypes.TestBech32NodeAddr20Bytes,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   hubtypes.TestBech32NodeAddr30Bytes,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   hubtypes.TestBech32NodeAddr20Bytes,
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   hubtypes.TestBech32NodeAddr20Bytes,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				From:   hubtypes.TestBech32NodeAddr20Bytes,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   hubtypes.TestBech32NodeAddr20Bytes,
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
		From        string
		NodeAddress string
		Hours       int64
		Gigabytes   int64
		Denom       string
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
				From: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr10Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr30Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: "",
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: "invalid",
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr10Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr30Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
		{
			"hours negative and gigabytes negative",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       -1000,
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"hours zero and gigabytes zero",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   0,
			},
			true,
		},
		{
			"hours positive and gigabytes positive",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   1000,
			},
			true,
		},
		{
			"hours negative",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       -1000,
			},
			true,
		},
		{
			"hours positive",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       0,
				Gigabytes:   1000,
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "",
			},
			false,
		},
		{
			"denom invalid",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "o",
			},
			true,
		},
		{
			"denom valid",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Gigabytes:   0,
				Denom:       "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:        tt.fields.From,
				NodeAddress: tt.fields.NodeAddress,
				Hours:       tt.fields.Hours,
				Gigabytes:   tt.fields.Gigabytes,
				Denom:       tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

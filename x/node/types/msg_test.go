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
			"empty from",
			fields{
				From: "",
			},
			true,
		},
		{
			"invalid from",
			fields{
				From: "sent",
			},
			true,
		},
		{
			"invalid from prefix",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From:           "sent1qypqxpq9qcrsszgslawd5s",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"nil gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive gigabyte_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"nil hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"empty denom hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive hourly_prices",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:           "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
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
			"empty from",
			fields{
				From: "",
			},
			true,
		},
		{
			"invalid from",
			fields{
				From: "sentnode",
			},
			true,
		},
		{
			"invalid from prefix",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgse4wwrm",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"nil gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive gigabyte_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"nil hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{},
			},
			true,
		},
		{
			"empty denom hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive hourly_prices",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL:      "https://remote.url:443",
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "",
			},
			false,
		},
		{
			"length 72 remote_url",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:           "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				GigabytePrices: nil,
				HourlyPrices:   nil,
				RemoteURL:      "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
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
			"empty from",
			fields{
				From: "",
			},
			true,
		},
		{
			"invalid from",
			fields{
				From: "sentnode",
			},
			true,
		},
		{
			"invalid prefix from",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgse4wwrm",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"unspecified status",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
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

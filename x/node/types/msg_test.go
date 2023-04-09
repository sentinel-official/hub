package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From             string
		PricePerGigabyte sdk.Coins
		PricePerHour     sdk.Coins
		RemoteURL        string
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
				From:             "sent1qypqxpq9qcrsszgslawd5s",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"nil price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive price_per_gigabyte",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"nil price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive price_per_hour",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:             "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRegisterRequest{
				From:             tt.fields.From,
				PricePerGigabyte: tt.fields.PricePerGigabyte,
				PricePerHour:     tt.fields.PricePerHour,
				RemoteURL:        tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From             string
		PricePerGigabyte sdk.Coins
		PricePerHour     sdk.Coins
		RemoteURL        string
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
				From:             "sentnode1qypqxpq9qcrsszgse4wwrm",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"nil price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive price_per_gigabyte",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"nil price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive price_per_hour",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "",
			},
			false,
		},
		{
			"length 72 remote_url",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url",
			},
			true,
		},
		{
			"remote_url with port",
			fields{
				From:             "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				PricePerGigabyte: nil,
				PricePerHour:     nil,
				RemoteURL:        "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateRequest{
				From:             tt.fields.From,
				PricePerGigabyte: tt.fields.PricePerGigabyte,
				PricePerHour:     tt.fields.PricePerHour,
				RemoteURL:        tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSetStatusRequest_ValidateBasic(t *testing.T) {
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
			m := &MsgSetStatusRequest{
				From:   tt.fields.From,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

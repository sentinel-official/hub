package types

import (
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
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
				From: "invalid",
			},
			true,
		},
		{
			"invalid prefix from",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From: "sent1qypqxpq9qcrsszgslawd5s",
			},
			true,
		},
		{
			"20 bytes from",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"30 bytes from",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
			},
			true,
		},
		{
			"empty provider and nil price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    nil,
			},
			true,
		},
		{
			"non-empty provider and non-nil price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"invalid prefix provider",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes provider",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes provider",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes provider",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"empty price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty remote_url",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "",
			},
			true,
		},
		{
			"length 72 remote_url",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url",
			},
			true,
		},
		{
			"non-empty remote_url port",
			fields{
				From:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRegisterRequest{
				From:      tt.fields.From,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Provider  string
		Price     sdk.Coins
		RemoteURL string
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
				From: "invalid",
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
				From: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			false,
		},
		{
			"empty provider and nil price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    nil,
			},
			false,
		},
		{
			"non-empty provider and non-nil price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"invalid prefix provider",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes provider",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			false,
		},
		{
			"20 bytes provider",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			false,
		},
		{
			"30 bytes provider",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			false,
		},
		{
			"empty price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				From:     "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider: "",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			false,
		},
		{
			"empty remote_url",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "",
			},
			false,
		},
		{
			"length 72 remote_url",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: strings.Repeat("r", 72),
			},
			true,
		},
		{
			"invalid remote_url",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "invalid",
			},
			true,
		},
		{
			"invalid remote_url scheme",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "tcp://remote.url:80",
			},
			true,
		},
		{
			"empty remote_url port",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url",
			},
			true,
		},
		{
			"non-empty remote_url port",
			fields{
				From:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Provider:  "",
				Price:     sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RemoteURL: "https://remote.url:443",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateRequest{
				From:      tt.fields.From,
				Provider:  tt.fields.Provider,
				Price:     tt.fields.Price,
				RemoteURL: tt.fields.RemoteURL,
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
				From: "invalid",
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
				From: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			true,
		},
		{
			"20 bytes from",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"30 bytes from",
			fields{
				From: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			true,
		},
		{
			"unknown status",
			fields{
				From:   "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Status: hubtypes.StatusUnknown,
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

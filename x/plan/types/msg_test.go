package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgAddRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From     string
		Prices   sdk.Coins
		Validity time.Duration
		Bytes    sdk.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				From: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From: "sentprov",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgsutj8xr",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"nil price",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   nil,
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"empty price",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"negative validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: -1000,
			},
			true,
		},
		{
			"zero validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 0,
			},
			true,
		},
		{
			"positive validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"negative bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"positive bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAddRequest{
				From:     tt.fields.From,
				Prices:   tt.fields.Prices,
				Validity: tt.fields.Validity,
				Bytes:    tt.fields.Bytes,
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
		ID     uint64
		Status hubtypes.Status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				From: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From: "sentprov",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"zero id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"unspecified status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSetStatusRequest{
				From:   tt.fields.From,
				ID:     tt.fields.ID,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgAddNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		ID      uint64
		Address string
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
				From: "sentprov",
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
				From:    "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"zero id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"empty address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAddNodeRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgRemoveNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		ID      uint64
		Address string
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
				From: "sentprov",
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
				From:    "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"zero id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"empty address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRemoveNodeRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

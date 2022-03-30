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
		Price    sdk.Coins
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
				From: "invalid",
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
				From: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"nil price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: nil,
			},
			true,
		},
		{
			"empty price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{sdk.Coin{Denom: ""}},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{sdk.Coin{Denom: "o"}},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				From:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"negative validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: -1000,
			},
			true,
		},
		{
			"zero validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 0,
			},
			true,
		},
		{
			"positive validity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"negative bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"positive bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Price:    sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
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
				Price:    tt.fields.Price,
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
		Id     uint64
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
				From: "invalid",
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
				From: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"zero id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:   1000,
			},
			true,
		},
		{
			"unknown status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:     1000,
				Status: hubtypes.StatusUnknown,
			},
			true,
		},
		{
			"active status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"inactive pending status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:     1000,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:     1000,
				Status: hubtypes.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSetStatusRequest{
				From:   tt.fields.From,
				Id:     tt.fields.Id,
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
		Id      uint64
		Address string
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
				From: "invalid",
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
				From: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"zero id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:   1000,
			},
			true,
		},
		{
			"empty address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAddNodeRequest{
				From:    tt.fields.From,
				Id:      tt.fields.Id,
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
		Id      uint64
		Address string
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
			"zero id",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:   1000,
			},
			true,
		},
		{
			"empty address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgRemoveNodeRequest{
				From:    tt.fields.From,
				Id:      tt.fields.Id,
				Address: tt.fields.Address,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

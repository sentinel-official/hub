package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgCreateRequest_ValidateBasic(t *testing.T) {
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
				From:     "sentprov1qypqxpq9qcrsszgsutj8xr",
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"price nil",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   nil,
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"price empty",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{},
			},
			true,
		},
		{
			"price empty denom",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"price empty amount",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"price invalid denom",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"price negative amount",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"price zero amount",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"price positive amount",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"validity negative",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				Validity: -1000,
			},
			true,
		},
		{
			"validity zero",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 0,
			},
			true,
		},
		{
			"validity positive",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
		{
			"bytes empty",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.Int{},
			},
			true,
		},
		{
			"bytes negative",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.NewInt(-1000),
			},
			true,
		},
		{
			"bytes zero",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.NewInt(0),
			},
			true,
		},
		{
			"bytes positive",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Validity: 1000,
				Bytes:    sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCreateRequest{
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

func TestMsgUpdateStatusRequest_ValidateBasic(t *testing.T) {
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
				From:   "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive pending",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:     1000,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
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
			m := &MsgUpdateStatusRequest{
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

func TestMsgLinkNodeRequest_ValidateBasic(t *testing.T) {
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
				From:    "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"id zero",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"address empty",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"address 30 bytes",
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
			m := &MsgLinkNodeRequest{
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

func TestMsgUnlinkNodeRequest_ValidateBasic(t *testing.T) {
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
				From:    "sentprov1qypqxpq9qcrsszgsutj8xr",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"id zero",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"address empty",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			false,
		},
		{
			"address 30 bytes",
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
			m := &MsgUnlinkNodeRequest{
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

func TestMsgSubscribeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From  string
		ID    uint64
		Denom string
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
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:  "sent1qypqxpq9qcrsszgslawd5s",
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"id zero",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:    1000,
				Denom: "",
			},
			false,
		},
		{
			"denom invalid",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:    1000,
				Denom: "o",
			},
			true,
		},
		{
			"denom one",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:    1000,
				Denom: "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:  tt.fields.From,
				ID:    tt.fields.ID,
				Denom: tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

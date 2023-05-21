package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgCancelRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From string
		ID   uint64
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
				ID:   1000,
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:   1000,
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ID:   1000,
			},
			false,
		},
		{
			"zero id",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:   1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCancelRequest{
				From: tt.fields.From,
				ID:   tt.fields.ID,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgAllocateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		ID      uint64
		Address string
		Bytes   sdk.Int
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
				From:    "sent1qypqxpq9qcrsszgslawd5s",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"zero id",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"empty address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgslawd5s",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"nil bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.Int{},
			},
			true,
		},
		{
			"negative bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"positive bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAllocateRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
				Bytes:   tt.fields.Bytes,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

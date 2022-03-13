package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgSubscribeToNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		Address string
		Deposit sdk.Coin
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
			"empty address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty denom deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"invalid denom deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"negative amount deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount deposit",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeToNodeRequest{
				From:    tt.fields.From,
				Address: tt.fields.Address,
				Deposit: tt.fields.Deposit,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSubscribeToPlanRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From  string
		Id    uint64
		Denom string
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
			false,
		},
		{
			"empty denom",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:    1000,
				Denom: "",
			},
			false,
		},
		{
			"invalid denom",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:    1000,
				Denom: "o",
			},
			true,
		},
		{
			"one denom",
			fields{
				From:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:    1000,
				Denom: "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeToPlanRequest{
				From:  tt.fields.From,
				Id:    tt.fields.Id,
				Denom: tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgCancelRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From string
		Id   uint64
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
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCancelRequest{
				From: tt.fields.From,
				Id:   tt.fields.Id,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgAddQuotaRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		Id      uint64
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
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"negative bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"positive bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAddQuotaRequest{
				From:    tt.fields.From,
				Id:      tt.fields.Id,
				Address: tt.fields.Address,
				Bytes:   tt.fields.Bytes,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateQuotaRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		Id      uint64
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
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"negative bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(0),
			},
			false,
		},
		{
			"positive bytes",
			fields{
				From:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Id:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bytes:   sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateQuotaRequest{
				From:    tt.fields.From,
				Id:      tt.fields.Id,
				Address: tt.fields.Address,
				Bytes:   tt.fields.Bytes,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

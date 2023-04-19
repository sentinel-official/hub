package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestQuota_GetAccountAddress(t *testing.T) {
	type fields struct {
		AccountAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty account address",
			fields{
				AccountAddress: "",
			},
			nil,
		},
		{
			"20 bytes account address",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Quota{
				AccountAddress: tt.fields.AccountAddress,
			}
			if got := m.GetAccountAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuota_Validate(t *testing.T) {
	type fields struct {
		AccountAddress string
		Allocated      sdk.Int
		Consumed       sdk.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty account address",
			fields{
				AccountAddress: "",
			},
			true,
		},
		{
			"invalid account address",
			fields{
				AccountAddress: "invalid",
			},
			true,
		},
		{
			"invalid prefix account address",
			fields{
				AccountAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes account address",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgslawd5s",
				Allocated:      sdk.NewInt(0),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"20 bytes account address",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(0),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"30 bytes account address",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				Allocated:      sdk.NewInt(0),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"nil allocated",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.Int{},
			},
			true,
		},
		{
			"negative allocated",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero allocated",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(0),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"positive allocated",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"nil consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.Int{},
			},
			true,
		},
		{
			"negative consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.NewInt(0),
			},
			false,
		},
		{
			"positive consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.NewInt(1000),
			},
			false,
		},
		{
			"allocated less than consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(1000),
				Consumed:       sdk.NewInt(2000),
			},
			true,
		},
		{
			"allocated equals to consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(2000),
				Consumed:       sdk.NewInt(2000),
			},
			false,
		},
		{
			"allocated greater than consumed",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Allocated:      sdk.NewInt(2000),
				Consumed:       sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Quota{
				AccountAddress: tt.fields.AccountAddress,
				Allocated:      tt.fields.Allocated,
				Consumed:       tt.fields.Consumed,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

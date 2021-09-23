package types

import (
	"reflect"
	"strings"
	"testing"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestProvider_GetAddress(t *testing.T) {
	type fields struct {
		Address     string
		Name        string
		Identity    string
		Website     string
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.ProvAddress
	}{
		{
			"empty",
			fields{
				Address: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			hubtypes.ProvAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Address:     tt.fields.Address,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
			}
			if got := p.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_Validate(t *testing.T) {
	type fields struct {
		Address     string
		Name        string
		Identity    string
		Website     string
		Description string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty address",
			fields{
				Address: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				Address: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			true,
		},
		{
			"30 bytes address",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			true,
		},
		{
			"empty name",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:    "",
			},
			true,
		},
		{
			"non-empty name",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:    "name",
			},
			false,
		},
		{
			"length 72 name",
			fields{
				Address: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:    strings.Repeat("n", 72),
			},
			true,
		},
		{
			"empty identity",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "",
			},
			false,
		},
		{
			"non-empty identity",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
			},
			false,
		},
		{
			"length 72 identity",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"empty website",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "",
			},
			false,
		},
		{
			"non-empty website",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
			},
			false,
		},
		{
			"length 72 website",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"invalid website",
			fields{
				Address:  "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"empty description",
			fields{
				Address:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: "",
			},
			false,
		},
		{
			"non-empty description",
			fields{
				Address:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: "description",
			},
			false,
		},
		{
			"length 264 description",
			fields{
				Address:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:        "name",
				Identity:    "identity",
				Website:     "https://website",
				Description: strings.Repeat("d", 264),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				Address:     tt.fields.Address,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

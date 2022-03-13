package types

import (
	"strings"
	"testing"
)

func TestMsgRegisterRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
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
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
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
			"empty name",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: "",
			},
			true,
		},
		{
			"non-empty name",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: "name",
			},
			false,
		},
		{
			"length 72 name",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"empty identity",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "",
			},
			false,
		},
		{
			"non-empty identity",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
			},
			false,
		},
		{
			"length 72 identity",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"empty website",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "",
			},
			false,
		},
		{
			"non-empty website",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
			},
			false,
		},
		{
			"length 72 website",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"invalid website",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"empty description",
			fields{
				From:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
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
				From:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
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
				From:        "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
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
			m := &MsgRegisterRequest{
				From:        tt.fields.From,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
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
			false,
		},
		{
			"20 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			false,
		},
		{
			"empty name",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: "",
			},
			false,
		},
		{
			"non-empty name",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: "name",
			},
			false,
		},
		{
			"length 72 name",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"empty identity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "",
			},
			false,
		},
		{
			"non-empty identity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
			},
			false,
		},
		{
			"length 72 identity",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"empty website",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "",
			},
			false,
		},
		{
			"non-empty website",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
			},
			false,
		},
		{
			"length 72 website",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"invalid website",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"empty description",
			fields{
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
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
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
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
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
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
			m := MsgUpdateRequest{
				From:        tt.fields.From,
				Name:        tt.fields.Name,
				Identity:    tt.fields.Identity,
				Website:     tt.fields.Website,
				Description: tt.fields.Description,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

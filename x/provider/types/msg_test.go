package types

import (
	"strings"
	"testing"

	hubtypes "github.com/sentinel-official/hub/types"
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
				From: "sent1qypqxpq9qcrsszgslawd5s",
				Name: "name",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: "name",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				Name: "name",
			},
			false,
		},
		{
			"name empty",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: "",
			},
			true,
		},
		{
			"name non-empty",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: strings.Repeat("n", 8),
			},
			false,
		},
		{
			"name length 72 chars",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"identity empty",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "",
			},
			false,
		},
		{
			"identity non-empty",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
			},
			false,
		},
		{
			"identity length 72 chars",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"website empty",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "",
			},
			false,
		},
		{
			"website non-empty",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
			},
			false,
		},
		{
			"website length 72 chars",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"website invalid",
			fields{
				From:     "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"description empty",
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
			"description non-empty",
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
			"description length 264 chars",
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
		Status      hubtypes.Status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"address empty",
			fields{
				From: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From: "sentprov1qypqxpq9qcrsszgsutj8xr",
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsh33zgx",
			},
			false,
		},
		{
			"name empty",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: "",
			},
			false,
		},
		{
			"name non-empty",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: "name",
			},
			false,
		},
		{
			"name length 72 chars",
			fields{
				From: "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"identity empty",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Identity: "",
			},
			false,
		},
		{
			"identity non-empty",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Identity: "identity",
			},
			false,
		},
		{
			"identity length 72 chars",
			fields{
				From:     "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"website empty",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Website: "",
			},
			false,
		},
		{
			"website non-empty",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Website: "https://website",
			},
			false,
		},
		{
			"website length 72 chars",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Website: strings.Repeat("w", 72),
			},
			true,
		},
		{
			"website invalid",
			fields{
				From:    "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Website: "invalid",
			},
			true,
		},
		{
			"description empty",
			fields{
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Description: "",
			},
			false,
		},
		{
			"description non-empty",
			fields{
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Description: "description",
			},
			false,
		},
		{
			"description length 264 chars",
			fields{
				From:        "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Description: strings.Repeat("d", 264),
			},
			true,
		},
		{
			"status unspecified",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Status: hubtypes.StatusUnspecified,
			},
			false,
		},
		{
			"status active",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   "sentprov1qypqxpq9qcrsszgszyfpx9q4zct3sxfq877k82",
				Status: hubtypes.StatusInactive,
			},
			false,
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
				Status:      tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

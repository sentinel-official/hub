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
				From: hubtypes.TestBech32ProvAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From: hubtypes.TestBech32AccAddr10Bytes,
				Name: "name",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				Name: "name",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From: hubtypes.TestBech32AccAddr30Bytes,
				Name: "name",
			},
			false,
		},
		{
			"name empty",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				Name: "",
			},
			true,
		},
		{
			"name non-empty",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				Name: strings.Repeat("n", 8),
			},
			false,
		},
		{
			"name length 72 chars",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"identity empty",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "",
			},
			false,
		},
		{
			"identity non-empty",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "identity",
			},
			false,
		},
		{
			"identity length 72 chars",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"website empty",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "",
			},
			false,
		},
		{
			"website non-empty",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "https://website",
			},
			false,
		},
		{
			"website length 72 chars",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  strings.Repeat("w", 72),
			},
			true,
		},
		{
			"website invalid",
			fields{
				From:     hubtypes.TestBech32AccAddr20Bytes,
				Name:     "name",
				Identity: "identity",
				Website:  "invalid",
			},
			true,
		},
		{
			"description empty",
			fields{
				From:        hubtypes.TestBech32AccAddr20Bytes,
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
				From:        hubtypes.TestBech32AccAddr20Bytes,
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
				From:        hubtypes.TestBech32AccAddr20Bytes,
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From: hubtypes.TestBech32ProvAddr10Bytes,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				From: hubtypes.TestBech32ProvAddr30Bytes,
			},
			false,
		},
		{
			"name empty",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				Name: "",
			},
			false,
		},
		{
			"name non-empty",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				Name: "name",
			},
			false,
		},
		{
			"name length 72 chars",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				Name: strings.Repeat("n", 72),
			},
			true,
		},
		{
			"identity empty",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Identity: "",
			},
			false,
		},
		{
			"identity non-empty",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Identity: "identity",
			},
			false,
		},
		{
			"identity length 72 chars",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Identity: strings.Repeat("i", 72),
			},
			true,
		},
		{
			"website empty",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				Website: "",
			},
			false,
		},
		{
			"website non-empty",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				Website: "https://website",
			},
			false,
		},
		{
			"website length 72 chars",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				Website: strings.Repeat("w", 72),
			},
			true,
		},
		{
			"website invalid",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				Website: "invalid",
			},
			true,
		},
		{
			"description empty",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				Description: "",
			},
			false,
		},
		{
			"description non-empty",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				Description: "description",
			},
			false,
		},
		{
			"description length 264 chars",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				Description: strings.Repeat("d", 264),
			},
			true,
		},
		{
			"status unspecified",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				Status: hubtypes.StatusUnspecified,
			},
			false,
		},
		{
			"status active",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
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

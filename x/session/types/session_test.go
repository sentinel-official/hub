package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestSession_GetAccountAddress(t *testing.T) {
	type fields struct {
		AccountAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty",
			fields{
				AccountAddress: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				AccountAddress: tt.fields.AccountAddress,
			}
			if got := m.GetAccountAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetNodeAddress(t *testing.T) {
	type fields struct {
		NodeAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty",
			fields{
				NodeAddress: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				NodeAddress: tt.fields.NodeAddress,
			}
			if got := m.GetNodeAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNodeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Validate(t *testing.T) {
	type fields struct {
		ID             uint64
		SubscriptionID uint64
		NodeAddress    string
		AccountAddress string
		Duration       time.Duration
		Bandwidth      hubtypes.Bandwidth
		Status         hubtypes.Status
		StatusAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero id",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"positive id",
			fields{
				ID: 1000,
			},
			true,
		},
		{
			"zero subscription",
			fields{
				ID:             1000,
				SubscriptionID: 0,
			},
			true,
		},
		{
			"positive subscription",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
			},
			true,
		},
		{
			"empty node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "",
			},
			true,
		},
		{
			"invalid node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "invalid",
			},
			true,
		},
		{
			"invalid prefix node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			true,
		},
		{
			"20 bytes node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"30 bytes node",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			true,
		},
		{
			"empty address",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative duration",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       -1000,
			},
			true,
		},
		{
			"zero duration",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       0,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive duration",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and negative download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"unspecified status",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"active status",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusActive,
			},
			true,
		},
		{
			"inactive pending status",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusActive,
			},
			true,
		},
		{
			"zero status at",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Time{},
			},
			true,
		},
		{
			"now status",
			fields{
				ID:             1000,
				SubscriptionID: 1000,
				NodeAddress:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				AccountAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:       1000,
				Bandwidth:      hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:         hubtypes.StatusActive,
				StatusAt:       time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				ID:             tt.fields.ID,
				SubscriptionID: tt.fields.SubscriptionID,
				NodeAddress:    tt.fields.NodeAddress,
				AccountAddress: tt.fields.AccountAddress,
				Duration:       tt.fields.Duration,
				Bandwidth:      tt.fields.Bandwidth,
				Status:         tt.fields.Status,
				StatusAt:       tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

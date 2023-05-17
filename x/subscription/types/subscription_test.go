package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestBaseSubscription_GetAddress(t *testing.T) {
	type fields struct {
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"address empty",
			fields{
				Address: "",
			},
			"",
		},
		{
			"address non-empty",
			fields{
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			"sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BaseSubscription{
				Address: tt.fields.Address,
			}
			if got := s.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSubscription_Validate(t *testing.T) {
	type fields struct {
		ID       uint64
		Address  string
		ExpiryAt time.Time
		Status   hubtypes.Status
		StatusAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"id zero",
			fields{
				ID: 0,
			},
			true,
		},
		{
			"id positive",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address empty",
			fields{
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				ID:      1000,
				Address: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgslawd5s",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"expiry_at empty",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ExpiryAt: time.Time{},
			},
			true,
		},
		{
			"expiry_at non-empty",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status unspecified",
			fields{
				ID:      1000,
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Status:  hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusInactivePending,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status inactive",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusInactive,
				StatusAt: time.Now(),
			},
			false,
		},
		{
			"status_at empty",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Time{},
			},
			true,
		},
		{
			"status_at non-empty",
			fields{
				ID:       1000,
				Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				ExpiryAt: time.Now(),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BaseSubscription{
				ID:       tt.fields.ID,
				Address:  tt.fields.Address,
				ExpiryAt: tt.fields.ExpiryAt,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNodeSubscription_GetNodeAddress(t *testing.T) {
	type fields struct {
		NodeAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"node_address empty",
			fields{
				NodeAddress: "",
			},
			nil,
		},
		{
			"node_address 20 bytes",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodeSubscription{
				NodeAddress: tt.fields.NodeAddress,
			}
			if got := s.GetNodeAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNodeAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeSubscription_Type(t *testing.T) {
	v := &NodeSubscription{}
	require.Equal(t, TypeNode, v.Type())
}

func TestNodeSubscription_Validate(t *testing.T) {
	type fields struct {
		NodeAddress string
		Gigabytes   int64
		Hours       int64
		Deposit     sdk.Coin
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"node_address empty",
			fields{
				NodeAddress: "",
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				NodeAddress: "invalid",
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				NodeAddress: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgse4wwrm",
				Gigabytes:   1000,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
				Gigabytes:   1000,
			},
			false,
		},
		{
			"gigabytes empty and hours empty",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   0,
				Hours:       0,
			},
			true,
		},
		{
			"gigabytes non-empty and hours non-empty",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Hours:       1000,
			},
			true,
		},
		{
			"gigabytes negative",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
			},
			false,
		},
		{
			"hours negative",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       -1000,
			},
			true,
		},
		{
			"hours positive",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Hours:       1000,
			},
			false,
		},
		{
			"deposit empty",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{},
			},
			false,
		},
		{
			"deposit empty denom",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			false,
		},
		{
			"deposit invalid denom",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "d", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				NodeAddress: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Gigabytes:   1000,
				Deposit:     sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodeSubscription{
				BaseSubscription: &BaseSubscription{
					ID:       1000,
					Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
					ExpiryAt: time.Now(),
					Status:   hubtypes.StatusActive,
					StatusAt: time.Now(),
				},
				NodeAddress: tt.fields.NodeAddress,
				Gigabytes:   tt.fields.Gigabytes,
				Hours:       tt.fields.Hours,
				Deposit:     tt.fields.Deposit,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlanSubscription_Type(t *testing.T) {
	v := &PlanSubscription{}
	require.Equal(t, TypePlan, v.Type())
}

func TestPlanSubscription_Validate(t *testing.T) {
	type fields struct {
		PlanID uint64
		Denom  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"plan_id zero",
			fields{
				PlanID: 0,
			},
			true,
		},
		{
			"plan_id positive",
			fields{
				PlanID: 1000,
				Denom:  "one",
			},
			false,
		},
		{
			"denom empty",
			fields{
				PlanID: 1000,
				Denom:  "",
			},
			false,
		},
		{
			"denom invalid",
			fields{
				PlanID: 1000,
				Denom:  "d",
			},
			true,
		},
		{
			"denom one",
			fields{
				PlanID: 1000,
				Denom:  "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &PlanSubscription{
				BaseSubscription: &BaseSubscription{
					ID:       1000,
					Address:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
					ExpiryAt: time.Now(),
					Status:   hubtypes.StatusActive,
					StatusAt: time.Now(),
				},
				PlanID: tt.fields.PlanID,
				Denom:  tt.fields.Denom,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

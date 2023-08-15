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
		want   sdk.AccAddress
	}{
		{
			"empty",
			fields{
				Address: hubtypes.TestAddrEmpty,
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
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
		ID         uint64
		Address    string
		InactiveAt time.Time
		Status     hubtypes.Status
		StatusAt   time.Time
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
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"address empty",
			fields{
				ID:      1000,
				Address: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"address invalid",
			fields{
				ID:      1000,
				Address: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr10Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr30Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"inactive_at empty",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr30Bytes,
				InactiveAt: hubtypes.TestTimeZero,
			},
			true,
		},
		{
			"inactive_at non-empty",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr30Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Status:  hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status inactive_pending",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusInactivePending,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status inactive",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusInactive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
		{
			"status_at empty",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeZero,
			},
			true,
		},
		{
			"status_at non-empty",
			fields{
				ID:         1000,
				Address:    hubtypes.TestBech32AccAddr20Bytes,
				InactiveAt: hubtypes.TestTimeNow,
				Status:     hubtypes.StatusActive,
				StatusAt:   hubtypes.TestTimeNow,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BaseSubscription{
				ID:         tt.fields.ID,
				Address:    tt.fields.Address,
				InactiveAt: tt.fields.InactiveAt,
				Status:     tt.fields.Status,
				StatusAt:   tt.fields.StatusAt,
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
				NodeAddress: hubtypes.TestAddrEmpty,
			},
			nil,
		},
		{
			"node_address 20 bytes",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
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
				NodeAddress: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				NodeAddress: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				NodeAddress: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr10Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr30Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
		{
			"gigabytes empty and hours empty",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   0,
				Hours:       0,
			},
			true,
		},
		{
			"gigabytes non-empty and hours non-empty",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Hours:       1000,
			},
			true,
		},
		{
			"gigabytes negative",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   -1000,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
		{
			"hours negative",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       -1000,
			},
			true,
		},
		{
			"hours positive",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Hours:       1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
		{
			"deposit empty",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinEmpty,
			},
			true,
		},
		{
			"deposit empty denom",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinEmptyDenom,
			},
			true,
		},
		{
			"deposit invalid denom",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinInvalidDenom,
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinEmptyAmount,
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinNegativeAmount,
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinZeroAmount,
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
				Gigabytes:   1000,
				Deposit:     hubtypes.TestCoinPositiveAmount,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodeSubscription{
				BaseSubscription: &BaseSubscription{
					ID:         1000,
					Address:    hubtypes.TestBech32AccAddr20Bytes,
					InactiveAt: hubtypes.TestTimeNow,
					Status:     hubtypes.StatusActive,
					StatusAt:   hubtypes.TestTimeNow,
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
			true,
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
					ID:         1000,
					Address:    hubtypes.TestBech32AccAddr20Bytes,
					InactiveAt: hubtypes.TestTimeNow,
					Status:     hubtypes.StatusActive,
					StatusAt:   hubtypes.TestTimeNow,
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

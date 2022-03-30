package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestSession_GetAddress(t *testing.T) {
	type fields struct {
		Id           uint64
		Subscription uint64
		Node         string
		Address      string
		Duration     time.Duration
		Bandwidth    hubtypes.Bandwidth
		Status       hubtypes.Status
		StatusAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
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
				Address: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				Id:           tt.fields.Id,
				Subscription: tt.fields.Subscription,
				Node:         tt.fields.Node,
				Address:      tt.fields.Address,
				Duration:     tt.fields.Duration,
				Bandwidth:    tt.fields.Bandwidth,
				Status:       tt.fields.Status,
				StatusAt:     tt.fields.StatusAt,
			}
			if got := m.GetAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_GetNode(t *testing.T) {
	type fields struct {
		Id           uint64
		Subscription uint64
		Node         string
		Address      string
		Duration     time.Duration
		Bandwidth    hubtypes.Bandwidth
		Status       hubtypes.Status
		StatusAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty",
			fields{
				Node: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Node: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				Id:           tt.fields.Id,
				Subscription: tt.fields.Subscription,
				Node:         tt.fields.Node,
				Address:      tt.fields.Address,
				Duration:     tt.fields.Duration,
				Bandwidth:    tt.fields.Bandwidth,
				Status:       tt.fields.Status,
				StatusAt:     tt.fields.StatusAt,
			}
			if got := m.GetNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Validate(t *testing.T) {
	type fields struct {
		Id           uint64
		Subscription uint64
		Node         string
		Address      string
		Duration     time.Duration
		Bandwidth    hubtypes.Bandwidth
		Status       hubtypes.Status
		StatusAt     time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero id",
			fields{
				Id: 0,
			},
			true,
		},
		{
			"positive id",
			fields{
				Id: 1000,
			},
			true,
		},
		{
			"zero subscription",
			fields{
				Id:           1000,
				Subscription: 0,
			},
			true,
		},
		{
			"positive subscription",
			fields{
				Id:           1000,
				Subscription: 1000,
			},
			true,
		},
		{
			"empty node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "",
			},
			true,
		},
		{
			"invalid node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "invalid",
			},
			true,
		},
		{
			"invalid prefix node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"10 bytes node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgse4wwrm",
			},
			true,
		},
		{
			"20 bytes node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"30 bytes node",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fsxqglcv",
			},
			true,
		},
		{
			"empty address",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "",
			},
			true,
		},
		{
			"invalid address",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "invalid",
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"20 bytes address",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative duration",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     -1000,
			},
			true,
		},
		{
			"zero duration",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     0,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive duration",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and negative download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"unknown status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.StatusUnknown,
			},
			true,
		},
		{
			"active status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.Active,
			},
			true,
		},
		{
			"inactive pending status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.Inactive,
			},
			true,
		},
		{
			"zero status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.Inactive,
				StatusAt:     time.Time{},
			},
			true,
		},
		{
			"now status",
			fields{
				Id:           1000,
				Subscription: 1000,
				Node:         "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Address:      "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Duration:     1000,
				Bandwidth:    hubtypes.Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
				Status:       hubtypes.Inactive,
				StatusAt:     time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Session{
				Id:           tt.fields.Id,
				Subscription: tt.fields.Subscription,
				Node:         tt.fields.Node,
				Address:      tt.fields.Address,
				Duration:     tt.fields.Duration,
				Bandwidth:    tt.fields.Bandwidth,
				Status:       tt.fields.Status,
				StatusAt:     tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

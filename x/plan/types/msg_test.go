package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

func TestMsgCreateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Duration  time.Duration
		Gigabytes int64
		Prices    sdk.Coins
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:      hubtypes.TestBech32ProvAddr10Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:      hubtypes.TestBech32ProvAddr30Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"duration negative",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Duration: -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Duration: 0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"gigabytes negative",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: -1000,
			},
			true,
		},
		{
			"gigabytes zero",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 0,
			},
			true,
		},
		{
			"gigabytes positive",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
		{
			"prices nil",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsNil,
			},
			true,
		},
		{
			"prices empty",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsEmpty,
			},
			true,
		},
		{
			"prices empty denom",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"prices empty amount",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"prices invalid denom",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"prices negative amount",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"prices zero amount",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"prices positive amount",
			fields{
				From:      hubtypes.TestBech32ProvAddr20Bytes,
				Duration:  1000,
				Gigabytes: 1000,
				Prices:    hubtypes.TestCoinsPositiveAmount,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCreateRequest{
				From:      tt.fields.From,
				Duration:  tt.fields.Duration,
				Gigabytes: tt.fields.Gigabytes,
				Prices:    tt.fields.Prices,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUpdateStatusRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From   string
		ID     uint64
		Status hubtypes.Status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:   hubtypes.TestBech32ProvAddr10Bytes,
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:   hubtypes.TestBech32ProvAddr30Bytes,
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status unspecified",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusUnspecified,
			},
			true,
		},
		{
			"status active",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusActive,
			},
			false,
		},
		{
			"status inactive pending",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"status inactive",
			fields{
				From:   hubtypes.TestBech32ProvAddr20Bytes,
				ID:     1000,
				Status: hubtypes.StatusInactive,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateStatusRequest{
				From:   tt.fields.From,
				ID:     tt.fields.ID,
				Status: tt.fields.Status,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgLinkNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
		ID          uint64
		NodeAddress string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr10Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr30Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgLinkNodeRequest{
				From:        tt.fields.From,
				ID:          tt.fields.ID,
				NodeAddress: tt.fields.NodeAddress,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUnlinkNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From        string
		ID          uint64
		NodeAddress string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"from invalid prefix",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr10Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr30Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"id zero",
			fields{
				From: hubtypes.TestBech32ProvAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address empty",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"node_address invalid",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"node_address invalid prefix",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"node_address 10 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"node_address 20 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"node_address 30 bytes",
			fields{
				From:        hubtypes.TestBech32ProvAddr20Bytes,
				ID:          1000,
				NodeAddress: hubtypes.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUnlinkNodeRequest{
				From:        tt.fields.From,
				ID:          tt.fields.ID,
				NodeAddress: tt.fields.NodeAddress,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgSubscribeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From  string
		ID    uint64
		Denom string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"from empty",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"from invalid",
			fields{
				From: hubtypes.TestAddrInvalid,
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
				From:  hubtypes.TestBech32AccAddr10Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:  hubtypes.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:  hubtypes.TestBech32AccAddr30Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"id zero",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"id positive",
			fields{
				From:  hubtypes.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
		{
			"denom empty",
			fields{
				From:  hubtypes.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "",
			},
			true,
		},
		{
			"denom invalid",
			fields{
				From:  hubtypes.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "o",
			},
			true,
		},
		{
			"denom one",
			fields{
				From:  hubtypes.TestBech32AccAddr20Bytes,
				ID:    1000,
				Denom: "one",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgSubscribeRequest{
				From:  tt.fields.From,
				ID:    tt.fields.ID,
				Denom: tt.fields.Denom,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

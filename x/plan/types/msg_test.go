package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgCreateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From     string
		Bytes    sdk.Int
		Duration time.Duration
		Prices   sdk.Coins
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:     hubtypes.TestBech32ProvAddr10Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:     hubtypes.TestBech32ProvAddr30Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
			},
			false,
		},
		{
			"bytes empty",
			fields{
				From:  hubtypes.TestBech32ProvAddr20Bytes,
				Bytes: sdk.Int{},
			},
			true,
		},
		{
			"bytes negative",
			fields{
				From:  hubtypes.TestBech32ProvAddr20Bytes,
				Bytes: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"bytes zero",
			fields{
				From:  hubtypes.TestBech32ProvAddr20Bytes,
				Bytes: sdk.NewInt(0),
			},
			true,
		},
		{
			"bytes positive",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
			},
			false,
		},
		{
			"duration negative",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: -1000,
			},
			true,
		},
		{
			"duration zero",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 0,
			},
			true,
		},
		{
			"duration positive",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
			},
			false,
		},
		{
			"prices nil",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   nil,
			},
			false,
		},
		{
			"prices empty",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{},
			},
			true,
		},
		{
			"prices empty denom",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"prices empty amount",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"prices invalid denom",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"prices negative amount",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"prices zero amount",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"prices positive amount",
			fields{
				From:     hubtypes.TestBech32ProvAddr20Bytes,
				Bytes:    sdk.NewInt(1000),
				Duration: 1000,
				Prices:   sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCreateRequest{
				From:     tt.fields.From,
				Bytes:    tt.fields.Bytes,
				Duration: tt.fields.Duration,
				Prices:   tt.fields.Prices,
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
		From    string
		ID      uint64
		Address string
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr10Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr30Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
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
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"address empty",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgLinkNodeRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgUnlinkNodeRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		ID      uint64
		Address string
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"from 10 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr10Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 20 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"from 30 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr30Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
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
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"address empty",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: "",
			},
			true,
		},
		{
			"address invalid",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: "invalid",
			},
			true,
		},
		{
			"address invalid prefix",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"address 10 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"address 20 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"address 30 bytes",
			fields{
				From:    hubtypes.TestBech32ProvAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUnlinkNodeRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
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
			false,
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

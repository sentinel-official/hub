package types

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestMsgCancelRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From string
		ID   uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty from",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid from",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix from",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr10Bytes,
				ID:   1000,
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				ID:   1000,
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr30Bytes,
				ID:   1000,
			},
			false,
		},
		{
			"zero id",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				ID:   1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgCancelRequest{
				From: tt.fields.From,
				ID:   tt.fields.ID,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgAllocateRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From    string
		ID      uint64
		Address string
		Bytes   sdkmath.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty from",
			fields{
				From: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid from",
			fields{
				From: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix from",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From:    hubtypes.TestBech32AccAddr10Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"20 bytes from",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"30 bytes from",
			fields{
				From:    hubtypes.TestBech32AccAddr30Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"zero id",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
				ID:   0,
			},
			true,
		},
		{
			"positive id",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"empty address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr10Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"20 bytes address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"30 bytes address",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr30Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"nil bytes",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.Int{},
			},
			true,
		},
		{
			"negative bytes",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(-1000),
			},
			true,
		},
		{
			"zero bytes",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(0),
			},
			false,
		},
		{
			"positive bytes",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
				Bytes:   sdkmath.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgAllocateRequest{
				From:    tt.fields.From,
				ID:      tt.fields.ID,
				Address: tt.fields.Address,
				Bytes:   tt.fields.Bytes,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

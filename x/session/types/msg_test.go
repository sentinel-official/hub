package types

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

func TestMsgStartRequest_ValidateBasic(t *testing.T) {
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
			},
			true,
		},
		{
			"20 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"30 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr30Bytes,
			},
			true,
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
			true,
		},
		{
			"empty node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestAddrEmpty,
			},
			true,
		},
		{
			"invalid node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestAddrInvalid,
			},
			true,
		},
		{
			"invalid prefix node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr10Bytes,
			},
			false,
		},
		{
			"20 bytes node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr20Bytes,
			},
			false,
		},
		{
			"30 bytes node",
			fields{
				From:    hubtypes.TestBech32AccAddr20Bytes,
				ID:      1000,
				Address: hubtypes.TestBech32NodeAddr30Bytes,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgStartRequest{
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

func TestMsgUpdateDetailsRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From      string
		Proof     Proof
		Signature []byte
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
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"10 bytes from",
			fields{
				From: hubtypes.TestBech32NodeAddr10Bytes,
			},
			true,
		},
		{
			"20 bytes from",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
			},
			true,
		},
		{
			"30 bytes from",
			fields{
				From: hubtypes.TestBech32NodeAddr30Bytes,
			},
			true,
		},
		{
			"zero proof->id",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID: 0,
				},
			},
			true,
		},
		{
			"positive proof->id",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
				},
			},
			false,
		},
		{
			"negative proof->bandwidth->upload and negative proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(-1000)},
				},
			},
			true,
		},
		{
			"negative proof->bandwidth->upload and zero proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(0)},
				},
			},
			true,
		},
		{
			"negative proof->bandwidth->upload and positive proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(-1000), Download: sdkmath.NewInt(1000)},
				},
			},
			true,
		},
		{
			"zero proof->bandwidth->upload and negative proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(-1000)},
				},
			},
			true,
		},
		{
			"zero proof->bandwidth->upload and zero proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
				},
			},
			false,
		},
		{
			"zero proof->bandwidth->upload and positive proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(1000)},
				},
			},
			false,
		},
		{
			"positive proof->bandwidth->upload and negative proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(-1000)},
				},
			},
			true,
		},
		{
			"positive proof->bandwidth->upload and zero proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(0)},
				},
			},
			false,
		},
		{
			"positive proof->bandwidth->upload and positive proof->bandwidth->download",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
				},
			},
			false,
		},
		{
			"negative proof->duration",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
					Duration:  -1000,
				},
			},
			true,
		},
		{
			"zero proof->duration",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
					Duration:  0,
				},
			},
			false,
		},
		{
			"positive proof->duration",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(0), Download: sdkmath.NewInt(0)},
					Duration:  1000,
				},
			},
			false,
		},
		{
			"nil signature",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
					Duration:  1000,
				},
				Signature: nil,
			},
			false,
		},
		{
			"empty signature",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
					Duration:  1000,
				},
				Signature: []byte{},
			},
			true,
		},
		{
			"32 byte signature",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
					Duration:  1000,
				},
				Signature: []byte{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
					0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
					0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31,
				},
			},
			true,
		},
		{
			"64 byte signature",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
					Duration:  1000,
				},
				Signature: []byte{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
					0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
					0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31,
					0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
					0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
					0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55,
					0x56, 0x57, 0x58, 0x59, 0x60, 0x61, 0x62, 0x63,
				},
			},
			false,
		},
		{
			"96 byte signature",
			fields{
				From: hubtypes.TestBech32NodeAddr20Bytes,
				Proof: Proof{
					ID:        1000,
					Bandwidth: hubtypes.Bandwidth{Upload: sdkmath.NewInt(1000), Download: sdkmath.NewInt(1000)},
					Duration:  1000,
				},
				Signature: []byte{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
					0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15,
					0x16, 0x17, 0x18, 0x19, 0x20, 0x21, 0x22, 0x23,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31,
					0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39,
					0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47,
					0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55,
					0x56, 0x57, 0x58, 0x59, 0x60, 0x61, 0x62, 0x63,
					0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x70, 0x71,
					0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79,
					0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
					0x88, 0x89, 0x90, 0x91, 0x92, 0x93, 0x94, 0x95,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgUpdateDetailsRequest{
				From:      tt.fields.From,
				Proof:     tt.fields.Proof,
				Signature: tt.fields.Signature,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMsgEndRequest_ValidateBasic(t *testing.T) {
	type fields struct {
		From   string
		ID     uint64
		Rating uint64
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
			},
			true,
		},
		{
			"20 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr20Bytes,
			},
			true,
		},
		{
			"30 bytes from",
			fields{
				From: hubtypes.TestBech32AccAddr30Bytes,
			},
			true,
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
		{
			"zero rating",
			fields{
				From:   hubtypes.TestBech32AccAddr20Bytes,
				ID:     1000,
				Rating: 0,
			},
			false,
		},
		{
			"5 rating",
			fields{
				From:   hubtypes.TestBech32AccAddr20Bytes,
				ID:     1000,
				Rating: 5,
			},
			false,
		},
		{
			"10 rating",
			fields{
				From:   hubtypes.TestBech32AccAddr20Bytes,
				ID:     1000,
				Rating: 10,
			},
			false,
		},
		{
			"15 rating",
			fields{
				From:   hubtypes.TestBech32AccAddr20Bytes,
				ID:     1000,
				Rating: 15,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MsgEndRequest{
				From:   tt.fields.From,
				ID:     tt.fields.ID,
				Rating: tt.fields.Rating,
			}
			if err := m.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

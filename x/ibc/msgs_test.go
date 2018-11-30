package ibc

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestMsgIBCTransaction_ValidateBasic(t *testing.T) {
	type fields struct {
		Relayer   csdkTypes.AccAddress
		Sequence  uint64
		IBCPacket sdkTypes.IBCPacket
	}
	tests := []struct {
		name   string
		fields fields
		want   csdkTypes.Error
	}{
		{"testSourceChainID", fields{
			Relayer:  accAddress1,
			Sequence: uint64(0),
			IBCPacket: sdkTypes.IBCPacket{
				SrcChainID:  "",
				DestChainID: "sentinel-hub",
				Message:     nil,
			},
		},
			errorEmptySrcChainID(),
		},

		{"testMsgValid", fields{
			Relayer:  accAddress1,
			Sequence: uint64(2),
			IBCPacket: sdkTypes.IBCPacket{
				SrcChainID:  "sentinel-vpn",
				DestChainID: "sentinel-hub",
				Message:     nil,
			},
		},
			nil,
		},

		{"testRelayAddress", fields{
			Relayer:  nil,
			Sequence: uint64(1),
			IBCPacket: sdkTypes.IBCPacket{
				SrcChainID:  "sentinel-vpn",
				DestChainID: "sentinel-hub",
				Message:     nil,
			},
		},
			errorEmptyRelayer(),
		},

		{"testDestinationChainID", fields{
			Relayer:  accAddress2,
			Sequence: uint64(0),
			IBCPacket: sdkTypes.IBCPacket{
				SrcChainID:  "sentinel-vpn",
				DestChainID: "",
				Message:     nil,
			},
		},
			errorEmptyDestChainID(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgIBCTransaction{
				Relayer:   tt.fields.Relayer,
				Sequence:  tt.fields.Sequence,
				IBCPacket: tt.fields.IBCPacket,
			}
			if got := msg.ValidateBasic(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgIBCTransaction.ValidateBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgIBCTransaction_GetSigners(t *testing.T) {
	msg := MsgIBCTransaction{
		csdkTypes.AccAddress("TestMsgIBCTransaction"),
		0,
		sdkTypes.IBCPacket{
			"sentinel-vpn",
			"sentinel-hub",
			nil,
		},
	}

	res := msg.GetSigners()
	require.Equal(t, fmt.Sprintf("%v", res), "[546573744D73674942435472616E73616374696F6E]")
}

func TestMsgIBCTransaction_GetSignBytes(t *testing.T) {

	msg := MsgIBCTransaction{
		csdkTypes.AccAddress("TestMsgIBCTransaction"),
		1,
		sdkTypes.IBCPacket{
			"sentinel-vpn",
			"sentinel-hub",
			nil,
		},
	}
	expected := `{"relayer":"cosmos123jhxazdwdn5jsjr23exzmnnv93hg6t0dcm0snjh","sequence":1,"ibc_packet":{"src_chain_id":"sentinel-vpn","dest_chain_id":"sentinel-hub","message":null}}`
	require.Equal(t, expected, string(msg.GetSignBytes()))
}

func TestMsgIBCTransaction_Route(t *testing.T) {
	msg := MsgIBCTransaction{
		csdkTypes.AccAddress("TestMsgIBCTransaction"),
		1,
		sdkTypes.IBCPacket{
			"sentinel-vpn",
			"sentinel-hub",
			nil,
		},
	}

	require.Equal(t, "ibc", msg.Route())
}

func TestMsgIBCTransaction_Type(t *testing.T) {
	msg := MsgIBCTransaction{
		csdkTypes.AccAddress("TestMsgIBCTransaction"),
		3,
		sdkTypes.IBCPacket{
			"sentinel-vpn",
			"sentinel-hub",
			nil,
		},
	}

	require.Equal(t, msg.Type(), "msg_ibc_transaction")
}

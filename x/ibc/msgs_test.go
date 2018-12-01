package ibc

import (
	"testing"

	"github.com/stretchr/testify/require"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestMsgIBCTransaction_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgIBCTransaction
		want csdkTypes.Error
	}{
		{"Relayer nil", MsgIBCTransaction{
			nil,
			uint64(0),
			sdkTypes.IBCPacket{"src_chain_id", "dest_chain_id", nil},
		}, errorInvalidRelayer()},
		{"Relayer empty", MsgIBCTransaction{
			csdkTypes.AccAddress{},
			uint64(0),
			sdkTypes.IBCPacket{"src_chain_id", "dest_chain_id", nil},
		}, errorInvalidRelayer()},
		{"Source chain ID empty", MsgIBCTransaction{
			accAddress,
			uint64(0),
			sdkTypes.IBCPacket{"", "dest_chain_id", nil},
		}, errorEmptySrcChainID()},
		{"Destination chain ID empty", MsgIBCTransaction{
			accAddress,
			uint64(0),
			sdkTypes.IBCPacket{"src_chain_id", "", nil},
		}, errorEmptyDestChainID()},
		{"Valid", MsgIBCTransaction{
			accAddress,
			uint64(0),
			sdkTypes.IBCPacket{"src_chain_id", "dest_chain_id", nil},
		}, nil},
	}

	for _, tc := range tests {
		require.Equal(t, tc.want, tc.msg.ValidateBasic())
	}
}

func TestMsgIBCTransaction_GetSigners(t *testing.T) {
	msg := msgIBCTransaction
	require.Equal(t, []csdkTypes.AccAddress{csdkTypes.AccAddress([]byte("acc_address"))}, msg.GetSigners())
}

func TestMsgIBCTransaction_GetSignBytes(t *testing.T) {
	msg := msgIBCTransaction
	require.Equal(t, `{"relayer":"cosmos1v93kxhmpv3j8yetnwvxysxeh","sequence":0,"ibc_packet":{"src_chain_id":"src_chain_id","dest_chain_id":"dest_chain_id","message":null}}`, string(msg.GetSignBytes()))
}

func TestMsgIBCTransaction_Route(t *testing.T) {
	msg := msgIBCTransaction
	require.Equal(t, sdkTypes.KeyIBC, msg.Route())
}

func TestMsgIBCTransaction_Type(t *testing.T) {
	msg := msgIBCTransaction
	require.Equal(t, "msg_ibc_transaction", msg.Type())
}

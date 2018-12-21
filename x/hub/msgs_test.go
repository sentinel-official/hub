package hub

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMsgLockCoins(t *testing.T) {
	tests := []struct {
		name         string
		msg          MsgLockCoins
		expectedPass bool
	}{
		{"Locker ID empty", MsgLockCoins{
			"",
			csdkTypes.Coins{coin(10, "x")},
			pubKey1,
			getMsgLockCoinsSignature("", csdkTypes.Coins{coin(10, "x")}),
		}, false},
		{"Coins nil", MsgLockCoins{
			"locker_id",
			nil,
			pubKey1,
			getMsgLockCoinsSignature("locker_id", nil),
		}, false},
		{"Coins len zero", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{},
			pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{}),
		}, false},
		{"Coins invalid", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(10, "y"), coin(10, "x")},
			pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(10, "y"), coin(10, "x")}),
		}, false},
		{"Coins negative", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(-10, "x")},
			pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(-10, "x")}),
		}, false},
		{"PubKey nil", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(10, "x")},
			nil,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(10, "x")}),
		}, false},
		{"Signature empty", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(10, "x")},
			pubKey1,
			[]byte(""),
		}, false},
		{"Signature invalid", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(10, "x")},
			pubKey1,
			[]byte("invalid_signature"),
		}, false},
		{"Valid", MsgLockCoins{
			"locker_id",
			csdkTypes.Coins{coin(10, "x")},
			pubKey1,
			getMsgLockCoinsSignature("locker_id", csdkTypes.Coins{coin(10, "x")}),
		}, true},
	}

	for _, tc := range tests {
		if tc.expectedPass {
			require.Nil(t, tc.msg.ValidateBasic())
		} else {
			require.NotNil(t, tc.msg.ValidateBasic())
		}
	}
}

func TestMsgReleaseCoins(t *testing.T) {
	tests := []struct {
		name         string
		msg          MsgReleaseCoins
		expectedPass bool
	}{
		{"Locker ID empty", MsgReleaseCoins{
			"",
			pubKey1,
			getMsgReleaseCoinsSignature(""),
		}, false},
		{"PubKey nil", MsgReleaseCoins{
			"locker_id",
			nil,
			getMsgReleaseCoinsSignature("locker_id"),
		}, false},
		{"Signature empty", MsgReleaseCoins{
			"locker_id",
			pubKey1,
			[]byte(""),
		}, false},
		{"Signature invalid", MsgReleaseCoins{
			"locker_id",
			pubKey1,
			[]byte("invalid_signature"),
		}, false},
		{"Valid", MsgReleaseCoins{
			"locker_id",
			pubKey1,
			getMsgReleaseCoinsSignature("locker_id"),
		}, true},
	}

	for _, tc := range tests {
		if tc.expectedPass {
			require.Nil(t, tc.msg.ValidateBasic())
		} else {
			require.NotNil(t, tc.msg.ValidateBasic())
		}
	}
}

func TestMsgReleaseCoinsToMany(t *testing.T) {
	tests := []struct {
		name         string
		msg          MsgReleaseCoinsToMany
		expectedPass bool
	}{
		{"Empty locker ID", MsgReleaseCoinsToMany{
			"",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, false},
		{"Addresses empty", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, false},
		{"Address empty", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, nil},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, nil},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, false},
		{"Empty shares", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{}),
		}, false},
		{"Share nil", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, nil},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, nil}),
		}, false},
		{"Share length zero", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, {}}),
		}, false},
		{"Share invalid", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{
				{coin(10, "x")},
				{coin(10, "y"), coin(10, "x")},
			},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{
					{coin(10, "x")},
					{coin(10, "y"), coin(10, "x")},
				}),
		}, false},
		{"Share negative", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(-10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(-10, "x")}}),
		}, false},
		{"Address and shares length mismatch", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, false},
		{"PubKey nil", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			nil,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, false},
		{"Signature empty", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			[]byte(""),
		}, false},
		{"Signature failed", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			[]byte("invalid_signature"),
		}, false},
		{"Valid", MsgReleaseCoinsToMany{
			"locker_id",
			[]csdkTypes.AccAddress{accAddress1, accAddress2},
			[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}},
			pubKey1,
			getMsgReleaseCoinsToManySignature("locker_id",
				[]csdkTypes.AccAddress{accAddress1, accAddress2},
				[]csdkTypes.Coins{{coin(10, "x")}, {coin(10, "x")}}),
		}, true},
	}

	for _, tc := range tests {
		if !tc.expectedPass {
			require.NotNil(t, tc.msg.ValidateBasic())
		} else {
			require.Nil(t, tc.msg.ValidateBasic())
		}
	}
}

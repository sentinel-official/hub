package hub

import (
	"github.com/stretchr/testify/require"
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func TestMsgLockCoins(t *testing.T) {
	type fields struct {
		LockerID  string
		Coins     csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{"test1", fields{LockerID: lockerId, Coins: coins1, PubKey: pk1, Signature: sign1}, true},
		{"test2", fields{LockerID: lockerId, Coins: coinsNeg, PubKey: pk1, Signature: sign1}, false},
		{"test3", fields{LockerID: lockerId2, Coins: coins1, PubKey: pk2, Signature: nil}, false},
		{"test4", fields{LockerID: "", Coins: nil, PubKey: pk2, Signature: sign1}, false},
	}

	for _, tc := range tests {
		msg := MsgLockCoins{tc.fields.LockerID, tc.fields.Coins, tc.fields.PubKey, tc.fields.Signature}

		if tc.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}
}

func TestMsgLockCoins_Verify(t *testing.T) {
	type fields struct {
		LockerID  string
		Coins     csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test1", fields{LockerID: lockerId, Coins: coins1, PubKey: pk1, Signature: sign1}, true},
		{"test2", fields{LockerID: lockerId, Coins: coins1, PubKey: pk2, Signature: nil}, false},
		{"test3", fields{LockerID: "", Coins: nil, PubKey: pk2, Signature: sign1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgLockCoins{
				LockerID:  tt.fields.LockerID,
				Coins:     tt.fields.Coins,
				PubKey:    tt.fields.PubKey,
				Signature: tt.fields.Signature,
			}

			if got := msg.Verify(); got != tt.want {
				t.Errorf("MsgLockCoins.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgReleaseCoins(t *testing.T) {
	type fields struct {
		LockerID  string
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{"test1", fields{lockerId2, pk2, sign2}, true},
		{"test2", fields{emptyLockerId, pk1, sign1}, false},
		{"test3", fields{emptyLockerId, emptyPubKey, sign1}, false},
		{"test4", fields{emptyLockerId, pk3, nil}, false},
	}

	for _, tc := range tests {
		msg := MsgReleaseCoins{tc.fields.LockerID, tc.fields.PubKey, tc.fields.Signature}

		if tc.expectedPass {
			require.Nil(t, msg.ValidateBasic())
		} else {
			require.NotNil(t, msg.ValidateBasic())
		}
	}
}

func TestMsgReleaseCoins_Verify(t *testing.T) {
	type fields struct {
		LockerID  string
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test1", fields{lockerId2, pk2, sign2}, true},
		{"test2", fields{lockerId, pk1, sign1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgReleaseCoins{
				LockerID:  tt.fields.LockerID,
				PubKey:    tt.fields.PubKey,
				Signature: tt.fields.Signature,
			}
			if got := msg.Verify(); got != tt.want {
				t.Errorf("MsgReleaseCoins.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMsgReleaseCoinsToMany(t *testing.T) {
	type fields struct {
		LockerID  string
		Addresses []csdkTypes.AccAddress
		Shares    []csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name         string
		fields       fields
		expectedPass bool
	}{
		{"test1", fields{emptyLockerId, []csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address()),
			csdkTypes.AccAddress(pk2.Address())}, []csdkTypes.Coins{coins1}, pk3, sign3}, false},

		{"test2", fields{lockerId3, []csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address()),
			csdkTypes.AccAddress(pk2.Address())}, []csdkTypes.Coins{{coin1}, coins1}, pk3, sign3}, true},

		{"test3", fields{lockerId, []csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address())},
			[]csdkTypes.Coins{{coin1}, coins1, coins2}, nil, sign1}, false},
	}

	for _, tc := range tests {
		msg := MsgReleaseCoinsToMany{tc.fields.LockerID, tc.fields.Addresses, tc.fields.Shares, tc.fields.PubKey, tc.fields.Signature}

		if !tc.expectedPass {
			require.NotNil(t, msg.ValidateBasic())
		} else {
			require.Nil(t, msg.ValidateBasic())
		}
	}
}

func TestMsgReleaseCoinsToMany_Verify(t *testing.T) {
	type fields struct {
		LockerID  string
		Addresses []csdkTypes.AccAddress
		Shares    []csdkTypes.Coins
		PubKey    crypto.PubKey
		Signature []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test1", fields{emptyLockerId, []csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address()),
			csdkTypes.AccAddress(pk2.Address())}, []csdkTypes.Coins{coins1}, pk3, sign3}, false},

		{"test2", fields{lockerId3, []csdkTypes.AccAddress{csdkTypes.AccAddress(pk1.Address()),
			csdkTypes.AccAddress(pk2.Address())}, []csdkTypes.Coins{{coin1}}, pk3, sign3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := MsgReleaseCoinsToMany{
				LockerID:  tt.fields.LockerID,
				Addresses: tt.fields.Addresses,
				Shares:    tt.fields.Shares,
				PubKey:    tt.fields.PubKey,
				Signature: tt.fields.Signature,
			}
			if got := msg.Verify(); got != tt.want {
				t.Errorf("MsgReleaseCoinsToMany.Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}

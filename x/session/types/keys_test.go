package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestSessionForAccountKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SessionForAccountKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SessionForAccountKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForAccountKey(addr, uint64(i))
		})
	}
}

func TestSessionForExpiryAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		at := time.Now()
		require.Equal(
			t,
			append(append(SessionForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SessionForExpiryAtKey(at, uint64(i)),
		)
	}
}

func TestSessionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(SessionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...),
			SessionKey(uint64(i)),
		)
	}
}

func TestIDFromSessionForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForAccountKey(key),
		)
	}
}

func TestIDFromSessionForExpiryAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SessionForExpiryAtKey(time.Now(), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForExpiryAtKey(key),
		)
	}
}

func TestIDFromSessionForNodeKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForNodeKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForNodeKey(key),
		)
	}
}

func TestIDFromSessionForAllocationKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SessionForAllocationKey(uint64(i+64), addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForAllocationKey(key),
		)
	}
}

func TestIDFromSessionForSubscriptionKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SessionForSubscriptionKey(uint64(i+64), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSessionForSubscriptionKey(key),
		)
	}
}

func TestSessionForNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SessionForNodeKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SessionForNodeKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForNodeKey(addr, uint64(i))
		})
	}
}

func TestSessionForAllocationKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(SessionForAllocationKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i+64))...),
				SessionForAllocationKey(uint64(i), addr, uint64(i+64)),
			)

			continue
		}

		require.Panics(t, func() {
			SessionForAllocationKey(uint64(i), addr, uint64(i+64))
		})
	}
}

func TestSessionForSubscriptionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SessionForSubscriptionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdk.Uint64ToBigEndian(uint64(i+64))...),
			SessionForSubscriptionKey(uint64(i), uint64(i+64)),
		)
	}
}

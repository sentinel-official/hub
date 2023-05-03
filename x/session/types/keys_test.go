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

func TestSessionExpiryAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		at := time.Now()
		require.Equal(
			t,
			append(append(SessionExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SessionExpiryAtKey(at, uint64(i)),
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

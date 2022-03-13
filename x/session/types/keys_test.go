package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActiveSessionForAddressKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(ActiveSessionForAddressKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				ActiveSessionForAddressKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveSessionForAddressKey(addr, 1000)
		})
	}
}

func TestInactiveSessionAtKey(t *testing.T) {
	var (
		at = time.Now()
	)

	require.Equal(
		t,
		append(append(InactiveSessionAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(1000)...),
		InactiveSessionAtKey(at, 1000),
	)
}

func TestInactiveSessionForAddressKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(InactiveSessionForAddressKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				InactiveSessionForAddressKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveSessionForAddressKey(addr, 1000)
		})
	}
}

func TestSessionKey(t *testing.T) {
	require.Equal(
		t,
		append(SessionKeyPrefix, sdk.Uint64ToBigEndian(1000)...),
		SessionKey(1000),
	)
}

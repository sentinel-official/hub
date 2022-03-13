package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestActiveSessionForAddressKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(ActiveSessionForAddressKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				ActiveSessionForAddressKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveSessionForAddressKey(address, 1000)
		})
	}
}

func TestIDFromStatusSessionAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 38 {
			require.Equal(
				t,
				sdk.BigEndianToUint64(key[30:]),
				IDFromStatusSessionAtKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromStatusSessionAtKey(key)
		})
	}
}

func TestIDFromStatusSessionForAddressKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 29 {
			require.Equal(
				t,
				sdk.BigEndianToUint64(key[21:]),
				IDFromStatusSessionForAddressKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromStatusSessionForAddressKey(key)
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
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(InactiveSessionForAddressKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				InactiveSessionForAddressKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveSessionForAddressKey(address, 1000)
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

package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestActiveSubscriptionForAddressKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(ActiveSubscriptionForAddressKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				ActiveSubscriptionForAddressKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveSubscriptionForAddressKey(address, 1000)
		})
	}
}

func TestIDFromInactiveSubscriptionAtKey(t *testing.T) {
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
				IDFromInactiveSubscriptionAtKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromInactiveSubscriptionAtKey(key)
		})
	}
}

func TestIDFromStatusSubscriptionForAddressKey(t *testing.T) {
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
				IDFromStatusSubscriptionForAddressKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromStatusSubscriptionForAddressKey(key)
		})
	}
}

func TestIDFromSubscriptionForNodeKey(t *testing.T) {
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
				IDFromSubscriptionForNodeKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromSubscriptionForNodeKey(key)
		})
	}
}

func TestIDFromSubscriptionForPlanKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 17 {
			require.Equal(
				t,
				sdk.BigEndianToUint64(key[9:]),
				IDFromSubscriptionForPlanKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromSubscriptionForPlanKey(key)
		})
	}
}

func TestInactiveSubscriptionAtKey(t *testing.T) {
	var (
		at = time.Now()
	)

	require.Equal(
		t,
		append(append(InactiveSubscriptionAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(1000)...),
		InactiveSubscriptionAtKey(at, 1000),
	)
}

func TestInactiveSubscriptionForAddressKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(InactiveSubscriptionForAddressKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				InactiveSubscriptionForAddressKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveSubscriptionForAddressKey(address, 1000)
		})
	}
}

func TestQuotaKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(1000)...), address...),
				QuotaKey(1000, address),
			)

			continue
		}

		require.Panics(t, func() {
			QuotaKey(1000, address)
		})
	}
}

func TestSubscriptionKey(t *testing.T) {
	require.Equal(
		t,
		append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(1000)...),
		SubscriptionKey(1000),
	)
}

package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActiveSubscriptionForAddressKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(ActiveSubscriptionForAddressKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				ActiveSubscriptionForAddressKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveSubscriptionForAddressKey(addr, 1000)
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
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(InactiveSubscriptionForAddressKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				InactiveSubscriptionForAddressKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveSubscriptionForAddressKey(addr, 1000)
		})
	}
}

func TestQuotaKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(1000)...), address.MustLengthPrefix(addr)...),
				QuotaKey(1000, addr),
			)

			continue
		}

		require.Panics(t, func() {
			QuotaKey(1000, addr)
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

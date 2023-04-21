package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestInactiveSubscriptionAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		at := time.Now()
		require.Equal(
			t,
			append(append(InactiveSubscriptionAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(uint64(i))...),
			InactiveSubscriptionAtKey(at, uint64(i)),
		)
	}
}

func TestQuotaKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(QuotaKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), address.MustLengthPrefix(addr)...),
				QuotaKey(uint64(i), addr),
			)

			continue
		}

		require.Panics(t, func() {
			QuotaKey(uint64(i), addr)
		})
	}
}

func TestSubscriptionForAccountKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SubscriptionForAccountKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SubscriptionForAccountKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SubscriptionForAccountKey(addr, uint64(i))
		})
	}
}

func TestSubscriptionForNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(SubscriptionForNodeKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(uint64(i))...),
				SubscriptionForNodeKey(addr, uint64(i)),
			)

			continue
		}

		require.Panics(t, func() {
			SubscriptionForNodeKey(addr, uint64(i))
		})
	}
}

func TestSubscriptionForPlanKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SubscriptionForPlanKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionForPlanKey(uint64(i), uint64(i)),
		)
	}
}

func TestSubscriptionKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(SubscriptionKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionKey(uint64(i)),
		)
	}
}

package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionForExpiryAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		at := time.Now()
		require.Equal(
			t,
			append(append(SubscriptionForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionForExpiryAtKey(at, uint64(i)),
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

func TestIDFromPayoutForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = PayoutForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromPayoutForAccountKey(key),
		)
	}
}

func TestIDFromPayoutForNodeKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = PayoutForNodeKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromPayoutForNodeKey(key),
		)
	}
}

func TestIDFromSubscriptionForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SubscriptionForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForAccountKey(key),
		)
	}
}

func TestIDFromSubscriptionForExpiryAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SubscriptionForExpiryAtKey(time.Now(), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForExpiryAtKey(key),
		)
	}
}

func TestIDFromSubscriptionForNodeKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = SubscriptionForNodeKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForNodeKey(key),
		)
	}
}

func TestIDFromSubscriptionForPlanKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SubscriptionForPlanKey(uint64(i+64), uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForPlanKey(key),
		)
	}
}

func TestPayoutForAccountKey(t *testing.T) {
	var (
		addr []byte
		id   uint64
	)

	for i := 1; i <= 512; i += 64 {
		id = uint64(i)
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(PayoutForAccountKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(id)...),
				PayoutForAccountKey(addr, id),
			)

			continue
		}

		require.Panics(t, func() {
			PayoutForAccountKey(addr, id)
		})
	}
}

func TestPayoutForNodeKey(t *testing.T) {
	var (
		addr []byte
		id   uint64
	)

	for i := 1; i <= 512; i += 64 {
		id = uint64(i)
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(PayoutForNodeKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(id)...),
				PayoutForNodeKey(addr, id),
			)

			continue
		}

		require.Panics(t, func() {
			PayoutForNodeKey(addr, id)
		})
	}
}

func TestPayoutKey(t *testing.T) {
	var (
		id uint64
	)

	for i := 1; i <= 512; i += 64 {
		id = uint64(i)
		require.Equal(
			t,
			append(PayoutKeyPrefix, sdk.Uint64ToBigEndian(id)...),
			PayoutKey(id),
		)
	}
}

package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestSubscriptionForInactiveAtKey(t *testing.T) {
	for i := 0; i < 512; i += 64 {
		require.Equal(
			t,
			append(append(SubscriptionForInactiveAtKeyPrefix, sdk.FormatTimeBytes(hubtypes.TestTimeNow)...), sdk.Uint64ToBigEndian(uint64(i))...),
			SubscriptionForInactiveAtKey(hubtypes.TestTimeNow, uint64(i)),
		)
	}
}

func TestAllocationKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(AllocationKeyPrefix, sdk.Uint64ToBigEndian(uint64(i))...), address.MustLengthPrefix(addr)...),
				AllocationKey(uint64(i), addr),
			)

			continue
		}

		require.Panics(t, func() {
			AllocationKey(uint64(i), addr)
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

func TestIDFromSubscriptionForInactiveAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = SubscriptionForInactiveAtKey(hubtypes.TestTimeNow, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromSubscriptionForInactiveAtKey(key),
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

func TestIDFromPayoutForAccountByNodeKey(t *testing.T) {
	var (
		accAddr  []byte
		nodeAddr []byte
		key      []byte
	)

	for i := 1; i <= 256; i += 64 {
		accAddr = make([]byte, i)
		_, _ = rand.Read(accAddr)

		nodeAddr = make([]byte, i)
		_, _ = rand.Read(nodeAddr)

		key = PayoutForAccountByNodeKey(accAddr, nodeAddr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromPayoutForAccountByNodeKey(key),
		)
	}
}

func TestIDFromPayoutForNextAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 1; i <= 256; i += 64 {
		key = PayoutForNextAtKey(hubtypes.TestTimeNow, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromPayoutForNextAtKey(key),
		)
	}
}

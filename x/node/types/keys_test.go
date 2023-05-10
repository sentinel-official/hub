package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestActiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(ActiveNodeKeyPrefix, address.MustLengthPrefix(addr)...),
				ActiveNodeKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveNodeKey(addr)
		})
	}
}

func TestAddressFromNodeForExpiryAtKey(t *testing.T) {
	var (
		at   = time.Now()
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = NodeForExpiryAtKey(at, addr)
		require.Equal(
			t,
			hubtypes.NodeAddress(addr),
			AddressFromNodeForExpiryAtKey(key),
		)
	}
}

func TestAddressFromNodeForPlanKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = NodeForPlanKey(uint64(i), addr)
		require.Equal(
			t,
			hubtypes.NodeAddress(addr),
			AddressFromNodeForPlanKey(key),
		)
	}
}

func TestIDFromLeaseForAccountKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = LeaseForAccountKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromLeaseForAccountKey(key),
		)
	}
}

func TestIDFromLeaseForNodeKey(t *testing.T) {
	var (
		addr []byte
		key  []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		key = LeaseForNodeKey(addr, uint64(i))
		require.Equal(
			t,
			uint64(i),
			IDFromLeaseForNodeKey(key),
		)
	}
}

func TestInactiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(InactiveNodeKeyPrefix, address.MustLengthPrefix(addr)...),
				InactiveNodeKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeKey(addr)
		})
	}
}

func TestLeaseForAccountKey(t *testing.T) {
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
				append(append(LeaseForAccountKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(id)...),
				LeaseForAccountKey(addr, id),
			)

			continue
		}

		require.Panics(t, func() {
			LeaseForAccountKey(addr, id)
		})
	}
}

func TestLeaseForNodeKey(t *testing.T) {
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
				append(append(LeaseForNodeKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(id)...),
				LeaseForNodeKey(addr, id),
			)

			continue
		}

		require.Panics(t, func() {
			LeaseForNodeKey(addr, id)
		})
	}
}

func TestLeaseKey(t *testing.T) {
	var (
		id uint64
	)

	for i := 1; i <= 512; i += 64 {
		id = uint64(i)
		require.Equal(
			t,
			append(LeaseKeyPrefix, sdk.Uint64ToBigEndian(id)...),
			LeaseKey(id),
		)
	}
}

func TestNodeForExpiryAtKey(t *testing.T) {
	var (
		at   = time.Now()
		addr []byte
	)

	for i := 1; i <= 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(NodeForExpiryAtKeyPrefix, sdk.FormatTimeBytes(at)...), address.MustLengthPrefix(addr)...),
				NodeForExpiryAtKey(at, addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForExpiryAtKey(at, addr)
		})
	}
}

func TestNodeForPlanKey(t *testing.T) {
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
				append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...), address.MustLengthPrefix(addr)...),
				NodeForPlanKey(id, addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForPlanKey(id, addr)
		})
	}
}

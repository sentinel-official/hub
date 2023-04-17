package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActiveNodeForPlanKey(t *testing.T) {
	var (
		addr []byte
		id   uint64
	)

	for i := 0; i < 512; i += 64 {
		id = uint64(i)
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...), 0x01), address.MustLengthPrefix(addr)...),
				ActiveNodeForPlanKey(id, addr),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveNodeForPlanKey(id, addr)
		})
	}
}

func TestActiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
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

func TestInactiveNodeAtKey(t *testing.T) {
	var (
		at   = time.Now()
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...), address.MustLengthPrefix(addr)...),
				InactiveNodeAtKey(at, addr),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeAtKey(at, addr)
		})
	}
}

func TestInactiveNodeForPlanKey(t *testing.T) {
	var (
		addr []byte
		id   uint64
	)

	for i := 0; i < 512; i += 64 {
		id = uint64(i)
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(id)...), 0x02), address.MustLengthPrefix(addr)...),
				InactiveNodeForPlanKey(id, addr),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeForPlanKey(id, addr)
		})
	}
}

func TestInactiveNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
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

package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActiveNodeForProviderKey(t *testing.T) {
	var (
		nodeAddr []byte
		provAddr []byte
	)

	for i := 0; i < 512; i += 64 {
		provAddr = make([]byte, i)
		_, _ = rand.Read(provAddr)

		for j := 0; j < 512; j += 64 {
			nodeAddr = make([]byte, j)
			_, _ = rand.Read(nodeAddr)

			if i < 256 && j < 256 {
				require.Equal(
					t,
					append(append(ActiveNodeForProviderKeyPrefix, address.MustLengthPrefix(provAddr)...), address.MustLengthPrefix(nodeAddr)...),
					ActiveNodeForProviderKey(provAddr, nodeAddr),
				)

				continue
			}

			require.Panics(t, func() {
				ActiveNodeForProviderKey(provAddr, nodeAddr)
			})
		}
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

func TestInactiveNodeForProviderKey(t *testing.T) {
	var (
		node     []byte
		provider []byte
	)

	for i := 0; i < 512; i += 64 {
		provider = make([]byte, i)
		_, _ = rand.Read(provider)

		for j := 0; j < 512; j += 64 {
			node = make([]byte, j)
			_, _ = rand.Read(node)

			if i < 256 && j < 256 {
				require.Equal(
					t,
					append(append(InactiveNodeForProviderKeyPrefix, address.MustLengthPrefix(provider)...), address.MustLengthPrefix(node)...),
					InactiveNodeForProviderKey(provider, node),
				)

				continue
			}

			require.Panics(t, func() {
				InactiveNodeForProviderKey(provider, node)
			})
		}
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

func TestNodeKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(NodeKeyPrefix, address.MustLengthPrefix(addr)...),
				NodeKey(addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeKey(addr)
		})
	}
}

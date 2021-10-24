package types

import (
	"crypto/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestActiveNodeForProviderKey(t *testing.T) {
	var (
		address  []byte
		provider []byte
	)

	for i := 0; i < 41; i++ {
		provider = make([]byte, i)
		_, _ = rand.Read(provider)

		for j := 0; j < 41; j++ {
			address = make([]byte, j)
			_, _ = rand.Read(address)

			if i == 20 && j == 20 {
				require.Equal(
					t,
					append(append(ActiveNodeForProviderKeyPrefix, provider...), address...),
					ActiveNodeForProviderKey(provider, address),
				)

				continue
			}

			require.Panics(t, func() {
				ActiveNodeForProviderKey(provider, address)
			})
		}
	}
}

func TestActiveNodeKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(ActiveNodeKeyPrefix, address...),
				ActiveNodeKey(address),
			)

			continue
		}

		require.Panics(t, func() {
			ActiveNodeKey(address)
		})
	}
}

func TestAddressFromStatusNodeAtKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 50 {
			require.Equal(
				t,
				key[30:],
				AddressFromStatusNodeAtKey(key).Bytes(),
			)

			continue
		}

		require.Panics(t, func() {
			AddressFromStatusNodeAtKey(key)
		})
	}
}

func TestAddressFromStatusNodeForProviderKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 41 {
			require.Equal(
				t,
				key[21:],
				AddressFromStatusNodeForProviderKey(key).Bytes(),
			)

			continue
		}

		require.Panics(t, func() {
			AddressFromStatusNodeForProviderKey(key)
		})
	}
}

func TestAddressFromStatusNodeKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 21 {
			require.Equal(
				t,
				key[1:],
				AddressFromStatusNodeKey(key).Bytes(),
			)

			continue
		}

		require.Panics(t, func() {
			AddressFromStatusNodeKey(key)
		})
	}
}

func TestInactiveNodeAtKey(t *testing.T) {
	var (
		at      = time.Now()
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...), address...),
				InactiveNodeAtKey(at, address),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeAtKey(at, address)
		})
	}
}

func TestInactiveNodeForProviderKey(t *testing.T) {
	var (
		address  []byte
		provider []byte
	)

	for i := 0; i < 41; i++ {
		provider = make([]byte, i)
		_, _ = rand.Read(provider)

		for j := 0; j < 41; j++ {
			address = make([]byte, j)
			_, _ = rand.Read(address)

			if i == 20 && j == 20 {
				require.Equal(
					t,
					append(append(InactiveNodeForProviderKeyPrefix, provider...), address...),
					InactiveNodeForProviderKey(provider, address),
				)

				continue
			}

			require.Panics(t, func() {
				InactiveNodeForProviderKey(provider, address)
			})
		}
	}
}

func TestInactiveNodeKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(InactiveNodeKeyPrefix, address...),
				InactiveNodeKey(address),
			)

			continue
		}

		require.Panics(t, func() {
			InactiveNodeKey(address)
		})
	}
}

func TestNodeKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(NodeKeyPrefix, address...),
				NodeKey(address),
			)

			continue
		}

		require.Panics(t, func() {
			NodeKey(address)
		})
	}
}

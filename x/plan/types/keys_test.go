package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestActivePlanForProviderKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(ActivePlanForProviderKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				ActivePlanForProviderKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActivePlanForProviderKey(address, 1000)
		})
	}
}

func TestActivePlanKey(t *testing.T) {
	require.Equal(
		t,
		append(ActivePlanKeyPrefix, sdk.Uint64ToBigEndian(1000)...),
		ActivePlanKey(1000),
	)
}

func TestAddressFromNodeForPlanKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 29 {
			require.Equal(
				t,
				key[9:],
				AddressFromNodeForPlanKey(key).Bytes(),
			)

			continue
		}

		require.Panics(t, func() {
			AddressFromNodeForPlanKey(key)
		})
	}
}

func TestCountForNodeByProviderKey(t *testing.T) {
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
					append(append(CountForNodeByProviderKeyPrefix, provider...), address...),
					CountForNodeByProviderKey(provider, address),
				)

				continue
			}

			require.Panics(t, func() {
				CountForNodeByProviderKey(provider, address)
			})
		}
	}
}

func TestIDFromStatusPlanForProviderKey(t *testing.T) {
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
				IDFromStatusPlanForProviderKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromStatusPlanForProviderKey(key)
		})
	}
}

func TestIDFromStatusPlanKey(t *testing.T) {
	var (
		key []byte
	)

	for i := 0; i < 60; i++ {
		key = make([]byte, i)
		_, _ = rand.Read(key)

		if i == 9 {
			require.Equal(
				t,
				sdk.BigEndianToUint64(key[1:]),
				IDFromStatusPlanKey(key),
			)

			continue
		}

		require.Panics(t, func() {
			IDFromStatusPlanKey(key)
		})
	}
}

func TestInactivePlanForProviderKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(InactivePlanForProviderKeyPrefix, address...), sdk.Uint64ToBigEndian(1000)...),
				InactivePlanForProviderKey(address, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactivePlanForProviderKey(address, 1000)
		})
	}
}

func TestInactivePlanKey(t *testing.T) {
	require.Equal(
		t,
		append(InactivePlanKeyPrefix, sdk.Uint64ToBigEndian(1000)...),
		InactivePlanKey(1000),
	)
}

func TestNodeForPlanKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(1000)...), address...),
				NodeForPlanKey(1000, address),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForPlanKey(1000, address)
		})
	}
}

func TestPlanKey(t *testing.T) {
	require.Equal(
		t,
		append(PlanKeyPrefix, sdk.Uint64ToBigEndian(1000)...),
		PlanKey(1000),
	)
}

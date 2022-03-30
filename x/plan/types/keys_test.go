package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/stretchr/testify/require"
)

func TestActivePlanForProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(ActivePlanForProviderKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				ActivePlanForProviderKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			ActivePlanForProviderKey(addr, 1000)
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

func TestCountForNodeByProviderKey(t *testing.T) {
	var (
		node     []byte
		provider []byte
	)

	for i := 0; i < 512; i++ {
		provider = make([]byte, i)
		_, _ = rand.Read(provider)

		for j := 0; j < 512; j++ {
			node = make([]byte, j)
			_, _ = rand.Read(node)

			if i < 256 && j < 256 {
				require.Equal(
					t,
					append(append(CountForNodeByProviderKeyPrefix, address.MustLengthPrefix(provider)...), address.MustLengthPrefix(node)...),
					CountForNodeByProviderKey(provider, node),
				)

				continue
			}

			require.Panics(t, func() {
				CountForNodeByProviderKey(provider, node)
			})
		}
	}
}

func TestInactivePlanForProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(InactivePlanForProviderKeyPrefix, address.MustLengthPrefix(addr)...), sdk.Uint64ToBigEndian(1000)...),
				InactivePlanForProviderKey(addr, 1000),
			)

			continue
		}

		require.Panics(t, func() {
			InactivePlanForProviderKey(addr, 1000)
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
		addr []byte
	)

	for i := 0; i < 512; i++ {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(NodeForPlanKeyPrefix, sdk.Uint64ToBigEndian(1000)...), address.MustLengthPrefix(addr)...),
				NodeForPlanKey(1000, addr),
			)

			continue
		}

		require.Panics(t, func() {
			NodeForPlanKey(1000, addr)
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

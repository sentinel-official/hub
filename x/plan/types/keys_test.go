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

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(PlanForProviderKeyPrefix, address.MustLengthPrefix(addr)...), 0x01), sdk.Uint64ToBigEndian(1000)...),
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
		append(append(PlanKeyPrefix, 0x01), sdk.Uint64ToBigEndian(1000)...),
		ActivePlanKey(1000),
	)
}

func TestInactivePlanForProviderKey(t *testing.T) {
	var (
		addr []byte
	)

	for i := 0; i < 512; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		if i < 256 {
			require.Equal(
				t,
				append(append(append(PlanForProviderKeyPrefix, address.MustLengthPrefix(addr)...), 0x02), sdk.Uint64ToBigEndian(1000)...),
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
		append(append(PlanKeyPrefix, 0x02), sdk.Uint64ToBigEndian(1000)...),
		InactivePlanKey(1000),
	)
}

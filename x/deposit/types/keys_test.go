package types

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDepositKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(DepositKeyPrefix, address...),
				DepositKey(address),
			)

			continue
		}

		require.Panics(t, func() {
			DepositKey(address)
		})
	}
}

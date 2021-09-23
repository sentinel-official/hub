package types

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProviderKey(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 41; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		if i == 20 {
			require.Equal(
				t,
				append(ProviderKeyPrefix, address...),
				ProviderKey(address),
			)

			continue
		}

		require.Panics(t, func() {
			ProviderKey(address)
		})
	}
}

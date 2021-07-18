package types

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwapKey(t *testing.T) {
	bytes := make([]byte, 32)
	_, _ = rand.Read(bytes)

	require.Equal(
		t,
		append(SwapKeyPrefix, bytes...),
		SwapKey(BytesToHash(bytes)),
	)
}

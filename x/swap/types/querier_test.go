package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(
		t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

func TestNewQuerySwapRequest(t *testing.T) {
	var (
		bytes []byte
	)

	for i := 0; i < 20; i++ {
		bytes = make([]byte, i)
		_, _ = rand.Read(bytes)

		require.Equal(
			t,
			&QuerySwapRequest{
				TxHash: BytesToHash(bytes).Bytes(),
			},
			NewQuerySwapRequest(BytesToHash(bytes)),
		)
	}
}

func TestNewQuerySwapsRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 20; i++ {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		require.Equal(
			t,
			&QuerySwapsRequest{
				Pagination: pagination,
			},
			NewQuerySwapsRequest(pagination),
		)
	}
}

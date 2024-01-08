package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

func TestNewQueryNodeRequest(t *testing.T) {
	var (
		addr []byte
	)

	for i := 1; i <= 256; i += 64 {
		addr = make([]byte, i)
		_, _ = rand.Read(addr)

		require.Equal(
			t,
			&QueryNodeRequest{
				Address: hubtypes.NodeAddress(addr).String(),
			},
			NewQueryNodeRequest(addr),
		)
	}
}

func TestNewQueryNodesRequest(t *testing.T) {
	var (
		status     hubtypes.Status
		pagination *query.PageRequest
	)

	for i := 1; i <= 16; i++ {
		status = hubtypes.Status(i % 4)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryNodesRequest{
				Status:     status,
				Pagination: pagination,
			},
			NewQueryNodesRequest(status, pagination),
		)
	}
}

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

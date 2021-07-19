package types

import (
	"crypto/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestNewQueryNodeRequest(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		require.Equal(
			t,
			&QueryNodeRequest{
				Address: hubtypes.NodeAddress(address).String(),
			},
			NewQueryNodeRequest(address),
		)
	}
}

func TestNewQueryNodesForProviderRequest(t *testing.T) {
	var (
		address    []byte
		status     hubtypes.Status
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		status = hubtypes.Status(i % 4)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(address)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QueryNodesForProviderRequest{
				Address:    hubtypes.ProvAddress(address).String(),
				Status:     status,
				Pagination: pagination,
			},
			NewQueryNodesForProviderRequest(address, status, pagination),
		)
	}
}

func TestNewQueryNodesRequest(t *testing.T) {
	var (
		status     hubtypes.Status
		pagination *query.PageRequest
	)

	for i := 0; i < 20; i++ {
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

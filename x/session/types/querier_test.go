package types

import (
	"crypto/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestNewQueryParamsRequest(t *testing.T) {
	require.Equal(t,
		&QueryParamsRequest{},
		NewQueryParamsRequest(),
	)
}

func TestNewQuerySessionRequest(t *testing.T) {
	for i := 0; i < 20; i++ {
		require.Equal(
			t,
			&QuerySessionRequest{
				Id: uint64(i),
			},
			NewQuerySessionRequest(uint64(i)),
		)
	}
}

func TestNewQuerySessionsForAccountRequest(t *testing.T) {
	var (
		addr       []byte
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		addr = make([]byte, i)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(addr)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySessionsForAccountRequest{
				Address:    sdk.AccAddress(addr).String(),
				Pagination: pagination,
			},
			NewQuerySessionsForAccountRequest(addr, pagination),
		)
	}
}

func TestNewQuerySessionsRequest(t *testing.T) {
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
			&QuerySessionsRequest{
				Pagination: pagination,
			},
			NewQuerySessionsRequest(pagination),
		)
	}
}

func TestNewQuerySessionsForNodeRequest(t *testing.T) {
	var (
		addr       []byte
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		addr = make([]byte, i)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(addr)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySessionsForNodeRequest{
				Address:    hubtypes.NodeAddress(addr).String(),
				Pagination: pagination,
			},
			NewQuerySessionsForNodeRequest(addr, pagination),
		)
	}
}

func TestNewQuerySessionsForQuotaRequest(t *testing.T) {
	var (
		addr       []byte
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		addr = make([]byte, i)
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(addr)
		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySessionsForQuotaRequest{
				Id:         uint64(i),
				Address:    sdk.AccAddress(addr).String(),
				Pagination: pagination,
			},
			NewQuerySessionsForQuotaRequest(uint64(i), addr, pagination),
		)
	}
}

func TestNewQuerySessionsForSubscriptionRequest(t *testing.T) {
	var (
		pagination *query.PageRequest
	)

	for i := 0; i < 40; i++ {
		pagination = &query.PageRequest{
			Key:        make([]byte, i),
			Offset:     uint64(i),
			Limit:      uint64(i),
			CountTotal: i/2 == 0,
		}

		_, _ = rand.Read(pagination.Key)

		require.Equal(
			t,
			&QuerySessionsForSubscriptionRequest{
				Id:         uint64(i),
				Pagination: pagination,
			},
			NewQuerySessionsForSubscriptionRequest(uint64(i), pagination),
		)
	}
}

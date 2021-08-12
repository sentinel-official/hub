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

func TestNewQuerySessionsForAddressRequest(t *testing.T) {
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
			&QuerySessionsForAddressRequest{
				Address:    sdk.AccAddress(address).String(),
				Status:     status,
				Pagination: pagination,
			},
			NewQuerySessionsForAddressRequest(address, status, pagination),
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

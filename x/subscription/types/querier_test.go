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

func TestNewQueryQuotaRequest(t *testing.T) {
	var (
		address []byte
	)

	for i := 0; i < 40; i++ {
		address = make([]byte, i)
		_, _ = rand.Read(address)

		require.Equal(
			t,
			&QueryQuotaRequest{
				Id:      uint64(i),
				Address: sdk.AccAddress(address).String(),
			},
			NewQueryQuotaRequest(uint64(i), address),
		)
	}
}

func TestNewQueryQuotasRequest(t *testing.T) {
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
			&QueryQuotasRequest{
				Id:         uint64(i),
				Pagination: pagination,
			},
			NewQueryQuotasRequest(uint64(i), pagination),
		)
	}
}

func TestNewQuerySubscriptionRequest(t *testing.T) {
	for i := 0; i < 20; i++ {
		require.Equal(
			t,
			&QuerySubscriptionRequest{
				Id: uint64(i),
			},
			NewQuerySubscriptionRequest(uint64(i)),
		)
	}
}

func TestNewQuerySubscriptionsForAddressRequest(t *testing.T) {
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
			&QuerySubscriptionsForAddressRequest{
				Address:    sdk.AccAddress(address).String(),
				Status:     status,
				Pagination: pagination,
			},
			NewQuerySubscriptionsForAddressRequest(address, status, pagination),
		)
	}
}

func TestNewQuerySubscriptionsRequest(t *testing.T) {
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
			&QuerySubscriptionsRequest{
				Pagination: pagination,
			},
			NewQuerySubscriptionsRequest(pagination),
		)
	}
}

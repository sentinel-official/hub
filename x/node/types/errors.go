// DO NOT COVER

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorDuplicate        = errors.Register(ModuleName, 201, "duplicate")
	ErrorInvalidGigabytes = errors.Register(ModuleName, 202, "invalid gigabytes")
	ErrorInvalidHours     = errors.Register(ModuleName, 203, "invalid hours")
	ErrorInvalidPrices    = errors.Register(ModuleName, 204, "invalid prices")
	ErrorNotFound         = errors.Register(ModuleName, 205, "not found")
)

func NewErrorInvalidGigabytePrices(prices sdk.Coins) error {
	return errors.Wrapf(ErrorInvalidPrices, "invalid gigabyte prices %s", prices)
}

func NewErrorInvalidHourlyPrices(prices sdk.Coins) error {
	return errors.Wrapf(ErrorInvalidPrices, "invalid hourly prices %s", prices)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNotFound, "node %s does not exist", addr)
}

func NewErrorDuplicateNode(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorDuplicate, "node %s already exists", addr)
}

func NewErrorHourlyPriceNotFound(denom string) error {
	return errors.Wrapf(ErrorNotFound, "hourly price for denom %s does not exist", denom)
}

func NewErrorGigabytePriceNotFound(denom string) error {
	return errors.Wrapf(ErrorNotFound, "gigabyte price for denom %s does not exist", denom)
}

func NewErrorInvalidLeaseGigabytes(gigabytes int64) error {
	return errors.Wrapf(ErrorInvalidGigabytes, "invalid lease gigabytes %d", gigabytes)
}

func NewErrorInvalidLeaseHours(hours int64) error {
	return errors.Wrapf(ErrorInvalidHours, "invalid lease hours %d", hours)
}

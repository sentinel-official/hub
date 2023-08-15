// DO NOT COVER

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode    = errors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidGigabytes = errors.Register(ModuleName, 202, "invalid gigabytes")
	ErrorInvalidHours     = errors.Register(ModuleName, 203, "invalid hours")
	ErrorInvalidPrices    = errors.Register(ModuleName, 204, "invalid prices")
	ErrorNodeNotFound     = errors.Register(ModuleName, 205, "node not found")
)

func NewErrorDuplicateNode(addr interface{}) error {
	return errors.Wrapf(ErrorDuplicateNode, "node %s already exist", addr)
}

func NewErrorInvalidGigabytes(gigabytes int64) error {
	return errors.Wrapf(ErrorInvalidGigabytes, "invalid gigabytes %d", gigabytes)
}

func NewErrorInvalidHours(hours int64) error {
	return errors.Wrapf(ErrorInvalidHours, "invalid hours %d", hours)
}

func NewErrorInvalidPrices(prices sdk.Coins) error {
	return errors.Wrapf(ErrorInvalidPrices, "invalid prices %s", prices)
}

func NewErrorNodeNotFound(addr interface{}) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

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
	ErrorInvalidPrices = errors.Register(ModuleName, 201, "invalid prices")
	ErrorDuplicateNode = errors.Register(ModuleName, 202, "duplicate node")
	ErrorNodeNotFound  = errors.Register(ModuleName, 203, "node not found")
)

func NewErrorInvalidGigabytePrices(prices sdk.Coins) error {
	return errors.Wrapf(ErrorInvalidPrices, "invalid gigabyte prices %s", prices)
}

func NewErrorInvalidHourlyPrices(prices sdk.Coins) error {
	return errors.Wrapf(ErrorInvalidPrices, "invalid hourly prices %s", prices)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorDuplicateNode(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorDuplicateNode, "node %s already exists", addr)
}

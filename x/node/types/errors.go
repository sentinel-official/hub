// DO NOT COVER

package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	ErrorInvalidMessage = sdkerrors.Register(ModuleName, 101, "invalid message")

	ErrorDuplicateNode    = sdkerrors.Register(ModuleName, 201, "duplicate node")
	ErrorInvalidGigabytes = sdkerrors.Register(ModuleName, 202, "invalid gigabytes")
	ErrorInvalidHours     = sdkerrors.Register(ModuleName, 203, "invalid hours")
	ErrorInvalidPrices    = sdkerrors.Register(ModuleName, 204, "invalid prices")
	ErrorNodeNotFound     = sdkerrors.Register(ModuleName, 205, "node not found")
)

func NewErrorDuplicateNode(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorDuplicateNode, "node %s already exist", addr)
}

func NewErrorInvalidGigabytes(gigabytes int64) error {
	return sdkerrors.Wrapf(ErrorInvalidGigabytes, "invalid gigabytes %d", gigabytes)
}

func NewErrorInvalidHours(hours int64) error {
	return sdkerrors.Wrapf(ErrorInvalidHours, "invalid hours %d", hours)
}

func NewErrorInvalidPrices(prices sdk.Coins) error {
	return sdkerrors.Wrapf(ErrorInvalidPrices, "invalid prices %s", prices)
}

func NewErrorNodeNotFound(addr interface{}) error {
	return sdkerrors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

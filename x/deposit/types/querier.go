package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	QueryDepositOfAddress = "deposit_of_address"
	QueryAllDeposits      = "all_deposits"
)

type QueryDepositOfAddressPrams struct {
	Address sdk.AccAddress
}

func NewQueryDepositOfAddressParams(address sdk.AccAddress) QueryDepositOfAddressPrams {
	return QueryDepositOfAddressPrams{
		Address: address,
	}
}

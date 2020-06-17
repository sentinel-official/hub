package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "provider"
	QuerierRoute = ModuleName
)

var (
	RouterKey = ""
)

var (
	ProvidersCountKey             = []byte{0x00}
	ProviderKeyPrefix             = []byte{0x01}
	ProviderIDForAddressKeyPrefix = []byte{0x02}
)

func ProviderKey(id hub.ProviderID) []byte {
	return append(ProviderKeyPrefix, id.Bytes()...)
}

func ProviderIDForAddressKey(address sdk.AccAddress) []byte {
	return append(ProviderIDForAddressKeyPrefix, address.Bytes()...)
}

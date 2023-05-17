package types

import (
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName = "provider"
)

var (
	ProviderKeyPrefix         = []byte{0x10}
	ActiveProviderKeyPrefix   = append(ProviderKeyPrefix, 0x01)
	InactiveProviderKeyPrefix = append(ProviderKeyPrefix, 0x02)
)

func ActiveProviderKey(addr hubtypes.ProvAddress) []byte {
	return append(ActiveProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveProviderKey(addr hubtypes.ProvAddress) (v []byte) {
	return append(InactiveProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

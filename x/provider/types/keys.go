package types

import (
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "provider"
	QuerierRoute = ModuleName
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgRegisterRequest = ModuleName + ":register"
	TypeMsgUpdateRequest   = ModuleName + ":update"
)

var (
	ProviderKeyPrefix = []byte{0x10}
)

func ProviderKey(addr hubtypes.ProvAddress) []byte {
	return append(ProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

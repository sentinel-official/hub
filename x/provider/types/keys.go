package types

import (
	"fmt"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName       = "provider"
	QuerierRoute     = ModuleName
	AddrLen      int = 20
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

func ProviderKey(address hubtypes.ProvAddress) []byte {
	v := append(ProviderKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

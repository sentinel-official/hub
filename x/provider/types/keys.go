package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "provider"
	QuerierRoute = ModuleName
	AddrLen      = 20
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

package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "node"
	QuerierRoute = ModuleName
	AddrLen	     = 20
)

var (
	ParamsSubspace = ModuleName
	RouterKey      = ModuleName
	StoreKey       = ModuleName
)

var (
	TypeMsgRegisterRequest  = ModuleName + ":register"
	TypeMsgUpdateRequest    = ModuleName + ":update"
	TypeMsgSetStatusRequest = ModuleName + ":set_status"
)

var (
	NodeKeyPrefix                    = []byte{0x10}
	ActiveNodeKeyPrefix              = []byte{0x20}
	InactiveNodeKeyPrefix            = []byte{0x21}
	ActiveNodeForProviderKeyPrefix   = []byte{0x30}
	InactiveNodeForProviderKeyPrefix = []byte{0x31}
	InactiveNodeAtKeyPrefix          = []byte{0x41}
)

func NodeKey(address hubtypes.NodeAddress) []byte {
	v := append(NodeKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func ActiveNodeKey(address hubtypes.NodeAddress) []byte {
	v := append(ActiveNodeKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func InactiveNodeKey(address hubtypes.NodeAddress) []byte {
	v := append(InactiveNodeKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func GetActiveNodeForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	v := append(ActiveNodeForProviderKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func ActiveNodeForProviderKey(provider hubtypes.ProvAddress, address hubtypes.NodeAddress) []byte {
	v := append(GetActiveNodeForProviderKeyPrefix(provider), address.Bytes()...)
	if len(v) != 1+2*AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+2*AddrLen))
	}

	return v
}

func GetInactiveNodeForProviderKeyPrefix(address hubtypes.ProvAddress) []byte {
	v := append(InactiveNodeForProviderKeyPrefix, address.Bytes()...)
	if len(v) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+AddrLen))
	}

	return v
}

func InactiveNodeForProviderKey(provider hubtypes.ProvAddress, address hubtypes.NodeAddress) []byte {
	v := append(GetInactiveNodeForProviderKeyPrefix(provider), address.Bytes()...)
	if len(v) != 1+2*AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+2*AddrLen))
	}

	return v
}

func GetInactiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveNodeAtKey(at time.Time, address hubtypes.NodeAddress) []byte {
	v := append(GetInactiveNodeAtKeyPrefix(at), address.Bytes()...)
	if len(v) != 1+29+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(v), 1+29+AddrLen))
	}

	return v
}

func AddressFromStatusNodeKey(key []byte) hubtypes.NodeAddress {
	if len(key) != 1+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+AddrLen))
	}

	return key[1:]
}

func AddressFromStatusNodeForProviderKey(key []byte) hubtypes.NodeAddress {
	if len(key) != 1+2*AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+2*AddrLen))
	}

	return key[1+sdk.AddrLen:]
}

func AddressFromStatusNodeAtKey(key []byte) hubtypes.NodeAddress {
	if len(key) != 1+29+AddrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 1+29+AddrLen))
	}

	return key[1+29:]
}

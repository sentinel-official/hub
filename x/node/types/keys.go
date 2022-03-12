package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	ModuleName   = "node"
	QuerierRoute = ModuleName
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

func NodeKey(addr hubtypes.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func ActiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeKey(addr hubtypes.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func GetActiveNodeForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(ActiveNodeForProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func ActiveNodeForProviderKey(provider hubtypes.ProvAddress, node hubtypes.NodeAddress) []byte {
	return append(GetActiveNodeForProviderKeyPrefix(provider), address.MustLengthPrefix(node.Bytes())...)
}

func GetInactiveNodeForProviderKeyPrefix(addr hubtypes.ProvAddress) []byte {
	return append(InactiveNodeForProviderKeyPrefix, address.MustLengthPrefix(addr.Bytes())...)
}

func InactiveNodeForProviderKey(provider hubtypes.ProvAddress, node hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeForProviderKeyPrefix(provider), address.MustLengthPrefix(node.Bytes())...)
}

func GetInactiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveNodeAtKey(at time.Time, addr hubtypes.NodeAddress) []byte {
	return append(GetInactiveNodeAtKeyPrefix(at), address.MustLengthPrefix(addr.Bytes())...)
}

func AddressFromStatusNodeKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | addrLen (1 byte) | addr

	addrLen := int(key[1])
	if len(key) != 2+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 2+addrLen))
	}

	return key[2:]
}

func AddressFromStatusNodeForProviderKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | providerLen (1 byte) | provider | nodeLen (1 byte) | node

	var (
		providerLen = int(key[1])
		nodeLen     = int(key[2+providerLen])
	)

	if len(key) != 3+providerLen+nodeLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 3+providerLen+nodeLen))
	}

	return key[3+providerLen:]
}

func AddressFromStatusNodeAtKey(key []byte) hubtypes.NodeAddress {
	// prefix (1 byte) | at (29 bytes) | addrLen (1 byte) | addr

	addrLen := int(key[30])
	if len(key) != 31+addrLen {
		panic(fmt.Errorf("invalid key length %d; expected %d", len(key), 31+addrLen))
	}

	return key[31:]
}

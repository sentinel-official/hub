package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuf "github.com/gogo/protobuf/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func NewDecoderStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.NodeKeyPrefix):
			var nodeA, nodeB types.Node
			cdc.MustUnmarshalBinaryBare(kvA.Value, &nodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &nodeB)
			return fmt.Sprintf("%s\n%s", nodeA, nodeB)

		case bytes.Equal(kvA.Key[:1], types.ActiveNodeKeyPrefix):
			var activeNodeA, activeNodeB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeNodeB)
			return fmt.Sprintf("%s\n%s", &activeNodeA, &activeNodeB)

		case bytes.Equal(kvA.Key[:1], types.InactiveNodeKeyPrefix):
			var inActiveNodeA, inActiveNodeB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inActiveNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inActiveNodeB)
			return fmt.Sprintf("%s\n%s", &inActiveNodeA, &inActiveNodeB)

		case bytes.Equal(kvA.Key[:1], types.ActiveNodeForProviderKeyPrefix):
			var activeNodeForProviderA, activeNodeForProviderB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeNodeForProviderB)
			return fmt.Sprintf("%s\n%s", &activeNodeForProviderA, &activeNodeForProviderB)

		case bytes.Equal(kvA.Key[:1], types.InactiveNodeForProviderKeyPrefix):
			var inactiveNodeForProviderA, inactiveNodeForProviderB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeForProviderB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeForProviderA, &inactiveNodeForProviderB)

		case bytes.Equal(kvA.Key[:1], types.InactiveNodeAtKeyPrefix):
			var inactiveNodeAtPrefixA, inactiveNodeAtPrefixB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeAtPrefixA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeAtPrefixB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeAtPrefixA, &inactiveNodeAtPrefixB)
		}

		panic(fmt.Sprintf("invalid node key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func NewStoreDecoder(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.NodeKeyPrefix):
			var nodeA, nodeB types.Node
			cdc.MustUnmarshalBinaryBare(kvA.Value, &nodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &nodeB)

			return fmt.Sprintf("%v\n%v", nodeA, nodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeKeyPrefix):
			var activeNodeA, activeNodeB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeNodeB)

			return fmt.Sprintf("%v\n%v", &activeNodeA, &activeNodeB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeKeyPrefix):
			var inactiveNodeA, inactiveNodeB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeA, &inactiveNodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeForProviderKeyPrefix):
			var activeNodeForProviderA, activeNodeForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &activeNodeForProviderA, &activeNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeForProviderKeyPrefix):
			var inactiveNodeForProviderA, inactiveNodeForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeForProviderA, &inactiveNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeAtKeyPrefix):
			var inactiveNodeAtA, inactiveNodeAtB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeAtA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeAtB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeAtA, &inactiveNodeAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

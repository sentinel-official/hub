package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.NodeKeyPrefix):
			var nodeA, nodeB types.Node
			cdc.MustUnmarshal(kvA.Value, &nodeA)
			cdc.MustUnmarshal(kvB.Value, &nodeB)

			return fmt.Sprintf("%v\n%v", nodeA, nodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeKeyPrefix):
			var activeNodeA, activeNodeB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &activeNodeA)
			cdc.MustUnmarshal(kvB.Value, &activeNodeB)

			return fmt.Sprintf("%v\n%v", &activeNodeA, &activeNodeB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeKeyPrefix):
			var inactiveNodeA, inactiveNodeB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveNodeA)
			cdc.MustUnmarshal(kvB.Value, &inactiveNodeB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeA, &inactiveNodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeForProviderKeyPrefix):
			var activeNodeForProviderA, activeNodeForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &activeNodeForProviderA)
			cdc.MustUnmarshal(kvB.Value, &activeNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &activeNodeForProviderA, &activeNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeForProviderKeyPrefix):
			var inactiveNodeForProviderA, inactiveNodeForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveNodeForProviderA)
			cdc.MustUnmarshal(kvB.Value, &inactiveNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeForProviderA, &inactiveNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeAtKeyPrefix):
			var inactiveNodeAtA, inactiveNodeAtB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveNodeAtA)
			cdc.MustUnmarshal(kvB.Value, &inactiveNodeAtB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeAtA, &inactiveNodeAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/node/types"
)

func NewStoreDecoder(appCodec codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.NodeKeyPrefix):
			var nodeA, nodeB types.Node
			appCodec.MustUnmarshal(kvA.Value, &nodeA)
			appCodec.MustUnmarshal(kvB.Value, &nodeB)

			return fmt.Sprintf("%v\n%v", nodeA, nodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeKeyPrefix):
			var activeNodeA, activeNodeB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &activeNodeA)
			appCodec.MustUnmarshal(kvB.Value, &activeNodeB)

			return fmt.Sprintf("%v\n%v", &activeNodeA, &activeNodeB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeKeyPrefix):
			var inactiveNodeA, inactiveNodeB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &inactiveNodeA)
			appCodec.MustUnmarshal(kvB.Value, &inactiveNodeB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeA, &inactiveNodeB)
		case bytes.Equal(kvA.Key[:1], types.ActiveNodeForProviderKeyPrefix):
			var activeNodeForProviderA, activeNodeForProviderB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &activeNodeForProviderA)
			appCodec.MustUnmarshal(kvB.Value, &activeNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &activeNodeForProviderA, &activeNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeForProviderKeyPrefix):
			var inactiveNodeForProviderA, inactiveNodeForProviderB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &inactiveNodeForProviderA)
			appCodec.MustUnmarshal(kvB.Value, &inactiveNodeForProviderB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeForProviderA, &inactiveNodeForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactiveNodeAtKeyPrefix):
			var inactiveNodeAtA, inactiveNodeAtB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &inactiveNodeAtA)
			appCodec.MustUnmarshal(kvB.Value, &inactiveNodeAtB)

			return fmt.Sprintf("%v\n%v", &inactiveNodeAtA, &inactiveNodeAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

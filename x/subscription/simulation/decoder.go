package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuf "github.com/gogo/protobuf/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func NewDecoderStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuf.Int64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)
			return fmt.Sprintf("%s\n%s", &countA, &countB)

		case bytes.Equal(kvA.Key[:1], types.SubscriptionKeyPrefix):
			var subscriptionA, subscriptionB types.Subscription
			cdc.MustUnmarshalBinaryBare(kvA.Value, &subscriptionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &subscriptionB)
			return fmt.Sprintf("%s\n%s", &subscriptionA, &subscriptionB)

		case bytes.Equal(kvA.Key[:1], types.SubscriptionForNodeKeyPrefix):
			var inActiveNodeA, inActiveNodeB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inActiveNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inActiveNodeB)
			return fmt.Sprintf("%s\n%s", &inActiveNodeA, &inActiveNodeB)

		case bytes.Equal(kvA.Key[:1], types.SubscriptionForPlanKeyPrefix):
			var activeNodeForProviderA, activeNodeForProviderB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeNodeForProviderB)
			return fmt.Sprintf("%s\n%s", &activeNodeForProviderA, &activeNodeForProviderB)

		case bytes.Equal(kvA.Key[:1], types.ActiveSubscriptionForAddressKeyPrefix):
			var inactiveNodeForProviderA, inactiveNodeForProviderB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeForProviderB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeForProviderA, &inactiveNodeForProviderB)

		case bytes.Equal(kvA.Key[:1], types.InactiveSubscriptionForAddressKeyPrefix):
			var inactiveNodeAtPrefixA, inactiveNodeAtPrefixB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeAtPrefixA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeAtPrefixB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeAtPrefixA, &inactiveNodeAtPrefixB)

		case bytes.Equal(kvA.Key[:1], types.InactiveSubscriptionAtKeyPrefix):
			var inactiveNodeAtPrefixA, inactiveNodeAtPrefixB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeAtPrefixA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeAtPrefixB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeAtPrefixA, &inactiveNodeAtPrefixB)

		case bytes.Equal(kvA.Key[:1], types.QuotaKeyPrefix):
			var inactiveNodeAtPrefixA, inactiveNodeAtPrefixB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveNodeAtPrefixA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveNodeAtPrefixB)
			return fmt.Sprintf("%s\n%s", &inactiveNodeAtPrefixA, &inactiveNodeAtPrefixB)
		}

		panic(fmt.Sprintf("invalid subscription key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

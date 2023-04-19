package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshal(kvA.Value, &countA)
			cdc.MustUnmarshal(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.SubscriptionKeyPrefix):
			var subscriptionA, subscriptionB types.Subscription
			cdc.UnmarshalInterface(kvA.Value, &subscriptionA)
			cdc.UnmarshalInterface(kvB.Value, &subscriptionB)

			return fmt.Sprintf("%v\n%v", &subscriptionA, &subscriptionB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSubscriptionAtKeyPrefix):
			var inactiveSubscriptionAtA, inactiveSubscriptionAtB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveSubscriptionAtA)
			cdc.MustUnmarshal(kvB.Value, &inactiveSubscriptionAtB)

			return fmt.Sprintf("%v\n%v", &inactiveSubscriptionAtA, &inactiveSubscriptionAtB)
		case bytes.Equal(kvA.Key[:1], types.QuotaKeyPrefix):
			var quotaA, quotaB types.Quota
			cdc.MustUnmarshal(kvA.Value, &quotaA)
			cdc.MustUnmarshal(kvB.Value, &quotaB)

			return fmt.Sprintf("%v\n%v", &quotaA, &quotaB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

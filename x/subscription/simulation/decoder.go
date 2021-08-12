package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func NewStoreDecoder(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.SubscriptionKeyPrefix):
			var subscriptionA, subscriptionB types.Subscription
			cdc.MustUnmarshalBinaryBare(kvA.Value, &subscriptionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &subscriptionB)

			return fmt.Sprintf("%v\n%v", &subscriptionA, &subscriptionB)

		case bytes.Equal(kvA.Key[:1], types.ActiveSubscriptionForAddressKeyPrefix):
			var activeSubscriptionForAddressA, activeSubscriptionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeSubscriptionForAddressA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeSubscriptionForAddressB)

			return fmt.Sprintf("%v\n%v", &activeSubscriptionForAddressA, &activeSubscriptionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSubscriptionForAddressKeyPrefix):
			var inactiveSubscriptionForAddressA, inactiveSubscriptionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveSubscriptionForAddressA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveSubscriptionForAddressB)

			return fmt.Sprintf("%v\n%v", &inactiveSubscriptionForAddressA, &inactiveSubscriptionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSubscriptionAtKeyPrefix):
			var inactiveSubscriptionAtA, inactiveSubscriptionAtB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveSubscriptionAtA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveSubscriptionAtB)

			return fmt.Sprintf("%v\n%v", &inactiveSubscriptionAtA, &inactiveSubscriptionAtB)
		case bytes.Equal(kvA.Key[:1], types.QuotaKeyPrefix):
			var quotaA, quotaB types.Quota
			cdc.MustUnmarshalBinaryBare(kvA.Value, &quotaA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &quotaB)

			return fmt.Sprintf("%v\n%v", &quotaA, &quotaB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

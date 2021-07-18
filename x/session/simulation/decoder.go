package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func NewStoreDecoder(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.SessionKeyPrefix):
			var sessionA, sessionB types.Session
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionB)

			return fmt.Sprintf("%v\n%v", &sessionA, &sessionB)
		case bytes.Equal(kvA.Key[:1], types.SessionForSubscriptionKeyPrefix):
			var sessionForSubscriptionA, sessionForSubscriptionB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionForSubscriptionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionForSubscriptionB)

			return fmt.Sprintf("%v\n%v", &sessionForSubscriptionA, &sessionForSubscriptionB)
		case bytes.Equal(kvA.Key[:1], types.SessionForNodeKeyPrefix):
			var sessionForNodeA, sessionForNodeB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionForNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionForNodeB)

			return fmt.Sprintf("%v\n%v", &sessionForNodeA, &sessionForNodeB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSessionForAddressKeyPrefix):
			var inactiveSessionForAddressA, inactiveSessionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveSessionForAddressA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionForAddressA, &inactiveSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.ActiveSessionForAddressKeyPrefix):
			var activeSessionForAddressA, activeSessionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeSessionForAddressA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &activeSessionForAddressA, &activeSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSessionAtKeyPrefix):
			var inactiveSessionAtA, inactiveSessionAtB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveSessionAtA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveSessionAtB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionAtA, &inactiveSessionAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

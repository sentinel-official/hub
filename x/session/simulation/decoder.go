package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuf "github.com/gogo/protobuf/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func NewDecoderStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuf.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)
			return fmt.Sprintf("%s\n%s", &countA, &countB)

		case bytes.Equal(kvA.Key[:1], types.SessionKeyPrefix):
			var sessionA, sessionB types.Session
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionB)
			return fmt.Sprintf("%s\n%s", &sessionA, &sessionB)

		case bytes.Equal(kvA.Key[:1], types.SessionForSubscriptionKeyPrefix):
			var sessionA, sessionB types.Session
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionB)
			return fmt.Sprintf("%s\n%s", &sessionA, &sessionB)

		case bytes.Equal(kvA.Key[:1], types.SessionForNodeKeyPrefix):
			var sessionForNodeA, sessionForNodeB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &sessionForNodeA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &sessionForNodeB)
			return fmt.Sprintf("%s\n%s", &sessionForNodeA, &sessionForNodeB)

		case bytes.Equal(kvA.Key[:1], types.InactiveSessionForAddressKeyPrefix):
			var inactiveSessionA, inactiveSessionB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactiveSessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactiveSessionB)
			return fmt.Sprintf("%s\n%s", &inactiveSessionA, &inactiveSessionB)

		case bytes.Equal(kvA.Key[:1], types.ActiveSessionForAddressKeyPrefix):
			var activeSessionA, activeSessionB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeSessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeSessionB)
			return fmt.Sprintf("%s\n%s", &activeSessionA, &activeSessionB)

		case bytes.Equal(kvA.Key[:1], types.ActiveSessionAtKeyPrefix):
			var activeSessionA, activeSessionB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activeSessionA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activeSessionB)
			return fmt.Sprintf("%s\n%s", &activeSessionA, &activeSessionB)
		}

		panic(fmt.Sprintf("invalid session key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

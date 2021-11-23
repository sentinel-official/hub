package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshal(kvA.Value, &countA)
			cdc.MustUnmarshal(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.SessionKeyPrefix):
			var sessionA, sessionB types.Session
			cdc.MustUnmarshal(kvA.Value, &sessionA)
			cdc.MustUnmarshal(kvB.Value, &sessionB)

			return fmt.Sprintf("%v\n%v", &sessionA, &sessionB)

		case bytes.Equal(kvA.Key[:1], types.InactiveSessionForAddressKeyPrefix):
			var inactiveSessionForAddressA, inactiveSessionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveSessionForAddressA)
			cdc.MustUnmarshal(kvB.Value, &inactiveSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionForAddressA, &inactiveSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.ActiveSessionForAddressKeyPrefix):
			var activeSessionForAddressA, activeSessionForAddressB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &activeSessionForAddressA)
			cdc.MustUnmarshal(kvB.Value, &activeSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &activeSessionForAddressA, &activeSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSessionAtKeyPrefix):
			var inactiveSessionAtA, inactiveSessionAtB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactiveSessionAtA)
			cdc.MustUnmarshal(kvB.Value, &inactiveSessionAtB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionAtA, &inactiveSessionAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

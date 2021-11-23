package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func NewStoreDecoder(appCodec codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			appCodec.MustUnmarshal(kvA.Value, &countA)
			appCodec.MustUnmarshal(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.SessionKeyPrefix):
			var sessionA, sessionB types.Session
			appCodec.MustUnmarshal(kvA.Value, &sessionA)
			appCodec.MustUnmarshal(kvB.Value, &sessionB)

			return fmt.Sprintf("%v\n%v", &sessionA, &sessionB)

		case bytes.Equal(kvA.Key[:1], types.InactiveSessionForAddressKeyPrefix):
			var inactiveSessionForAddressA, inactiveSessionForAddressB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &inactiveSessionForAddressA)
			appCodec.MustUnmarshal(kvB.Value, &inactiveSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionForAddressA, &inactiveSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.ActiveSessionForAddressKeyPrefix):
			var activeSessionForAddressA, activeSessionForAddressB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &activeSessionForAddressA)
			appCodec.MustUnmarshal(kvB.Value, &activeSessionForAddressB)

			return fmt.Sprintf("%v\n%v", &activeSessionForAddressA, &activeSessionForAddressB)
		case bytes.Equal(kvA.Key[:1], types.InactiveSessionAtKeyPrefix):
			var inactiveSessionAtA, inactiveSessionAtB protobuftypes.BoolValue
			appCodec.MustUnmarshal(kvA.Value, &inactiveSessionAtA)
			appCodec.MustUnmarshal(kvB.Value, &inactiveSessionAtB)

			return fmt.Sprintf("%v\n%v", &inactiveSessionAtA, &inactiveSessionAtB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/plan/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshal(kvA.Value, &countA)
			cdc.MustUnmarshal(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.PlanKeyPrefix):
			var planA, planB types.Plan
			cdc.MustUnmarshal(kvA.Value, &planA)
			cdc.MustUnmarshal(kvB.Value, &planB)

			return fmt.Sprintf("%v\n%v", &planA, &planB)
		case bytes.Equal(kvA.Key[:1], types.ActivePlanKeyPrefix):
			var activePlanA, activePlanB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &activePlanA)
			cdc.MustUnmarshal(kvB.Value, &activePlanB)

			return fmt.Sprintf("%v\n%v", &activePlanA, &activePlanB)
		case bytes.Equal(kvA.Key[:1], types.InactivePlanKeyPrefix):
			var inactivePlanA, inactivePlanB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &inactivePlanA)
			cdc.MustUnmarshal(kvB.Value, &inactivePlanB)

			return fmt.Sprintf("%v\n%v", &inactivePlanA, &inactivePlanB)
		case bytes.Equal(kvA.Key[:1], types.PlanForProviderKeyPrefix):
			var activePlanForProviderA, activePlanForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshal(kvA.Value, &activePlanForProviderA)
			cdc.MustUnmarshal(kvB.Value, &activePlanForProviderB)

			return fmt.Sprintf("%v\n%v", &activePlanForProviderA, &activePlanForProviderB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

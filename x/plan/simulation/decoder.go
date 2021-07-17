package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuftypes "github.com/gogo/protobuf/types"

	"github.com/sentinel-official/hub/x/plan/types"
)

func NewStoreDecoder(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuftypes.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)

			return fmt.Sprintf("%v\n%v", &countA, &countB)
		case bytes.Equal(kvA.Key[:1], types.PlanKeyPrefix):
			var planA, planB types.Plan
			cdc.MustUnmarshalBinaryBare(kvA.Value, &planA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &planB)

			return fmt.Sprintf("%v\n%v", &planA, &planB)
		case bytes.Equal(kvA.Key[:1], types.ActivePlanKeyPrefix):
			var activePlanA, activePlanB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activePlanA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activePlanB)

			return fmt.Sprintf("%v\n%v", &activePlanA, &activePlanB)
		case bytes.Equal(kvA.Key[:1], types.InactivePlanKeyPrefix):
			var inactivePlanA, inactivePlanB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactivePlanA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactivePlanB)

			return fmt.Sprintf("%v\n%v", &inactivePlanA, &inactivePlanB)
		case bytes.Equal(kvA.Key[:1], types.ActivePlanForProviderKeyPrefix):
			var activePlanForProviderA, activePlanForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &activePlanForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activePlanForProviderB)

			return fmt.Sprintf("%v\n%v", &activePlanForProviderA, &activePlanForProviderB)
		case bytes.Equal(kvA.Key[:1], types.InactivePlanForProviderKeyPrefix):
			var inactivePlanForProviderA, inactivePlanForProviderB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inactivePlanForProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactivePlanForProviderB)

			return fmt.Sprintf("%v\n%v", &inactivePlanForProviderA, &inactivePlanForProviderB)
		case bytes.Equal(kvA.Key[:1], types.NodeForPlanKeyPrefix):
			var nodeForPlanA, nodeForPlanB protobuftypes.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &nodeForPlanA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &nodeForPlanB)

			return fmt.Sprintf("%v\n%v", &nodeForPlanA, &nodeForPlanB)
		case bytes.Equal(kvA.Key[:1], types.CountForNodeByProviderKeyPrefix):
			var countForNodeByProviderA, countForNodeByProviderB protobuftypes.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countForNodeByProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countForNodeByProviderB)

			return fmt.Sprintf("%v\n%v", &countForNodeByProviderA, &countForNodeByProviderB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

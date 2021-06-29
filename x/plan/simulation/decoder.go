package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	protobuf "github.com/gogo/protobuf/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func NewDecoderStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.CountKey):
			var countA, countB protobuf.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countB)
			return fmt.Sprintf("%s\n%s", &countA, &countB)

		case bytes.Equal(kvA.Key[:1], types.PlanKeyPrefix):
			var planA, planB types.Plan
			cdc.MustUnmarshalBinaryBare(kvA.Value, &planA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &planB)
			return fmt.Sprintf("%s\n%s", &planA, &planB)

		case bytes.Equal(kvA.Key[:1], types.ActivePlanKeyPrefix):
			var acivePlanA, activePlanB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &acivePlanA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &activePlanB)
			return fmt.Sprintf("%s\n%s", &acivePlanA, &activePlanB)

		case bytes.Equal(kvA.Key[:1], types.InactivePlanKeyPrefix):
			var inacivePlanA, inactivePlanB protobuf.BoolValue
			cdc.MustUnmarshalBinaryBare(kvA.Value, &inacivePlanA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &inactivePlanB)
			return fmt.Sprintf("%s\n%s", &inacivePlanA, &inactivePlanB)

		case bytes.Equal(kvA.Key[:1], types.CountForNodeByProviderKeyPrefix):
			var countForNodeByProviderA, countForNodeByProviderB protobuf.UInt64Value
			cdc.MustUnmarshalBinaryBare(kvA.Value, &countForNodeByProviderA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &countForNodeByProviderB)
			return fmt.Sprintf("%s\n%s", &countForNodeByProviderA, &countForNodeByProviderB)
		}

		panic(fmt.Sprintf("invalid plan key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

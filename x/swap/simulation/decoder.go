package simulation

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/sentinel-official/hub/x/swap/types"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		if bytes.Equal(kvA.Key[:1], types.SwapKeyPrefix) {
			var swapA, swapB types.MsgSwapRequest
			cdc.MustUnmarshalBinaryBare(kvA.Value, &swapA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &swapB)
			return fmt.Sprintf("%s\n%s", swapA, swapB)
		}

		panic(fmt.Sprintf("invalid swap key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

package simulation

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		if bytes.Equal(kvA.Key[:1], swaptypes.SwapKeyPrefix) {
			var swapA, swapB swaptypes.MsgSwapRequest
			err := cdc.UnmarshalBinaryBare(kvA.Value, &swapA)
			if err != nil {
				fmt.Println(err)
			}
			err = cdc.UnmarshalBinaryBare(kvB.Value, &swapB)
			if err != nil {
				fmt.Println(err)
			}
			result := fmt.Sprintf("%v\n%v", swapA, swapB)
			return result
		}

		panic(fmt.Sprintf("invalid staking key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

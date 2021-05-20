package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		if bytes.Equal(kvA.Key[:1], types.DepositKeyPrefix) {
			var depositA, depositB types.Deposit
			cdc.MustUnmarshalBinaryBare(kvA.Value, &depositA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &depositB)
			return fmt.Sprintf("%s\n%s", depositA, depositB)
		}

		panic(fmt.Sprintf("invalid deposit key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

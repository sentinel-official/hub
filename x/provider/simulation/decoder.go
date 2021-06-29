package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/sentinel-official/hub/x/provider/types"
)

func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	decoderFn := func(kvA, kvB kv.Pair) string {
		if bytes.Equal(kvA.Key[:1], types.ProviderKeyPrefix) {
			var providerA, providerB types.Provider
			cdc.MustUnmarshalBinaryBare(kvA.Value, &providerA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &providerB)
			return fmt.Sprintf("%s\n%s", providerA, providerB)
		}

		panic(fmt.Sprintf("invalid provider key prefix: %X", kvA.Key[:1]))
	}

	return decoderFn
}

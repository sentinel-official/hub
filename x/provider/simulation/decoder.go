package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/sentinel-official/hub/x/provider/types"
)

func NewStoreDecoder(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		if bytes.Equal(kvA.Key[:1], types.ProviderKeyPrefix) {
			var providerA, providerB types.Provider
			cdc.MustUnmarshal(kvA.Value, &providerA)
			cdc.MustUnmarshal(kvB.Value, &providerB)

			return fmt.Sprintf("%v\n%v", providerA, providerB)
		}

		panic(fmt.Sprintf("invalid key prefix %X", kvA.Key[:1]))
	}
}

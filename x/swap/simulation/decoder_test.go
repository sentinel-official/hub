package simulation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/sentinel-official/hub/x/swap/types"
)

func TestDecoderStore(t *testing.T) {
	cdc, _ := simapp.MakeCodecs()
	dec := NewDecodeStore(cdc)
	swapperPK := ed25519.GenPrivKey().PubKey()
	swapperAddr := sdk.AccAddress(swapperPK.Address())
	receiver := sdk.AccAddress(swapperPK.Address())
	txHash := types.EthereumHash{}
	txHash.SetBytes([]byte("6ca7abe5be0ee9bacf582a42f70a2c384b2121ffa6cc33f3ae0b7e41d3480dbe"))
	amount := sdk.NewInt(1000)

	swap := types.NewMsgSwapRequest(swapperAddr, txHash, receiver, amount)

	KVPair := kv.Pair{
		Key: types.SwapKeyPrefix, Value: cdc.MustMarshalBinaryBare(swap),
	}

	tests := []struct {
		name string
		want string
	}{
		{"SwapMessage", fmt.Sprintf("%s\n%s", *swap, *swap)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, dec(KVPair, KVPair), tt.name)
		})
	}
}

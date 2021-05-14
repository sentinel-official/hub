package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

func init() {
	sdk.PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

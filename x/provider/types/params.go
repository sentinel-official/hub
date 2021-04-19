package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
)

var (
	KeyDeposit = []byte("Deposit")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (p *Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Deposit: %s
`), p.Deposit)
}

func (p *Params) Validate() error {
	if !p.Deposit.IsValid() {
		return fmt.Errorf("deposit should be valid")
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyDeposit,
			Value: &p.Deposit,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
	}
}

func NewParams(deposit sdk.Coin) Params {
	return Params{
		Deposit: deposit,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit: DefaultDeposit,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

package types

import (
	"fmt"

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

func (p *Params) Validate() error {
	if p.Deposit.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !p.Deposit.IsValid() {
		return fmt.Errorf("invalid deposit %s", p.Deposit)
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyDeposit,
			Value: &p.Deposit,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(sdk.Coin)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value.IsNegative() {
					return fmt.Errorf("deposit cannot be negative")
				}
				if !value.IsValid() {
					return fmt.Errorf("invalid deposit %s", value)
				}

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

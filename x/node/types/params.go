package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultInactiveDuration = 5 * time.Minute
)

var (
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
)

var (
	KeyDeposit          = []byte("Deposit")
	KeyInactiveDuration = []byte("InactiveDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (p *Params) Validate() error {
	if !p.Deposit.IsValid() {
		return fmt.Errorf("deposit should be valid")
	}
	if p.InactiveDuration <= 0 {
		return fmt.Errorf("inactive_duration should be positive")
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

				if !value.IsValid() {
					return fmt.Errorf("deposit value should be valid")
				}

				return nil
			},
		},
		{
			Key:   KeyInactiveDuration,
			Value: &p.InactiveDuration,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(time.Duration)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value <= 0 {
					return fmt.Errorf("inactive duration value should be positive")
				}

				return nil
			},
		},
	}
}

func NewParams(deposit sdk.Coin, inactiveDuration time.Duration) Params {
	return Params{
		Deposit:          deposit,
		InactiveDuration: inactiveDuration,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:          DefaultDeposit,
		InactiveDuration: DefaultInactiveDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

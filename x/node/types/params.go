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

func (m *Params) Validate() error {
	if m.Deposit.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !m.Deposit.IsValid() {
		return fmt.Errorf("invalid deposit %s", m.Deposit)
	}
	if m.InactiveDuration < 0 {
		return fmt.Errorf("inactive_duration cannot be negative")
	}
	if m.InactiveDuration == 0 {
		return fmt.Errorf("inactive_duration cannot be zero")
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyDeposit,
			Value: &m.Deposit,
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
		{
			Key:   KeyInactiveDuration,
			Value: &m.InactiveDuration,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(time.Duration)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value < 0 {
					return fmt.Errorf("inactive_duration cannot be negative")
				}
				if value == 0 {
					return fmt.Errorf("inactive_duration cannot be zero")
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

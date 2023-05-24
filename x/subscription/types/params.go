package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultExpiryDuration = 2 * time.Minute
)

var (
	KeyExpiryDuration = []byte("ExpiryDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateExpiryDuration(m.ExpiryDuration); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyExpiryDuration,
			Value:       &m.ExpiryDuration,
			ValidatorFn: validateExpiryDuration,
		},
	}
}

func NewParams(expiryDuration time.Duration) Params {
	return Params{
		ExpiryDuration: expiryDuration,
	}
}

func DefaultParams() Params {
	return Params{
		ExpiryDuration: DefaultExpiryDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func validateExpiryDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("expiry_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("expiry_duration cannot be zero")
	}

	return nil
}

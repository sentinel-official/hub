package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultInactiveDuration = 10 * time.Minute
)

var (
	KeyInactiveDuration = []byte("InactiveDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateInactiveDuration(m.InactiveDuration); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyInactiveDuration,
			Value:       &m.InactiveDuration,
			ValidatorFn: validateInactiveDuration,
		},
	}
}

func NewParams(inactiveDuration time.Duration) Params {
	return Params{
		InactiveDuration: inactiveDuration,
	}
}

func DefaultParams() Params {
	return Params{
		InactiveDuration: DefaultInactiveDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func validateInactiveDuration(v interface{}) error {
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
}

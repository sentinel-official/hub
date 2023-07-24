package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultInactivePendingDuration = 2 * time.Minute
)

var (
	KeyInactivePendingDuration = []byte("InactivePendingDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateInactivePendingDuration(m.InactivePendingDuration); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyInactivePendingDuration,
			Value:       &m.InactivePendingDuration,
			ValidatorFn: validateInactivePendingDuration,
		},
	}
}

func NewParams(inactivePendingDuration time.Duration) Params {
	return Params{
		InactivePendingDuration: inactivePendingDuration,
	}
}

func DefaultParams() Params {
	return Params{
		InactivePendingDuration: DefaultInactivePendingDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func validateInactivePendingDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("inactive_pending_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("inactive_pending_duration cannot be zero")
	}

	return nil
}

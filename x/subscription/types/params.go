package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultStatusChangeDelay = 2 * time.Minute
)

var (
	KeyStatusChangeDelay = []byte("StatusChangeDelay")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyStatusChangeDelay,
			Value:       &m.StatusChangeDelay,
			ValidatorFn: validateStatusChangeDelay,
		},
	}
}

func NewParams(statusChangeDelay time.Duration) Params {
	return Params{
		StatusChangeDelay: statusChangeDelay,
	}
}

func DefaultParams() Params {
	return Params{
		StatusChangeDelay: DefaultStatusChangeDelay,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func validateStatusChangeDelay(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("status_change_delay cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("status_change_delay cannot be zero")
	}

	return nil
}

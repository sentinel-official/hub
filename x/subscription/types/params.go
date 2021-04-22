package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultInactiveDuration = 10 * time.Minute
)

var (
	KeyInactiveDuration = []byte("InactiveDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (p *Params) Validate() error {
	if p.InactiveDuration <= 0 {
		return fmt.Errorf("inactive_duration should be positive")
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyInactiveDuration,
			Value: &p.InactiveDuration,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
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

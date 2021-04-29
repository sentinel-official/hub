package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultInactiveDuration         = 5 * time.Minute
	DefaultProofVerificationEnabled = false
)

var (
	KeyInactiveDuration         = []byte("InactiveDuration")
	KeyProofVerificationEnabled = []byte("ProofVerificationEnabled")
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
		{
			Key:   KeyProofVerificationEnabled,
			Value: &p.ProofVerificationEnabled,
			ValidatorFn: func(v interface{}) error {
				_, ok := v.(bool)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				return nil
			},
		},
	}
}

func NewParams(inactiveDuration time.Duration, proofVerificationEnabled bool) Params {
	return Params{
		InactiveDuration:         inactiveDuration,
		ProofVerificationEnabled: proofVerificationEnabled,
	}
}

func DefaultParams() Params {
	return Params{
		InactiveDuration:         DefaultInactiveDuration,
		ProofVerificationEnabled: DefaultProofVerificationEnabled,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultInactivePendingDuration  = 1 * time.Minute
	DefaultProofVerificationEnabled = false
)

var (
	KeyInactivePendingDuration  = []byte("InactivePendingDuration")
	KeyProofVerificationEnabled = []byte("ProofVerificationEnabled")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateInactivePendingDuration(m.InactivePendingDuration); err != nil {
		return err
	}
	if err := validateProofVerificationEnabled(m.ProofVerificationEnabled); err != nil {
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
		{
			Key:         KeyProofVerificationEnabled,
			Value:       &m.ProofVerificationEnabled,
			ValidatorFn: validateProofVerificationEnabled,
		},
	}
}

func NewParams(inactivePendingDuration time.Duration, proofVerificationEnabled bool) Params {
	return Params{
		InactivePendingDuration:  inactivePendingDuration,
		ProofVerificationEnabled: proofVerificationEnabled,
	}
}

func DefaultParams() Params {
	return Params{
		InactivePendingDuration:  DefaultInactivePendingDuration,
		ProofVerificationEnabled: DefaultProofVerificationEnabled,
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

func validateProofVerificationEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	return nil
}

package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultStatusChangeDelay        = 1 * time.Minute
	DefaultProofVerificationEnabled = false
)

var (
	KeyStatusChangeDelay        = []byte("StatusChangeDelay")
	KeyProofVerificationEnabled = []byte("ProofVerificationEnabled")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateStatusChangeDelay(m.StatusChangeDelay); err != nil {
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
			Key:         KeyStatusChangeDelay,
			Value:       &m.StatusChangeDelay,
			ValidatorFn: validateStatusChangeDelay,
		},
		{
			Key:         KeyProofVerificationEnabled,
			Value:       &m.ProofVerificationEnabled,
			ValidatorFn: validateProofVerificationEnabled,
		},
	}
}

func NewParams(statusChangeDelay time.Duration, proofVerificationEnabled bool) Params {
	return Params{
		StatusChangeDelay:        statusChangeDelay,
		ProofVerificationEnabled: proofVerificationEnabled,
	}
}

func DefaultParams() Params {
	return Params{
		StatusChangeDelay:        DefaultStatusChangeDelay,
		ProofVerificationEnabled: DefaultProofVerificationEnabled,
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

func validateProofVerificationEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	return nil
}

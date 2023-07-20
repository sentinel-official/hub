package types

import (
	"fmt"
	"time"

	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultExpiryDuration           = 5 * time.Minute
	DefaultProofVerificationEnabled = false
)

var (
	KeyExpiryDuration           = []byte("ExpiryDuration")
	KeyProofVerificationEnabled = []byte("ProofVerificationEnabled")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateExpiryDuration(m.ExpiryDuration); err != nil {
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
			Key:         KeyExpiryDuration,
			Value:       &m.ExpiryDuration,
			ValidatorFn: validateExpiryDuration,
		},
		{
			Key:         KeyProofVerificationEnabled,
			Value:       &m.ProofVerificationEnabled,
			ValidatorFn: validateProofVerificationEnabled,
		},
	}
}

func NewParams(expiryDuration time.Duration, proofVerificationEnabled bool) Params {
	return Params{
		ExpiryDuration:           expiryDuration,
		ProofVerificationEnabled: proofVerificationEnabled,
	}
}

func DefaultParams() Params {
	return Params{
		ExpiryDuration:           DefaultExpiryDuration,
		ProofVerificationEnabled: DefaultProofVerificationEnabled,
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

func validateProofVerificationEnabled(v interface{}) error {
	_, ok := v.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	return nil
}

package types

import (
	"fmt"
	"strings"
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

var _ params.ParamSet = (*Params)(nil)

func (p *Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Inactive duration:          %s
Proof verification enabled: %t
`), p.InactiveDuration, p.ProofVerificationEnabled)
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
		{
			Key:   KeyProofVerificationEnabled,
			Value: &p.ProofVerificationEnabled,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
	}
}

func (p *Params) Validate() error {
	if p.InactiveDuration <= 0 {
		return fmt.Errorf("inactive_duration should be positive")
	}

	return nil
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

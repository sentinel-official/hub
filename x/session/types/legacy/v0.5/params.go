package v05

import (
	"time"
)

type (
	Params struct {
		InactiveDuration         time.Duration `json:"inactive_duration"`
		ProofVerificationEnabled bool          `json:"proof_verification_enabled"`
	}
)

package v0_5

import (
	"time"
)

type (
	Params struct {
		InactiveDuration time.Duration `json:"inactive_duration"`
	}
)

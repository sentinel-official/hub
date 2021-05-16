package v05

import (
	"time"
)

type (
	Params struct {
		InactiveDuration time.Duration `json:"inactive_duration"`
	}
)

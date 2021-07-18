package simulation

import (
	"math"
)

const (
	MaxSessionDuration          = 1 << 18
	MaxSessionBandwidthUpload   = math.MaxInt32
	MaxSessionBandwidthDownload = math.MaxInt32
)

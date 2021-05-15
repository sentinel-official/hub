package types

import (
	"time"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewProof(id uint64, duration time.Duration, bandwidth hubtypes.Bandwidth) Proof {
	return Proof{
		Id:        id,
		Duration:  duration,
		Bandwidth: bandwidth,
	}
}

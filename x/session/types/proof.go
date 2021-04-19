package types

import (
	"time"

	hubtypes "github.com/sentinel-official/hub/types"
)

func NewProof(channel, subscription uint64, node hubtypes.NodeAddress, duration time.Duration, bandwidth hubtypes.Bandwidth) Proof {
	return Proof{
		Channel:      channel,
		Subscription: subscription,
		Node:         node.String(),
		Duration:     duration,
		Bandwidth:    bandwidth,
	}
}

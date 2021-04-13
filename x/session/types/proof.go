package types

import (
	"time"

	hub "github.com/sentinel-official/hub/types"
)

func NewProof(channel, subscription uint64, node hub.NodeAddress, duration time.Duration, bandwidth hub.Bandwidth) Proof {
	return Proof{
		Channel:      channel,
		Subscription: subscription,
		Node:         node.String(),
		Duration:     duration,
		Bandwidth:    bandwidth,
	}
}

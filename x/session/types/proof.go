package types

import (
	"encoding/json"
	"time"

	hub "github.com/sentinel-official/hub/types"
)

type Proof struct {
	Channel      uint64          `json:"channel,omitempty"`
	Subscription uint64          `json:"subscription"`
	Node         hub.NodeAddress `json:"node"`
	Duration     time.Duration   `json:"duration"`
	Bandwidth    hub.Bandwidth   `json:"bandwidth"`
}

func (p Proof) Bytes() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return bytes
}

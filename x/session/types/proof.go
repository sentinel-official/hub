package types

import (
	"encoding/json"
	"time"

	hub "github.com/sentinel-official/hub/types"
)

type Proof struct {
	Identity  uint64          `json:"identity"`
	Address   hub.NodeAddress `json:"address"`
	Duration  time.Duration   `json:"duration"`
	Bandwidth hub.Bandwidth   `json:"bandwidth"`
}

func (p Proof) Bytes() []byte {
	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	return bytes
}

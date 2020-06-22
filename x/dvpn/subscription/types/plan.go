package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Plan struct {
	ID           uint64          `json:"id"`
	Provider     hub.ProvAddress `json:"provider"`
	Price        sdk.Coins       `json:"price"`
	Duration     time.Duration   `json:"duration"`
	MaxBandwidth hub.Bandwidth   `json:"max_bandwidth"`
	MaxDuration  time.Duration   `json:"max_duration"`
	Status       hub.Status      `json:"status"`
	StatusAt     int64           `json:"status_at"`
}

func (p Plan) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
ID: %d
Provider: %s
Price: %s
Duration: %s
Max bandwidth: %s
Max duration: %s
Status: %s
Status at: %d
`, p.ID, p.Provider, p.Price, p.Duration, p.MaxBandwidth, p.MaxDuration, p.Status, p.StatusAt))
}

type Plans []Plan

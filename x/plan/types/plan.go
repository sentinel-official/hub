package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Plan struct {
	ID        uint64          `json:"id"`
	Provider  hub.ProvAddress `json:"provider"`
	Price     sdk.Coins       `json:"price"`
	Validity  time.Duration   `json:"validity"`
	Bandwidth hub.Bandwidth   `json:"bandwidth"`
	Duration  time.Duration   `json:"duration"`
	Status    hub.Status      `json:"status"`
	StatusAt  time.Time       `json:"status_at"`
}

func (p Plan) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
ID:        %d
Provider:  %s
Price:     %s
Validity:  %s
Bandwidth: %s
Duration:  %s
Status:    %s
Status at: %s
`, p.ID, p.Provider, p.Price, p.Validity, p.Bandwidth, p.Duration, p.Status, p.StatusAt))
}

func (p Plan) PriceForDenom(d string) (sdk.Coin, bool) {
	for _, coin := range p.Price {
		if coin.Denom == d {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

type Plans []Plan

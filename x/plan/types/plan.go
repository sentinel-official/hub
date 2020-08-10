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
Status:    %s
Status at: %s
`, p.ID, p.Provider, p.Price, p.Validity, p.Bandwidth, p.Status, p.StatusAt))
}

func (p Plan) PriceForDenom(d string) (sdk.Coin, bool) {
	for _, coin := range p.Price {
		if coin.Denom == d {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (p Plan) Validate() error {
	if p.ID == 0 {
		return fmt.Errorf("id should not be zero")
	}
	if p.Provider == nil || p.Provider.Empty() {
		return fmt.Errorf("provider should not be nil and empty")
	}
	if p.Price != nil && !p.Price.IsValid() {
		return fmt.Errorf("price should be nil or valid")
	}
	if p.Validity <= 0 {
		return fmt.Errorf("validity should be positive")
	}
	if !p.Bandwidth.IsValid() {
		return fmt.Errorf("bandwidth should be positive")
	}
	if !p.Status.IsValid() {
		return fmt.Errorf("status should be valid")
	}
	if p.StatusAt.IsZero() {
		return fmt.Errorf("status_at should not be zero")
	}

	return nil
}

type Plans []Plan

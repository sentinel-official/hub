package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

func (n Node) String() string {
	if n.Provider == "" {
		return strings.TrimSpace(fmt.Sprintf(`
Address:    %s
Price:      %s
Remote URL: %s
Status:     %s
Status at:  %s
`, n.Address, n.Price, n.RemoteURL, n.Status, n.StatusAt))
	}

	return strings.TrimSpace(fmt.Sprintf(`
Address:    %s
Provider:   %s
Remote URL: %s
Status:     %s
Status at:  %s
`, n.Address, n.Provider, n.RemoteURL, n.Status, n.StatusAt))
}

func (n Node) Validate() error {
	if _, err := hub.NodeAddressFromBech32(n.Address); err != nil {
		return fmt.Errorf("address should not be nil or empty")
	}
	if (n.Provider != "" && n.Price != nil) ||
		(n.Provider == "" && n.Price == nil) {
		return fmt.Errorf("either provider or price should be nil")
	}
	if _, err := hub.ProvAddressFromBech32(n.Provider); err != nil {
		return fmt.Errorf("provider should not be empty")
	}
	if n.Price != nil && !n.Price.IsValid() {
		return fmt.Errorf("price should be valid")
	}
	if len(n.RemoteURL) == 0 || len(n.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length should be between 1 and 64")
	}
	if !n.Status.Equal(hub.StatusActive) && !n.Status.Equal(hub.StatusInactive) {
		return fmt.Errorf("status should be either active or inactive")
	}
	if n.StatusAt.IsZero() {
		return fmt.Errorf("status_at should not be zero")
	}

	return nil
}

func (n Node) PriceForDenom(d string) (sdk.Coin, bool) {
	for _, coin := range n.Price {
		if coin.Denom == d {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (n Node) BytesForCoin(coin sdk.Coin) (sdk.Int, error) {
	price, found := n.PriceForDenom(coin.Denom)
	if !found {
		return sdk.ZeroInt(), fmt.Errorf("price does not exist")
	}

	x := hub.Gigabyte.Quo(price.Amount)
	if x.IsPositive() {
		return coin.Amount.Mul(x), nil
	}

	y := sdk.NewDecFromInt(price.Amount).
		QuoInt(hub.Gigabyte).
		Ceil().TruncateInt()

	return coin.Amount.Quo(y), nil
}

type Nodes []Node

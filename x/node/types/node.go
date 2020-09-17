package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	CategoryUnknown Category = iota + 0x00
	CategoryOpenVPN
	CategoryWireGuard
)

type Category byte

func CategoryFromString(s string) Category {
	switch s {
	case "OpenVPN":
		return CategoryOpenVPN
	case "WireGuard":
		return CategoryWireGuard
	default:
		return CategoryUnknown
	}
}

func (n Category) Equal(v Category) bool {
	return n == v
}

func (n Category) String() string {
	switch n {
	case CategoryOpenVPN:
		return "OpenVPN"
	case CategoryWireGuard:
		return "WireGuard"
	default:
		return "Unknown"
	}
}

func (n Category) IsValid() bool {
	return n == CategoryOpenVPN ||
		n == CategoryWireGuard
}

type Node struct {
	Address       hub.NodeAddress `json:"address"`
	Provider      hub.ProvAddress `json:"provider,omitempty"`
	Price         sdk.Coins       `json:"price,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      Category        `json:"category"`
	Status        hub.Status      `json:"status"`
	StatusAt      time.Time       `json:"status_at"`
}

func (n Node) String() string {
	if n.Provider == nil {
		return strings.TrimSpace(fmt.Sprintf(`
Address:        %s
Price:          %s
Internet speed: %s
Remote URL:     %s
Version:        %s
Category:       %s
Status:         %s
Status at:      %s
`, n.Address, n.Price, n.InternetSpeed, n.RemoteURL, n.Version, n.Category, n.Status, n.StatusAt))
	}

	return strings.TrimSpace(fmt.Sprintf(`
Address:        %s
Provider:       %s
Internet speed: %s
Remote URL:     %s
Version:        %s
Category:       %s
Status:         %s
Status at:      %s
`, n.Address, n.Provider, n.InternetSpeed, n.RemoteURL, n.Version, n.Category, n.Status, n.StatusAt))
}

func (n Node) Validate() error {
	if n.Address == nil || n.Address.Empty() {
		return fmt.Errorf("address should not be nil or empty")
	}
	if (n.Provider != nil && n.Price != nil) ||
		(n.Provider == nil && n.Price == nil) {
		return fmt.Errorf("either provider or price should be nil")
	}
	if n.Provider != nil && n.Provider.Empty() {
		return fmt.Errorf("provider should not be empty")
	}
	if n.Price != nil && !n.Price.IsValid() {
		return fmt.Errorf("price should be valid")
	}
	if !n.InternetSpeed.IsValid() {
		return fmt.Errorf("internet_speed should be valid")
	}
	if len(n.RemoteURL) == 0 || len(n.RemoteURL) > 64 {
		return fmt.Errorf("remote_url length should be (0, 64]")
	}
	if len(n.Version) == 0 || len(n.Version) > 64 {
		return fmt.Errorf("version length should be (0, 64]")
	}
	if !n.Category.IsValid() {
		return fmt.Errorf("category should be valid")
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

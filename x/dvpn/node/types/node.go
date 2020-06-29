package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	CategoryUnknown = iota + 0x00
	CategoryOpenVPN
	CategoryWireGuard
)

type NodeCategory byte

func NodeCategoryFromString(s string) NodeCategory {
	switch s {
	case "OpenVPN":
		return CategoryOpenVPN
	case "WireGuard":
		return CategoryWireGuard
	default:
		return CategoryUnknown
	}
}

func (n NodeCategory) Equal(v NodeCategory) bool {
	return n == v
}

func (n NodeCategory) String() string {
	switch n {
	case CategoryOpenVPN:
		return "OpenVPN"
	case CategoryWireGuard:
		return "WireGuard"
	default:
		return "Unknown"
	}
}

func (n NodeCategory) IsValid() bool {
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
	Category      NodeCategory    `json:"category"`
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

func (n Node) PriceForDenom(s string) (sdk.Coin, bool) {
	for _, coin := range n.Price {
		if coin.Denom == s {
			return coin, true
		}
	}

	return sdk.Coin{}, false
}

func (n Node) BandwidthForCoin(coin sdk.Coin) (hub.Bandwidth, error) {
	price, found := n.PriceForDenom(coin.Denom)
	if !found {
		return hub.Bandwidth{}, fmt.Errorf("price does not exist")
	}

	bytes := coin.Amount.
		Mul(hub.Gigabyte.QuoRaw(2)).
		Quo(price.Amount)

	return hub.NewBandwidth(bytes, bytes), nil
}

type Nodes []Node

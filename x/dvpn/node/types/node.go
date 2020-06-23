package types

import (
	"fmt"
	"strings"

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
	PricePerGB    sdk.Coins       `json:"price_per_gb,omitempty"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      NodeCategory    `json:"category"`
	Status        hub.Status      `json:"status"`
	StatusAt      int64           `json:"status_at"`
}

func (n Node) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address: %s
Provider: %s
Price per Gigabyte: %s
Internet speed: %s
Remote URL: %s
Version: %s
Category: %s
Status: %s
Status at: %d
`, n.Address, n.Provider, n.PricePerGB, n.InternetSpeed, n.RemoteURL, n.Version, n.Category, n.Status, n.StatusAt))
}

type Nodes []Node

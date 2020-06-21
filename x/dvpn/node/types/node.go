package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	CategoryUnknown = iota + 0x01
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

const (
	StatusUnknown = iota + 0x01
	StatusActive
	StatusInactive
)

type NodeStatus byte

func NodeStatusFromString(s string) NodeStatus {
	switch s {
	case "Active":
		return StatusActive
	case "Inactive":
		return StatusInactive
	default:
		return StatusUnknown
	}
}

func (n NodeStatus) Equal(v NodeStatus) bool {
	return n == v
}

func (n NodeStatus) String() string {
	switch n {
	case StatusActive:
		return "Active"
	case StatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

type NodeBandwidthSpeed struct {
	Upload   uint64 `json:"upload"`
	Download uint64 `json:"download"`
}

func NewNodeBandwidthSpeed(upload, download uint64) NodeBandwidthSpeed {
	return NodeBandwidthSpeed{
		Upload:   upload,
		Download: download,
	}
}

func (n NodeBandwidthSpeed) IsAnyZero() bool {
	return n.Upload == 0 || n.Download == 0
}

func (n NodeBandwidthSpeed) String() string {
	return fmt.Sprintf("%d↑, %d↓", n.Upload, n.Download)
}

type Node struct {
	Address          hub.NodeAddress    `json:"address"`
	Provider         hub.ProvAddress    `json:"provider"`
	PricePerGB       sdk.Coins          `json:"price_per_gb"`
	RemoteURL        string             `json:"remote_url"`
	Version          string             `json:"version"`
	BandwidthSpeed   NodeBandwidthSpeed `json:"bandwidth_speed"`
	Category         NodeCategory       `json:"category"`
	Status           NodeStatus         `json:"status"`
	StatusModifiedAt int64              `json:"status_modified_at"`
}

func (n Node) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address: %s
Provider: %s
Price per Gigabyte: %s
Remote URL: %s
Version: %s
Bandwidth speed: %s
Category: %s
Status: %s
Status modified at: %d
`, n.Address, n.Provider, n.PricePerGB, n.RemoteURL, n.Version,
		n.BandwidthSpeed, n.Category, n.Status, n.StatusModifiedAt))
}

type Nodes []Node

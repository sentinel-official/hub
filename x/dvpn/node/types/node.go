package types

import (
	"fmt"
	hub "github.com/sentinel-official/hub/types"
	"strings"
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

const (
	StatusUnknown = iota + 0x00
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

func (n NodeStatus) IsValid() bool {
	return n == StatusActive || n == StatusInactive
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

type Node struct {
	Address       hub.NodeAddress `json:"address"`
	Provider      hub.ProvAddress `json:"provider"`
	InternetSpeed hub.Bandwidth   `json:"internet_speed"`
	RemoteURL     string          `json:"remote_url"`
	Version       string          `json:"version"`
	Category      NodeCategory    `json:"category"`
	Status        NodeStatus      `json:"status"`
	StatusAt      int64           `json:"status_at"`
}

func (n Node) String() string {
	return strings.TrimSpace(fmt.Sprintf(`
Address: %s
Provider: %s
Internet speed: %s
Remote URL: %s
Version: %s
Category: %s
Status: %s
Status at: %d
`, n.Address, n.Provider, n.InternetSpeed, n.RemoteURL, n.Version, n.Category, n.Status, n.StatusAt))
}

type Nodes []Node

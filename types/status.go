package types

import (
	"strings"
)

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusInactivePending:
		return "inactive_pending"
	case StatusInactive:
		return "inactive"
	default:
		return "unspecified"
	}
}

func (s Status) IsValid() bool {
	return s == StatusActive ||
		s == StatusInactivePending ||
		s == StatusInactive
}

func (s Status) Equal(v Status) bool {
	return s == v
}

func (s Status) IsOneOf(items ...Status) bool {
	for _, item := range items {
		if s.Equal(item) {
			return true
		}
	}

	return false
}

func StatusFromString(s string) Status {
	s = strings.ToLower(s)
	switch s {
	case "active":
		return StatusActive
	case "inactive_pending":
		return StatusInactivePending
	case "inactive":
		return StatusInactive
	default:
		return StatusUnspecified
	}
}

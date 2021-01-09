package types

const (
	StatusUnknown Status = iota + 0x00
	StatusActive
	StatusInactivePending
	StatusInactive
)

type Status byte

func StatusFromString(s string) Status {
	switch s {
	case "Active":
		return StatusActive
	case "InactivePending":
		return StatusInactivePending
	case "Inactive":
		return StatusInactive
	default:
		return StatusUnknown
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

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "Active"
	case StatusInactivePending:
		return "InactivePending"
	case StatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

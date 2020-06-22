package types

const (
	StatusUnknown = iota + 0x00
	StatusActive
	StatusInactive
)

type Status byte

func StatusFromString(s string) Status {
	switch s {
	case "Active":
		return StatusActive
	case "Inactive":
		return StatusInactive
	default:
		return StatusUnknown
	}
}

func (s Status) IsValid() bool {
	return s == StatusActive || s == StatusInactive
}

func (s Status) Equal(v Status) bool {
	return s == v
}

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "Active"
	case StatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

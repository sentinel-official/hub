package types

const (
	StatusUnknown Status = iota + 0x00
	StatusActive
	StatusCancel
	StatusInactive
)

type Status byte

func StatusFromString(s string) Status {
	switch s {
	case "Active":
		return StatusActive
	case "Cancel":
		return StatusCancel
	case "Inactive":
		return StatusInactive
	default:
		return StatusUnknown
	}
}

func (s Status) IsValid() bool {
	return s == StatusActive ||
		s == StatusCancel ||
		s == StatusInactive
}

func (s Status) Equal(v Status) bool {
	return s == v
}

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "Active"
	case StatusCancel:
		return "Cancel"
	case StatusInactive:
		return "Inactive"
	default:
		return "Unknown"
	}
}

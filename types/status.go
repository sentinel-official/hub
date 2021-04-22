package types

const (
	StatusUnknown Status = iota
	StatusActive
	StatusInactivePending
	StatusInactive
)

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

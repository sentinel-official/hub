package types

import (
	"github.com/spf13/pflag"
)

const (
	FlagStatus = "status"
)

func StatusFromFlags(flags *pflag.FlagSet) (Status, error) {
	s, err := flags.GetString(FlagStatus)
	if err != nil {
		return StatusUnspecified, err
	}
	if s == "" {
		return StatusUnspecified, nil
	}

	return StatusFromString(s), nil
}

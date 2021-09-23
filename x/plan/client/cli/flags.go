package cli

import (
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagProvider = "provider"
	flagStatus   = "status"
)

func GetProvider(flags *pflag.FlagSet) (hubtypes.ProvAddress, error) {
	s, err := flags.GetString(flagProvider)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return hubtypes.ProvAddressFromBech32(s)
}

func GetStatus(flags *pflag.FlagSet) (hubtypes.Status, error) {
	s, err := flags.GetString(flagStatus)
	if err != nil {
		return hubtypes.Unspecified, err
	}
	if s == "" {
		return hubtypes.Unspecified, nil
	}

	return hubtypes.StatusFromString(s), nil
}

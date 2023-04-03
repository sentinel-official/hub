package cli

import (
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagName        = "name"
	flagIdentity    = "identity"
	flagWebsite     = "website"
	flagDescription = "description"
	flagStatus      = "status"
)

func GetStatus(flags *pflag.FlagSet) (hubtypes.Status, error) {
	s, err := flags.GetString(flagStatus)
	if err != nil {
		return hubtypes.StatusUnspecified, err
	}
	if s == "" {
		return hubtypes.StatusUnspecified, nil
	}

	return hubtypes.StatusFromString(s), nil
}

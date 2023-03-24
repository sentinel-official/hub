package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagProvider  = "provider"
	flagPlan      = "plan"
	flagPrice     = "price"
	flagRemoteURL = "remote-url"
	flagStatus    = "status"
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

func GetPrice(flags *pflag.FlagSet) (sdk.Coins, error) {
	s, err := flags.GetString(flagPrice)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(s)
}

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

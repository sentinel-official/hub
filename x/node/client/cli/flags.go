package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagGigabytes      = "gigabytes"
	flagGigabytePrices = "gigabyte-prices"
	flagHours          = "hours"
	flagHourlyPrices   = "hourly-prices"
	flagPlan           = "plan"
	flagRemoteURL      = "remote-url"
)

func GetGigabytePrices(flags *pflag.FlagSet) (sdk.Coins, error) {
	s, err := flags.GetString(flagGigabytePrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(s)
}

func GetHourlyPrices(flags *pflag.FlagSet) (sdk.Coins, error) {
	s, err := flags.GetString(flagHourlyPrices)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.ParseCoinsNormalized(s)
}

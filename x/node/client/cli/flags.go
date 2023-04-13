package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagGigabytePrices = "gigabyte-prices"
	flagHourlyPrices   = "hourly-prices"
	flagPlan           = "plan"
	flagProvider       = "provider"
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

func GetProviderAddr(flags *pflag.FlagSet) (hubtypes.ProvAddress, error) {
	s, err := flags.GetString(flagProvider)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return hubtypes.ProvAddressFromBech32(s)
}

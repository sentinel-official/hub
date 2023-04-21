package cli

import (
	"encoding/base64"

	hubtypes "github.com/sentinel-official/hub/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"
)

const (
	flagAccountAddress = "account-addr"
	flagNodeAddress    = "node-addr"
	flagRating         = "rating"
	flagSignature      = "signature"
	flagSubscriptionID = "subscription-id"
)

func GetAccountAddress(flags *pflag.FlagSet) (sdk.AccAddress, error) {
	s, err := flags.GetString(flagAccountAddress)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.AccAddressFromBech32(s)
}

func GetNodeAddress(flags *pflag.FlagSet) (hubtypes.NodeAddress, error) {
	s, err := flags.GetString(flagNodeAddress)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return hubtypes.NodeAddressFromBech32(s)
}

func GetSignature(flags *pflag.FlagSet) ([]byte, error) {
	s, err := flags.GetString(flagSignature)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return base64.StdEncoding.DecodeString(s)
}

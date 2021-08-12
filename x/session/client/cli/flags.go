package cli

import (
	"encoding/base64"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagAddress   = "address"
	flagRating    = "rating"
	flagSignature = "signature"
	flagStatus    = "status"
)

func GetAddress(flags *pflag.FlagSet) (sdk.AccAddress, error) {
	s, err := flags.GetString(flagAddress)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	return sdk.AccAddressFromBech32(s)
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

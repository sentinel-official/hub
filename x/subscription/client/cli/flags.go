package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/pflag"

	hubtypes "github.com/sentinel-official/hub/types"
)

const (
	flagBytes          = "bytes"
	flagAccountAddress = "account-addr"
	flagNodeAddress    = "node-addr"
	flagPlanID         = "plan-id"
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

package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return flags.GetCommands(
		queryDeposit(cdc),
		queryDeposits(cdc),
	)
}

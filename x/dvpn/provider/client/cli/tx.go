package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetTxRegisterProviderCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "register",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

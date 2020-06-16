package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryProviderCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

func GetQueryProvidersCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}

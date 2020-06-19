package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/provider/client/common"
)

func getQueryProviderCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "provider",
		Short: "Query provider",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.ProvAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			provider, err := common.QueryProvider(ctx, address)
			if err != nil {
				return err
			}

			fmt.Println(provider)
			return nil
		},
	}
}

func getQueryProvidersCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "providers",
		Short: "Query providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			providers, err := common.QueryProviders(ctx)
			if err != nil {
				return err
			}

			for _, provider := range providers {
				fmt.Println(provider)
			}

			return nil
		},
	}
}

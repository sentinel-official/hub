package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/client/common"
)

func queryProvider(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "provider",
		Short: "Query a provider",
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

func queryProviders(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "providers",
		Short: "Query providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)

			skip, err := cmd.Flags().GetInt(flagSkip)
			if err != nil {
				return err
			}

			limit, err := cmd.Flags().GetInt(flagLimit)
			if err != nil {
				return err
			}

			providers, err := common.QueryProviders(ctx, skip, limit)
			if err != nil {
				return err
			}

			for _, provider := range providers {
				fmt.Printf("%s\n\n", provider)
			}

			return nil
		},
	}

	cmd.Flags().Int(flagSkip, 0, "skip")
	cmd.Flags().Int(flagLimit, 25, "limit")

	return cmd
}

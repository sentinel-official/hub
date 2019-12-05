package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func QueryFreeClientsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-clients",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			if id != "" {
				freeClients, err := common.QueryQueryFreeClientsOfNode(ctx, id)
				if err != nil {
					return err
				}

				for _, freeClient := range freeClients {
					fmt.Println(freeClient)
				}

				return nil
			} else if address != "" {
				freeNodes, err := common.QueryFreeNodesOfClient(ctx, address)
				if err != nil {
					return err
				}

				for _, freeClient := range freeNodes {
					fmt.Println(freeClient)
				}

				return nil
			} else {
				freeClients, err := common.QueryAllFreeClients(ctx)
				if err != nil {
					return err
				}

				for _, freeClient := range freeClients {
					fmt.Println(freeClient)
				}

				return nil
			}
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagAddress, "", "Account address")

	return cmd
}

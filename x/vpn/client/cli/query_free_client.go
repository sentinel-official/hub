package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/sentinel-official/hub/x/vpn/client/common"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func QueryFreeClientsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "free-clients",
		Short: "Query subscriptions",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ctx := context.NewCLIContext().WithCodec(cdc)

			id := viper.GetString(flagNodeID)
			address := viper.GetString(flagAddress)

			var freeClients []types.FreeClient
			if id != "" {
				freeClients, err = common.QueryQueryFreeClientsOfNode(ctx, id)
			} else if address != "" {
				freeClients, err = common.QueryFreeNodesOfClient(ctx, address)
			} else {
				freeClients, err = common.QueryAllFreeClients(ctx)
			}

			if err != nil {
				return err
			}

			for _, freeClient := range freeClients {
				fmt.Println(freeClient)
			}

			return nil
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagAddress, "", "Account address")

	return cmd
}

package cli

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	
	"github.com/sentinel-official/hub/x/deposit/client/common"
)

func QueryDepositsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "Query deposits",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			address := viper.GetString(flagAddress)
			
			if address != "" {
				deposit, err := common.QueryDepositOfAddress(ctx, address)
				if err != nil {
					return err
				}
				
				fmt.Println(deposit)
				return nil
			}
			
			deposits, err := common.QueryAllDeposits(ctx)
			if err != nil {
				return err
			}
			
			for _, deposit := range deposits {
				fmt.Println(deposit)
			}
			
			return nil
		},
	}
	
	cmd.Flags().String(flagAddress, "", "Account address")
	
	return client.GetCommands(cmd)[0]
}

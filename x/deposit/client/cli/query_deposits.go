package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/client/common"
)

func QueryDepositsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "Query deposits",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			address := viper.GetString(flagAddress)

			if address != "" {
				_address, err := csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				deposit, err := common.QueryDepositOfAddress(cliCtx, cdc, _address)
				if err != nil {
					return err
				}

				fmt.Println(deposit)

				return nil
			}

			deposits, err := common.QueryAllDeposits(cliCtx, cdc)
			if err != nil {
				return err
			}

			for _, _deposit := range deposits {
				fmt.Println(_deposit)
			}

			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "Account address")

	return client.GetCommands(cmd)[0]
}

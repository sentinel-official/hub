package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/client/common"
)

func QueryDepositsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposits",
		Short: "Query deposits",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			address := viper.GetString(flagAddress)

			var res []byte
			var err error

			if address != "" {
				var _address csdkTypes.AccAddress

				_address, err = csdkTypes.AccAddressFromBech32(address)
				if err != nil {
					return err
				}

				res, err = common.QueryDepositsOfAddress(cliCtx, cdc, _address)
				if err != nil {
					return err
				}
				if res == nil {
					return fmt.Errorf("no deposit found")
				}

				var _deposit deposit.Deposit
				if err = cdc.UnmarshalJSON(res, &_deposit); err != nil {
					return err
				}

				fmt.Println(_deposit)

				return nil
			}

			res, err = common.QueryAllDeposits(cliCtx)
			if err != nil {
				return err
			}
			if string(res) == "[]" || string(res) == "null" {
				return fmt.Errorf("no deposits found")
			}

			var deposits []deposit.Deposit
			if err = cdc.UnmarshalJSON(res, &deposits); err != nil {
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

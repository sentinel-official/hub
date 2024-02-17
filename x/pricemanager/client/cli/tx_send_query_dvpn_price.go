package cli

import (
	"github.com/sentinel-official/hub/v12/x/pricemanager/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdSendQueryDVPNPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-query-dvpn-price [channel-id] [pool-id] [base-asset-denom] [quote-asset-denom]",
		Short: "Query the balances of an account on the remote chain via ICQ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgSendQueryDVPNPrice(
				clientCtx.GetFromAddress().String(),
				args[0], // channel id
				args[1], // pool id
				args[2], // base asset denomination
				args[3], // quote asset denomination
				pageReq,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "send query dvpn price")

	return cmd
}

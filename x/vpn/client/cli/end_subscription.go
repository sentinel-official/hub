package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func EndSubscriptionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "end",
		Short: "End subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)
			
			id, err := hub.NewSubscriptionIDFromString(args[0])
			if err != nil {
				return err
			}
			
			fromAddress := ctx.GetFromAddress()
			
			msg := types.NewMsgEndSubscription(fromAddress, id)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	
	return cmd
}

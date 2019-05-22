package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func StartSubscriptionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start subscription",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := uint64(viper.GetInt64(flagNodeID))
			deposit := viper.GetString(flagDeposit)

			parsedDeposit, err := csdkTypes.ParseCoin(deposit)
			if err != nil {
				return err
			}

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgStartSubscription(fromAddress, nodeID, parsedDeposit)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().Uint64(flagNodeID, 0, "Node ID")
	cmd.Flags().String(flagDeposit, "", "Deposit")

	_ = cmd.MarkFlagRequired(flagNodeID)
	_ = cmd.MarkFlagRequired(flagDeposit)

	return cmd
}

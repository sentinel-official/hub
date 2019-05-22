// nolint: dupl
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

func DeregisterNodeTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deregister",
		Short: "Deregister node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := uint64(viper.GetInt64(flagNodeID))
			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgDeregisterNode(fromAddress, nodeID)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().Uint64(flagNodeID, 0, "Node ID")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

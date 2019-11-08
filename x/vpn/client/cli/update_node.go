package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func UpdateNodeInfoTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-info",
		Short: "Update info of the node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			nodeID := hub.NewIDFromString(viper.GetString(flagNodeID))
			_type := viper.GetString(flagType)
			version := viper.GetString(flagVersion)
			moniker := viper.GetString(flagMoniker)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			internetSpeed := hub.Bandwidth{
				Upload:   sdk.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: sdk.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			encryption := viper.GetString(flagEncryption)

			parsedPricesPerGB, err := sdk.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress := ctx.GetFromAddress()

			msg := types.NewMsgUpdateNodeInfo(fromAddress, nodeID,
				_type, version, moniker, parsedPricesPerGB, internetSpeed, encryption)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagType, "", "VPN node type")
	cmd.Flags().String(flagVersion, "", "VPN node version")
	cmd.Flags().String(flagMoniker, "", "Moniker")
	cmd.Flags().String(flagPricesPerGB, "", "Prices per GB")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncryption, "", "VPN encryption method")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

func UpdateNodeStatusTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-status",
		Short: "Update status of the node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			nodeID := hub.NewIDFromString(viper.GetString(flagNodeID))
			status := strings.ToUpper(args[0])

			fromAddress := ctx.GetFromAddress()

			msg := types.NewMsgUpdateNodeStatus(fromAddress, nodeID, status)
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func UpdateNodeInfoTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-info",
		Short: "Update info of the node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := sdk.NewIDFromString(viper.GetString(flagNodeID))
			_type := viper.GetString(flagType)
			version := viper.GetString(flagVersion)
			moniker := viper.GetString(flagMoniker)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			internetSpeed := sdk.Bandwidth{
				Upload:   csdk.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: csdk.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			encryption := viper.GetString(flagEncryption)

			parsedPricesPerGB, err := csdk.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgUpdateNodeInfo(fromAddress, nodeID,
				_type, version, moniker, parsedPricesPerGB, internetSpeed, encryption)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdk.Msg{msg}, false)
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
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := sdk.NewIDFromString(viper.GetString(flagNodeID))
			status := strings.ToUpper(args[0])

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgUpdateNodeStatus(fromAddress, nodeID, status)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

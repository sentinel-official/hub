package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func UpdateNodeDetailsTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-details",
		Short: "Update details of the node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := sdkTypes.NewIDFromString(viper.GetString(flagNodeID))
			moniker := viper.GetString(flagMoniker)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			internetSpeed := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			encryptionMethod := viper.GetString(flagEncryption)
			type_ := viper.GetString(flagType)
			version := viper.GetString(flagVersion)

			parsedPricesPerGB, err := csdkTypes.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgUpdateNodeInfo(fromAddress, nodeID,
				moniker, type_, version, parsedPricesPerGB, internetSpeed,
				encryptionMethod)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().String(flagMoniker, "", "Moniker")
	cmd.Flags().String(flagType, "", "VPN node type")
	cmd.Flags().String(flagVersion, "", "VPN node version")
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

			nodeID := sdkTypes.NewIDFromString(viper.GetString(flagNodeID))
			status := strings.ToUpper(args[0])

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgUpdateNodeStatus(fromAddress, nodeID, status)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

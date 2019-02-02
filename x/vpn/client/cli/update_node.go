package cli

import (
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func UpdateNodeInfoTxCmd(cdc *codec.Codec) *cobra.Command {
	updateNodeCmd := &cobra.Command{
		Use:   "update",
		Short: "Update node information",
	}

	updateNodeCmd.AddCommand(client.PostCommands(
		updateNodeDetailsTxCmd(cdc),
		updateNodeStatusTxCmd(cdc),
	)...)

	return updateNodeCmd
}

func updateNodeDetailsTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "details",
		Short: "Update details of the node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := viper.GetString(flagNodeID)
			apiPort := viper.GetInt(flagAPIPort)
			uploadSpeed := viper.GetInt64(flagUploadSpeed)
			downloadSpeed := viper.GetInt64(flagDownloadSpeed)
			encMethod := viper.GetString(flagEncMethod)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			version := viper.GetString(flagVersion)

			parsedPricesPerGB, err := csdkTypes.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateNodeDetails(fromAddress, nodeID,
				uint16(apiPort), uint64(uploadSpeed), uint64(downloadSpeed),
				encMethod, parsedPricesPerGB, version)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")
	cmd.Flags().Int64(flagAPIPort, 8000, "Node API port")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncMethod, "", "VPN tunnel encryption method")
	cmd.Flags().String(flagPricesPerGB, "100sent,1000sut", "Prices for one GB of data")
	cmd.Flags().String(flagVersion, "", "Node version")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

func updateNodeStatusTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Update status of the node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			nodeID := viper.GetString(flagNodeID)
			status := strings.ToUpper(args[0])

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateNodeStatus(fromAddress, nodeID, status)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagNodeID, "", "Node ID")

	_ = cmd.MarkFlagRequired(flagNodeID)

	return cmd
}

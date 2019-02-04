package cli

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func RegisterNodeTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			amountToLock := viper.GetString(flagAmountToLock)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			uploadSpeed := csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed))
			downloadSpeed := csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed))
			apiPort := vpn.NewAPIPort(uint32(viper.GetInt(flagAPIPort)))
			encMethod := viper.GetString(flagEncMethod)
			nodeType := viper.GetString(flagNodeType)
			version := viper.GetString(flagVersion)

			parsedAmountToLock, err := csdkTypes.ParseCoin(amountToLock)
			if err != nil {
				return err
			}

			parsedPricesPerGB, err := csdkTypes.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			msg := vpn.NewMsgRegisterNode(fromAddress,
				parsedAmountToLock, parsedPricesPerGB, uploadSpeed, downloadSpeed,
				apiPort, encMethod, nodeType, version)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagAmountToLock, "1000sent", "Locking amount to register node")
	cmd.Flags().Int64(flagAPIPort, 8000, "Node API port")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncMethod, "", "VPN tunnel encryption method")
	cmd.Flags().String(flagPricesPerGB, "100sent,1000sut", "Prices for one GB of data")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().String(flagNodeType, "OpenVPN", "Type of VPN node")

	_ = cmd.MarkFlagRequired(flagAmountToLock)
	_ = cmd.MarkFlagRequired(flagAPIPort)
	_ = cmd.MarkFlagRequired(flagUploadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagEncMethod)
	_ = cmd.MarkFlagRequired(flagPricesPerGB)
	_ = cmd.MarkFlagRequired(flagVersion)
	_ = cmd.MarkFlagRequired(flagNodeType)

	return cmd
}

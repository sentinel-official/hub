package cli

import (
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

func RegisterNodeTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register node details",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			moniker := viper.GetString(flagMoniker)
			amountToLock := viper.GetString(flagAmountToLock)
			pricesPerGB := viper.GetString(flagPricesPerGB)
			netSpeed := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			apiPort := uint16(viper.GetInt(flagAPIPort))
			encryptionMethod := viper.GetString(flagEncryptionMethod)
			type_ := viper.GetString(flagType)
			version := viper.GetString(flagVersion)

			parsedAmountToLock, err := csdkTypes.ParseCoin(amountToLock)
			if err != nil {
				return err
			}

			parsedPricesPerGB, err := csdkTypes.ParseCoins(pricesPerGB)
			if err != nil {
				return err
			}

			fromAddress := cliCtx.GetFromAddress()

			msg := vpn.NewMsgRegisterNode(fromAddress,
				moniker, parsedAmountToLock, parsedPricesPerGB, netSpeed,
				apiPort, encryptionMethod, type_, version)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagMoniker, "", "Node moniker")
	cmd.Flags().String(flagAmountToLock, "1000sent", "Locking amount to register node")
	cmd.Flags().Int64(flagAPIPort, 8000, "Node API port")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncryptionMethod, "", "VPN tunnel encryption method")
	cmd.Flags().String(flagPricesPerGB, "100sent,1000sut", "Prices for one GB of data")
	cmd.Flags().String(flagType, "OpenVPN", "Type of VPN node")
	cmd.Flags().String(flagVersion, "", "Node version")

	_ = cmd.MarkFlagRequired(flagMoniker)
	_ = cmd.MarkFlagRequired(flagAmountToLock)
	_ = cmd.MarkFlagRequired(flagAPIPort)
	_ = cmd.MarkFlagRequired(flagUploadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagEncryptionMethod)
	_ = cmd.MarkFlagRequired(flagPricesPerGB)
	_ = cmd.MarkFlagRequired(flagVersion)
	_ = cmd.MarkFlagRequired(flagType)

	return cmd
}

package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/x/vpn"
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

			msg := vpn.NewMsgRegisterNode(cliCtx.FromAddress, _type, version,
				moniker, parsedPricesPerGB, internetSpeed, encryption)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagType, "", "VPN node type")
	cmd.Flags().String(flagVersion, "", "VPN node version")
	cmd.Flags().String(flagMoniker, "", "Moniker")
	cmd.Flags().String(flagPricesPerGB, "", "Prices per GB")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet download speed in bytes/sec")
	cmd.Flags().String(flagEncryption, "", "VPN encryption method")

	_ = cmd.MarkFlagRequired(flagType)
	_ = cmd.MarkFlagRequired(flagVersion)
	_ = cmd.MarkFlagRequired(flagMoniker)
	_ = cmd.MarkFlagRequired(flagUploadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagEncryption)
	_ = cmd.MarkFlagRequired(flagPricesPerGB)

	return cmd
}

package cli

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func SignSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-bandwidth",
		Short: "Sign session bandwidth details",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			sessionID := viper.GetString(flagSessionID)
			uploadSpeed := csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed))
			downloadSpeed := csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed))

			bandwidth := types.Bandwidth{
				Upload:   uploadSpeed,
				Download: downloadSpeed,
			}
			signBytes, err := common.GetSessionBandwidthSignBytes(cliCtx, cdc, sessionID, bandwidth)
			if err != nil {
				return err
			}

			fromName, err := cliCtx.GetFromName()
			if err != nil {
				return err
			}

			password, err := keys.GetPassphrase(fromName)
			if err != nil {
				return err
			}

			kb, err := keys.GetKeyBase()
			if err != nil {
				return err
			}

			signature, _, err := kb.Sign(fromName, password, signBytes)
			if err != nil {
				return err
			}

			fmt.Println(base64.StdEncoding.EncodeToString(signature))

			return nil
		},
	}

	cmd.Flags().String(flagSessionID, "", "Session ID")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")

	_ = cmd.MarkFlagRequired(flagSessionID)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)

	return cmd
}

func UpdateSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bandwidth",
		Short: "Update session bandwidth details",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithAccountDecoder(cdc).WithCodec(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			sessionID := vpn.NewSessionID(viper.GetString(flagSessionID))
			upload := csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed))
			download := csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed))
			clientSign := viper.GetString(flagClientSign)
			nodeOwnerSign := viper.GetString(flagNodeOwnerSign)

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			clientSignBytes, err := base64.StdEncoding.DecodeString(clientSign)
			if err != nil {
				return err
			}

			nodeOwnerSignBytes, err := base64.StdEncoding.DecodeString(nodeOwnerSign)
			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateSessionBandwidth(fromAddress, sessionID, upload, download,
				clientSignBytes, nodeOwnerSignBytes)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagSessionID, "", "Session ID")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().String(flagClientSign, "", "Bandwidth signature of the client")
	cmd.Flags().String(flagNodeOwnerSign, "", "Bandwidth signature of the node owner")

	_ = cmd.MarkFlagRequired(flagSessionID)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagClientSign)
	_ = cmd.MarkFlagRequired(flagNodeOwnerSign)

	return cmd
}

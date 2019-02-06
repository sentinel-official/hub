package cli

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

	"github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func SignSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-session",
		Short: "Sign session with updated bandwidth details ",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			sessionID := viper.GetString(flagSessionID)
			upload := csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed))
			download := csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed))

			bandwidth := types.Bandwidth{
				Upload:   upload,
				Download: download,
			}

			bytes, err := common.GetSessionBandwidthSignBytes(cliCtx, cdc, sessionID, bandwidth)
			if err != nil {
				return err
			}

			res, err := common.MakeSignature(cliCtx, bytes)
			if err != nil {
				return err
			}
			fmt.Println(base64.StdEncoding.EncodeToString(res))

			return nil
		},
	}

	cmd.Flags().String(flagSessionID, "", "session to sign")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")

	_ = cmd.MarkFlagRequired(flagSessionID)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)

	return cmd
}

func UpdateSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-session",
		Short: "Update bandwidth details in active session ",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithAccountDecoder(cdc).WithCodec(cdc)
			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			fromAddress, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			sessionID := vpn.NewSessionID(viper.GetString(flagSessionID))
			upload := csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed))
			download := csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed))

			clientSign := viper.GetString(flagClientSign)
			clientSignBytes, err := base64.StdEncoding.DecodeString(clientSign)
			if err != nil {
				return err
			}

			nodeOwnerSign := viper.GetString(flagNodeOwnerSign)
			nodeOwnerSignBytes, err := base64.StdEncoding.DecodeString(nodeOwnerSign)
			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateSessionBandwidth(fromAddress, sessionID, upload, download, clientSignBytes, nodeOwnerSignBytes)
			if cliCtx.GenerateOnly {
				return utils.PrintUnsignedStdTx(os.Stdout, txBldr, cliCtx, []csdkTypes.Msg{msg}, false)
			}

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []csdkTypes.Msg{msg})
		},
	}

	cmd.Flags().String(flagSessionID, "", "Details of session  ")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().String(flagClientSign, "", "verify  and update bandwidth")
	cmd.Flags().String(flagNodeOwnerSign, "", "verify and update bandwidth")

	_ = cmd.MarkFlagRequired(flagSessionID)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagClientSign)
	_ = cmd.MarkFlagRequired(flagNodeOwnerSign)

	return cmd
}

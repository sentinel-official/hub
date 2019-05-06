package cli

import (
	"encoding/base64"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
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
			bandwidth := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			signBytes, err := common.GetSessionBandwidthSignDataBytes(cliCtx, cdc, sessionID, bandwidth)
			if err != nil {
				return err
			}

			fromName := cliCtx.GetFromName()

			password, err := keys.GetPassphrase(fromName)
			if err != nil {
				return err
			}

			kb, err := keys.NewKeyBaseFromHomeFlag()
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

			sessionID := sdkTypes.NewID(viper.GetString(flagSessionID))
			bandwidth := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUploadSpeed)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownloadSpeed)),
			}
			clientSign := viper.GetString(flagClientSign)
			nodeOwnerSign := viper.GetString(flagNodeOwnerSign)

			fromAddress := cliCtx.GetFromAddress()

			clientSignBytes, err := base64.StdEncoding.DecodeString(clientSign)
			if err != nil {
				return err
			}

			nodeOwnerSignBytes, err := base64.StdEncoding.DecodeString(nodeOwnerSign)
			if err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateSessionBandwidth(fromAddress, sessionID, bandwidth,
				nodeOwnerSignBytes, clientSignBytes)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagSessionID, "", "Session ID")
	cmd.Flags().Int64(flagDownloadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().Int64(flagUploadSpeed, 0, "Internet upload speed in bytes/sec")
	cmd.Flags().String(flagNodeOwnerSign, "", "Bandwidth signature of the node owner")
	cmd.Flags().String(flagClientSign, "", "Bandwidth signature of the client")

	_ = cmd.MarkFlagRequired(flagSessionID)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagNodeOwnerSign)
	_ = cmd.MarkFlagRequired(flagClientSign)

	return cmd
}

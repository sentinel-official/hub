package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func SignSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-bandwidth",
		Short: "Sign session bandwidth",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			id := sdkTypes.NewIDFromString(viper.GetString(flagSubscriptionID))
			bandwidth := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUpload)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownload)),
			}

			msg := vpn.MsgUpdateSessionInfo{
				SubscriptionID: id,
				Bandwidth:      bandwidth,
			}

			passphrase, err := keys.GetPassphrase(cliCtx.FromName)
			if err != nil {
				return err
			}

			kb, err := keys.NewKeyBaseFromHomeFlag()
			if err != nil {
				return err
			}

			sigBytes, pubKey, err := kb.Sign(cliCtx.FromName, passphrase, msg.GetSignBytes())
			if err != nil {
				return err
			}

			stdSignature := auth.StdSignature{
				PubKey:    pubKey,
				Signature: sigBytes,
			}

			bz, err := cdc.MarshalJSON(stdSignature)
			if err != nil {
				return err
			}

			fmt.Println(string(bz))

			return nil
		},
	}

	cmd.Flags().String(flagSubscriptionID, "", "Subscription ID")
	cmd.Flags().Int64(flagUpload, 0, "Upload in in bytes")
	cmd.Flags().Int64(flagDownload, 0, "Download in bytes")

	_ = cmd.MarkFlagRequired(flagSubscriptionID)
	_ = cmd.MarkFlagRequired(flagUpload)
	_ = cmd.MarkFlagRequired(flagDownload)

	return cmd
}

func UpdateSessionInfoTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-session-info",
		Short: "Update session info",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authTxBuilder.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			id := sdkTypes.NewIDFromString(viper.GetString(flagSubscriptionID))
			bandwidth := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUpload)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownload)),
			}
			nodeOwnerSignatureStr := viper.GetString(flagNodeOwnerSign)
			clientSignatureStr := viper.GetString(flagClientSign)

			var nodeOwnerSignature auth.StdSignature
			if err := cdc.UnmarshalJSON([]byte(nodeOwnerSignatureStr), &nodeOwnerSignature); err != nil {
				return err
			}

			var clientSignature auth.StdSignature
			if err := cdc.UnmarshalJSON([]byte(clientSignatureStr), &clientSignature); err != nil {
				return err
			}

			msg := vpn.NewMsgUpdateSessionInfo(cliCtx.FromAddress, id, bandwidth, nodeOwnerSignature, clientSignature)

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []csdkTypes.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagSubscriptionID, "", "Subscription ID")
	cmd.Flags().Int64(flagUpload, 0, "Upload in in bytes")
	cmd.Flags().Int64(flagDownload, 0, "Download in bytes")
	cmd.Flags().String(flagNodeOwnerSign, "", "Signature of the node owner")
	cmd.Flags().String(flagClientSign, "", "Signature of the client")

	_ = cmd.MarkFlagRequired(flagSubscriptionID)
	_ = cmd.MarkFlagRequired(flagUpload)
	_ = cmd.MarkFlagRequired(flagDownload)
	_ = cmd.MarkFlagRequired(flagNodeOwnerSign)
	_ = cmd.MarkFlagRequired(flagClientSign)

	return cmd
}

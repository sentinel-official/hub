package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/client/common"
	"github.com/sentinel-official/hub/x/vpn/types"
)

// nolint:funlen
func SignSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-bandwidth",
		Short: "Sign session bandwidth",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			_id := viper.GetString(flagSubscriptionID)
			bandwidth := hub.Bandwidth{
				Upload:   sdk.NewInt(viper.GetInt64(flagUpload)),
				Download: sdk.NewInt(viper.GetInt64(flagDownload)),
			}

			scs, err := common.QuerySessionsCountOfSubscription(ctx, _id)
			if err != nil {
				return err
			}

			id, err := hub.NewSubscriptionIDFromString(_id)
			if err != nil {
				return err
			}

			data := hub.NewBandwidthSignatureData(id, scs, bandwidth).Bytes()

			passphrase, err := keys.GetPassphrase(ctx.FromName)
			if err != nil {
				return err
			}

			kb, err := keys.NewKeyBaseFromHomeFlag()
			if err != nil {
				return err
			}

			sigBytes, pubKey, err := kb.Sign(ctx.FromName, passphrase, data)
			if err != nil {
				return err
			}

			stdSignature := auth.StdSignature{
				PubKey:    pubKey,
				Signature: sigBytes,
			}

			bytes, err := cdc.MarshalJSON(stdSignature)
			if err != nil {
				return err
			}

			fmt.Println(string(bytes))
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
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			id, err := hub.NewSubscriptionIDFromString(viper.GetString(flagSubscriptionID))
			if err != nil {
				return err
			}
			bandwidth := hub.Bandwidth{
				Upload:   sdk.NewInt(viper.GetInt64(flagUpload)),
				Download: sdk.NewInt(viper.GetInt64(flagDownload)),
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

			msg := types.NewMsgUpdateSessionInfo(ctx.FromAddress, id, bandwidth, nodeOwnerSignature, clientSignature)

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
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

package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authTxBuilder "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func SignSessionBandwidthTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-bandwidth",
		Short: "Sign session bandwidth",
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

			_, sign, err := common.BuildMsgUpdateSessionInfoAndSign(txBldr, cliCtx, cdc, id, bandwidth)
			if err != nil {
				return err
			}

			bz, err := cdc.MarshalJSON(sign)
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
			clientSignStr := viper.GetString(flagClientSign)
			bandwidth := sdkTypes.Bandwidth{
				Upload:   csdkTypes.NewInt(viper.GetInt64(flagUpload)),
				Download: csdkTypes.NewInt(viper.GetInt64(flagDownload)),
			}

			var clientSign auth.StdSignature
			if err := cdc.UnmarshalJSON([]byte(clientSignStr), &clientSign); err != nil {
				return err
			}

			stdSignMsg, sign, err := common.BuildMsgUpdateSessionInfoAndSign(txBldr, cliCtx, cdc, id, bandwidth)
			if err != nil {
				return err
			}

			stdTx := auth.NewStdTx(stdSignMsg.Msgs, stdSignMsg.Fee, []auth.StdSignature{sign, clientSign}, stdSignMsg.Memo)
			stdTxBytes, err := txBldr.TxEncoder()(stdTx)
			if err != nil {
				return err
			}

			txRes, err := cliCtx.BroadcastTx(stdTxBytes)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(txRes)
		},
	}

	cmd.Flags().String(flagSubscriptionID, "", "Subscription ID")
	cmd.Flags().Int64(flagUpload, 0, "Upload in in bytes")
	cmd.Flags().Int64(flagDownload, 0, "Download in bytes")
	cmd.Flags().String(flagClientSign, "", "Signature of the client")

	_ = cmd.MarkFlagRequired(flagSubscriptionID)
	_ = cmd.MarkFlagRequired(flagUpload)
	_ = cmd.MarkFlagRequired(flagDownload)
	_ = cmd.MarkFlagRequired(flagClientSign)

	return cmd
}

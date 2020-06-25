package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

func txRegisterNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			provider, err := hub.ProvAddressFromBech32(s)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			price, err := sdk.ParseCoins(s)
			if err != nil {
				return err
			}

			upload, err := cmd.Flags().GetUint64(flagUploadSpeed)
			if err != nil {
				return err
			}

			download, err := cmd.Flags().GetUint64(flagDownloadSpeed)
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			version, err := cmd.Flags().GetString(flagVersion)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagCategory)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterNode(ctx.FromAddress, provider, price,
				hub.NewBandwidth(upload, download), remoteURL, version, types.NodeCategoryFromString(s))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "Node provider address")
	cmd.Flags().String(flagPrice, "", "Node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "Node remove URL")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().Uint64(flagUploadSpeed, 0, "Node upload speed")
	cmd.Flags().Uint64(flagDownloadSpeed, 0, "Node download speed")
	cmd.Flags().String(flagCategory, "", "Node category")

	_ = cmd.MarkFlagRequired(flagRemoteURL)
	_ = cmd.MarkFlagRequired(flagVersion)
	_ = cmd.MarkFlagRequired(flagUploadSpeed)
	_ = cmd.MarkFlagRequired(flagDownloadSpeed)
	_ = cmd.MarkFlagRequired(flagCategory)

	return cmd
}

func txUpdateNodeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update node",
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			var provider hub.ProvAddress
			if s == "-" {
				provider = hub.EmptyProviderAddress
			} else {
				provider, err = hub.ProvAddressFromBech32(s)
				if err != nil {
					return err
				}
			}

			s, err = cmd.Flags().GetString(flagPrice)
			if err != nil {
				return err
			}

			var price sdk.Coins
			if s == "-" {
				price = hub.EmptyCoins
			} else {
				price, err = sdk.ParseCoins(s)
				if err != nil {
					return err
				}
			}

			upload, err := cmd.Flags().GetUint64(flagUploadSpeed)
			if err != nil {
				return err
			}

			download, err := cmd.Flags().GetUint64(flagDownloadSpeed)
			if err != nil {
				return err
			}

			s, err = cmd.Flags().GetString(flagCategory)
			if err != nil {
				return err
			}

			remoteURL, err := cmd.Flags().GetString(flagRemoteURL)
			if err != nil {
				return err
			}

			version, err := cmd.Flags().GetString(flagVersion)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNode(ctx.FromAddress.Bytes(), provider, price,
				hub.NewBandwidth(upload, download), remoteURL, version, types.NodeCategoryFromString(s))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "Node provider address")
	cmd.Flags().String(flagPrice, "", "Node price per Gigabyte")
	cmd.Flags().String(flagRemoteURL, "", "Node remove URL")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().Uint64(flagUploadSpeed, 0, "Node upload speed")
	cmd.Flags().Uint64(flagDownloadSpeed, 0, "Node download speed")
	cmd.Flags().String(flagCategory, "", "Node category")

	return cmd
}

func txSetNodeStatusCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status-set",
		Short: "Set node status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			msg := types.NewMsgSetNodeStatus(ctx.FromAddress.Bytes(), hub.StatusFromString(args[0]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

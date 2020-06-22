package cli

import (
	"fmt"

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

			msg := types.NewMsgRegisterNode(ctx.FromAddress, provider,
				hub.NewBandwidth(upload, download), remoteURL, version, types.NodeCategoryFromString(s))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "Node provider address")
	cmd.Flags().String(flagRemoteURL, "", "Node remove URL")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().Uint64(flagUploadSpeed, 0, "Node upload speed")
	cmd.Flags().Uint64(flagDownloadSpeed, 0, "Node download speed")
	cmd.Flags().String(flagCategory, "", "Node category")

	_ = cmd.MarkFlagRequired(flagProvider)
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
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			if !address.Equals(ctx.FromAddress) {
				return fmt.Errorf("node address is not equal to from address")
			}

			s, err := cmd.Flags().GetString(flagProvider)
			if err != nil {
				return err
			}

			provider, err := hub.ProvAddressFromBech32(s)
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

			msg := types.NewMsgUpdateNode(ctx.FromAddress.Bytes(), provider,
				hub.NewBandwidth(upload, download), remoteURL, version, types.NodeCategoryFromString(s))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProvider, "", "Node provider address")
	cmd.Flags().String(flagRemoteURL, "", "Node remove URL")
	cmd.Flags().String(flagVersion, "", "Node version")
	cmd.Flags().Uint64(flagUploadSpeed, 0, "Node upload speed")
	cmd.Flags().Uint64(flagDownloadSpeed, 0, "Node download speed")
	cmd.Flags().String(flagCategory, "", "Node category")

	return cmd
}

func txSetNodeStatusCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-status",
		Short: "Set node status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContext().WithCodec(cdc)

			address, err := hub.NodeAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			if !address.Equals(ctx.FromAddress) {
				return fmt.Errorf("node address is not equal to from address")
			}

			msg := types.NewMsgSetNodeStatus(ctx.FromAddress.Bytes(), hub.StatusFromString(args[1]))
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

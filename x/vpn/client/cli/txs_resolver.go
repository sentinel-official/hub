package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/sentinel-official/hub/x/vpn/types"
	"github.com/spf13/cobra"
)

func RegisterResolverTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [commission-rate]",
		Short: "Register node as Resolver",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			commission, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return err
			}

			if commission.LT(sdk.ZeroDec()) || commission.GT(sdk.OneDec()) {
				return fmt.Errorf("commission rate %s : between 0 and 1 ", commission.String())
			}

			msg := types.NewMsgRegisterResolver(ctx.GetFromAddress(), commission)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	return cmd
}

func UpdateResolverInfoTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [commission-rate]",
		Short: "Update the info of Resolver node",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			commission, err := sdk.NewDecFromStr(args[0])
			if err != nil {
				return err
			}

			if commission.LT(sdk.ZeroDec()) || commission.GT(sdk.OneDec()) {
				return fmt.Errorf("commission rate %s : between 0 and 1 ", commission.String())
			}

			msg := types.NewMsgUpdateResolverInfo(ctx.GetFromAddress(), commission)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	return cmd
}

func DeregisterResolverTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "de-register [address]",
		Short: "Deregister from resolver node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCLIContext().WithCodec(cdc)
			txb := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeregisterResolver(ctx.GetFromAddress())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
	return cmd
}

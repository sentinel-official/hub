package cli

import (
	"bufio"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func txUpsert(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "upsert [subscription] [address] [duration] [upload] [download] (signature)",
		Short: "Add or update a session",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			buffer := bufio.NewReader(cmd.InOrStdin())
			txb := auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContextWithInput(buffer).WithCodec(cdc)

			subscription, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			address, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			duration, err := time.ParseDuration(args[2])
			if err != nil {
				return err
			}

			upload, err := strconv.ParseInt(args[3], 10, 64)
			if err != nil {
				return err
			}

			download, err := strconv.ParseInt(args[4], 10, 64)
			if err != nil {
				return err
			}

			var signature []byte
			if len(args[5]) > 0 {
				signature, err = hex.DecodeString(args[5])
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpsert(ctx.FromAddress.Bytes(), subscription, address,
				duration, hub.NewBandwidthFromInt64(upload, download), signature)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}
}

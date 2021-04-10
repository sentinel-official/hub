package cli

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func txUpsert() *cobra.Command {
	return &cobra.Command{
		Use:   "upsert [subscription] [address] [duration] [upload] [download] (channel) (signature)",
		Short: "Add or update a session",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

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

			var channel uint64
			if len(args) > 5 && args[5] != "" {
				channel, err = strconv.ParseUint(args[3], 10, 64)
				if err != nil {
					return err
				}
			}

			var signature []byte = nil
			if len(args) > 6 && args[6] != "" {
				signature, err = hex.DecodeString(args[6])
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpsertRequest(
				types.NewProof(
					channel,
					subscription,
					ctx.FromAddress.Bytes(),
					duration,
					hub.NewBandwidthFromInt64(upload, download),
				), address, signature)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
}

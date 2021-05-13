package cli

import (
	"encoding/hex"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func txUpsert() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id] [upload] [download] [duration] (signature)",
		Short: "Add or update a session",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[1], 10, 64)
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

			var signature []byte = nil
			if len(args) > 5 && args[5] != "" {
				signature, err = hex.DecodeString(args[6])
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgUpdateRequest(
				ctx.FromAddress.Bytes(),
				types.Proof{
					Id:        id,
					Duration:  duration,
					Bandwidth: hubtypes.NewBandwidthFromInt64(upload, download),
				}, signature)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

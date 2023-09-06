// DO NOT COVER

package cli

import (
	"encoding/hex"
	"strconv"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/v1/x/swap/types"
)

func txSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [txHash] [receiver] [amount]",
		Short: "Swap from SENT to DVPN",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txHash, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			receiver, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			amount, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwapRequest(
				ctx.FromAddress,
				types.BytesToHash(txHash),
				receiver,
				sdkmath.NewInt(amount),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

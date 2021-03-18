package cli

import (
	"bufio"
	"encoding/hex"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/swap/types"
)

func txSwap(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [txHash] [receiver] [amount]",
		Short: "Swap from SENT to DVPN",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			buffer := bufio.NewReader(cmd.InOrStdin())
			txb := auth.NewTxBuilderFromCLI(buffer).WithTxEncoder(utils.GetTxEncoder(cdc))
			ctx := context.NewCLIContextWithInput(buffer).WithCodec(cdc)

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

			msg := types.NewMsgSwap(ctx.FromAddress, types.BytesToHash(txHash), receiver, sdk.NewInt(amount))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(ctx, txb, []sdk.Msg{msg})
		},
	}

	return cmd
}

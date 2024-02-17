package cli

import (
	"fmt"
	"strconv"

	"github.com/sentinel-official/hub/v12/x/pricemanager/types"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

var _ = strconv.Itoa(0)

func CmdQueryState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-state [sequence]",
		Short: "Returns the request and response of an ICQ query given the packet sequence",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			sequence, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid sequence: %w", err)
			}
			params := &types.QueryQueryStateRequest{
				Sequence: sequence,
			}

			res, err := queryClient.QueryState(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

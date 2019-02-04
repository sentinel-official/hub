package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func QuerySessionCmd(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Args:  cobra.ExactArgs(1),
		Short: "Get details of a session",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			sessionID := args[0]

			key, err := cdc.MarshalBinaryLengthPrefixed(sessionID)
			if err != nil {
				return err
			}
			res, err := cliCtx.QueryStore(key, storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no session found")
			}

			var session vpn.SessionDetails
			if err := cdc.UnmarshalBinaryLengthPrefixed(res, &session); err != nil {
				return err
			}

			sessionData, err := cdc.MarshalJSONIndent(session, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sessionData))

			return nil
		},
	}

	return cmd
}

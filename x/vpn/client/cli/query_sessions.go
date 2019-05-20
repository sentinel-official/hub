package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func QuerySessionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Get session",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			res, err := common.QuerySession(cliCtx, cdc, args[0])
			if err != nil {
				return err
			}

			sessionData, err := cdc.MarshalJSONIndent(res, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sessionData))

			return nil
		},
	}

	return cmd
}

func QuerySessionsCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sessions",
		Short: "Get sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			var sessions []vpn.Session
			res, err := cliCtx.QuerySubspace(vpn.SessionKeyPrefix, vpn.StoreKeySession)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("no sessions found")
			}

			for _, kv := range res {
				var session vpn.Session
				if err := cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &session); err != nil {
					return err
				}

				sessions = append(sessions, session)
			}

			sessionsData, err := cdc.MarshalJSONIndent(sessions, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sessionsData))

			return nil
		},
	}

	return cmd
}

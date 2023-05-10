// DO NOT COVER

package cli

import (
	"github.com/spf13/cobra"

	depositcli "github.com/sentinel-official/hub/x/deposit/client/cli"
	nodecli "github.com/sentinel-official/hub/x/node/client/cli"
	plancli "github.com/sentinel-official/hub/x/plan/client/cli"
	providercli "github.com/sentinel-official/hub/x/provider/client/cli"
	sessioncli "github.com/sentinel-official/hub/x/session/client/cli"
	subscriptioncli "github.com/sentinel-official/hub/x/subscription/client/cli"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the VPN module",
	}

	cmd.AddCommand(depositcli.GetQueryCommands()...)
	cmd.AddCommand(providercli.GetQueryCommands()...)
	cmd.AddCommand(nodecli.GetQueryCommands()...)
	cmd.AddCommand(plancli.GetQueryCommands()...)
	cmd.AddCommand(subscriptioncli.GetQueryCommands()...)
	cmd.AddCommand(sessioncli.GetQueryCommands()...)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	cmd.AddCommand(providercli.GetTxCommands()...)
	cmd.AddCommand(nodecli.GetTxCommands()...)
	cmd.AddCommand(plancli.GetTxCommands()...)
	cmd.AddCommand(subscriptioncli.GetTxCommands()...)
	cmd.AddCommand(sessioncli.GetTxCommands()...)

	return cmd
}

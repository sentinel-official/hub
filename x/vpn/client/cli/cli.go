package cli

import (
	"github.com/spf13/cobra"

	deposit "github.com/sentinel-official/hub/x/deposit/client/cli"
	node "github.com/sentinel-official/hub/x/node/client/cli"
	plan "github.com/sentinel-official/hub/x/plan/client/cli"
	provider "github.com/sentinel-official/hub/x/provider/client/cli"
	session "github.com/sentinel-official/hub/x/session/client/cli"
	subscription "github.com/sentinel-official/hub/x/subscription/client/cli"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the VPN module",
	}

	cmd.AddCommand(deposit.GetQueryCommands()...)
	cmd.AddCommand(provider.GetQueryCommands()...)
	cmd.AddCommand(node.GetQueryCommands()...)
	cmd.AddCommand(plan.GetQueryCommands()...)
	cmd.AddCommand(subscription.GetQueryCommands()...)
	cmd.AddCommand(session.GetQueryCommands()...)

	return cmd
}

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	cmd.AddCommand(provider.GetTxCommands()...)
	cmd.AddCommand(node.GetTxCommands()...)
	cmd.AddCommand(plan.GetTxCommands()...)
	cmd.AddCommand(subscription.GetTxCommands()...)
	cmd.AddCommand(session.GetTxCommands()...)

	return cmd
}

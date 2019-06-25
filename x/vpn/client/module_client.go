package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	vpnCli "github.com/sentinel-official/sentinel-hub/x/vpn/client/cli"
)

type ModuleClient struct {
	cdc *codec.Codec
}

func NewModuleClient(cdc *codec.Codec) ModuleClient {
	return ModuleClient{cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	vpnQueryCmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the vpn module",
	}

	vpnQueryCmd.AddCommand(client.GetCommands(
		vpnCli.QueryNodeCmd(mc.cdc),
		vpnCli.QueryNodesCmd(mc.cdc),
		vpnCli.QuerySubscriptionCmd(mc.cdc),
		vpnCli.QuerySubscriptionsCmd(mc.cdc),
		vpnCli.QuerySessionCmd(mc.cdc),
		vpnCli.QuerySessionsCmd(mc.cdc),
	)...)

	return vpnQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	vpnTxCmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	vpnTxCmd.AddCommand(nodeTxCmd(mc.cdc),
		subscriptionTxCmd(mc.cdc),
		sessionTxCmd(mc.cdc))

	return vpnTxCmd
}

func nodeTxCmd(cdc *codec.Codec) *cobra.Command {
	nodeTxCmd := &cobra.Command{
		Use:   "node",
		Short: "Node transactions subcommands",
	}

	nodeTxCmd.AddCommand(client.PostCommands(
		vpnCli.RegisterNodeTxCmd(cdc),
		vpnCli.UpdateNodeInfoTxCmd(cdc),
		vpnCli.UpdateNodeStatusTxCmd(cdc),
		vpnCli.DeregisterNodeTxCmd(cdc),
	)...)

	return nodeTxCmd
}

func subscriptionTxCmd(cdc *codec.Codec) *cobra.Command {
	subscriptionTxCmd := &cobra.Command{
		Use:   "subscription",
		Short: "Client subscription subcommands",
	}

	subscriptionTxCmd.AddCommand(client.PostCommands(
		vpnCli.StartSubscriptionTxCmd(cdc),
		vpnCli.EndSubscriptionTxCmd(cdc),
	)...)

	return subscriptionTxCmd
}

func sessionTxCmd(cdc *codec.Codec) *cobra.Command {
	sessionTxCmd := &cobra.Command{
		Use:   "session",
		Short: "Session transactions subcommands",
	}

	sessionTxCmd.AddCommand(client.PostCommands(
		vpnCli.SignSessionBandwidthTxCmd(cdc),
		vpnCli.UpdateSessionInfoTxCmd(cdc),
	)...)

	return sessionTxCmd
}

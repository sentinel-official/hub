package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the vpn module",
	}
	
	cmd.AddCommand(client.GetCommands(
		QueryNodeCmd(cdc),
		QueryNodesCmd(cdc),
		QuerySubscriptionCmd(cdc),
		QuerySubscriptionsCmd(cdc),
		QuerySessionCmd(cdc),
		QuerySessionsCmd(cdc),
		QueryFreeClientsCmd(cdc),
		QueryFreeNodesCmd(cdc),
		QueryResolversOfNodeCmd(cdc),
		QueryNodesOfResolverCmd(cdc),
		QueryResolversCmd(cdc),
		QueryParams(cdc),
	)...)
	
	return cmd
}

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}
	
	cmd.AddCommand(
		nodeTxCmd(cdc),
		subscriptionTxCmd(cdc),
		sessionTxCmd(cdc),
		resolverTxCmd(cdc),
	)
	
	return cmd
}

func nodeTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node transactions subcommands",
	}
	
	cmd.AddCommand(client.PostCommands(
		RegisterNodeTxCmd(cdc),
		UpdateNodeInfoTxCmd(cdc),
		AddFreeClientTxCmd(cdc),
		RemoveFreeClientTxCmd(cdc),
		RegisterVPNOnResolverTxCmd(cdc),
		RemoveVPNOnResolverTxCmd(cdc),
		DeregisterNodeTxCmd(cdc),
	)...)
	
	return cmd
}

func subscriptionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Client subscription subcommands",
	}
	
	cmd.AddCommand(client.PostCommands(
		StartSubscriptionTxCmd(cdc),
		EndSubscriptionTxCmd(cdc),
	)...)
	
	return cmd
}

func sessionTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session transactions subcommands",
	}
	
	cmd.AddCommand(client.PostCommands(
		SignSessionBandwidthTxCmd(cdc),
		UpdateSessionInfoTxCmd(cdc),
	)...)
	
	return cmd
}

func resolverTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resolver",
		Short: "Resolver node subcommands",
	}
	
	cmd.AddCommand(client.PostCommands(
		RegisterResolverTxCmd(cdc),
		UpdateResolverInfoTxCmd(cdc),
		DeregisterResolverTxCmd(cdc),
	)...)
	
	return cmd
}

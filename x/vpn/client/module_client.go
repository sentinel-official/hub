package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"

	vpnCli "github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/cli"
)

type ModuleClient struct {
	nodeStoreKey    string
	sessionStoreKey string
	cdc             *amino.Codec
}

func NewModuleClient(nodeStoreKey, sessionStoreKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{
		nodeStoreKey,
		sessionStoreKey,
		cdc,
	}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	vpnQueryCmd := &cobra.Command{
		Use:   "vpn",
		Short: "Querying commands for the vpn module",
	}

	vpnQueryCmd.AddCommand(client.GetCommands(
		vpnCli.QueryNodeCmd(mc.cdc),
		vpnCli.QueryNodesCmd(mc.cdc),
		vpnCli.QuerySessionCmd(mc.sessionStoreKey, mc.cdc),
	)...)
	return vpnQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	vpnTxCmd := &cobra.Command{
		Use:   "vpn",
		Short: "VPN transactions subcommands",
	}

	vpnTxCmd.AddCommand(client.PostCommands(vpnCli.RegisterNodeTxCmd(mc.cdc))...)
	vpnTxCmd.AddCommand(vpnCli.UpdateNodeInfoTxCmd(mc.cdc))
	vpnTxCmd.AddCommand(client.PostCommands(vpnCli.DeregisterNodeTxCmd(mc.cdc))...)
	vpnTxCmd.AddCommand(client.PostCommands(vpnCli.InitSessionTxCmd(mc.cdc))...)

	return vpnTxCmd
}

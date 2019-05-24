package client

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/client/cli"
)

type ModuleClient struct {
	cdc *codec.Codec
}

func NewModuleClient(cdc *codec.Codec) ModuleClient {
	return ModuleClient{
		cdc,
	}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	return cli.QueryDepositsCmd(mc.cdc)
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	return &cobra.Command{}
}

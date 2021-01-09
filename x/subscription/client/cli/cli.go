package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return flags.GetCommands(
		querySubscription(cdc),
		querySubscriptions(cdc),
		queryQuota(cdc),
		queryQuotas(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(flags.PostCommands(
		txSubscribeToNode(cdc),
		txSubscribeToPlan(cdc),
		txCancel(cdc),
		txAddQuota(cdc),
		txUpdateQuota(cdc),
	)...)

	return []*cobra.Command{cmd}
}

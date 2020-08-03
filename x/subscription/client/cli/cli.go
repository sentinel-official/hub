package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCommands(cdc *codec.Codec) []*cobra.Command {
	return client.GetCommands(
		querySubscriptionCmd(cdc),
		querySubscriptionsCmd(cdc),
		queryQuotaCmd(cdc),
		queryQuotasCmd(cdc),
	)
}

func GetTxCommands(cdc *codec.Codec) []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(client.PostCommands(
		txStartPlanSubscription(cdc),
		txStartNodeSubscription(cdc),
		txAddQuotaForSubscription(cdc),
		txUpdateQuotaForSubscription(cdc),
		txEndSubscription(cdc),
	)...)

	return []*cobra.Command{cmd}
}

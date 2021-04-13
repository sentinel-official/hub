package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		querySubscription(),
		querySubscriptions(),
		queryQuota(),
		queryQuotas(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "subscription",
		Short: "Subscription module sub-commands",
	}

	cmd.AddCommand(
		txSubscribeToNode(),
		txSubscribeToPlan(),
		txCancel(),
		txAddQuota(),
		txUpdateQuota(),
	)

	return []*cobra.Command{cmd}
}

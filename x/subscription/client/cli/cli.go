package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		querySubscription(),
		querySubscriptions(),
		queryParams(),
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
		txCancel(),
		txShare(),
	)

	return []*cobra.Command{cmd}
}

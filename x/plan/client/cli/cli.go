// DO NOT COVER

package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryPlan(),
		queryPlans(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan",
		Short: "Plan module sub-commands",
	}

	cmd.AddCommand(
		txCreate(),
		txSetStatus(),
		txLinkNode(),
		txUnlinkNode(),
	)

	return []*cobra.Command{cmd}
}

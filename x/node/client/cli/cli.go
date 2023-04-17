package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryNode(),
		queryNodes(),
		queryLease(),
		queryLeases(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node module sub-commands",
	}

	cmd.AddCommand(
		txRegister(),
		txUpdateDetails(),
		txUpdateStatus(),
	)

	return []*cobra.Command{cmd}
}

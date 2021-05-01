package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryNode(),
		queryNodes(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node module sub-commands",
	}

	cmd.AddCommand(
		txRegister(),
		txUpdate(),
		txSetStatus(),
	)

	return []*cobra.Command{cmd}
}

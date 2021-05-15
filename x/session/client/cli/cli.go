package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		querySession(),
		querySessions(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "session",
		Short: "Session module sub-commands",
	}

	cmd.AddCommand(
		txStart(),
		txUpdate(),
		txEnd(),
	)

	return []*cobra.Command{cmd}
}

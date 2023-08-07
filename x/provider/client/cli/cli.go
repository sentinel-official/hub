// DO NOT COVER

package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryProvider(),
		queryProviders(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "provider",
		Short: "Provider module sub-commands",
	}

	cmd.AddCommand(
		txRegister(),
		txUpdate(),
	)

	return []*cobra.Command{cmd}
}

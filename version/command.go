package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	flagLong = "long"
)

// nolint:gochecknoglobals
var (
	Cmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		RunE: func(_ *cobra.Command, _ []string) error {
			info := newInfo()

			if !viper.GetBool(flagLong) {
				fmt.Println(info.Version)
				return nil
			}

			if viper.GetString(cli.OutputFlag) != "json" {
				fmt.Print(info)
				return nil
			}

			bz, err := json.Marshal(info)
			if err != nil {
				return err
			}

			fmt.Println(string(bz))
			return nil
		},
	}
)

// nolint:gochecknoinits
func init() {
	Cmd.Flags().Bool(flagLong, false, "Print long version information")
}

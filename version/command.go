package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("sentinel-sdk:", Version)
			fmt.Println("git commit:", Commit)
			fmt.Println("vendor hash:", VendorDirHash)
			fmt.Printf("go version %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
)

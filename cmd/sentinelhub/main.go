package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	sentinelhub "github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/cmd/sentinelhub/cmd"
)

func main() {
	root, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(root, sentinelhub.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

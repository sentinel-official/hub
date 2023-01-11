package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/sentinel-official/hub"
)

func main() {
	root, _ := NewRootCmd()
	if err := servercmd.Execute(root, hub.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}

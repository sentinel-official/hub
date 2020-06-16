package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	provider "github.com/sentinel-official/hub/x/dvpn/provider/client/rest"
)

func RegisterRoutes(context context.CLIContext, router *mux.Router) {
	provider.RegisterRoutes(context, router)
}

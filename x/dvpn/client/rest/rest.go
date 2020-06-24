package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	node "github.com/sentinel-official/hub/x/dvpn/node/client/rest"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/client/rest"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/client/rest"
)

func RegisterRoutes(context context.CLIContext, router *mux.Router) {
	provider.RegisterRoutes(context, router)
	node.RegisterRoutes(context, router)
	subscription.RegisterRoutes(context, router)
}

package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

func RegisterRoutes(context context.CLIContext, router *mux.Router) {
	provider.RegisterRoutes(context, router)
	node.RegisterRoutes(context, router)
	subscription.RegisterRoutes(context, router)
}

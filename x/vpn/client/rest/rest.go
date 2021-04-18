package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	deposit "github.com/sentinel-official/hub/x/deposit/client/rest"
	node "github.com/sentinel-official/hub/x/node/client/rest"
	plan "github.com/sentinel-official/hub/x/plan/client/rest"
	provider "github.com/sentinel-official/hub/x/provider/client/rest"
	session "github.com/sentinel-official/hub/x/session/client/rest"
	subscription "github.com/sentinel-official/hub/x/subscription/client/rest"
)

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	deposit.RegisterRoutes(ctx, router)
	provider.RegisterRoutes(ctx, router)
	node.RegisterRoutes(ctx, router)
	plan.RegisterRoutes(ctx, router)
	subscription.RegisterRoutes(ctx, router)
	session.RegisterRoutes(ctx, router)
}

package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/plans", queryPlans(ctx)).
		Methods(http.MethodGet)
	router.HandleFunc("/plans/{id}", queryPlan(ctx)).
		Methods(http.MethodGet)
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
}

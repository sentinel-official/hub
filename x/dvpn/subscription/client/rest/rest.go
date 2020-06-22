package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", queryPlansHandlerFunc(ctx)).
		Methods("GET")
	router.HandleFunc("/plans/{id}", queryPlanHandlerFunc(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", txAddPlanHandlerFunc(ctx)).
		Methods("GET")
	router.HandleFunc("/plans/{id}", txSetPlanStatusHandlerFunc(ctx)).
		Methods("GET")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

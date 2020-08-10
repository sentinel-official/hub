package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", queryPlans(ctx)).
		Methods("GET")
	router.HandleFunc("/plans/{id}", queryPlan(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", txAdd(ctx)).
		Methods("POST")
	router.HandleFunc("/plans/{id}/status", txSetStatus(ctx)).
		Methods("PUT")
	router.HandleFunc("/plans/{id}/nodes", txAddNode(ctx)).
		Methods("POST")
	router.HandleFunc("/plans/{id}/nodes/{address}", txRemoveNode(ctx)).
		Methods("DELETE")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

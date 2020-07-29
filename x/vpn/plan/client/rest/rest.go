package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", queryPlansHandlerFunc(ctx)).
		Methods("GET")
	router.HandleFunc("/plans/{address}", queryPlansForProviderHandlerFunc(ctx)).
		Methods("GET")
	router.HandleFunc("/plans/{address}/{id}", queryPlanHandlerFunc(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/plans", txAddPlanHandlerFunc(ctx)).
		Methods("POST")
	router.HandleFunc("/plans/{address}/{id}/status", txSetPlanStatusHandlerFunc(ctx)).
		Methods("PUT")
	router.HandleFunc("/plans/{address}/{id}/nodes", txAddNodeHandlerFunc(ctx)).
		Methods("POST")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

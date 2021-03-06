package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/providers", queryProviders(ctx)).
		Methods("GET")
	router.HandleFunc("/providers/{address}", queryProvider(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/providers", txRegister(ctx)).
		Methods("POST")
	router.HandleFunc("/providers/{address}", txUpdate(ctx)).
		Methods("PUT")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

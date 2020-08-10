package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/nodes", queryNodes(ctx)).
		Methods("GET")
	router.HandleFunc("/nodes/{address}", queryNode(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/nodes", txRegister(ctx)).
		Methods("POST")
	router.HandleFunc("/nodes/{address}", txUpdate(ctx)).
		Methods("PUT")
	router.HandleFunc("/nodes/{address}/status", txSetStatus(ctx)).
		Methods("PUT")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

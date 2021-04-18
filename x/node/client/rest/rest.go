package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/nodes", queryNodes(ctx)).
		Methods("GET")
	router.HandleFunc("/nodes/{address}", queryNode(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/nodes", txRegister(ctx)).
		Methods("POST")
	router.HandleFunc("/nodes/{address}", txUpdate(ctx)).
		Methods("PUT")
	router.HandleFunc("/nodes/{address}/status", txSetStatus(ctx)).
		Methods("PUT")
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

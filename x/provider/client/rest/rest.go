package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/providers", queryProviders(ctx)).
		Methods("GET")
	router.HandleFunc("/providers/{address}", queryProvider(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/providers", txRegister(ctx)).
		Methods("POST")
	router.HandleFunc("/providers/{address}", txUpdate(ctx)).
		Methods("PUT")
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

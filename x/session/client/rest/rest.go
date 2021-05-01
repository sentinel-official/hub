package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/sessions", querySessions(ctx)).
		Methods("GET")
	router.HandleFunc("/sessions/{id}", querySession(ctx)).
		Methods("GET")
}

func registerTxRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/sessions", txUpsert(ctx)).
		Methods("POST")
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
	registerTxRoutes(ctx, router)
}

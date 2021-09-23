package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/nodes", queryNodes(ctx)).
		Methods(http.MethodGet)
	router.HandleFunc("/nodes/{address}", queryNode(ctx)).
		Methods(http.MethodGet)
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
}

package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/deposits", queryDeposits(ctx)).
		Methods("GET")
	router.HandleFunc("/deposits/{address}", queryDeposit(ctx)).
		Methods("GET")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
}

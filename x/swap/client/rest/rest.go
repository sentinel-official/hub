package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx context.CLIContext, router *mux.Router) {
	router.HandleFunc("/swaps", querySwaps(ctx)).
		Methods("GET")
	router.HandleFunc("/swaps/{txHash}", querySwap(ctx)).
		Methods("GET")
}

func RegisterRoutes(ctx context.CLIContext, router *mux.Router) {
	registerQueryRoutes(ctx, router)
}

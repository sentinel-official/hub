package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(ctx, r)
}

func registerQueryRoutes(ctx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/deposits", getAllDeposits(ctx)).
		Methods("GET")
	r.HandleFunc("/deposits/{address}", getDepositOfAddressHandlerFunc(ctx)).
		Methods("GET")
}

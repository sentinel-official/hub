package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	registerQueryRoutes(cliCtx, r, cdc)
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/deposits", getAllDeposits(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/deposits/{address}", getDepositOfAddressHandlerFunc(cliCtx, cdc)).
		Methods("GET")
}

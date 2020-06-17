package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(context context.CLIContext, router *mux.Router) {

}

func registerTxRoutes(context context.CLIContext, router *mux.Router) {

}

func RegisterRoutes(context context.CLIContext, router *mux.Router) {
	registerQueryRoutes(context, router)
	registerTxRoutes(context, router)
}

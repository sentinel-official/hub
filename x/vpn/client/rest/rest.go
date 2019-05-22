package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	registerTxRoutes(cliCtx, r, cdc)
	registerQueryRoutes(cliCtx, r, cdc)
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/nodes", registerNodeHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/nodes/{nodeID}/deregister", deregisterNodeHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/nodes/{nodeID}/details", updateNodeHandlerFunc(cliCtx, cdc)).
		Methods("PUT")
	r.HandleFunc("/nodes/{nodeID}/status", updateNodeStatusHandlerFunc(cliCtx, cdc)).
		Methods("PUT")

	r.HandleFunc("/subscribe", startSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/subscribe/{subscriptionID}", endSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("PUT")

	r.HandleFunc("/sessions/{sessionID}/bandwidth/sign", signSessionBandwidthHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/sessions/{sessionID}/bandwidth", updateSessionInfoHandlerFunc(cliCtx, cdc)).
		Methods("POST")

}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/nodes", getNodesHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{nodeID}", getNodeHandlerFunc(cliCtx, cdc)).
		Methods("GET")

	r.HandleFunc("/subscribe/{subscriptionID}", getSubscribeHandlerFunc(cliCtx, cdc)).
		Methods("GET")
}

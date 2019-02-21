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
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/deregister", deregisterNodeHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/details", updateNodeDetailsHandlerFunc(cliCtx, cdc)).
		Methods("PUT")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/status", updateNodeStatusHandlerFunc(cliCtx, cdc)).
		Methods("PUT")

	r.HandleFunc("/sessions", initSessionHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/sessions/{sessionID:[^/]+/[^/]+}/bandwidth/sign", signSessionBandwidthHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/sessions/{sessionID:[^/]+/[^/]+}/bandwidth", updateSessionBandwidthHandlerFunc(cliCtx, cdc)).
		Methods("PUT")
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/nodes", getNodesHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}", getNodeHandlerFunc(cliCtx, cdc)).
		Methods("GET")

	r.HandleFunc("/sessions/{sessionID:[^/]+/[^/]+}", getSessionHandlerFunc(cliCtx, cdc)).
		Methods("GET")
}

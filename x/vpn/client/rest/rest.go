package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	registerTxRoutes(cliCtx, r, cdc, kb)
	registerQueryRoutes(cliCtx, r, cdc, kb)
}

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	r.HandleFunc("/nodes", registerNodeHandlerFunc(cliCtx, cdc, kb)).
		Methods("POST")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/deregister", deregisterNodeHandlerFunc(cliCtx, cdc, kb)).
		Methods("POST")
	r.HandleFunc("/sessions", initSessionHandlerFunc(cliCtx, cdc, kb)).
		Methods("POST")

	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/details", updateNodeDetailsHandlerFunc(cliCtx, cdc, kb)).
		Methods("PUT")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/status", updateNodeStatusHandlerFunc(cliCtx, cdc, kb)).
		Methods("PUT")
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	r.HandleFunc("/nodes", getNodesHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}", getNodeHandlerFunc(cliCtx, cdc)).
		Methods("GET")
}

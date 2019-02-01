package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	r.HandleFunc("/nodes", registerNodeHandlerFunc(cliCtx, cdc, kb)).
		Methods("POST")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}", updateNodeHandlerFunc(cliCtx, cdc, kb)).
		Methods("PUT")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}/deregister", deregisterNodeHandlerFunc(cliCtx, cdc, kb)).
		Methods("POST")

	r.HandleFunc("/nodes", getNodes(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{nodeID:[^/]+/[^/]+}", getNode(cliCtx, cdc)).
		Methods("GET")
}

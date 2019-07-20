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
	r.HandleFunc("/nodes/{id}", deregisterNodeHandlerFunc(cliCtx, cdc)).
		Methods("DELETE")
	r.HandleFunc("/nodes/{id}/info", updateNodeInfoHandlerFunc(cliCtx, cdc)).
		Methods("PUT")
	r.HandleFunc("/nodes/{id}/status", updateNodeStatusHandlerFunc(cliCtx, cdc)).
		Methods("PUT")
	r.HandleFunc("/nodes/{id}/subscriptions", startSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("POST")

	r.HandleFunc("/subscriptions/{id}", endSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("DELETE")
	r.HandleFunc("/subscriptions/{id}/sessions/bandwidth/sign", signSessionBandwidthHandlerFunc(cliCtx, cdc)).
		Methods("POST")
	r.HandleFunc("/subscriptions/{id}/sessions", updateSessionInfoHandlerFunc(cliCtx, cdc)).
		Methods("PUT")

}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec) {
	r.HandleFunc("/nodes", getAllNodesHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}", getNodeHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}/subscriptions", getSubscriptionsOfNodeHandlerFunc(cliCtx, cdc)).
		Methods("GET")

	r.HandleFunc("/subscriptions", getAllSubscriptionsHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/subscriptions/{id}", getSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/subscriptions/{id}/sessions", getSessionsOfSubscriptionHandlerFunc(cliCtx, cdc)).
		Methods("GET")

	r.HandleFunc("/sessions", getAllSessionsHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/sessions/{id}", getSessionHandlerFunc(cliCtx, cdc)).
		Methods("GET")

	r.HandleFunc("/accounts/{address}/subscriptions", getSubscriptionsOfAddressHandlerFunc(cliCtx, cdc)).
		Methods("GET")
	r.HandleFunc("/accounts/{address}/nodes", getNodesOfAddressHandlerFunc(cliCtx, cdc)).
		Methods("GET")

}

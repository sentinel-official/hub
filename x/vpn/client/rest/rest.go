package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.CLIContext, r *mux.Router) {
	registerTxRoutes(ctx, r)
	registerQueryRoutes(ctx, r)
}

func registerTxRoutes(ctx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nodes", registerNodeHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/nodes/{id}", deregisterNodeHandlerFunc(ctx)).
		Methods("DELETE")
	r.HandleFunc("/nodes/{id}/info", updateNodeInfoHandlerFunc(ctx)).
		Methods("PUT")
	r.HandleFunc("/nodes/{id}/status", updateNodeStatusHandlerFunc(ctx)).
		Methods("PUT")
	r.HandleFunc("/nodes/{id}/subscriptions", startSubscriptionHandlerFunc(ctx)).
		Methods("POST")

	r.HandleFunc("/subscriptions/{id}", endSubscriptionHandlerFunc(ctx)).
		Methods("DELETE")
	r.HandleFunc("/subscriptions/{id}/sessions/bandwidth/sign", signSessionBandwidthHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/subscriptions/{id}/sessions", updateSessionInfoHandlerFunc(ctx)).
		Methods("PUT")

}

func registerQueryRoutes(ctx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nodes", getAllNodesHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}", getNodeHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}/subscriptions", getSubscriptionsOfNodeHandlerFunc(ctx)).
		Methods("GET")

	r.HandleFunc("/subscriptions", getAllSubscriptionsHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/subscriptions/{id}", getSubscriptionHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/subscriptions/{id}/sessions", getSessionsOfSubscriptionHandlerFunc(ctx)).
		Methods("GET")

	r.HandleFunc("/sessions", getAllSessionsHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/sessions/{id}", getSessionHandlerFunc(ctx)).
		Methods("GET")

	r.HandleFunc("/accounts/{address}/subscriptions", getSubscriptionsOfAddressHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/accounts/{address}/nodes", getNodesOfAddressHandlerFunc(ctx)).
		Methods("GET")

}

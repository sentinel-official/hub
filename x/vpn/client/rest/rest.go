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
	r.HandleFunc("/nodes/{id}/add-free-client", addFreeClientHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/nodes/{id}/remove-free-client/{address}", removeFreeClientHandlerFunc(ctx)).
		Methods("DELETE")
	r.HandleFunc("/nodes/{id}/register-vpn-on-resolver", registerVPNOnResolverHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/nodes/{id}/deregister-vpn-on-resolver/{address}", deregisterVPNOnResolverHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/nodes/{id}/subscriptions", startSubscriptionHandlerFunc(ctx)).
		Methods("POST")
	
	r.HandleFunc("/subscriptions/{id}", endSubscriptionHandlerFunc(ctx)).
		Methods("DELETE")
	r.HandleFunc("/subscriptions/{id}/sessions/bandwidth/sign", signSessionBandwidthHandlerFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/subscriptions/{id}/sessions", updateSessionInfoHandlerFunc(ctx)).
		Methods("PUT")
	
	r.HandleFunc("/resolver", registerResolverHandleFunc(ctx)).
		Methods("POST")
	r.HandleFunc("/resolver/update", updateResolverHandleFunc(ctx)).
		Methods("PUT")
	r.HandleFunc("/resolver/de-register", deregisterResolverHandleFunc(ctx)).
		Methods("DELETE")
	
}

func registerQueryRoutes(ctx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/nodes", getAllNodesHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}", getNodeHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}/free-clients", getFreeClientsOfNodeHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/nodes/{id}/resolvers", getResolversOfNodeHandlerFunc(ctx)).
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
	r.HandleFunc("/accounts/{address}/free-nodes", getFreeNodesOfClientHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/accounts/{address}/resolver-nodes", getNodesOfResolverHandlerFunc(ctx)).
		Methods("GET")
	
	r.HandleFunc("/resolvers", getResolversHandlerFunc(ctx)).
		Methods("GET")
	r.HandleFunc("/vpn/params", getParamsHandlerFunc(ctx)).
		Methods("GET")
	
}

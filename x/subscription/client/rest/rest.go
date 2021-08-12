package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func registerQueryRoutes(ctx client.Context, router *mux.Router) {
	router.HandleFunc("/subscriptions", querySubscriptions(ctx)).
		Methods(http.MethodGet)
	router.HandleFunc("/subscriptions/{id}", querySubscription(ctx)).
		Methods(http.MethodGet)
	router.HandleFunc("/subscriptions/{id}/quotas", queryQuotas(ctx)).
		Methods(http.MethodGet)
	router.HandleFunc("/subscriptions/{id}/quotas/{address}", queryQuota(ctx)).
		Methods(http.MethodGet)
}

func RegisterRoutes(ctx client.Context, router *mux.Router) {
	registerQueryRoutes(ctx, router)
}

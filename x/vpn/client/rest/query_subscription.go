package rest

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	
	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func getSubscriptionHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		subscription, err := common.QuerySubscription(ctx, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, subscription)
	}
}

func getSubscriptionsOfNodeHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		subscriptions, err := common.QuerySubscriptionsOfNode(ctx, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, subscriptions)
	}
}

func getSubscriptionsOfAddressHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		subscriptions, err := common.QuerySubscriptionsOfAddress(ctx, vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, subscriptions)
	}
}

func getAllSubscriptionsHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subscriptions, err := common.QueryAllSubscriptions(ctx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, subscriptions)
	}
}

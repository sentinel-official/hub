// nolint:dupl
package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/sentinel-hub/x/vpn/client/common"
)

func getSubscriptionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		subscription, err := common.QuerySubscription(cliCtx, cdc, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, subscription, cliCtx.Indent)
	}
}

func getSubscriptionsOfNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		subscriptions, err := common.QuerySubscriptionsOfNode(cliCtx, cdc, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, subscriptions, cliCtx.Indent)
	}
}

func getSubscriptionsOfAddressHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		subscriptions, err := common.QuerySubscriptionsOfAddress(cliCtx, cdc, vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, subscriptions, cliCtx.Indent)
	}
}

func getAllSubscriptionsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subscriptions, err := common.QueryAllSubscriptions(cliCtx, cdc)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, subscriptions, cliCtx.Indent)
	}
}

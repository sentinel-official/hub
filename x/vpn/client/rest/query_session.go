package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func getAllSessionsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := common.QueryAllSessions(cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no sessions found")
			return
		}

		var sessions []vpn.Session
		if err := cdc.UnmarshalJSON(res, &sessions); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func getSessionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["sessionID"])
		res, err := common.QuerySession(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no sessions found")
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func getSessionsOfSubscriptionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["subscriptionID"])
		res, err := common.QuerySessionsOfSubscription(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no sessions found")
			return
		}

		var subscriptions []vpn.Subscription
		if err := cdc.UnmarshalJSON(res, &subscriptions); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

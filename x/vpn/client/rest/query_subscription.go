package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

// nolint:dupl
func getAllSubscriptionsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := common.QueryAllSubscriptions(cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no subscriptions found")
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

// nolint:dupl
func getSubscriptionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["subscriptionID"])
		res, err := common.QuerySubscription(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no subscriptions found")
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

// nolint:dupl
func getSubscriptionsOfNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["nodeID"])
		res, err := common.QuerySubscriptionsOfNode(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no subscriptions found")
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

func getSubscriptionsOfAddressHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var err error

		address := vars["address"]
		_address, err := csdkTypes.AccAddressFromBech32(address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := common.QuerySubscriptionsOfAddress(cliCtx, cdc, _address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no subscriptions found")
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

package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		res, err := qc.QuerySubscription(context.Background(),
			types.NewQuerySubscriptionRequest(id))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func querySubscriptions(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			query  = r.URL.Query()
			status = hub.StatusFromString(query.Get("status"))
			qc     = types.NewQueryServiceClient(ctx)
		)

		if query.Get("address") != "" {
			address, err := sdk.AccAddressFromBech32(query.Get("address"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err := qc.QuerySubscriptionsForAddress(context.Background(),
				types.NewQuerySubscriptionsForAddressRequest(address, status, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		} else if query.Get("plan") != "" {
			id, err := strconv.ParseUint(query.Get("plan"), 10, 64)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err := qc.QuerySubscriptionsForPlan(context.Background(),
				types.NewQuerySubscriptionsForPlanRequest(id, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		} else if query.Get("node") != "" {
			address, err := hub.NodeAddressFromBech32(query.Get("node"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err := qc.QuerySubscriptionsForNode(context.Background(),
				types.NewQuerySubscriptionsForNodeRequest(address, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		} else {
			res, err := qc.QuerySubscriptions(context.Background(),
				types.NewQuerySubscriptionsRequest(nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		}
	}
}

func queryQuota(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		address, err := sdk.AccAddressFromBech32(vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		res, err := qc.QueryQuota(context.Background(),
			types.NewQueryQuotaRequest(id, address))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryQuotas(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			vars = mux.Vars(r)
			qc   = types.NewQueryServiceClient(ctx)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := qc.QueryQuotas(context.Background(),
			types.NewQueryQuotasRequest(id, nil))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/subscription/client/common"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		subscription, err := common.QuerySubscription(ctx, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, subscription)
	}
}

func querySubscriptions(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		skip, limit, err := utils.ParseQuery(query)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			plan          uint64
			address       sdk.AccAddress
			node          hub.NodeAddress
			subscriptions types.Subscriptions
			status        = hub.StatusFromString(query.Get("status"))
		)

		if query.Get("address") != "" {
			address, err = sdk.AccAddressFromBech32(query.Get("address"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			subscriptions, err = common.QuerySubscriptionsForAddress(ctx, address, status, skip, limit)
		} else if query.Get("plan") != "" {
			plan, err = strconv.ParseUint(query.Get("plan"), 10, 64)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			subscriptions, err = common.QuerySubscriptionsForPlan(ctx, plan, skip, limit)
		} else if query.Get("node") != "" {
			node, err = hub.NodeAddressFromBech32(query.Get("node"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			subscriptions, err = common.QuerySubscriptionsForNode(ctx, node, skip, limit)
		} else {
			subscriptions, err = common.QuerySubscriptions(ctx, skip, limit)
		}

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, subscriptions)
	}
}

func queryQuota(ctx context.CLIContext) http.HandlerFunc {
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

		quota, err := common.QueryQuota(ctx, id, address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, quota)
	}
}

func queryQuotas(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		skip, limit, err := utils.ParseQuery(query)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		quotas, err := common.QueryQuotas(ctx, id, skip, limit)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, quotas)
	}
}

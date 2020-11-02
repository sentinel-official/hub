package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/plan/client/common"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		plan, err := common.QueryPlan(ctx, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, plan)
	}
}

func queryPlans(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		skip, limit, err := utils.ParseQuery(query)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			provider hub.ProvAddress
			plans    types.Plans
		)

		if query.Get("provider") != "" {
			provider, err = hub.ProvAddressFromBech32(query.Get("provider"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			plans, err = common.QueryPlansForProvider(ctx, provider, skip, limit)
		} else {
			plans, err = common.QueryPlans(ctx, skip, limit)
		}

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, plans)
	}
}

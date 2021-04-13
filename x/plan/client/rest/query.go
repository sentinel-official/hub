package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx client.Context) http.HandlerFunc {
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

		res, err := qc.QueryPlan(context.Background(),
			types.NewQueryPlanRequest(id))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryPlans(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			query  = r.URL.Query()
			status = hub.StatusFromString(query.Get("status"))
			qc     = types.NewQueryServiceClient(ctx)
		)

		if query.Get("provider") != "" {
			address, err := hub.ProvAddressFromBech32(query.Get("provider"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err := qc.QueryPlansForProvider(context.Background(), types.NewQueryPlansForProviderRequest(address, status, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		}

		res, err := qc.QueryPlans(context.Background(), types.NewQueryPlansRequest(status, nil))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

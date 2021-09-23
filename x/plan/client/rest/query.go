package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/plan/types"
)

func queryPlan(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryPlan(
			context.Background(),
			types.NewQueryPlanRequest(
				id,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryPlans(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc     = types.NewQueryServiceClient(ctx)
			query  = r.URL.Query()
			status = hubtypes.StatusFromString(query.Get("status"))
		)

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if query.Get("provider") != "" {
			address, err := hubtypes.ProvAddressFromBech32(query.Get("provider"))
			if rest.CheckBadRequestError(w, err) {
				return
			}

			res, err := qc.QueryPlansForProvider(
				context.Background(),
				types.NewQueryPlansForProviderRequest(
					address,
					status,
					&sdkquery.PageRequest{
						Offset: uint64(page * limit),
						Limit:  uint64(limit),
					},
				),
			)
			if rest.CheckInternalServerError(w, err) {
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		}

		res, err := qc.QueryPlans(
			context.Background(),
			types.NewQueryPlansRequest(
				status,
				&sdkquery.PageRequest{
					Offset: uint64(page * limit),
					Limit:  uint64(limit),
				},
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

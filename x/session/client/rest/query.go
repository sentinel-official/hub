package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession(ctx client.Context) http.HandlerFunc {
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

		res, err := qc.QuerySession(context.Background(),
			types.NewQuerySessionRequest(id))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func querySessions(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			query = r.URL.Query()
			qc    = types.NewQueryServiceClient(ctx)
		)

		if query.Get("address") != "" {
			address, err := sdk.AccAddressFromBech32(query.Get("address"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			var status hubtypes.Status
			if query.Get("active") != "" {
				status = hubtypes.StatusActive
			}

			res, err := qc.QuerySessionsForAddress(context.Background(),
				types.NewQuerySessionsForAddressRequest(address, status, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		} else {
			res, err := qc.QuerySessions(context.Background(), types.NewQuerySessionsRequest(nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		}
	}
}

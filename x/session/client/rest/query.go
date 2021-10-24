package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QuerySession(
			context.Background(),
			types.NewQuerySessionRequest(
				id,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func querySessions(ctx client.Context) http.HandlerFunc {
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

		if query.Get("address") != "" {
			address, err := sdk.AccAddressFromBech32(query.Get("address"))
			if rest.CheckBadRequestError(w, err) {
				return
			}

			res, err := qc.QuerySessionsForAddress(
				context.Background(),
				types.NewQuerySessionsForAddressRequest(
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

		res, err := qc.QuerySessions(
			context.Background(),
			types.NewQuerySessionsRequest(
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

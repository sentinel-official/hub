package rest

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		address, err := hubtypes.NodeAddressFromBech32(vars["address"])
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryNode(
			context.Background(),
			types.NewQueryNodeRequest(
				address,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryNodes(ctx client.Context) http.HandlerFunc {
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

			res, err := qc.QueryNodesForProvider(
				context.Background(),
				types.NewQueryNodesForProviderRequest(
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
		} else if query.Get("plan") != "" {
			rest.PostProcessResponse(w, ctx, nil)
			return
		}

		res, err := qc.QueryNodes(
			context.Background(),
			types.NewQueryNodesRequest(
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

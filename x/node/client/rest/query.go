package rest

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := hub.NodeAddressFromBech32(vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		res, err := qc.QueryNode(context.Background(),
			types.NewQueryNodeRequest(address))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryNodes(ctx client.Context) http.HandlerFunc {
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

			res, err := qc.QueryNodesForProvider(context.Background(),
				types.NewQueryNodesForProviderRequest(address, status, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		} else if query.Get("plan") != "" {
			rest.PostProcessResponse(w, ctx, nil)
			return
		} else {
			res, err := qc.QueryNodes(context.Background(),
				types.NewQueryNodesRequest(status, nil))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			rest.PostProcessResponse(w, ctx, res)
			return
		}
	}
}

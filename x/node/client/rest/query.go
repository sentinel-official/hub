package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/utils"
	"github.com/sentinel-official/hub/x/node/client/common"
	"github.com/sentinel-official/hub/x/node/types"
)

func queryNode(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := hub.NodeAddressFromBech32(vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		node, err := common.QueryNode(ctx, address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, node)
	}
}

func queryNodes(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		skip, limit, err := utils.ParseQuery(query)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			plan     uint64
			provider hub.ProvAddress
			nodes    types.Nodes
			status   = hub.StatusFromString(query.Get("status"))
		)

		if query.Get("provider") != "" {
			provider, err = hub.ProvAddressFromBech32(query.Get("provider"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			nodes, err = common.QueryNodesForProvider(ctx, provider, status, skip, limit)
		} else if query.Get("plan") != "" {
			plan, err = strconv.ParseUint(query.Get("plan"), 10, 64)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			nodes, err = common.QueryNodesForPlan(ctx, plan, skip, limit)
		} else {
			nodes, err = common.QueryNodes(ctx, status, skip, limit)
		}

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, nodes)
	}
}

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
	"github.com/sentinel-official/hub/x/session/client/common"
	"github.com/sentinel-official/hub/x/session/types"
)

func querySession(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		session, err := common.QuerySession(ctx, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, session)
	}
}

func querySessions(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		page, limit, err := utils.ParseQuery(query)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			subscription uint64
			address      sdk.AccAddress
			node         hub.NodeAddress
			sessions     types.Sessions
		)

		if query.Get("subscription") != "" {
			subscription, err = strconv.ParseUint(query.Get("subscription"), 10, 64)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			sessions, err = common.QuerySessionsForSubscription(ctx, subscription, page, limit)
		} else if query.Get("node") != "" {
			node, err = hub.NodeAddressFromBech32(query.Get("node"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			sessions, err = common.QuerySessionsForNode(ctx, node, page, limit)
		} else if query.Get("address") != "" {
			address, err = sdk.AccAddressFromBech32(query.Get("address"))
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			sessions, err = common.QuerySessionsForAddress(ctx, address, page, limit)
		} else {
			sessions, err = common.QuerySessions(ctx, page, limit)
		}

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, sessions)
	}
}

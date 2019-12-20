package rest

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	
	"github.com/sentinel-official/hub/x/vpn/client/common"
)

func getNodeHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		node, err := common.QueryNode(ctx, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, node)
	}
}

func getNodesOfAddressHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		nodes, err := common.QueryNodesOfAddress(ctx, vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, nodes)
	}
}

func getAllNodesHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nodes, err := common.QueryAllNodes(ctx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, nodes)
	}
}

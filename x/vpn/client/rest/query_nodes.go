package rest

import (
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func getNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		vars := mux.Vars(r)

		nodeID := vars["nodeID"]
		if len(nodeID) == 0 {
			err := errors.New("nodeID is empty")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := common.QueryNode(cliCtx, cdc, nodeID)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if res == nil || len(res) == 0 {
			err := errors.New("no node found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, res, cliCtx.Indent)
		return
	}
}

func getNodesHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var res []byte
		var err error

		owner := r.URL.Query().Get("owner")
		if owner == "" {
			res, err = common.QueryNodes(cliCtx)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			owner, err := csdkTypes.AccAddressFromBech32(owner)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err = common.QueryNodesOfOwner(cliCtx, cdc, owner)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		if res == nil || len(res) == 0 {
			err := errors.New("no nodes found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, res, cliCtx.Indent)
		return
	}
}

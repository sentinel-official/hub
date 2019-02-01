package rest

import (
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

func getNodes(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		var res []byte
		var err error

		ownerAddress := r.URL.Query().Get("address")
		if len(ownerAddress) > 0 {
			res, err = queryNodesofOwner(cliCtx, cdc, ownerAddress)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

		} else {
			res, err = cliCtx.QueryWithData("/custom/vpn/nodes", nil)
			if err != nil {
				utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

		}

		if string(res) == "null" {
			err = errors.New("details are not found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, res, cliCtx.Indent)
		return
	}
}

func getNode(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		vars := mux.Vars(r)

		nodeID := vars["nodeID"]
		if len(nodeID) == 0 {
			err := errors.New("nodeID is empty")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		params := vpn.NewQueryNodeParams(nodeID)

		bz, err := cdc.MarshalJSON(params)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := cliCtx.QueryWithData("/custom/vpn/node", bz)
		if res == nil {
			err = errors.New("details are not found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.PostProcessResponse(w, cdc, res, cliCtx.Indent)
		return
	}
}

func queryNodesofOwner(cliCtx context.CLIContext, cdc *codec.Codec, ownerAddress string) ([]byte, error) {

	address, err := types.AccAddressFromBech32(ownerAddress)
	if err != nil {
		return nil, err
	}

	params := vpn.NewQueryNodesOfOwnerParams(address)

	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	res, err := cliCtx.QueryWithData("/custom/vpn/nodesOfOwner", bz)
	if err != nil {
		err = errors.New("details are not found")
		return nil, err
	}
	return res, nil
}

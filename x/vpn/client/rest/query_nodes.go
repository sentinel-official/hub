package rest

import (
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func getNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		res, err := common.QueryNode(cliCtx, cdc, sdkTypes.NewIDFromString(vars["nodeID"]))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

func getNodesHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res []byte

		address := r.URL.Query().Get("address")
		if len(address) == 0 {
			kvs, err := cliCtx.QuerySubspace(vpn.NodeKeyPrefix, vpn.StoreKeyNode)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			if len(kvs) == 0 {
				err := errors.New("no nodes found")
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			var nodes []vpn.Node
			for _, kv := range kvs {
				var node vpn.Node
				if err := cdc.UnmarshalBinaryLengthPrefixed(kv.Value, &node); err != nil {
					rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
					return
				}

				nodes = append(nodes, node)
			}

			if res, err = cdc.MarshalJSON(nodes); err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			address, err := csdkTypes.AccAddressFromBech32(address)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			res, err = common.QueryNodesOfAddress(cliCtx, cdc, address)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			if string(res) == "null" {
				err := errors.New("no nodes found")
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		rest.PostProcessResponse(w, cdc, res, cliCtx.Indent)
	}
}

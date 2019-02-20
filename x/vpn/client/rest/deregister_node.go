package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgDeregisterNode struct {
	BaseReq rest.BaseReq `json:"base_req"`
}

func deregisterNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgDeregisterNode

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly).WithSimulation(req.BaseReq.Simulate)

		info, err := kb.Get(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := sdkTypes.NewID(vars["nodeID"])

		msg := vpn.NewMsgDeregisterNode(info.GetAddress(), id)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

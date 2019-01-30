package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgDeregisterNode struct {
	BaseReq utils.BaseReq `json:"base_req"`
}

func deregisterNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgDeregisterNode

		vars := mux.Vars(r)
		id := vars["nodeID"]

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly)
		cliCtx.WithSimulation(req.BaseReq.Simulate)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		msg := vpn.NewMsgDeregisterNode(info.GetAddress(), id)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
		return
	}
}

package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgInitSession struct {
	BaseReq      utils.BaseReq `json:"base_req"`
	AmountToLock string        `json:"amount_to_lock"`
	NodeID       string        `json:"node_id"`
}

func initSessionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgInitSession

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly).WithSimulation(req.BaseReq.Simulate)

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		amountToLock, err := csdkTypes.ParseCoin(req.AmountToLock)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := vpn.NewMsgInitSession(info.GetAddress(), req.NodeID, amountToLock)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
		return
	}
}

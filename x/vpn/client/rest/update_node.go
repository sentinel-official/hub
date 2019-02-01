package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgUpdateNode struct {
	BaseReq     utils.BaseReq      `json:"base_req"`
	APIPort     uint16             `json:"api_port"`
	NetSpeed    sdkTypes.Bandwidth `json:"net_speed"`
	EncMethod   string             `json:"enc_method"`
	PerGBAmount string             `json:"per_gb_amount"`
	Version     string             `json:"version"`
}

func updateNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNode

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		vars := mux.Vars(r)
		id := vars["nodeID"]

		cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly)
		cliCtx.WithSimulation(req.BaseReq.Simulate)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		}

		perGBAmount, err := csdkTypes.ParseCoins(req.PerGBAmount)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		msg := vpn.NewMsgUpdateNode(info.GetAddress(), id, req.APIPort,
			req.NetSpeed.Upload, req.NetSpeed.Download, req.EncMethod,
			perGBAmount, req.Version)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())

		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
		return
	}
}

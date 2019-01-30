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

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, kb keys.Keybase) {
	r.HandleFunc("/nodes", registerNodeHandlerFunc(cliCtx, cdc, kb), ).Methods("POST")
}

type msgRegisterNode struct {
	BaseReq      utils.BaseReq      `json:"base_req"`
	AmountToLock string             `json:"amount_to_lock,required"`
	APIPort      uint16             `json:"api_port,required"`
	NetSpeed     sdkTypes.Bandwidth `json:"net_speed,required"`
	EncMethod    string             `json:"enc_method,required"`
	PerGBAmount  string             `json:"per_gb_amount,required"`
	Version      string             `json:"version,required"`
}

func registerNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRegisterNode

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		cliCtx = cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly)
		cliCtx = cliCtx.WithSimulation(req.BaseReq.Simulate)

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		cliCtx = cliCtx.WithFrom(req.BaseReq.Name)

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		amount, err := csdkTypes.ParseCoin(req.AmountToLock)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		perGBAmount, err := csdkTypes.ParseCoins(req.PerGBAmount)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		msg := vpn.NewMsgRegisterNode(info.GetAddress(),
			req.APIPort, req.NetSpeed.Upload, req.NetSpeed.Download,
			req.EncMethod, perGBAmount, req.Version, amount)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
		return
	}

}

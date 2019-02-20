package rest

import (
	"net/http"
	"strings"

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

type msgUpdateNodeDetails struct {
	BaseReq     rest.BaseReq       `json:"base_req"`
	APIPort     uint32             `json:"api_port"`
	NetSpeed    sdkTypes.Bandwidth `json:"net_speed"`
	EncMethod   string             `json:"enc_method"`
	PricesPerGB string             `json:"prices_per_gb"`
	Version     string             `json:"version"`
}

func updateNodeDetailsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNodeDetails

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

		pricesPerGB, err := csdkTypes.ParseCoins(req.PricesPerGB)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := sdkTypes.NewID(vars["nodeID"])
		apiPort := vpn.NewAPIPort(req.APIPort)

		msg := vpn.NewMsgUpdateNodeDetails(info.GetAddress(), id,
			pricesPerGB, req.NetSpeed.Upload, req.NetSpeed.Download,
			apiPort, req.EncMethod, req.Version)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

type msgUpdateNodeStatus struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Status  string       `json:"status"`
}

func updateNodeStatusHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNodeStatus

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
		status := strings.ToUpper(req.Status)

		msg := vpn.NewMsgUpdateNodeStatus(info.GetAddress(), id, status)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

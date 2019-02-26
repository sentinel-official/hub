package rest

import (
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgUpdateNodeDetails struct {
	BaseReq     rest.BaseReq       `json:"base_req"`
	APIPort     uint16             `json:"api_port"`
	NetSpeed    sdkTypes.Bandwidth `json:"net_speed"`
	Encryption  string             `json:"encryption"`
	PricesPerGB string             `json:"prices_per_gb"`
	Version     string             `json:"version"`
}

func updateNodeDetailsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNodeDetails

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, fromName, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		pricesPerGB, err := csdkTypes.ParseCoins(req.PricesPerGB)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := sdkTypes.NewID(vars["nodeID"])

		cliCtx = cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly).WithSimulation(req.BaseReq.Simulate).
			WithFromName(fromName).WithFromAddress(fromAddress)

		msg := vpn.NewMsgUpdateNodeDetails(fromAddress, id,
			pricesPerGB, req.NetSpeed.Upload, req.NetSpeed.Download,
			req.APIPort, req.Encryption, req.Version)
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

func updateNodeStatusHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNodeStatus

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, fromName, err := context.GetFromFields(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := sdkTypes.NewID(vars["nodeID"])
		status := strings.ToUpper(req.Status)

		cliCtx = cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly).WithSimulation(req.BaseReq.Simulate).
			WithFromName(fromName).WithFromAddress(fromAddress)

		msg := vpn.NewMsgUpdateNodeStatus(fromAddress, id, status)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

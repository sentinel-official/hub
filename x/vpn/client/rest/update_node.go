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

type msgUpdateNode struct {
	BaseReq       rest.BaseReq       `json:"base_req"`
	Moniker       string             `json:"moniker"`
	PricesPerGB   string             `json:"prices_per_gb"`
	InternetSpeed sdkTypes.Bandwidth `json:"internet_speed"`
	Encryption    string             `json:"encryption"`
	Type          string             `json:"type"`
	Version       string             `json:"version"`
}

func updateNodeInfoHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNode

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := csdkTypes.AccAddressFromBech32(req.BaseReq.From)
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
		id := sdkTypes.NewIDFromString(vars["nodeID"])
		msg := vpn.NewMsgUpdateNodeInfo(fromAddress, id, req.Type, req.Version,
			req.Moniker, pricesPerGB, req.InternetSpeed, req.Encryption)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []csdkTypes.Msg{msg})
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

		fromAddress, err := csdkTypes.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id := sdkTypes.NewIDFromString(vars["nodeID"])
		status := strings.ToUpper(req.Status)

		msg := vpn.NewMsgUpdateNodeStatus(fromAddress, id, status)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []csdkTypes.Msg{msg})
	}
}

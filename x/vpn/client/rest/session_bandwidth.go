package rest

import (
	"encoding/base64"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

type msgSignSessionBandwidth struct {
	BaseReq   utils.BaseReq      `json:"base_req"`
	Bandwidth sdkTypes.Bandwidth `json:"bandwidth"`
}

func signSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgSignSessionBandwidth
		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		signBytes, err := common.GetSessionBandwidthSignBytes(cliCtx, cdc, vars["sessionID"], req.Bandwidth)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		signature, _, err := kb.Sign(req.BaseReq.Name, req.BaseReq.Password, signBytes)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		_, _ = w.Write([]byte(base64.StdEncoding.EncodeToString(signature)))

	}
}

type msgUpdateSessionBandwidth struct {
	BaseReq       utils.BaseReq      `json:"base_req"`
	ClientSign    string             `json:"client_sign"`
	NodeOwnerSign string             `json:"node_owner_sign"`
	Bandwidth     sdkTypes.Bandwidth `json:"bandwidth"`
}

func updateSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateSessionBandwidth

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

		clientSign, err := base64.StdEncoding.DecodeString(req.ClientSign)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		nodeOwnerSign, err := base64.StdEncoding.DecodeString(req.NodeOwnerSign)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		sessionID := vpn.NewSessionID(vars["sessionID"])

		msg := vpn.NewMsgUpdateSessionBandwidth(info.GetAddress(), sessionID,
			req.Bandwidth.Upload, req.Bandwidth.Download, clientSign, nodeOwnerSign)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

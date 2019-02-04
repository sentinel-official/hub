package rest

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

type msgUpdateSessionBandwidth struct {
	BaseReq    utils.BaseReq      `json:"base_req"`
	ClientSign string             `json:"client_sign"`
	Bandwidth  sdkTypes.Bandwidth `json:"bandwidth"`
}

func updateSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateSessionBandwidth
		var sessionDetails types.SessionDetails

		vars := mux.Vars(r)
		sessionID := vars["sessionID"]

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		cliCtx.WithSimulation(req.BaseReq.Simulate).WithGenerateOnly(req.BaseReq.GenerateOnly)

		clientSignBytes, err := base64.StdEncoding.DecodeString(req.ClientSign)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		sessionBytes, err := cdc.MarshalBinaryLengthPrefixed(sessionID)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		resSession, err := cliCtx.QueryStore(sessionBytes, vpn.StoreKeySession)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(resSession) == 0 {
			err := errors.New("no session found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := cdc.UnmarshalBinaryLengthPrefixed(resSession, &sessionDetails); err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if (sessionDetails.Status != types.StatusInit) && (sessionDetails.Status != types.StatusActive) {
			err := errors.New("session is not active")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		sign, err := types.NewBandwidthSign(sessionID, req.Bandwidth, info.GetAddress(), sessionDetails.Client).GetBytes()
		signature, _, err := kb.Sign(req.BaseReq.Name, req.BaseReq.Password, sign)

		msg := vpn.NewMsgUpdateSessionBandwidth(info.GetAddress(), sessionID, req.Bandwidth.Upload, req.Bandwidth.Download, clientSignBytes, signature)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

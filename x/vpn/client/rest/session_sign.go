package rest

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

type msgSignSession struct {
	BaseReq   utils.BaseReq      `json:"base_req"`
	Bandwidth sdkTypes.Bandwidth `json:"bandwidth"`
}

func signSessionByClientHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgSignSession
		var sessionDetails types.SessionDetails
		var nodeDetails vpn.NodeDetails

		vars := mux.Vars(r)
		sessionID := vars["sessionID"]

		if err := utils.ReadRESTReq(w, r, cdc, &req); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		info, err := kb.Get(req.BaseReq.Name)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		sessionBytes, err := cdc.MarshalBinaryLengthPrefixed(sessionID)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		res, err := cliCtx.QueryStore(sessionBytes, vpn.StoreKeySession)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if len(res) == 0 {
			err := errors.New("no session found")
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := cdc.UnmarshalBinaryLengthPrefixed(res, &sessionDetails); err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		nodeIDBytes, err := cdc.MarshalBinaryLengthPrefixed(sessionDetails.NodeID)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		resNode, err := cliCtx.QueryStore(nodeIDBytes, vpn.StoreKeyNode)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(res) == 0 {
			err := errors.New("node not found")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := cdc.UnmarshalBinaryLengthPrefixed(resNode, &nodeDetails); err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if (sessionDetails.Status != types.StatusInit) && (sessionDetails.Status != types.StatusActive) {
			err := errors.New("session is not active")
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		sign, err := vpn.NewBandwidthSign(sessionID, req.Bandwidth, nodeDetails.Owner, info.GetAddress()).GetBytes()
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		signature, _, err := kb.Sign(req.BaseReq.Name, req.BaseReq.Password, sign)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Write([]byte(base64.StdEncoding.EncodeToString(signature)))

	}

}

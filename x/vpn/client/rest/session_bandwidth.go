package rest

import (
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientKeys "github.com/cosmos/cosmos-sdk/client/keys"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

type msgSignSessionBandwidth struct {
	BaseReq   rest.BaseReq       `json:"base_req"`
	Password  string             `json:"password"`
	Bandwidth sdkTypes.Bandwidth `json:"bandwidth"`
}

func signSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgSignSessionBandwidth

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["sessionID"])
		if err != nil {
			return
		}

		signBytes, err := common.GetBandwidthSignDataBytes(cliCtx, cdc, sdkTypes.NewIDFromUInt64(uint64(id)), req.Bandwidth)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		keybase, err := clientKeys.NewKeyBaseFromHomeFlag()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		fromAddress, fromName, err := context.GetFromFields(req.BaseReq.From, false)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx = cliCtx.WithGenerateOnly(true).WithSimulation(req.BaseReq.Simulate).
			WithFromName(fromName).WithFromAddress(fromAddress)

		signature, _, err := keybase.Sign(fromName, req.Password, signBytes)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		_, _ = w.Write([]byte(base64.StdEncoding.EncodeToString(signature)))
	}
}

type msgUpdateSessionBandwidthInfo struct {
	BaseReq       rest.BaseReq       `json:"base_req"`
	Consumed      sdkTypes.Bandwidth `json:"consumed"`
	NodeOwnerSign string             `json:"node_owner_sign"`
	ClientSign    string             `json:"client_sign"`
}

func updateSessionInfoHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateSessionBandwidthInfo

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

		nodeOwnerSign, err := hex.DecodeString(req.NodeOwnerSign)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientSign, err := hex.DecodeString(req.ClientSign)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["subscriptionID"])
		if err != nil {
			return
		}

		msg := vpn.NewMsgUpdateSessionInfo(fromAddress,
			sdkTypes.NewIDFromUInt64(uint64(id)), req.Consumed, nodeOwnerSign, clientSign)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []csdkTypes.Msg{msg})
	}
}

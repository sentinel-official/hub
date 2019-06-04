// nolint
package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

type msgSignSessionBandwidth struct {
	From      string             `json:"from"`
	Password  string             `json:"password"`
	Bandwidth sdkTypes.Bandwidth `json:"bandwidth"`
}

func signSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgSignSessionBandwidth
		vars := mux.Vars(r)

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		scs, err := common.QuerySessionsCountOfSubscription(cliCtx, cdc, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id := sdkTypes.NewIDFromString(vars["id"])
		data := vpn.NewBandwidthSignatureData(id, scs, req.Bandwidth).Bytes()

		kb, err := keys.NewKeyBaseFromHomeFlag()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		sigBytes, pubKey, err := kb.Sign(req.From, req.Password, data)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		stdSignature := auth.StdSignature{
			PubKey:    pubKey,
			Signature: sigBytes,
		}

		bz, err := cdc.MarshalJSON(stdSignature)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		w.Write(bz)
		return
	}
}

type msgUpdateSessionBandwidthInfo struct {
	BaseReq       rest.BaseReq       `json:"base_req"`
	Bandwidth     sdkTypes.Bandwidth `json:"bandwidth"`
	NodeOwnerSign auth.StdSignature  `json:"node_owner_sign"`
	ClientSign    auth.StdSignature  `json:"client_sign"`
}

func updateSessionInfoHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateSessionBandwidthInfo
		vars := mux.Vars(r)

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

		id := sdkTypes.NewIDFromString(vars["id"])
		msg := vpn.NewMsgUpdateSessionInfo(fromAddress, id, req.Bandwidth, req.NodeOwnerSign, req.ClientSign)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []csdkTypes.Msg{msg})
	}
}

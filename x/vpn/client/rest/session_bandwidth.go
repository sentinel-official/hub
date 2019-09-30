package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/client/common"
	"github.com/sentinel-official/hub/x/vpn/types"
)

type msgSignSessionBandwidth struct {
	From      string        `json:"from"`
	Password  string        `json:"password"`
	Bandwidth hub.Bandwidth `json:"bandwidth"`
}

func signSessionBandwidthHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgSignSessionBandwidth
		vars := mux.Vars(r)

		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}

		scs, err := common.QuerySessionsCountOfSubscription(ctx, vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := hub.NewSubscriptionIDFromString(vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		data := types.NewBandwidthSignatureData(id, scs, req.Bandwidth).Bytes()

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

		bz, err := ctx.Codec.MarshalJSON(stdSignature)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, _ = w.Write(bz)
	}
}

type msgUpdateSessionBandwidthInfo struct {
	BaseReq       rest.BaseReq      `json:"base_req"`
	Bandwidth     hub.Bandwidth     `json:"bandwidth"`
	NodeOwnerSign auth.StdSignature `json:"node_owner_sign"`
	ClientSign    auth.StdSignature `json:"client_sign"`
}

func updateSessionInfoHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateSessionBandwidthInfo
		vars := mux.Vars(r)

		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddress, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		id, err := hub.NewSubscriptionIDFromString(vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgUpdateSessionInfo(fromAddress, id, req.Bandwidth, req.NodeOwnerSign, req.ClientSign)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

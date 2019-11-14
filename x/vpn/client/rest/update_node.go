package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

type msgUpdateNode struct {
	BaseReq       rest.BaseReq  `json:"base_req"`
	Moniker       string        `json:"moniker"`
	PricesPerGB   string        `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth `json:"internet_speed"`
	Encryption    string        `json:"encryption"`
	Type          string        `json:"type"`
	Version       string        `json:"version"`
}

func updateNodeInfoHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateNode

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

		pricesPerGB, err := sdk.ParseCoins(req.PricesPerGB)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		id, err := hub.NewNodeIDFromString(vars["id"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgUpdateNodeInfo(fromAddress, id, req.Type, req.Version,
			req.Moniker, pricesPerGB, req.InternetSpeed, req.Encryption)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

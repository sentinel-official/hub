package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/x/vpn"
)

type msgRegisterNode struct {
	BaseReq       rest.BaseReq  `json:"base_req"`
	Type          string        `json:"type"`
	Version       string        `json:"version"`
	Moniker       string        `json:"moniker"`
	PricesPerGB   string        `json:"prices_per_gb"`
	InternetSpeed hub.Bandwidth `json:"internet_speed"`
	Encryption    string        `json:"encryption"`
}

func registerNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRegisterNode

		if !rest.ReadRESTReq(w, r, cdc, &req) {
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

		msg := vpn.NewMsgRegisterNode(fromAddress, req.Type, req.Version,
			req.Moniker, pricesPerGB, req.InternetSpeed, req.Encryption)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.WriteGenerateStdTxResponse(w, cdc, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

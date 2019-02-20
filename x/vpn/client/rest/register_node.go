package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	clientKeys "github.com/cosmos/cosmos-sdk/client/keys"
	clientRest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgRegisterNode struct {
	BaseReq      rest.BaseReq       `json:"base_req"`
	AmountToLock string             `json:"amount_to_lock"`
	PricesPerGB  string             `json:"prices_per_gb"`
	NetSpeed     sdkTypes.Bandwidth `json:"net_speed"`
	APIPort      uint32             `json:"api_port"`
	EncMethod    string             `json:"enc_method"`
	Version      string             `json:"version"`
	NodeType     string             `json:"node_type"`
}

func registerNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRegisterNode

		if !rest.ReadRESTReq(w, r, cdc, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		cliCtx.WithGenerateOnly(req.BaseReq.GenerateOnly).WithSimulation(req.BaseReq.Simulate)

		keybase, err := clientKeys.NewKeyBaseFromHomeFlag()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		info, err := keybase.Get(req.BaseReq.From)
		if err != nil {

			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		amountToLock, err := csdkTypes.ParseCoin(req.AmountToLock)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		pricesPerGB, err := csdkTypes.ParseCoins(req.PricesPerGB)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		apiPort := vpn.NewAPIPort(req.APIPort)

		msg := vpn.NewMsgRegisterNode(info.GetAddress(),
			amountToLock, pricesPerGB, req.NetSpeed.Upload, req.NetSpeed.Download,
			apiPort, req.EncMethod, req.NodeType, req.Version)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		clientRest.CompleteAndBroadcastTxREST(w, cliCtx, req.BaseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

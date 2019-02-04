package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type msgRegisterNode struct {
	BaseReq      utils.BaseReq      `json:"base_req"`
	AmountToLock string             `json:"amount_to_lock"`
	APIPort      uint16             `json:"api_port"`
	NetSpeed     sdkTypes.Bandwidth `json:"net_speed"`
	EncMethod    string             `json:"enc_method"`
	PricesPerGB  string             `json:"prices_per_gb"`
	Version      string             `json:"version"`
	NodeType     string             `json:"node_type"`
}

func registerNodeHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec, kb keys.Keybase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRegisterNode

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

		amountToLock, err := csdkTypes.ParseCoin(req.AmountToLock)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		pricesPerGB, err := csdkTypes.ParseCoins(req.PricesPerGB)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := vpn.NewMsgRegisterNode(info.GetAddress(),
			req.APIPort, req.NetSpeed.Upload, req.NetSpeed.Download,
			req.EncMethod, pricesPerGB, req.Version, req.NodeType, amountToLock)
		if err := msg.ValidateBasic(); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.CompleteAndBroadcastTxREST(w, r, cliCtx, baseReq, []csdkTypes.Msg{msg}, cdc)
	}
}

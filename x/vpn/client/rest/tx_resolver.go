package rest

import (
	"fmt"
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

type msgRegisterResolver struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Commission string       `json:"commission"`
}

func registerResolverHandleFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgRegisterResolver
		
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		
		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		commission, err := sdk.NewDecFromStr(req.Commission)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		if commission.LT(sdk.ZeroDec()) || commission.GT(sdk.OneDec()) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("commission %s: between 0 and 1", commission.String()))
			return
		}
		
		msg := types.NewMsgRegisterResolver(addr, commission)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type msgUpdateResolver struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Commission string       `json:"commission"`
	ResolverID string       `json:"resolver_id"`
}

func updateResolverHandleFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgUpdateResolver
		
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		
		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		resolverID, err := hub.NewResolverIDFromString(req.ResolverID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		commission, err := sdk.NewDecFromStr(req.Commission)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		if commission.LT(sdk.ZeroDec()) || commission.GT(sdk.OneDec()) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("commission %s: between 0 and 1", commission.String()))
			return
		}
		
		msg := types.NewMsgUpdateResolverInfo(addr, resolverID, commission)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type msgDeregisterResolver struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	ResolverID string       `json:"resolver_id"`
}

func deregisterResolverHandleFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req msgDeregisterResolver
		
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		
		addr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		resolverID, err := hub.NewResolverIDFromString(req.ResolverID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		msg := types.NewMsgDeregisterResolver(addr, resolverID)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

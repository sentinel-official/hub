package rest

import (
	"net/http"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	
	"github.com/sentinel-official/hub/x/deposit/client/common"
)

func getDepositOfAddressHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		
		deposit, err := common.QueryDepositOfAddress(ctx, vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, deposit)
	}
}

func getAllDeposits(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deposits, err := common.QueryAllDeposits(ctx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		
		rest.PostProcessResponse(w, ctx, deposits)
	}
}

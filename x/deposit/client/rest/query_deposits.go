package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/client/common"
)

func getDepositsOfAddressHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		address, err := csdkTypes.AccAddressFromBech32(vars["address"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := common.QueryDepositsOfAddress(cliCtx, cdc, address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no deposit found")
			return
		}

		var _deposit deposit.Deposit
		if err := cdc.UnmarshalJSON(res, &_deposit); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, _deposit, cliCtx.Indent)
	}
}

// nolint:dupl
func getAllDeposits(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := common.QueryAllDeposits(cliCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if string(res) == "[]" || string(res) == "null" {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "no deposits found")
			return
		}

		var deposits []deposit.Deposit
		if err := cdc.UnmarshalJSON(res, &deposits); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, deposits, cliCtx.Indent)
	}
}

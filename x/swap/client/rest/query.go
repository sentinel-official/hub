package rest

import (
	"context"
	"encoding/hex"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/hub/x/swap/types"
)

func querySwap(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		txHash, err := hex.DecodeString(vars["txHash"])
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		res, err := qc.QuerySwap(context.Background(),
			types.NewQuerySwapRequest(types.BytesToHash(txHash)))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func querySwaps(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		res, err := qc.QuerySwaps(context.Background(),
			types.NewQuerySwapsRequest(nil))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

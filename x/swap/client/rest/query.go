package rest

import (
	"context"
	"encoding/hex"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/hub/x/swap/types"
)

func querySwap(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		txHash, err := hex.DecodeString(vars["txHash"])
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QuerySwap(
			context.Background(),
			types.NewQuerySwapRequest(
				types.BytesToHash(txHash),
			),
		)
		if rest.CheckInternalServerError(w, err) {
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

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QuerySwaps(
			context.Background(),
			types.NewQuerySwapsRequest(
				&sdkquery.PageRequest{
					Offset: uint64(page * limit),
					Limit:  uint64(limit),
				},
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

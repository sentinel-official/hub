package rest

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func queryDeposit(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		address, err := sdk.AccAddressFromBech32(vars["address"])
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryDeposit(
			context.Background(),
			types.NewQueryDepositRequest(
				address,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryDeposits(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryDeposits(
			context.Background(),
			types.NewQueryDepositsRequest(
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

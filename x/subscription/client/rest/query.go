package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func querySubscription(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QuerySubscription(
			context.Background(),
			types.NewQuerySubscriptionRequest(
				id,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func querySubscriptions(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc     = types.NewQueryServiceClient(ctx)
			query  = r.URL.Query()
			status = hubtypes.StatusFromString(query.Get("status"))
		)

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		if query.Get("address") != "" {
			address, err := sdk.AccAddressFromBech32(query.Get("address"))
			if rest.CheckBadRequestError(w, err) {
				return
			}

			res, err := qc.QuerySubscriptionsForAddress(
				context.Background(),
				types.NewQuerySubscriptionsForAddressRequest(
					address,
					status,
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
			return
		}

		res, err := qc.QuerySubscriptions(
			context.Background(),
			types.NewQuerySubscriptionsRequest(
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

func queryQuota(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		address, err := sdk.AccAddressFromBech32(vars["address"])
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryQuota(
			context.Background(),
			types.NewQueryQuotaRequest(
				id,
				address,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryQuotas(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryQuotas(
			context.Background(),
			types.NewQueryQuotasRequest(
				id,
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

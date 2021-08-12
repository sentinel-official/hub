package rest

import (
	"context"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkquery "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

func queryProvider(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc   = types.NewQueryServiceClient(ctx)
			vars = mux.Vars(r)
		)

		address, err := hubtypes.ProvAddressFromBech32(vars["address"])
		if rest.CheckBadRequestError(w, err) {
			return
		}

		res, err := qc.QueryProvider(
			context.Background(),
			types.NewQueryProviderRequest(
				address,
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, res)
	}
}

func queryProviders(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			qc = types.NewQueryServiceClient(ctx)
		)

		_, page, limit, err := rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		providers, err := qc.QueryProviders(
			context.Background(),
			types.NewQueryProvidersRequest(
				&sdkquery.PageRequest{
					Offset: uint64(page * limit),
					Limit:  uint64(limit),
				},
			),
		)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponse(w, ctx, providers)
	}
}

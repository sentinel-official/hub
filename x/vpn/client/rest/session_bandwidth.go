// nolint
package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type msgSignSessionBandwidth struct {
	BaseReq   rest.BaseReq       `json:"base_req"`
	Password  string             `json:"password"`
	Bandwidth sdkTypes.Bandwidth `json:"bandwidth"`
}

func signSessionBandwidthHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

type msgUpdateSessionBandwidthInfo struct {
	BaseReq    rest.BaseReq       `json:"base_req"`
	ClientSign string             `json:"client_sign"`
	Bandwidth  sdkTypes.Bandwidth `json:"bandwidth"`
}

func updateSessionInfoHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

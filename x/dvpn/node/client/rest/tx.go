package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func txRegisterNodeHandlerFunc(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

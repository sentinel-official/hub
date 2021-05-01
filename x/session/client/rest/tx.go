package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

func txUpsert(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

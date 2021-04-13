package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

func txAdd(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txSetStatus(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txAddNode(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txRemoveNode(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

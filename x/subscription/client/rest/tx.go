package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
)

func txSubscribeToNode(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txSubscribeToPlan(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txCancel(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txAddQuota(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func txUpdateQuota(_ client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

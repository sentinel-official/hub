package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	// HandlerFunc defines the Hotfix handler function
	HandlerFunc func(ctx sdk.Context) error

	// Hotfix defines the Hotfix type
	Hotfix struct {
		Name    string      `json:"name"`
		Height  int64       `json:"height"`
		Handler HandlerFunc `json:"handler"`
	}
)

// NewHotfix initializes Hotfix and returns the reference
func NewHotfix() *Hotfix {
	return &Hotfix{}
}

// WithName sets the name for Hotfix
func (h *Hotfix) WithName(v string) *Hotfix {
	h.Name = v
	return h
}

// WithHeight sets the height for Hotfix
func (h *Hotfix) WithHeight(v int64) *Hotfix {
	h.Height = v
	return h
}

// WithHandler sets the handler for Hotfix
func (h *Hotfix) WithHandler(v HandlerFunc) *Hotfix {
	h.Handler = v
	return h
}

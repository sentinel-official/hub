package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	HandlerFunc func(ctx sdk.Context) error

	Hotfix struct {
		Name    string      `json:"name"`
		Height  int64       `json:"height"`
		Handler HandlerFunc `json:"handler"`
	}
)

func NewHotfix() *Hotfix {
	return &Hotfix{}
}

func (h *Hotfix) WithName(v string) *Hotfix {
	h.Name = v
	return h
}

func (h *Hotfix) WithHeight(v int64) *Hotfix {
	h.Height = v
	return h
}

func (h *Hotfix) WithHandler(v HandlerFunc) *Hotfix {
	h.Handler = v
	return h
}

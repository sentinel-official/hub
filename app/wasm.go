package app

import (
	"strings"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

const (
	DefaultWasmProposals = "*"
)

func GetWasmEnabledProposals(v string) []wasmtypes.ProposalType {
	if v == "" {
		return wasmtypes.DisableAllProposals
	}
	if v == "*" {
		return wasmtypes.EnableAllProposals
	}

	proposals, err := wasmtypes.ConvertToProposals(strings.Split(v, ","))
	if err != nil {
		panic(err)
	}

	return proposals
}

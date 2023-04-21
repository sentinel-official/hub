package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	ErrorInvalidMessage = errors.Register(ModuleName, 101, "invalid message")
)

var (
	ErrorProviderNotFound = errors.Register(ModuleName, 201, "provider not found")
	ErrorNodeNotFound     = errors.Register(ModuleName, 202, "node not found")
	ErrorPlanNotFound     = errors.Register(ModuleName, 203, "plan not found")
	ErrorUnauthorized     = errors.Register(ModuleName, 204, "unauthorized")
	ErrorDuplicateNode    = errors.Register(ModuleName, 205, "duplicate node")
)

func NewErrorProviderNotFound(addr hubtypes.ProvAddress) error {
	return errors.Wrapf(ErrorProviderNotFound, "provider %s does not exist", addr)
}

func NewErrorNodeNotFound(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorNodeNotFound, "node %s does not exist", addr)
}

func NewErrorPlanNotFound(id uint64) error {
	return errors.Wrapf(ErrorPlanNotFound, "plan %d does not exist", id)
}

func NewErrorDuplicateNode(addr hubtypes.NodeAddress) error {
	return errors.Wrapf(ErrorDuplicateNode, "node %d already exists", addr)
}

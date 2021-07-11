package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorInvalidField       = errors.Register(ModuleName, 101, "invalid field")
	ErrorInvalidFrom        = errors.Register(ModuleName, 102, "invalid from")
	ErrorInvalidName        = errors.Register(ModuleName, 103, "invalid name")
	ErrorInvalidIdentity    = errors.Register(ModuleName, 104, "invalid identity")
	ErrorInvalidWebsite     = errors.Register(ModuleName, 105, "invalid website")
	ErrorInvalidDescription = errors.Register(ModuleName, 106, "invalid description")
)

var (
	ErrorDuplicateProvider    = errors.Register(ModuleName, 201, "duplicate provider")
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 202, "provider does not exist")
)

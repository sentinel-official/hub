package v06

import (
	"github.com/sentinel-official/hub/x/subscription/types"
	legacy "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.5"
)

func MigrateParams(params legacy.Params) types.Params {
	return types.Params{
		InactiveDuration: params.InactiveDuration,
	}
}

package v06

import (
	"github.com/sentinel-official/hub/x/node/types"
	legacy "github.com/sentinel-official/hub/x/node/types/legacy/v0.5"
)

func MigrateParams(params legacy.Params) types.Params {
	return types.NewParams(
		params.Deposit,
		params.InactiveDuration,
	)
}

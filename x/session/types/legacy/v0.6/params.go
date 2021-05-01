package v0_6

import (
	"github.com/sentinel-official/hub/x/session/types"
	legacy "github.com/sentinel-official/hub/x/session/types/legacy/v0.5"
)

func MigrateParams(params legacy.Params) types.Params {
	return types.NewParams(
		params.InactiveDuration,
		params.ProofVerificationEnabled,
	)
}

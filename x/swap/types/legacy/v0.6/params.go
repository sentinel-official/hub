package v06

import (
	"github.com/sentinel-official/hub/x/swap/types"
	legacy "github.com/sentinel-official/hub/x/swap/types/legacy/v0.5"
)

func MigrateParams(params legacy.Params) types.Params {
	return types.NewParams(
		params.SwapEnabled,
		params.SwapDenom,
		params.ApproveBy.String(),
	)
}

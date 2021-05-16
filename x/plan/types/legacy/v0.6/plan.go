package v06

import (
	hubtypes "github.com/sentinel-official/hub/types/legacy/v0.6"
	"github.com/sentinel-official/hub/x/plan/types"
	legacy "github.com/sentinel-official/hub/x/plan/types/legacy/v0.5"
)

func MigratePlan(item legacy.Plan) types.Plan {
	return types.Plan{
		Id:       item.ID,
		Provider: item.Provider.String(),
		Price:    item.Price,
		Validity: item.Validity,
		Bytes:    item.Bytes,
		Status:   hubtypes.MigrateStatus(item.Status),
		StatusAt: item.StatusAt,
	}
}

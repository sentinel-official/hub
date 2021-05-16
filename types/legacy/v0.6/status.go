package v06

import (
	"github.com/sentinel-official/hub/types"
	v05 "github.com/sentinel-official/hub/types/legacy/v0.5"
)

func MigrateStatus(v v05.Status) types.Status {
	return types.Status(v)
}

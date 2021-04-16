package v0_6

import (
	"github.com/sentinel-official/hub/types"
	v05 "github.com/sentinel-official/hub/types/legacy/v0.5"
)

func MigrateBandwidth(v v05.Bandwidth) types.Bandwidth {
	return types.Bandwidth{
		Upload:   v.Upload,
		Download: v.Download,
	}
}

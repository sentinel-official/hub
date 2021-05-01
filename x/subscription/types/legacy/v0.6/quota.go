package v0_6

import (
	"github.com/sentinel-official/hub/x/subscription/types"
	legacy "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.5"
)

func MigrateQuota(item legacy.Quota) types.Quota {
	return types.Quota{
		Address:   item.Address.String(),
		Allocated: item.Allocated,
		Consumed:  item.Consumed,
	}
}

func MigrateQuotas(items legacy.Quotas) types.Quotas {
	var quotas types.Quotas
	for _, item := range items {
		quotas = append(quotas, MigrateQuota(item))
	}

	return quotas
}

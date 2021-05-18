package v06

import (
	"github.com/sentinel-official/hub/x/provider/types"
	legacy "github.com/sentinel-official/hub/x/provider/types/legacy/v0.5"
)

func MigrateProvider(item legacy.Provider) types.Provider {
	return types.Provider{
		Address:     item.Address.String(),
		Name:        item.Name,
		Identity:    item.Identity,
		Website:     item.Website,
		Description: item.Description,
	}
}

func MigrateProviders(items legacy.Providers) types.Providers {
	var providers types.Providers
	for _, item := range items {
		providers = append(providers, MigrateProvider(item))
	}

	return providers
}

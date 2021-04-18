package v0_6

import (
	"github.com/sentinel-official/hub/x/deposit/types"
	legacy "github.com/sentinel-official/hub/x/deposit/types/legacy/v0.5"
)

func MigrateDeposit(item legacy.Deposit) types.Deposit {
	return types.Deposit{
		Address: item.Address.String(),
		Coins:   item.Coins,
	}
}

func MigrateDeposits(items legacy.Deposits) types.Deposits {
	var deposits types.Deposits
	for _, item := range items {
		deposits = append(deposits, MigrateDeposit(item))
	}

	return deposits
}

package deposit

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

func InitGenesis(ctx csdkTypes.Context, k Keeper, data types.GenesisState) {
	for _, deposit := range data {
		k.SetDeposit(ctx, deposit)
	}
}

func ExportGenesis(ctx csdkTypes.Context, k Keeper) types.GenesisState {
	deposits := k.GetAllDeposits(ctx)

	return types.NewGenesisState(deposits)
}

func ValidateGenesis(data types.GenesisState) error {
	addressMap := make(map[string]bool, len(data))
	for _, deposit := range data {
		if deposit.Address == nil {
			return fmt.Errorf("address value is nil for deposit %s", deposit)
		}

		addressStr := deposit.Address.String()
		if addressMap[addressStr] {
			return fmt.Errorf("duplicate address for deposit %s", deposit)
		}

		addressMap[addressStr] = true
		if deposit.Coins == nil {
			return fmt.Errorf("coins value is nil for deposit %s", deposit)
		}

		if !deposit.Coins.IsValid() {
			return fmt.Errorf("invalid coins for deposit %s", deposit)
		}
	}

	return nil
}

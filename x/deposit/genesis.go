package deposit

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	"github.com/sentinel-official/hub/x/deposit/types"
)

func InitGenesis(ctx sdk.Context, k Keeper, data types.GenesisState) {
	for _, deposit := range data {
		k.SetDeposit(ctx, deposit)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) types.GenesisState {
	deposits := k.GetAllDeposits(ctx)
	
	return types.NewGenesisState(deposits)
}

func ValidateGenesis(data types.GenesisState) error {
	addressMap := make(map[string]bool, len(data))
	for _, deposit := range data {
		if err := deposit.IsValid(); err != nil {
			return fmt.Errorf("%s for the %s", err.Error(), deposit)
		}
		
		addressStr := deposit.Address.String()
		if addressMap[addressStr] {
			return fmt.Errorf("duplicate address for the %s", deposit)
		}
		
		addressMap[addressStr] = true
	}
	
	return nil
}

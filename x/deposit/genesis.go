package deposit

import (
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
	return nil
}

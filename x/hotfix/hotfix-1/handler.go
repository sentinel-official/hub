package hotfix1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/sentinel-official/hub/x/hotfix/types"
)

func Handler(ak authkeeper.AccountKeeper) types.HandlerFunc {
	return func(ctx sdk.Context) error {
		address, err := sdk.AccAddressFromBech32("sent1vv8kmwrs24j5emzw8dp7k8satgea62l7knegd7")
		if err != nil {
			return err
		}

		account := ak.GetAccount(ctx, address)
		if account == nil {
			return nil
		}

		if err := account.SetPubKey(nil); err != nil {
			return err
		}

		ak.SetAccount(ctx, account)
		return nil
	}
}

package upgrade1

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Handler returns UpgradeHandler function
// Sets the public key of multi-signature accounts to nil value
func Handler(ak authkeeper.AccountKeeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan) {
		ak.IterateAccounts(ctx, func(account authtypes.AccountI) (stop bool) {
			pubKey := account.GetPubKey()
			if pubKey == nil {
				return false
			}

			if _, ok := pubKey.(*multisig.LegacyAminoPubKey); !ok {
				return false
			}

			account.SetPubKey(nil)
			ak.SetAccount(ctx, account)

			return false
		})
	}
}

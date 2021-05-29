package upgrade

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/upgrade/expected"
	"github.com/sentinel-official/hub/x/upgrade/types"
	upgrade1 "github.com/sentinel-official/hub/x/upgrade/upgrade-1"
)

func BeginBlock(ctx sdk.Context, ak expected.AccountKeeper, bk expected.BankKeeper, sk stakingkeeper.Keeper) []abcitypes.ValidatorUpdate {
	switch ctx.BlockHeight() {
	case upgrade1.Height:
		cacheCtx, writeCache := ctx.CacheContext()
		if err := upgrade1.Handler(cacheCtx, ak, bk, sk); err == nil {
			writeCache()
		} else {
			ctx.Logger().With("module", types.ModuleName).
				Error(fmt.Sprintf("failed to apply upgrade %s", upgrade1.Name), "cause", err)
		}
	}

	return nil
}

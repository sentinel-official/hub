package hotfix

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/hotfix/expected"
	hotfix1 "github.com/sentinel-official/hub/x/hotfix/hotfix-1"
	"github.com/sentinel-official/hub/x/hotfix/types"
)

func BeginBlock(ctx sdk.Context, ak expected.AccountKeeper, bk expected.BankKeeper, sk stakingkeeper.Keeper) []abcitypes.ValidatorUpdate {
	switch ctx.BlockHeight() {
	case hotfix1.Height:
		cacheCtx, writeCache := ctx.CacheContext()
		if err := hotfix1.Handler(cacheCtx, ak, bk, sk); err != nil {
			ctx.Logger().With("module", types.ModuleName).
				Error(fmt.Sprintf("failed to apply hotfix %s", hotfix1.Name), "cause", err)
		} else {
			writeCache()
		}
	}

	return nil
}

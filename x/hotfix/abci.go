package hotfix

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/hotfix/types"
)

func BeginBlock(ctx sdk.Context, registry *types.Registry) []abcitypes.ValidatorUpdate {
	var (
		height = ctx.BlockHeight()
		hotfix = registry.Hotfix(height)
	)

	if hotfix != nil {
		cacheCtx, writeCache := ctx.CacheContext()
		if err := hotfix.Handler(cacheCtx); err != nil {
			ctx.Logger().With("module", types.ModuleName).
				Error(fmt.Sprintf("failed to apply hotfix %d", hotfix.Name), "cause", err)
		} else {
			writeCache()
		}
	}

	return nil
}

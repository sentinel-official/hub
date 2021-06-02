package hotfix

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/hotfix/types"
)

func BeginBlock(ctx sdk.Context, registry *types.Registry) []abcitypes.ValidatorUpdate {
	var (
		height = ctx.BlockHeight()
		hotfix = registry.Hotfix(height)
		logger = ctx.Logger().With("module", types.ModuleName)
	)

	if hotfix != nil {
		cacheCtx, writeCache := ctx.CacheContext()
		if err := hotfix.Handler(cacheCtx); err != nil {
			logger.Error("failed to apply the hotfix", "name", hotfix.Name, "cause", err)
		} else {
			writeCache()
			logger.Info("successfully applied the hotfix", "name", hotfix.Name)
		}
	}

	return nil
}

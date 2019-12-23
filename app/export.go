package app

import (
	"encoding/json"
	"log"
	
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	tm "github.com/tendermint/tendermint/types"
)

func (app *HubApp) ExportAppStateAndValidators(forZeroHeight bool,
	jailWhiteList []string) (json.RawMessage, []tm.GenesisValidator, error) {
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})
	
	if forZeroHeight {
		app.prepForZeroHeightGenesis(ctx, jailWhiteList)
	}
	
	state := app.mm.ExportGenesis(ctx)
	appState, err := codec.MarshalJSONIndent(app.cdc, state)
	if err != nil {
		return nil, nil, err
	}
	
	validators := staking.WriteValidators(ctx, app.stakingKeeper)
	return appState, validators, nil
}

// nolint:funlen
func (app *HubApp) prepForZeroHeightGenesis(
	ctx sdk.Context, jailWhiteList []string) {
	applyWhiteList := false
	if len(jailWhiteList) > 0 {
		applyWhiteList = true
	}
	
	whiteListMap := make(map[string]bool)
	for _, addr := range jailWhiteList {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}
		
		whiteListMap[addr] = true
	}
	
	app.crisisKeeper.AssertInvariants(ctx)
	
	app.stakingKeeper.IterateValidators(ctx, func(_ int64, val staking.ValidatorI) (stop bool) {
		_, _ = app.distributionKeeper.WithdrawValidatorCommission(ctx, val.GetOperator())
		return false
	})
	
	delegations := app.stakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range delegations {
		_, _ = app.distributionKeeper.WithdrawDelegationRewards(ctx,
			delegation.DelegatorAddress, delegation.ValidatorAddress)
	}
	
	app.distributionKeeper.DeleteAllValidatorSlashEvents(ctx)
	app.distributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)
	
	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)
	
	app.stakingKeeper.IterateValidators(ctx, func(_ int64, val staking.ValidatorI) (stop bool) {
		scraps := app.distributionKeeper.GetValidatorOutstandingRewards(ctx, val.GetOperator())
		feePool := app.distributionKeeper.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps)
		app.distributionKeeper.SetFeePool(ctx, feePool)
		
		app.distributionKeeper.Hooks().AfterValidatorCreated(ctx, val.GetOperator())
		return false
	})
	
	for _, del := range delegations {
		app.distributionKeeper.Hooks().BeforeDelegationCreated(ctx, del.DelegatorAddress, del.ValidatorAddress)
		app.distributionKeeper.Hooks().AfterDelegationModified(ctx, del.DelegatorAddress, del.ValidatorAddress)
	}
	
	ctx = ctx.WithBlockHeight(height)
	
	app.stakingKeeper.IterateRedelegations(ctx, func(_ int64, red staking.Redelegation) (stop bool) {
		for i := range red.Entries {
			red.Entries[i].CreationHeight = 0
		}
		app.stakingKeeper.SetRedelegation(ctx, red)
		return false
	})
	
	app.stakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, ubd staking.UnbondingDelegation) (stop bool) {
		for i := range ubd.Entries {
			ubd.Entries[i].CreationHeight = 0
		}
		app.stakingKeeper.SetUnbondingDelegation(ctx, ubd)
		return false
	})
	
	store := ctx.KVStore(app.keys[staking.StoreKey])
	iter := sdk.KVStoreReversePrefixIterator(store, staking.ValidatorsKey)
	defer iter.Close()
	
	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(iter.Key()[1:])
		validator, found := app.stakingKeeper.GetValidator(ctx, addr)
		if !found {
			panic("expected validator, not found")
		}
		
		validator.UnbondingHeight = 0
		if applyWhiteList && !whiteListMap[addr.String()] {
			validator.Jailed = true
		}
		
		app.stakingKeeper.SetValidator(ctx, validator)
	}
	
	_ = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	
	app.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashing.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.slashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)
}

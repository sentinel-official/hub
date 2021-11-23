package hub

import (
	"encoding/json"
	"log"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func (a *App) ExportAppStateAndValidators(
	forZeroHeight bool, jailAllowedAddrs []string,
) (servertypes.ExportedApp, error) {
	ctx := a.NewContext(true, tmproto.Header{Height: a.LastBlockHeight()})

	height := a.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		a.prepForZeroHeightGenesis(ctx, jailAllowedAddrs)
	}

	genState := a.manager.ExportGenesis(ctx, a.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators, err := staking.WriteValidators(ctx, a.stakingKeeper)
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: a.BaseApp.GetConsensusParams(ctx),
	}, err
}

func (a *App) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
	applyAllowedAddrs := false

	if len(jailAllowedAddrs) > 0 {
		applyAllowedAddrs = true
	}

	allowedAddrsMap := make(map[string]bool)

	for _, addr := range jailAllowedAddrs {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}
		allowedAddrsMap[addr] = true
	}

	a.crisisKeeper.AssertInvariants(ctx)

	a.stakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		_, _ = a.distributionKeeper.WithdrawValidatorCommission(ctx, val.GetOperator())
		return false
	})

	dels := a.stakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range dels {
		valAddr, err := sdk.ValAddressFromBech32(delegation.ValidatorAddress)
		if err != nil {
			panic(err)
		}

		delAddr, err := sdk.AccAddressFromBech32(delegation.DelegatorAddress)
		if err != nil {
			panic(err)
		}
		_, _ = a.distributionKeeper.WithdrawDelegationRewards(ctx, delAddr, valAddr)
	}

	a.distributionKeeper.DeleteAllValidatorSlashEvents(ctx)

	a.distributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	a.stakingKeeper.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
		scraps := a.distributionKeeper.GetValidatorOutstandingRewardsCoins(ctx, val.GetOperator())
		feePool := a.distributionKeeper.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps...)
		a.distributionKeeper.SetFeePool(ctx, feePool)

		a.distributionKeeper.Hooks().AfterValidatorCreated(ctx, val.GetOperator())
		return false
	})

	for _, del := range dels {
		valAddr, err := sdk.ValAddressFromBech32(del.ValidatorAddress)
		if err != nil {
			panic(err)
		}
		delAddr, err := sdk.AccAddressFromBech32(del.DelegatorAddress)
		if err != nil {
			panic(err)
		}
		a.distributionKeeper.Hooks().BeforeDelegationCreated(ctx, delAddr, valAddr)
		a.distributionKeeper.Hooks().AfterDelegationModified(ctx, delAddr, valAddr)
	}

	ctx = ctx.WithBlockHeight(height)

	a.stakingKeeper.IterateRedelegations(ctx, func(_ int64, red stakingtypes.Redelegation) (stop bool) {
		for i := range red.Entries {
			red.Entries[i].CreationHeight = 0
		}
		a.stakingKeeper.SetRedelegation(ctx, red)
		return false
	})

	a.stakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, ubd stakingtypes.UnbondingDelegation) (stop bool) {
		for i := range ubd.Entries {
			ubd.Entries[i].CreationHeight = 0
		}
		a.stakingKeeper.SetUnbondingDelegation(ctx, ubd)
		return false
	})

	store := ctx.KVStore(a.keys[stakingtypes.StoreKey])
	iter := sdk.KVStoreReversePrefixIterator(store, stakingtypes.ValidatorsKey)
	counter := int16(0)

	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(iter.Key()[1:])
		validator, found := a.stakingKeeper.GetValidator(ctx, addr)
		if !found {
			panic("expected validator, not found")
		}

		validator.UnbondingHeight = 0
		if applyAllowedAddrs && !allowedAddrsMap[addr.String()] {
			validator.Jailed = true
		}

		a.stakingKeeper.SetValidator(ctx, validator)
		counter++
	}

	iter.Close()

	_, err := a.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		log.Fatal(err)
	}

	a.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, info slashingtypes.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			a.slashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)
}

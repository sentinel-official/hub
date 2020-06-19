package hub

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

func (a *App) ExportAppStateAndValidators(zeroHeight bool,
	jailWhitelist []string) (json.RawMessage, []tm.GenesisValidator, error) {
	ctx := a.NewContext(true, abci.Header{Height: a.LastBlockHeight()})

	if zeroHeight {
		a.prepForZeroHeightGenesis(ctx, jailWhitelist)
	}

	state, err := codec.MarshalJSONIndent(a.cdc, a.manager.ExportGenesis(ctx))
	if err != nil {
		return nil, nil, err
	}

	return state, staking.WriteValidators(ctx, a.stakingKeeper), nil
}

func (a *App) prepForZeroHeightGenesis(ctx sdk.Context, jailWhitelist []string) {
	jail := false
	if len(jailWhitelist) > 0 {
		jail = true
	}

	jailed := make(map[string]bool)
	for _, address := range jailWhitelist {
		_, err := sdk.ValAddressFromBech32(address)
		if err != nil {
			log.Fatal(err)
		}

		jailed[address] = true
	}

	a.crisisKeeper.AssertInvariants(ctx)

	a.stakingKeeper.IterateValidators(ctx, func(_ int64, item staking.ValidatorI) (stop bool) {
		_, _ = a.distributionKeeper.WithdrawValidatorCommission(ctx, item.GetOperator())
		return false
	})

	delegations := a.stakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range delegations {
		_, _ = a.distributionKeeper.WithdrawDelegationRewards(ctx,
			delegation.DelegatorAddress, delegation.ValidatorAddress)
	}

	a.distributionKeeper.DeleteAllValidatorSlashEvents(ctx)
	a.distributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	a.stakingKeeper.IterateValidators(ctx, func(_ int64, item staking.ValidatorI) (stop bool) {
		scraps := a.distributionKeeper.GetValidatorOutstandingRewards(ctx, item.GetOperator())
		feePool := a.distributionKeeper.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(scraps)
		a.distributionKeeper.SetFeePool(ctx, feePool)

		a.distributionKeeper.Hooks().AfterValidatorCreated(ctx, item.GetOperator())
		return false
	})

	for _, delegation := range delegations {
		a.distributionKeeper.Hooks().
			BeforeDelegationCreated(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
		a.distributionKeeper.Hooks().
			AfterDelegationModified(ctx, delegation.DelegatorAddress, delegation.ValidatorAddress)
	}

	ctx = ctx.WithBlockHeight(height)

	a.stakingKeeper.IterateRedelegations(ctx, func(_ int64, item staking.Redelegation) (stop bool) {
		for i := range item.Entries {
			item.Entries[i].CreationHeight = 0
		}

		a.stakingKeeper.SetRedelegation(ctx, item)
		return false
	})

	a.stakingKeeper.IterateUnbondingDelegations(ctx, func(_ int64, item staking.UnbondingDelegation) (stop bool) {
		for i := range item.Entries {
			item.Entries[i].CreationHeight = 0
		}

		a.stakingKeeper.SetUnbondingDelegation(ctx, item)
		return false
	})

	store := ctx.KVStore(a.keys[staking.StoreKey])
	iterator := sdk.KVStoreReversePrefixIterator(store, staking.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		address := sdk.ValAddress(iterator.Key()[1:])
		validator, found := a.stakingKeeper.GetValidator(ctx, address)
		if !found {
			panic("expected validator not found")
		}

		validator.UnbondingHeight = 0
		if jail && !jailed[address.String()] {
			validator.Jailed = true
		}

		a.stakingKeeper.SetValidator(ctx, validator)
	}

	_ = a.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	a.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr sdk.ConsAddress, item slashing.ValidatorSigningInfo) (stop bool) {
			item.StartHeight = 0
			a.slashingKeeper.SetValidatorSigningInfo(ctx, addr, item)
			return false
		},
	)
}

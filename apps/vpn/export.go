package vpn

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmTypes "github.com/tendermint/tendermint/types"
)

func (app *VPN) ExportAppStateAndValidators(forZeroHeight bool) (
	appState json.RawMessage, validators []tmTypes.GenesisValidator, err error) {

	ctx := app.NewContext(true, abciTypes.Header{Height: app.LastBlockHeight()})

	if forZeroHeight {
		app.prepForZeroHeightGenesis(ctx)
	}

	accounts := []GenesisAccount{}
	appendAccount := func(acc auth.Account) (stop bool) {
		account := NewGenesisAccountI(acc)
		accounts = append(accounts, account)
		return false
	}
	app.accountKeeper.IterateAccounts(ctx, appendAccount)

	genState := NewGenesisState(accounts,
		auth.ExportGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper),
		staking.ExportGenesis(ctx, app.stakingKeeper),
		slashing.ExportGenesis(ctx, app.slashingKeeper),
		distribution.ExportGenesis(ctx, app.distributionKeeper),
		gov.ExportGenesis(ctx, app.govKeeper),
		mint.ExportGenesis(ctx, app.mintKeeper))
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}
	validators = staking.WriteValidators(ctx, app.stakingKeeper)
	return appState, validators, nil
}

func (app *VPN) prepForZeroHeightGenesis(ctx csdkTypes.Context) {
	app.assertRuntimeInvariantsOnContext(ctx)
	app.stakingKeeper.IterateValidators(ctx, func(_ int64, val csdkTypes.Validator) (stop bool) {
		_ = app.distributionKeeper.WithdrawValidatorCommission(ctx, val.GetOperator())
		return false
	})

	dels := app.stakingKeeper.GetAllDelegations(ctx)
	for _, delegation := range dels {
		_ = app.distributionKeeper.WithdrawDelegationRewards(ctx, delegation.DelegatorAddr, delegation.ValidatorAddr)
	}
	app.distributionKeeper.DeleteAllValidatorSlashEvents(ctx)
	app.distributionKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	height := ctx.BlockHeight()
	ctx = ctx.WithBlockHeight(0)

	app.stakingKeeper.IterateValidators(ctx, func(_ int64, val csdkTypes.Validator) (stop bool) {
		app.distributionKeeper.Hooks().AfterValidatorCreated(ctx, val.GetOperator())
		return false
	})

	for _, del := range dels {
		app.distributionKeeper.Hooks().BeforeDelegationCreated(ctx, del.DelegatorAddr, del.ValidatorAddr)
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

	store := ctx.KVStore(app.keyStaking)
	iter := csdkTypes.KVStoreReversePrefixIterator(store, staking.ValidatorsKey)
	counter := int16(0)

	var valConsAddrs []csdkTypes.ConsAddress
	for ; iter.Valid(); iter.Next() {
		addr := csdkTypes.ValAddress(iter.Key()[1:])
		validator, found := app.stakingKeeper.GetValidator(ctx, addr)
		if !found {
			panic("expected validator, not found")
		}

		validator.BondHeight = 0
		validator.UnbondingHeight = 0
		valConsAddrs = append(valConsAddrs, validator.ConsAddress())

		app.stakingKeeper.SetValidator(ctx, validator)
		counter++
	}
	iter.Close()

	app.slashingKeeper.IterateValidatorSigningInfos(
		ctx,
		func(addr csdkTypes.ConsAddress, info slashing.ValidatorSigningInfo) (stop bool) {
			info.StartHeight = 0
			app.slashingKeeper.SetValidatorSigningInfo(ctx, addr, info)
			return false
		},
	)
}

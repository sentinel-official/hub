package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/sentinel-official/hub/x/deposit"
)

func RegisterInvariants(c types.CrisisKeeper, k keeper.Keeper, f types.FeeCollectionKeeper,
	d types.DistributionKeeper, am auth.AccountKeeper, dk deposit.Keeper) {

	c.RegisterRoute(types.ModuleName, "supply",
		SupplyInvariants(k, f, d, am, dk))
	c.RegisterRoute(types.ModuleName, "nonnegative-power",
		keeper.NonNegativePowerInvariant(k))
	c.RegisterRoute(types.ModuleName, "positive-delegation",
		keeper.PositiveDelegationInvariant(k))
	c.RegisterRoute(types.ModuleName, "delegator-shares",
		keeper.DelegatorSharesInvariant(k))
}

func SupplyInvariants(k keeper.Keeper, f types.FeeCollectionKeeper,
	d types.DistributionKeeper, am auth.AccountKeeper, dk deposit.Keeper) sdk.Invariant {

	return func(ctx sdk.Context) error {
		pool := k.GetPool(ctx)

		loose := sdk.ZeroDec()
		bonded := sdk.ZeroDec()
		am.IterateAccounts(ctx, func(acc auth.Account) bool {
			loose = loose.Add(acc.GetCoins().AmountOf(k.BondDenom(ctx)).ToDec())
			return false
		})
		k.IterateUnbondingDelegations(ctx, func(_ int64, ubd types.UnbondingDelegation) bool {
			for _, entry := range ubd.Entries {
				loose = loose.Add(entry.Balance.ToDec())
			}
			return false
		})
		k.IterateValidators(ctx, func(_ int64, validator sdk.Validator) bool {
			switch validator.GetStatus() {
			case sdk.Bonded:
				bonded = bonded.Add(validator.GetBondedTokens().ToDec())
			case sdk.Unbonding, sdk.Unbonded:
				loose = loose.Add(validator.GetTokens().ToDec())
			}
			loose = loose.Add(d.GetValidatorOutstandingRewardsCoins(ctx, validator.GetOperator()).AmountOf(k.BondDenom(ctx)))
			return false
		})

		loose = loose.Add(f.GetCollectedFees(ctx).AmountOf(k.BondDenom(ctx)).ToDec())

		loose = loose.Add(d.GetFeePoolCommunityCoins(ctx).AmountOf(k.BondDenom(ctx)))

		dk.IterateDeposits(ctx, func(index int64, dep deposit.Deposit) (stop bool) {
			loose = loose.Add(dep.Coins.AmountOf(k.BondDenom(ctx)).ToDec())
			return false
		})

		if !pool.NotBondedTokens.ToDec().Equal(loose) {
			return fmt.Errorf("loose token invariance:\n"+
				"\tpool.NotBondedTokens: %v\n"+
				"\tsum of account tokens: %v", pool.NotBondedTokens, loose)
		}

		if !pool.BondedTokens.ToDec().Equal(bonded) {
			return fmt.Errorf("bonded token invariance:\n"+
				"\tpool.BondedTokens: %v\n"+
				"\tsum of account tokens: %v", pool.BondedTokens, bonded)
		}

		return nil
	}
}

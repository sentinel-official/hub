package upgrade1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sentinel-official/hub/x/upgrade/expected"
)

func Handler(ctx sdk.Context, ak expected.AccountKeeper, bk expected.BankKeeper, sk stakingkeeper.Keeper) error {
	return migrateVestingAccounts(ctx, ak, bk, sk)
}

func migrateVestingAccounts(ctx sdk.Context, ak expected.AccountKeeper, bk expected.BankKeeper, sk stakingkeeper.Keeper) (err error) {
	ak.IterateAccounts(ctx, func(account authtypes.AccountI) bool {
		vestingAccount, ok := resetVestingDelegatedBalances(account)
		if !ok {
			return false
		}

		var (
			balances             = bk.GetAllBalances(ctx, account.GetAddress())
			delegations          = sdk.NewCoins()
			unbondingDelegations = sdk.NewCoins()
		)

		delegations, err = getDelegatorDelegationsSum(ctx, sk, account.GetAddress())
		if err != nil {
			return true
		}

		unbondingDelegations, err = getDelegatorUnbondingDelegationsSum(ctx, sk, account.GetAddress())
		if err != nil {
			return true
		}

		delegations = delegations.Add(unbondingDelegations...)

		for _, coin := range delegations {
			balances = balances.Add(coin)
		}

		vestingAccount.TrackDelegation(ctx.BlockTime(), balances, delegations)
		ak.SetAccount(ctx, vestingAccount)

		return false
	})

	return err
}

func getDelegatorDelegationsSum(c sdk.Context, sk stakingkeeper.Keeper, address sdk.AccAddress) (sdk.Coins, error) {
	var (
		ctx   = sdk.WrapSDKContext(c)
		qs    = stakingkeeper.Querier{Keeper: sk}
		coins = sdk.NewCoins()
	)

	res, err := qs.DelegatorDelegations(ctx,
		&stakingtypes.QueryDelegatorDelegationsRequest{
			DelegatorAddr: address.String(),
		},
	)

	switch status.Code(err) {
	case codes.OK:
		for _, delegation := range res.DelegationResponses {
			coins = coins.Add(delegation.GetBalance())
		}
	case codes.NotFound:
		return coins, nil
	default:
		return nil, err
	}

	return coins, nil
}

func getDelegatorUnbondingDelegationsSum(c sdk.Context, sk stakingkeeper.Keeper, address sdk.AccAddress) (sdk.Coins, error) {
	var (
		ctx   = sdk.WrapSDKContext(c)
		qs    = stakingkeeper.Querier{Keeper: sk}
		denom = sk.BondDenom(c)
		coins = sdk.NewCoins()
	)

	res, err := qs.DelegatorUnbondingDelegations(ctx,
		&stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
			DelegatorAddr: address.String(),
		},
	)

	switch status.Code(err) {
	case codes.OK:
		for _, delegation := range res.UnbondingResponses {
			for _, entry := range delegation.Entries {
				coins = coins.Add(sdk.NewCoin(denom, entry.Balance))
			}
		}
	case codes.NotFound:
		return coins, nil
	default:
		return nil, err
	}

	return coins, nil
}

func resetVestingDelegatedBalances(account authtypes.AccountI) (exported.VestingAccount, bool) {
	var (
		free    = sdk.NewCoins()
		vesting = sdk.NewCoins()
	)

	switch vestingAccount := account.(type) {
	case *vestingtypes.ContinuousVestingAccount:
		vestingAccount.DelegatedFree = free
		vestingAccount.DelegatedVesting = vesting
		return vestingAccount, true
	case *vestingtypes.DelayedVestingAccount:
		vestingAccount.DelegatedFree = free
		vestingAccount.DelegatedVesting = vesting
		return vestingAccount, true
	case *vestingtypes.PeriodicVestingAccount:
		vestingAccount.DelegatedFree = free
		vestingAccount.DelegatedVesting = vesting
		return vestingAccount, true
	default:
		return nil, false
	}
}

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit      = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	DefaultStakingShare = sdk.NewDecWithPrec(1, 1)
)

var (
	KeyDeposit      = []byte("Deposit")
	KeyStakingShare = []byte("StakingShare")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if m.Deposit.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !m.Deposit.IsValid() {
		return fmt.Errorf("invalid deposit %s", m.Deposit)
	}
	if m.StakingShare.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if m.StakingShare.GT(sdk.NewDec(1)) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyDeposit,
			Value: &m.Deposit,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(sdk.Coin)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value.IsNegative() {
					return fmt.Errorf("deposit cannot be negative")
				}
				if !value.IsValid() {
					return fmt.Errorf("invalid deposit %s", value)
				}

				return nil
			},
		},
		{
			Key:   KeyStakingShare,
			Value: &m.StakingShare,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(sdk.Dec)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value.IsNegative() {
					return fmt.Errorf("staking_share cannot be negative")
				}
				if value.GT(sdk.NewDec(1)) {
					return fmt.Errorf("staking_share cannot be greater than 1")
				}

				return nil
			},
		},
	}
}

func NewParams(deposit sdk.Coin, stakingShare sdk.Dec) Params {
	return Params{
		Deposit:      deposit,
		StakingShare: stakingShare,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:      DefaultDeposit,
		StakingShare: DefaultStakingShare,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

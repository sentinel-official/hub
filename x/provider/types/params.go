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
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateStakingShare(m.StakingShare); err != nil {
		return err
	}

	return nil
}

func (m *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:         KeyDeposit,
			Value:       &m.Deposit,
			ValidatorFn: validateDeposit,
		},
		{
			Key:         KeyStakingShare,
			Value:       &m.StakingShare,
			ValidatorFn: validateStakingShare,
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

func validateDeposit(v interface{}) error {
	value, ok := v.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("deposit cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !value.IsValid() {
		return fmt.Errorf("invalid deposit %s", value)
	}

	return nil
}

func validateStakingShare(v interface{}) error {
	value, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("staking_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("staking_share cannot be negative")
	}
	if value.GT(sdk.NewDec(1)) {
		return fmt.Errorf("staking_share cannot be greater than 1")
	}

	return nil
}

package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit          = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	DefaultInactiveDuration = 5 * time.Minute
	DefaultMaxPrice         = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultMinPrice         = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10)))
	DefaultStakingShare     = sdk.NewDecWithPrec(1, 1)
)

var (
	KeyDeposit          = []byte("Deposit")
	KeyInactiveDuration = []byte("InactiveDuration")
	KeyMaxPrice         = []byte("MaxPrice")
	KeyMinPrice         = []byte("MinPrice")
	KeyStakingShare     = []byte("StakingShare")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateInactiveDuration(m.InactiveDuration); err != nil {
		return err
	}
	if err := validateMaxPrice(m.MaxPrice); err != nil {
		return err
	}
	if err := validateMinPrice(m.MinPrice); err != nil {
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
			Key:         KeyInactiveDuration,
			Value:       &m.InactiveDuration,
			ValidatorFn: validateInactiveDuration,
		},
		{
			Key:         KeyMaxPrice,
			Value:       &m.MaxPrice,
			ValidatorFn: validateMaxPrice,
		},
		{
			Key:         KeyMinPrice,
			Value:       &m.MinPrice,
			ValidatorFn: validateMinPrice,
		},
		{
			Key:         KeyStakingShare,
			Value:       &m.StakingShare,
			ValidatorFn: validateStakingShare,
		},
	}
}

func NewParams(deposit sdk.Coin, inactiveDuration time.Duration, maxPrice, minPrice sdk.Coins, stakingShare sdk.Dec) Params {
	return Params{
		Deposit:          deposit,
		InactiveDuration: inactiveDuration,
		MaxPrice:         maxPrice,
		MinPrice:         minPrice,
		StakingShare:     stakingShare,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:          DefaultDeposit,
		InactiveDuration: DefaultInactiveDuration,
		MaxPrice:         DefaultMaxPrice,
		MinPrice:         DefaultMinPrice,
		StakingShare:     DefaultStakingShare,
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

func validateInactiveDuration(v interface{}) error {
	value, ok := v.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("inactive_duration cannot be negative")
	}
	if value == 0 {
		return fmt.Errorf("inactive_duration cannot be zero")
	}

	return nil
}

func validateMaxPrice(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("max_price cannot be nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("max_price must be valid")
	}

	return nil
}

func validateMinPrice(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("min_price cannot be nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("min_price must be valid")
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

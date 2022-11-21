package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultInactiveDuration = 5 * time.Minute
)

var (
	DefaultDeposit      = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	DefaultMaxPrice     = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultMinPrice     = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10)))
	DefaultStakingShare = sdk.NewDecWithPrec(1, 1)
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
	if m.Deposit.IsNegative() {
		return fmt.Errorf("deposit cannot be negative")
	}
	if !m.Deposit.IsValid() {
		return fmt.Errorf("invalid deposit %s", m.Deposit)
	}
	if m.InactiveDuration < 0 {
		return fmt.Errorf("inactive_duration cannot be negative")
	}
	if m.InactiveDuration == 0 {
		return fmt.Errorf("inactive_duration cannot be zero")
	}
	if m.MaxPrice != nil {
		if !m.MaxPrice.IsValid() {
			return fmt.Errorf("max_price must be valid")
		}
	}
	if m.MinPrice != nil {
		if !m.MinPrice.IsValid() {
			return fmt.Errorf("min_price must be valid")
		}
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
			Key:   KeyInactiveDuration,
			Value: &m.InactiveDuration,
			ValidatorFn: func(v interface{}) error {
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
			},
		},
		{
			Key:   KeyMaxPrice,
			Value: &m.MaxPrice,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(sdk.Coins)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value != nil {
					if !value.IsValid() {
						return fmt.Errorf("max_price must be valid")
					}
				}

				return nil
			},
		},
		{
			Key:   KeyMinPrice,
			Value: &m.MinPrice,
			ValidatorFn: func(v interface{}) error {
				value, ok := v.(sdk.Coins)
				if !ok {
					return fmt.Errorf("invalid parameter type %T", v)
				}

				if value != nil {
					if !value.IsValid() {
						return fmt.Errorf("min_price must be valid")
					}
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

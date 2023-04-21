package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit      = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
	DefaultRevenueShare = sdk.NewDecWithPrec(1, 1)
)

var (
	KeyDeposit      = []byte("Deposit")
	KeyRevenueShare = []byte("RevenueShare")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (m *Params) Validate() error {
	if err := validateDeposit(m.Deposit); err != nil {
		return err
	}
	if err := validateRevenueShare(m.RevenueShare); err != nil {
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
			Key:         KeyRevenueShare,
			Value:       &m.RevenueShare,
			ValidatorFn: validateRevenueShare,
		},
	}
}

func NewParams(deposit sdk.Coin, revenueShare sdk.Dec) Params {
	return Params{
		Deposit:      deposit,
		RevenueShare: revenueShare,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:      DefaultDeposit,
		RevenueShare: DefaultRevenueShare,
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

func validateRevenueShare(v interface{}) error {
	value, ok := v.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value.IsNil() {
		return fmt.Errorf("revenue_share cannot be nil")
	}
	if value.IsNegative() {
		return fmt.Errorf("revenue_share cannot be negative")
	}
	if value.GT(sdk.NewDec(1)) {
		return fmt.Errorf("revenue_share cannot be greater than 1")
	}

	return nil
}

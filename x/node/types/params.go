package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultDeposit                        = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))
	DefaultInactiveDuration               = 30 * time.Second
	DefaultMaxGigabytePrices              = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultMinGigabytePrices              = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
	DefaultMaxHourlyPrices                = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	DefaultMinHourlyPrices                = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1)))
	DefaultMaxSubscriptionGigabytes int64 = 10
	DefaultMinSubscriptionGigabytes int64 = 1
	DefaultMaxSubscriptionHours     int64 = 10
	DefaultMinSubscriptionHours     int64 = 1
	DefaultRevenueShare                   = sdk.NewDecWithPrec(1, 1)
)

var (
	KeyDeposit                  = []byte("Deposit")
	KeyInactiveDuration         = []byte("InactiveDuration")
	KeyMaxGigabytePrices        = []byte("MaxGigabytePrices")
	KeyMinGigabytePrices        = []byte("MinGigabytePrices")
	KeyMaxHourlyPrices          = []byte("MaxHourlyPrices")
	KeyMinHourlyPrices          = []byte("MinHourlyPrices")
	KeyMaxSubscriptionGigabytes = []byte("MaxSubscriptionGigabytes")
	KeyMinSubscriptionGigabytes = []byte("MinSubscriptionGigabytes")
	KeyMaxSubscriptionHours     = []byte("MaxSubscriptionHours")
	KeyMinSubscriptionHours     = []byte("MinSubscriptionHours")
	KeyRevenueShare             = []byte("RevenueShare")
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
	if err := validateMaxGigabytePrices(m.MaxGigabytePrices); err != nil {
		return err
	}
	if err := validateMinGigabytePrices(m.MinGigabytePrices); err != nil {
		return err
	}
	if err := validateMaxHourlyPrices(m.MaxHourlyPrices); err != nil {
		return err
	}
	if err := validateMinHourlyPrices(m.MinHourlyPrices); err != nil {
		return err
	}
	if err := validateMaxSubscriptionGigabytes(m.MaxSubscriptionGigabytes); err != nil {
		return err
	}
	if err := validateMinSubscriptionGigabytes(m.MinSubscriptionGigabytes); err != nil {
		return err
	}
	if err := validateMaxSubscriptionHours(m.MaxSubscriptionHours); err != nil {
		return err
	}
	if err := validateMinSubscriptionHours(m.MinSubscriptionHours); err != nil {
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
			Key:         KeyInactiveDuration,
			Value:       &m.InactiveDuration,
			ValidatorFn: validateInactiveDuration,
		},
		{
			Key:         KeyMaxGigabytePrices,
			Value:       &m.MaxGigabytePrices,
			ValidatorFn: validateMaxGigabytePrices,
		},
		{
			Key:         KeyMinGigabytePrices,
			Value:       &m.MinGigabytePrices,
			ValidatorFn: validateMinGigabytePrices,
		},
		{
			Key:         KeyMaxHourlyPrices,
			Value:       &m.MaxHourlyPrices,
			ValidatorFn: validateMaxHourlyPrices,
		},
		{
			Key:         KeyMinHourlyPrices,
			Value:       &m.MinHourlyPrices,
			ValidatorFn: validateMinHourlyPrices,
		},
		{
			Key:         KeyMaxSubscriptionGigabytes,
			Value:       &m.MaxSubscriptionGigabytes,
			ValidatorFn: validateMaxSubscriptionGigabytes,
		},
		{
			Key:         KeyMinSubscriptionGigabytes,
			Value:       &m.MinSubscriptionGigabytes,
			ValidatorFn: validateMinSubscriptionGigabytes,
		},
		{
			Key:         KeyMaxSubscriptionHours,
			Value:       &m.MaxSubscriptionHours,
			ValidatorFn: validateMaxSubscriptionHours,
		},
		{
			Key:         KeyMinSubscriptionHours,
			Value:       &m.MinSubscriptionHours,
			ValidatorFn: validateMinSubscriptionHours,
		},
		{
			Key:         KeyRevenueShare,
			Value:       &m.RevenueShare,
			ValidatorFn: validateRevenueShare,
		},
	}
}

func NewParams(
	deposit sdk.Coin, inactiveDuration time.Duration, maxGigabytePrices, minGigabytePrices,
	maxHourlyPrices, minHourlyPrices sdk.Coins, maxSubscriptionGigabytes, minSubscriptionGigabytes int64,
	maxSubscriptionHours, minSubscriptionHours int64, revenueShare sdk.Dec,
) Params {
	return Params{
		Deposit:                  deposit,
		InactiveDuration:         inactiveDuration,
		MaxGigabytePrices:        maxGigabytePrices,
		MinGigabytePrices:        minGigabytePrices,
		MaxHourlyPrices:          maxHourlyPrices,
		MinHourlyPrices:          minHourlyPrices,
		MaxSubscriptionGigabytes: maxSubscriptionGigabytes,
		MinSubscriptionGigabytes: minSubscriptionGigabytes,
		MaxSubscriptionHours:     maxSubscriptionHours,
		MinSubscriptionHours:     minSubscriptionHours,
		RevenueShare:             revenueShare,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultDeposit,
		DefaultInactiveDuration,
		DefaultMaxGigabytePrices,
		DefaultMinGigabytePrices,
		DefaultMaxHourlyPrices,
		DefaultMinHourlyPrices,
		DefaultMaxSubscriptionGigabytes,
		DefaultMinSubscriptionGigabytes,
		DefaultMaxSubscriptionHours,
		DefaultMinSubscriptionHours,
		DefaultRevenueShare,
	)
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

func validateMaxGigabytePrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("max_gigabyte_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("max_gigabyte_prices must be valid")
	}

	return nil
}

func validateMinGigabytePrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("min_gigabyte_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("min_gigabyte_prices must be valid")
	}

	return nil
}

func validateMaxHourlyPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("max_hourly_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("max_hourly_prices must be valid")
	}

	return nil
}

func validateMinHourlyPrices(v interface{}) error {
	value, ok := v.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value == nil {
		return nil
	}
	if value.IsAnyNil() {
		return fmt.Errorf("min_hourly_prices cannot contain nil")
	}
	if !value.IsValid() {
		return fmt.Errorf("min_hourly_prices must be valid")
	}

	return nil
}

func validateMaxSubscriptionGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_subscription_gigabytes cannot be negative")
	}

	return nil
}

func validateMinSubscriptionGigabytes(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_subscription_gigabytes cannot be negative")
	}

	return nil
}

func validateMaxSubscriptionHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("max_subscription_hours cannot be negative")
	}

	return nil
}

func validateMinSubscriptionHours(v interface{}) error {
	value, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type %T", v)
	}

	if value < 0 {
		return fmt.Errorf("min_subscription_hours cannot be negative")
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
	if value.GT(sdk.OneDec()) {
		return fmt.Errorf("revenue_share cannot be greater than 1")
	}

	return nil
}

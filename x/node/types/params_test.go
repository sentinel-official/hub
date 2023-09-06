package types

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit                  sdk.Coin
		ActiveDuration           time.Duration
		MaxGigabytePrices        sdk.Coins
		MinGigabytePrices        sdk.Coins
		MaxHourlyPrices          sdk.Coins
		MinHourlyPrices          sdk.Coins
		MaxSubscriptionGigabytes int64
		MinSubscriptionGigabytes int64
		MaxSubscriptionHours     int64
		MinSubscriptionHours     int64
		StakingShare             sdkmath.LegacyDec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"deposit empty",
			fields{
				Deposit: hubtypes.TestCoinEmpty,
			},
			true,
		},
		{
			"deposit empty denom",
			fields{
				Deposit: hubtypes.TestCoinEmptyDenom,
			},
			true,
		},
		{
			"deposit invalid denom",
			fields{
				Deposit: hubtypes.TestCoinInvalidDenom,
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				Deposit: hubtypes.TestCoinEmptyAmount,
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				Deposit: hubtypes.TestCoinNegativeAmount,
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				Deposit: hubtypes.TestCoinZeroAmount,
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"active_duration negative",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: -1000,
			},
			true,
		},
		{
			"active_duration zero",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: 0,
			},
			true,
		},
		{
			"active_duration positive",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices nil",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxGigabytePrices:        nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxGigabytePrices:        hubtypes.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty denom",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MaxGigabytePrices: hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"max_gigabyte_prices invalid denom",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MaxGigabytePrices: hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"max_gigabyte_prices empty amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MaxGigabytePrices: hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"max_gigabyte_prices negative amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MaxGigabytePrices: hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"max_gigabyte_prices zero amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MaxGigabytePrices: hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"max_gigabyte_prices positive amount",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxGigabytePrices:        hubtypes.TestCoinsPositiveAmount,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices nil",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinGigabytePrices:        nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinGigabytePrices:        hubtypes.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty denom",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MinGigabytePrices: hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"min_gigabyte_prices invalid denom",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MinGigabytePrices: hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"min_gigabyte_prices empty amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MinGigabytePrices: hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"min_gigabyte_prices negative amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MinGigabytePrices: hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"min_gigabyte_prices zero amount",
			fields{
				Deposit:           hubtypes.TestCoinPositiveAmount,
				ActiveDuration:    1000,
				MinGigabytePrices: hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"min_gigabyte_prices positive amount",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinGigabytePrices:        hubtypes.TestCoinsPositiveAmount,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices nil",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxHourlyPrices:          nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxHourlyPrices:          hubtypes.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty denom",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MaxHourlyPrices: hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"max_hourly_prices invalid denom",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MaxHourlyPrices: hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"max_hourly_prices empty amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MaxHourlyPrices: hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"max_hourly_prices negative amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MaxHourlyPrices: hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"max_hourly_prices zero amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MaxHourlyPrices: hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"max_hourly_prices positive amount",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxHourlyPrices:          hubtypes.TestCoinsPositiveAmount,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices nil",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinHourlyPrices:          nil,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinHourlyPrices:          hubtypes.TestCoinsEmpty,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty denom",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MinHourlyPrices: hubtypes.TestCoinsEmptyDenom,
			},
			true,
		},
		{
			"min_hourly_prices invalid denom",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MinHourlyPrices: hubtypes.TestCoinsInvalidDenom,
			},
			true,
		},
		{
			"min_hourly_prices empty amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MinHourlyPrices: hubtypes.TestCoinsEmptyAmount,
			},
			true,
		},
		{
			"min_hourly_prices negative amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MinHourlyPrices: hubtypes.TestCoinsNegativeAmount,
			},
			true,
		},
		{
			"min_hourly_prices zero amount",
			fields{
				Deposit:         hubtypes.TestCoinPositiveAmount,
				ActiveDuration:  1000,
				MinHourlyPrices: hubtypes.TestCoinsZeroAmount,
			},
			true,
		},
		{
			"min_hourly_prices positive amount",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MinHourlyPrices:          hubtypes.TestCoinsPositiveAmount,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_gigabytes negative",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"max_subscription_gigabytes zero",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"max_subscription_gigabytes positive",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_gigabytes negative",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"min_subscription_gigabytes zero",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_gigabytes positive",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_hours negative",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     -1000,
			},
			true,
		},
		{
			"max_subscription_hours zero",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"max_subscription_hours positive",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_hours negative",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     -1000,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_hours zero",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     0,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_hours positive",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share empty",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyDec{},
			},
			true,
		},
		{
			"staking_share -10",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-10, 0),
			},
			true,
		},
		{
			"staking_share -1",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-1, 0),
			},
			true,
		},
		{
			"staking_share -0.5",
			fields{
				Deposit:        hubtypes.TestCoinPositiveAmount,
				ActiveDuration: 1000,
				StakingShare:   sdkmath.LegacyNewDecWithPrec(-5, 1),
			},
			true,
		},
		{
			"staking_share 0",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(0, 0),
			},
			false,
		},
		{
			"staking_share 0.5",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(5, 1),
			},
			false,
		},
		{
			"staking_share 1",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share 10",
			fields{
				Deposit:                  hubtypes.TestCoinPositiveAmount,
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				MinSubscriptionGigabytes: 1,
				MaxSubscriptionHours:     1000,
				MinSubscriptionHours:     1,
				StakingShare:             sdkmath.LegacyNewDecWithPrec(10, 0),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				Deposit:                  tt.fields.Deposit,
				ActiveDuration:           tt.fields.ActiveDuration,
				MaxGigabytePrices:        tt.fields.MaxGigabytePrices,
				MinGigabytePrices:        tt.fields.MinGigabytePrices,
				MaxHourlyPrices:          tt.fields.MaxHourlyPrices,
				MinHourlyPrices:          tt.fields.MinHourlyPrices,
				MaxSubscriptionGigabytes: tt.fields.MaxSubscriptionGigabytes,
				MinSubscriptionGigabytes: tt.fields.MinSubscriptionGigabytes,
				MaxSubscriptionHours:     tt.fields.MaxSubscriptionHours,
				MinSubscriptionHours:     tt.fields.MinSubscriptionHours,
				StakingShare:             tt.fields.StakingShare,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

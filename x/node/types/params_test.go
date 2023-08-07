package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
		StakingShare             sdk.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"deposit empty",
			fields{
				Deposit: sdk.Coin{},
			},
			true,
		},
		{
			"deposit empty denom",
			fields{
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"deposit invalid denom",
			fields{
				Deposit: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"deposit empty amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"deposit negative amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"deposit zero amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"deposit positive amount",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"active_duration negative",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: -1000,
			},
			true,
		},
		{
			"active_duration zero",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 0,
			},
			true,
		},
		{
			"active_duration positive",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices nil",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: nil,
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{},
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_gigabyte_prices empty denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"max_gigabyte_prices invalid denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"max_gigabyte_prices empty amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"max_gigabyte_prices negative amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"max_gigabyte_prices zero amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"max_gigabyte_prices positive amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices nil",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: nil,
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{},
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_gigabyte_prices empty denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"min_gigabyte_prices invalid denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"min_gigabyte_prices empty amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"min_gigabyte_prices negative amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"min_gigabyte_prices zero amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"min_gigabyte_prices positive amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:    1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:      sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices nil",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: nil,
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{},
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_hourly_prices empty denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"max_hourly_prices invalid denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"max_hourly_prices empty amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"max_hourly_prices negative amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"max_hourly_prices zero amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"max_hourly_prices positive amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MaxHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices nil",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: nil,
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{},
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_hourly_prices empty denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"min_hourly_prices invalid denom",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"min_hourly_prices empty amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"min_hourly_prices negative amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"min_hourly_prices zero amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"min_hourly_prices positive amount",
			fields{
				Deposit:         sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:  1000,
				MinHourlyPrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:    sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_gigabytes negative",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"max_subscription_gigabytes zero",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 0,
				StakingShare:             sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_gigabytes positive",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MaxSubscriptionGigabytes: 1000,
				StakingShare:             sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_gigabytes negative",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MinSubscriptionGigabytes: -1000,
			},
			true,
		},
		{
			"min_subscription_gigabytes zero",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MinSubscriptionGigabytes: 0,
				StakingShare:             sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_gigabytes positive",
			fields{
				Deposit:                  sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:           1000,
				MinSubscriptionGigabytes: 1000,
				StakingShare:             sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_hours negative",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MaxSubscriptionHours: -1000,
			},
			true,
		},
		{
			"max_subscription_hours zero",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MaxSubscriptionHours: 0,
				StakingShare:         sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"max_subscription_hours positive",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MaxSubscriptionHours: 1000,
				StakingShare:         sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_hours negative",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MinSubscriptionHours: -1000,
				StakingShare:         sdk.NewDecWithPrec(1, 0),
			},
			true,
		},
		{
			"min_subscription_hours zero",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MinSubscriptionHours: 0,
				StakingShare:         sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"min_subscription_hours positive",
			fields{
				Deposit:              sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration:       1000,
				MinSubscriptionHours: 1000,
				StakingShare:         sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share empty",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.Dec{},
			},
			true,
		},
		{
			"staking_share -10",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(-10, 0),
			},
			true,
		},
		{
			"staking_share -1",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(-1, 0),
			},
			true,
		},
		{
			"staking_share -0.5",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(-5, 1),
			},
			true,
		},
		{
			"staking_share 0",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(0, 0),
			},
			false,
		},
		{
			"staking_share 0.5",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(5, 1),
			},
			false,
		},
		{
			"staking_share 1",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(1, 0),
			},
			false,
		},
		{
			"staking_share 10",
			fields{
				Deposit:        sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				ActiveDuration: 1000,
				StakingShare:   sdk.NewDecWithPrec(10, 0),
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

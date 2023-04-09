package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit           sdk.Coin
		InactiveDuration  time.Duration
		MaxGigabytePrices sdk.Coins
		MaxHourlyPrices   sdk.Coins
		MinGigabytePrices sdk.Coins
		MinHourlyPrices   sdk.Coins
		RevenueShare      sdk.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty deposit",
			fields{
				Deposit: sdk.Coin{},
			},
			true,
		},
		{
			"empty deposit denom",
			fields{
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"invalid deposit denom",
			fields{
				Deposit: sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"empty deposit amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"negative deposit amount",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero deposit amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
				InactiveDuration: 1,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"positive deposit amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"negative inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: -1000,
			},
			true,
		},
		{
			"zero inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 0,
			},
			true,
		},
		{
			"positive inactive duration",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"nil max gigabyte prices",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: nil,
				RevenueShare:      sdk.NewDec(0),
			},
			false,
		},
		{
			"empty max gigabyte prices",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty max gigabyte prices denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid max gigabyte prices denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty max gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative max gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero max gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive max gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MaxGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:      sdk.NewDec(0),
			},
			false,
		},
		{
			"nil max hourly prices",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  nil,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"empty max hourly prices",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{},
			},
			true,
		},
		{
			"empty max hourly prices denom",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid max hourly prices denom",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty max hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative max hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero max hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive max hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"nil min gigabyte prices",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: nil,
				RevenueShare:      sdk.NewDec(0),
			},
			false,
		},
		{
			"empty min gigabyte prices",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{},
			},
			true,
		},
		{
			"empty min gigabyte prices denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid min gigabyte prices denom",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty min gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative min gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero min gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive min gigabyte prices amount",
			fields{
				Deposit:           sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration:  1000,
				MinGigabytePrices: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:      sdk.NewDec(0),
			},
			false,
		},
		{
			"nil min hourly prices",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  nil,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"empty min hourly prices",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{},
			},
			true,
		},
		{
			"empty min hourly prices denom",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid min hourly prices denom",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "o", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty min hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative min hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero min hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive min hourly prices amount",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MinHourlyPrices:  sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"empty revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.Dec{},
			},
			true,
		},
		{
			"negative revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDec(-1),
			},
			true,
		},
		{
			"zero revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"less than 1 revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDecWithPrec(1, 1),
			},
			false,
		},
		{
			"equals to 1 revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDec(1),
			},
			false,
		},
		{
			"greater than 1 revenue share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				RevenueShare:     sdk.NewDec(2),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Params{
				Deposit:           tt.fields.Deposit,
				InactiveDuration:  tt.fields.InactiveDuration,
				MaxGigabytePrices: tt.fields.MaxGigabytePrices,
				MaxHourlyPrices:   tt.fields.MaxHourlyPrices,
				MinGigabytePrices: tt.fields.MinGigabytePrices,
				MinHourlyPrices:   tt.fields.MinHourlyPrices,
				RevenueShare:      tt.fields.RevenueShare,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package types

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit          sdk.Coin
		InactiveDuration time.Duration
		MaxPrice         sdk.Coins
		MinPrice         sdk.Coins
		StakingShare     sdk.Dec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty denom deposit",
			fields{
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"invalid denom deposit",
			fields{
				Deposit: sdk.Coin{Denom: "0", Amount: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"empty amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.Int{}},
			},
			true,
		},
		{
			"negative amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount deposit",
			fields{
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
			},
			true,
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
			},
			true,
		},
		{
			"nil max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         nil,
			},
			true,
		},
		{
			"empty max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{},
			},
			true,
		},
		{
			"empty denom coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid denom coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "0", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty amount coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative amount coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount coin max price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"nil min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         nil,
			},
			true,
		},
		{
			"empty min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{},
			},
			true,
		},
		{
			"empty denom coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"invalid denom coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "0", Amount: sdk.NewInt(1000)}},
			},
			true,
		},
		{
			"empty amount coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.Int{}}},
			},
			true,
		},
		{
			"negative amount coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)}},
			},
			true,
		},
		{
			"zero amount coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}},
			},
			true,
		},
		{
			"positive amount coin min price",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"empty staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.Dec{},
			},
			true,
		},
		{
			"less than 0 staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDec(-1),
			},
			true,
		},
		{
			"equals to 0 staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDec(0),
			},
			false,
		},
		{
			"less than 1 staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDecWithPrec(1, 1),
			},
			false,
		},
		{
			"equals to 1 staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDec(1),
			},
			false,
		},
		{
			"greater than 1 staking share",
			fields{
				Deposit:          sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				InactiveDuration: 1000,
				MaxPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				MinPrice:         sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)}},
				StakingShare:     sdk.NewDec(2),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Params{
				Deposit:          tt.fields.Deposit,
				InactiveDuration: tt.fields.InactiveDuration,
				MaxPrice:         tt.fields.MaxPrice,
				MinPrice:         tt.fields.MinPrice,
				StakingShare:     tt.fields.StakingShare,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

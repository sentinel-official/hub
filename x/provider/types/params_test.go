package types

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		Deposit      sdk.Coin
		StakingShare sdkmath.LegacyDec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"empty denom deposit",
			fields{
				Deposit: hubtypes.TestCoinEmptyDenom,
			},
			true,
		},
		{
			"invalid denom deposit",
			fields{
				Deposit: hubtypes.TestCoinInvalidDenom,
			},
			true,
		},
		{
			"empty amount deposit",
			fields{
				Deposit: hubtypes.TestCoinEmptyAmount,
			},
			true,
		},
		{
			"negative amount deposit",
			fields{
				Deposit: hubtypes.TestCoinNegativeAmount,
			},
			true,
		},
		{
			"zero amount deposit",
			fields{
				Deposit:      hubtypes.TestCoinZeroAmount,
				StakingShare: sdkmath.LegacyNewDec(0),
			},
			false,
		},
		{
			"positive amount deposit",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDec(0),
			},
			false,
		},
		{
			"empty staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyDec{},
			},
			true,
		},
		{
			"less than 0 staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDec(-1),
			},
			true,
		},
		{
			"equals to 0 staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDec(0),
			},
			false,
		},
		{
			"less than 1 staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDecWithPrec(1, 1),
			},
			false,
		},
		{
			"equals to 1 staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDec(1),
			},
			false,
		},
		{
			"greater than 1 staking share",
			fields{
				Deposit:      hubtypes.TestCoinPositiveAmount,
				StakingShare: sdkmath.LegacyNewDec(2),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Params{
				Deposit:      tt.fields.Deposit,
				StakingShare: tt.fields.StakingShare,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

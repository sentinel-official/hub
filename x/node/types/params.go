package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	DefaultInactiveDuration = 5 * time.Minute
)

var (
	DefaultDeposit = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))
)

var (
	KeyDeposit          = []byte("Deposit")
	KeyInactiveDuration = []byte("InactiveDuration")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

type Params struct {
	Deposit          sdk.Coin      `json:"deposit"`
	InactiveDuration time.Duration `json:"inactive_duration"`
}

func (p Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Deposit:           %s
Inactive duration: %s
`), p.Deposit, p.InactiveDuration)
}

func (p Params) Validate() error {
	if !p.Deposit.IsValid() {
		return fmt.Errorf("deposit should be valid")
	}
	if p.InactiveDuration <= 0 {
		return fmt.Errorf("inactive_duration should be positive")
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeyDeposit,
			Value: &p.Deposit,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
		{
			Key:   KeyInactiveDuration,
			Value: &p.InactiveDuration,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
	}
}

func NewParams(deposit sdk.Coin, inactiveDuration time.Duration) Params {
	return Params{
		Deposit:          deposit,
		InactiveDuration: inactiveDuration,
	}
}

func DefaultParams() Params {
	return Params{
		Deposit:          DefaultDeposit,
		InactiveDuration: DefaultInactiveDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

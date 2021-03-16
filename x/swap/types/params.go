package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/rand"
)

const (
	DefaultSwapEnabled = false
	DefaultSwapDenom   = "tsent"
)

var (
	DefaultApproveBy = sdk.AccAddress(rand.Bytes(20))
)

var (
	KeySwapEnabled = []byte("SwapEnabled")
	KeySwapDenom   = []byte("SwapDenom")
	KeyApproveBy   = []byte("ApproveBy")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	SwapEnabled bool           `json:"swap_enabled"`
	SwapDenom   string         `json:"swap_denom"`
	ApproveBy   sdk.AccAddress `json:"approve_by"`
}

func (p Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Swap enabled: %s
Swap denom  : %s
Approve by  : %s
`), p.SwapEnabled, p.SwapDenom, p.ApproveBy)
}

func (p Params) Validate() error {
	if err := sdk.ValidateDenom(p.SwapDenom); err != nil {
		return err
	}
	if p.ApproveBy == nil || p.ApproveBy.Empty() {
		return fmt.Errorf("approve_by should not nil or empty")
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{
			Key:   KeySwapEnabled,
			Value: &p.SwapEnabled,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
		{
			Key:   KeySwapDenom,
			Value: &p.SwapDenom,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
		{
			Key:   KeyApproveBy,
			Value: &p.ApproveBy,
			ValidatorFn: func(_ interface{}) error {
				return nil
			},
		},
	}
}

func NewParams(swapEnabled bool, swapDenom string, approveBy sdk.AccAddress) Params {
	return Params{
		SwapEnabled: swapEnabled,
		SwapDenom:   swapDenom,
		ApproveBy:   approveBy,
	}
}

func DefaultParams() Params {
	return Params{
		SwapEnabled: DefaultSwapEnabled,
		SwapDenom:   DefaultSwapDenom,
		ApproveBy:   DefaultApproveBy,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

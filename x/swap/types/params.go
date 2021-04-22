package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultSwapEnabled = false
	DefaultSwapDenom   = "tsent"
	DefaultApproveBy   = ""
)

var (
	KeySwapEnabled = []byte("SwapEnabled")
	KeySwapDenom   = []byte("SwapDenom")
	KeyApproveBy   = []byte("ApproveBy")
)

var (
	_ params.ParamSet = (*Params)(nil)
)

func (p *Params) Validate() error {
	if err := sdk.ValidateDenom(p.SwapDenom); err != nil {
		return err
	}

	approveBy, err := sdk.AccAddressFromBech32(p.ApproveBy)
	if err != nil {
		return err
	}
	if approveBy == nil || approveBy.Empty() {
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

func NewParams(swapEnabled bool, swapDenom, approveBy string) Params {
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

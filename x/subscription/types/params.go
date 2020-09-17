package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	DefaultCancelDuration = 10 * time.Minute
)

var (
	KeyCancelDuration = []byte("CancelDuration")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	CancelDuration time.Duration `json:"cancel_duration"`
}

func (p Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Cancel duration: %s
`), p.CancelDuration)
}

func (p Params) Validate() error {
	if p.CancelDuration <= 0 {
		return fmt.Errorf("cancel_duration should be positive")
	}

	return nil
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyCancelDuration, Value: &p.CancelDuration},
	}
}

func NewParams(cancelDuration time.Duration) Params {
	return Params{
		CancelDuration: cancelDuration,
	}
}

func DefaultParams() Params {
	return Params{
		CancelDuration: DefaultCancelDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

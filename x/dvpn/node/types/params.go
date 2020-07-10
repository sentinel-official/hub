package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/x/params"
)

const (
	DefaultInactiveDuration = 30 * time.Minute
)

var (
	KeyInactiveDuration = []byte("InactiveDuration")
)

var _ params.ParamSet = (*Params)(nil)

type Params struct {
	InactiveDuration time.Duration `json:"inactive_duration"`
}

func (p Params) String() string {
	return fmt.Sprintf(strings.TrimSpace(`
Inactive duration: %s
`), p.InactiveDuration)
}

func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{Key: KeyInactiveDuration, Value: &p.InactiveDuration},
	}
}

func NewParams(inactiveDuration time.Duration) Params {
	return Params{
		InactiveDuration: inactiveDuration,
	}
}

func DefaultParams() Params {
	return Params{
		InactiveDuration: DefaultInactiveDuration,
	}
}

func ParamsKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

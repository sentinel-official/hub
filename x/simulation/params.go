package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/simulation"
)

type ParamChange struct {
	subspace string
	key      string
	simValue simulation.SimValFn
}

func NewSimParamChange(subspace, key string, simVal simulation.SimValFn) simulation.ParamChange {
	return ParamChange{
		subspace: subspace,
		key:      key,
		simValue: simVal,
	}
}

func (pc ParamChange) Subspace() string {
	return pc.subspace
}

func (pc ParamChange) Key() string {
	return pc.key
}

func (pc ParamChange) SimValue() simulation.SimValFn {
	return pc.simValue
}

func (pc ParamChange) ComposedKey() string {
	return fmt.Sprintf("%s/%s", pc.subspace, pc.key)
}

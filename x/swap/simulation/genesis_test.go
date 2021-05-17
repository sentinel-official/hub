package simulation

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocdc "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sentinel-official/hub/x/swap/types"
	"github.com/stretchr/testify/require"
)

func TestRandomizedGenesisState(t *testing.T) {
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	cryptocdc.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := &module.SimulationState{
		AppParams:    make(sdksimulation.AppParams),
		Cdc:          cdc,
		Rand:         r,
		GenState:     make(map[string]json.RawMessage),
		Accounts:     sdksimulation.RandomAccounts(r, 3),
		InitialStake: 1000,
		NumBonded:    0,
		GenTimestamp: time.Time{},
		UnbondTime:   0,
		ParamChanges: []sdksimulation.ParamChange{},
		Contents:     []sdksimulation.WeightedProposalContent{},
	}

	RandomizedGenesisState(simState)

	var swapGenesis types.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[types.ModuleName], &swapGenesis)

	require.Equal(t, "sent1grdunxx5jxd0ja75wt508sn6v39p70hhw53zs8", swapGenesis.Swaps[0].Receiver)
	require.Equal(t, []byte{
		0x62, 0x33, 0x37, 0x63, 0x34, 0x62, 0x33, 0x36, 0x32, 0x39, 0x38, 0x66, 0x32, 0x30, 0x61, 0x63, 0x37, 0x66,
		0x63, 0x35, 0x64, 0x65, 0x32, 0x62, 0x34, 0x61, 0x63, 0x31, 0x31, 0x36, 0x37, 0x64, 0x37, 0x61, 0x30, 0x34,
		0x34, 0x34, 0x30, 0x32, 0x61, 0x34, 0x30, 0x66, 0x32, 0x62, 0x35, 0x66, 0x31, 0x65, 0x31, 0x62, 0x62, 0x37,
		0x66, 0x30, 0x65, 0x65, 0x35, 0x66, 0x36, 0x33, 0x34, 0x36,
	}, swapGenesis.Swaps[0].TxHash)
	require.Equal(t, sdk.Coin{Denom: "sent", Amount: sdk.NewInt(1000)}, swapGenesis.Swaps[0].Amount)
	require.Equal(t, false, swapGenesis.Params.SwapEnabled)
	require.Equal(t, "cxgdXhhuTSkuxK", swapGenesis.Params.SwapDenom)
}

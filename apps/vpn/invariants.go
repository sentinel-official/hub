package vpn

import (
	"fmt"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankSim "github.com/cosmos/cosmos-sdk/x/bank/simulation"
	distributionSim "github.com/cosmos/cosmos-sdk/x/distribution/simulation"
	mockSim "github.com/cosmos/cosmos-sdk/x/mock/simulation"
	stakingSim "github.com/cosmos/cosmos-sdk/x/staking/simulation"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func (app *VPN) runtimeInvariants() []mockSim.Invariant {
	return []mockSim.Invariant{
		bankSim.NonnegativeBalanceInvariant(app.accountKeeper),
		// stakingSim.SupplyInvariants(app.bankKeeper, app.stakingKeeper, app.feeCollectionKeeper, app.distributionKeeper, app.accountKeeper),
		stakingSim.NonNegativePowerInvariant(app.stakingKeeper),
		distributionSim.NonNegativeOutstandingInvariant(app.distributionKeeper),
	}
}

func (app *VPN) assertRuntimeInvariants() {
	ctx := app.NewContext(false, abciTypes.Header{Height: app.LastBlockHeight() + 1})
	app.assertRuntimeInvariantsOnContext(ctx)
}

func (app *VPN) assertRuntimeInvariantsOnContext(ctx csdkTypes.Context) {
	start := time.Now()
	invariants := app.runtimeInvariants()
	for _, inv := range invariants {
		if err := inv(ctx); err != nil {
			panic(fmt.Errorf("invariant broken: %s", err))
		}
	}
	end := time.Now()
	diff := end.Sub(start)
	app.BaseApp.Logger.With("module", "invariants").Info("Asserted all invariants", "duration", diff)
}

package hub

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func init() {
	simapp.GetSimulatorFlags()
}

func TestAppExport(t *testing.T) {
	var (
		db      = tmdb.NewMemDB()
		logger  = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		app     = NewApp(logger, db, nil, true, map[int64]bool{}, 0)
		genesis = ModuleBasics.DefaultGenesis()
	)

	state, err := codec.MarshalJSONIndent(app.cdc, genesis)
	require.NoError(t, err)

	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: state,
		},
	)
	app.Commit()

	app = NewApp(logger, db, nil, true, map[int64]bool{}, 0)
	_, _, err = app.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func fauxMerkleModeOpt(app *baseapp.BaseApp) {
	app.SetFauxMerkleMode()
}

func TestFullAppSimulation(t *testing.T) {
	config, db, dir, logger, skip, err := simapp.SetupSimulation("leveldb-simulation", "simulation")
	if skip {
		t.Skip("skipping application simulation")
	}
	require.NoError(t, err, "simulation setup failed")

	defer func() {
		require.NoError(t, db.Close())
		require.NoError(t, os.RemoveAll(dir))
	}()

	app := NewApp(logger, db, nil, true, map[int64]bool{}, simapp.FlagPeriodValue, fauxMerkleModeOpt)
	require.Equal(t, appName, app.Name())

	_, params, err := simulation.SimulateFromSeed(
		t, os.Stdout, app.BaseApp, simapp.AppStateFn(app.Codec(), app.SimulationManager()),
		simapp.SimulationOperations(app, app.Codec(), config),
		app.ModuleAccountAddrs(), config,
	)

	require.NoError(t, simapp.CheckExportSimulation(app, config, params))
	require.NoError(t, err)

	if config.Commit {
		simapp.PrintStats(db)
	}
}

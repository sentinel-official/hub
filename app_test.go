package hub

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"

	node "github.com/sentinel-official/hub/x/node/simulation"
	plan "github.com/sentinel-official/hub/x/plan/simulation"
	provider "github.com/sentinel-official/hub/x/provider/simulation"
	session "github.com/sentinel-official/hub/x/session/simulation"
	subscription "github.com/sentinel-official/hub/x/subscription/simulation"
)

var (
	genesisFile        string
	paramsFile         string
	exportParamsPath   string
	exportParamsHeight int
	exportStatePath    string
	exportStatsPath    string
	seed               int64
	initialBlockHeight int
	numBlocks          int
	blockSize          int
	enabled            bool
	verbose            bool
	lean               bool
	commit             bool
	period             int
	onOperation        bool
	allInvariants      bool
	genesisTime        int64
)

func init() {
	flag.StringVar(&genesisFile, "Genesis", "", "custom simulation genesis file; cannot be used with params file")
	flag.StringVar(&paramsFile, "Params", "", "custom simulation params file which overrides any random params; cannot be used with genesis")
	flag.StringVar(&exportParamsPath, "ExportParamsPath", "", "custom file path to save the exported params JSON")
	flag.IntVar(&exportParamsHeight, "ExportParamsHeight", 0, "height to which export the randomly generated params")
	flag.StringVar(&exportStatePath, "ExportStatePath", "", "custom file path to save the exported app state JSON")
	flag.StringVar(&exportStatsPath, "ExportStatsPath", "", "custom file path to save the exported simulation statistics JSON")
	flag.Int64Var(&seed, "Seed", time.Now().UnixNano(), "simulation random seed")
	flag.IntVar(&initialBlockHeight, "InitialBlockHeight", 1, "initial block to start the simulation")
	flag.IntVar(&numBlocks, "NumBlocks", 500, "number of new blocks to simulate from the initial block height")
	flag.IntVar(&blockSize, "BlockSize", 200, "operations per block")
	flag.BoolVar(&enabled, "Enabled", false, "enable the simulation")
	flag.BoolVar(&verbose, "Verbose", false, "verbose log output")
	flag.BoolVar(&lean, "Lean", false, "lean simulation log output")
	flag.BoolVar(&commit, "Commit", false, "have the simulation commit")
	flag.IntVar(&period, "Period", 1, "run slow invariants only once every period assertions")
	flag.BoolVar(&onOperation, "SimulateEveryOperation", false, "run slow invariants every operation")
	flag.BoolVar(&allInvariants, "PrintAllInvariants", false, "print all invariants if a broken invariant is found")
	flag.Int64Var(&genesisTime, "GenesisTime", 0, "override genesis UNIX time instead of using a random UNIX time")
}

func TestFullAppSimulation(t *testing.T) {
	if !enabled {
		t.Skip("Skipping application simulation")
	}

	logger := log.NewNopLogger()
	if verbose {
		logger = log.TestingLogger()
	}

	directory, err := ioutil.TempDir("", "go_leveldb_simulation")
	require.NoError(t, err)

	db, err := sdk.NewLevelDB("go_leveldb", directory)
	require.NoError(t, err)

	defer func() {
		db.Close()
		require.NoError(t, os.RemoveAll(directory))
	}()

	app := NewApp(logger, db, nil, true, 0, func(app *baseapp.BaseApp) {
		app.SetFauxMerkleMode()
	})
	_, params, err := simulation.SimulateFromSeed(simulateFromSeed(t, app))

	if exportStatePath != "" {
		fmt.Println("Exporting app state")
		state, _, err := app.ExportAppStateAndValidators(false, nil)
		require.NoError(t, err)
		require.NoError(t, ioutil.WriteFile(exportStatePath, state, 0644))
	}
	if exportParamsPath != "" {
		fmt.Println("Exporting simulation params")
		bytes, err := json.MarshalIndent(params, "", " ")
		require.NoError(t, err)
		require.NoError(t, ioutil.WriteFile(exportParamsPath, bytes, 0644))
	}

	require.NoError(t, err)

	if commit {
		fmt.Println("\nGoLevelDB Stats")
		fmt.Println(db.Stats()["leveldb.stats"])
		fmt.Println("GoLevelDB cached block size", db.Stats()["leveldb.cachedblock"])
	}
}

func invariants(app *App) sdk.Invariants {
	if period == 1 {
		return app.crisisKeeper.Invariants()
	}
	return simulation.PeriodicInvariants(app.crisisKeeper.Invariants(), period, 0)
}

func simulateFromSeed(t testing.TB, app *App) (
	testing.TB, io.Writer, *baseapp.BaseApp, simulation.AppStateFn, int64, simulation.WeightedOperations,
	sdk.Invariants, int, int, int, int, string, bool, bool, bool, bool, bool, map[string]bool) {
	return t, os.Stdout, app.BaseApp, appState, seed, operations(app), invariants(app), initialBlockHeight, numBlocks,
		exportParamsHeight, blockSize, exportStatsPath, exportParamsPath != "", commit, lean, onOperation,
		allInvariants, app.ModuleAccountAddresses()
}

func appState(r *rand.Rand, accounts []simulation.Account) (json.RawMessage, []simulation.Account, string, time.Time) {
	timestamp := simulation.RandTimestamp(r)
	if genesisTime != 0 {
		timestamp = time.Unix(genesisTime, 0)
	}

	cdc := MakeCodec()
	params := make(simulation.AppParams)

	switch {
	case paramsFile != "" && genesisFile != "":
		panic("cannot provide both a genesis file and a params file")
	case genesisFile != "":
		doc, accounts := stateFromGenesisFile(r)
		if genesisTime == 0 {
			timestamp = doc.GenesisTime
		}

		return doc.AppState, accounts, doc.ChainID, timestamp
	case paramsFile != "":
		bytes, err := ioutil.ReadFile(paramsFile)
		if err != nil {
			panic(err)
		}

		cdc.MustUnmarshalJSON(bytes, &params)
	}

	state, accounts, chain := randomizedAppState(r, accounts, timestamp, params)
	return state, accounts, chain, timestamp
}

func randomizedAppState(r *rand.Rand, accounts []simulation.Account, timestamp time.Time,
	params simulation.AppParams) (json.RawMessage, []simulation.Account, string) {
	var (
		amount  int64
		bonded  int64
		cdc     = MakeCodec()
		genesis = ModuleBasics.DefaultGenesis()
	)

	params.GetOrGenerate(cdc, simapp.StakePerAccount, &amount, r, func(r *rand.Rand) { amount = int64(r.Intn(1e15)) })
	params.GetOrGenerate(cdc, simapp.InitiallyBondedValidators, &amount, r, func(r *rand.Rand) { bonded = int64(r.Intn(250)) })

	if bonded > int64(len(accounts)) {
		bonded = int64(len(accounts))
	}

	fmt.Printf(`Selected randomly generated parameters for simulated genesis:
{
  "stake_per_account": "%v",
  "initially_bonded_validators": "%v"
}
`, amount, bonded)

	simapp.GenGenesisAccounts(cdc, r, accounts, timestamp, amount, bonded, genesis)
	simapp.GenAuthGenesisState(cdc, r, params, genesis)
	simapp.GenBankGenesisState(cdc, r, params, genesis)
	simapp.GenSupplyGenesisState(cdc, amount, bonded, int64(len(accounts)), genesis)
	simapp.GenGovGenesisState(cdc, r, params, genesis)
	simapp.GenMintGenesisState(cdc, r, params, genesis)
	simapp.GenDistrGenesisState(cdc, r, params, genesis)
	simapp.GenSlashingGenesisState(cdc, r, simapp.GenStakingGenesisState(
		cdc, r, accounts, amount, int64(len(accounts)), bonded, params, genesis), params, genesis)

	state, err := cdc.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return state, accounts, "simulation"
}

func stateFromGenesisFile(r *rand.Rand) (doc tm.GenesisDoc, accounts []simulation.Account) {
	bytes, err := ioutil.ReadFile(genesisFile)
	if err != nil {
		panic(err)
	}

	cdc := MakeCodec()
	cdc.MustUnmarshalJSON(bytes, &genesisFile)

	var state map[string]json.RawMessage
	cdc.MustUnmarshalJSON(doc.AppState, &state)

	genesisAccounts := genaccounts.GetGenesisStateFromAppState(cdc, state)
	accounts = make([]simulation.Account, len(genesisAccounts))
	for i, acc := range genesisAccounts {
		privateKeySeed := make([]byte, 15)
		r.Read(privateKeySeed)

		privateKey := secp256k1.GenPrivKeySecp256k1(privateKeySeed)
		accounts[i] = simulation.Account{PrivKey: privateKey, PubKey: privateKey.PubKey(), Address: acc.Address}
	}

	return doc, accounts
}

func operations(app *App) []simulation.WeightedOperation {
	var (
		cdc    = MakeCodec()
		params = make(simulation.AppParams)
	)

	if paramsFile != "" {
		bytes, err := ioutil.ReadFile(paramsFile)
		if err != nil {
			panic(err)
		}

		cdc.MustUnmarshalJSON(bytes, &params)
	}

	return []simulation.WeightedOperation{
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "provider:weight_msg_register", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: provider.SimulateMsgRegister(app.vpnKeeper.Provider),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "provider:weight_msg_update", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: provider.SimulateMsgUpdate(app.vpnKeeper.Provider),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "plan:weight_msg_add", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: plan.SimulateMsgAdd(app.vpnKeeper.Provider, app.vpnKeeper.Plan),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "plan:weight_msg_set_status", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: plan.SimulateMsgSetStatus(app.vpnKeeper.Provider, app.vpnKeeper.Plan),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_register", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: node.SimulateMsgRegister(app.vpnKeeper.Provider, app.vpnKeeper.Node),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_update", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: node.SimulateMsgUpdate(app.vpnKeeper.Provider, app.vpnKeeper.Node),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "node:weight_msg_set_status", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: node.SimulateMsgSetStatus(app.vpnKeeper.Node),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "plan:weight_msg_add_node", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: plan.SimulateMsgAddNode(app.vpnKeeper.Provider, app.vpnKeeper.Node, app.vpnKeeper.Plan),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "plan:weight_msg_remove_node", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: plan.SimulateMsgRemoveNode(app.vpnKeeper.Provider, app.vpnKeeper.Node, app.vpnKeeper.Plan),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "subscription:weight_msg_subscribe_to_plan", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: subscription.SimulateMsgSubscribeToPlan(app.vpnKeeper.Plan, app.vpnKeeper.Subscription),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "subscription:weight_msg_subscribe_to_node", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: subscription.SimulateMsgSubscribeToNode(app.vpnKeeper.Node, app.vpnKeeper.Subscription),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "subscription:weight_msg_cancel", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: subscription.SimulateMsgCancel(app.vpnKeeper.Subscription),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "subscription:weight_msg_add_quota", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: subscription.SimulateMsgAddQuota(app.vpnKeeper.Subscription),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "subscription:weight_msg_update_quota", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: subscription.SimulateMsgUpdateQuota(app.vpnKeeper.Subscription),
		},
		{
			Weight: func(_ *rand.Rand) (v int) {
				params.GetOrGenerate(cdc, "session:weight_msg_upsert", &v, nil,
					func(_ *rand.Rand) { v = 100 },
				)
				return v
			}(nil),
			Op: session.SimulateUpsert(app.vpnKeeper.Node, app.vpnKeeper.Subscription, app.vpnKeeper.Session),
		},
	}
}

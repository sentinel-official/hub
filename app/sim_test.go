package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authSim "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankSim "github.com/cosmos/cosmos-sdk/x/bank/simulation"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionSim "github.com/cosmos/cosmos-sdk/x/distribution/simulation"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govSim "github.com/cosmos/cosmos-sdk/x/gov/simulation"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingSim "github.com/cosmos/cosmos-sdk/x/slashing/simulation"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingSim "github.com/cosmos/cosmos-sdk/x/staking/simulation"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/deposit"
	_staking "github.com/sentinel-official/hub/x/staking"
	"github.com/sentinel-official/hub/x/vpn"
	vpnSim "github.com/sentinel-official/hub/x/vpn/simulation"
)

func getRandomGenesisState(r *rand.Rand, accounts []simulation.Account,
	genesisTimestamp time.Time) (json.RawMessage, []simulation.Account, string) {

	var genesisAccounts []GenesisAccount

	amount := int64(r.Intn(1000000000000))
	numInitiallyBonded := int64(r.Intn(250))
	accountsLen := int64(len(accounts))
	if numInitiallyBonded > accountsLen {
		numInitiallyBonded = accountsLen
	}

	fmt.Printf("Selected randomly generated parameters for simulated genesis:\n"+
		"\t{amount of stake per account: %v, initially bonded validators: %v}\n",
		amount, numInitiallyBonded)

	genesisAccounts, accounts = setAccounts(r, genesisAccounts, accounts, amount, numInitiallyBonded, genesisTimestamp)
	randomStakingGenesis := getRandomStakingGenesis(r, amount, numInitiallyBonded, accounts, accountsLen)

	genesis := GenesisState{
		Accounts:     genesisAccounts,
		Auth:         getRandomAuthGenesis(r),
		Bank:         getRandomBankGenesis(r),
		Staking:      randomStakingGenesis,
		Mint:         getRandomMintGenesis(r),
		Distribution: getRandomDistributionGenesis(r),
		Slashing:     getRandomSlashingGenesis(r, randomStakingGenesis),
		Gov:          getRandomGovGenesis(r),
		Deposit:      getRandomDepositGenesis(r, accounts),
		VPN:          getRandomVPNGenesis(r, accounts),
	}

	appState, err := MakeCodec().MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return appState, accounts, "hub"
}

func setWeightedOperations(app *HubApp) []simulation.WeightedOperation {
	return []simulation.WeightedOperation{
		{5, authSim.SimulateDeductFee(app.accountKeeper, app.feeCollectionKeeper)},
		{100, bankSim.SimulateMsgSend(app.accountKeeper, app.bankKeeper)},
		{10, bankSim.SimulateSingleInputMsgMultiSend(app.accountKeeper, app.bankKeeper)},
		{50, distributionSim.SimulateMsgSetWithdrawAddress(app.accountKeeper, app.distributionKeeper)},
		{50, distributionSim.SimulateMsgWithdrawDelegatorReward(app.accountKeeper, app.distributionKeeper)},
		{50, distributionSim.SimulateMsgWithdrawValidatorCommission(app.accountKeeper, app.distributionKeeper)},
		{5, govSim.SimulateSubmittingVotingAndSlashingForProposal(app.govKeeper)},
		{100, govSim.SimulateMsgDeposit(app.govKeeper)},
		{100, stakingSim.SimulateMsgCreateValidator(app.accountKeeper, app.stakingKeeper)},
		{5, stakingSim.SimulateMsgEditValidator(app.stakingKeeper)},
		{100, stakingSim.SimulateMsgDelegate(app.accountKeeper, app.stakingKeeper)},
		{100, stakingSim.SimulateMsgUndelegate(app.accountKeeper, app.stakingKeeper)},
		{100, stakingSim.SimulateMsgBeginRedelegate(app.accountKeeper, app.stakingKeeper)},
		{100, slashingSim.SimulateMsgUnjail(app.slashingKeeper)},
		{100, vpnSim.SimulateMsgUpdateNodeStatus(app.vpnKeeper)},
		{100, vpnSim.SimulateMsgUpdateNodeInfo(app.vpnKeeper)},
		{100, vpnSim.SimulateMsgRegisterNode(app.vpnKeeper, app.accountKeeper)},
		{100, vpnSim.SimulateMsgStartSubscription(app.vpnKeeper, app.accountKeeper)},
		{100, vpnSim.SimulateMsgEndSubscription(app.vpnKeeper)},
		{100, vpnSim.SimulateMsgUpdateSessionInfo(app.vpnKeeper)},
	}
}

func setInvariants(app *HubApp) []sdk.Invariant {
	return []sdk.Invariant{
		simulation.PeriodicInvariant(bank.NonnegativeBalanceInvariant(app.accountKeeper), period, 0),
		simulation.PeriodicInvariant(distribution.AllInvariants(app.distributionKeeper, app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(_staking.SupplyInvariants(app.stakingKeeper, app.feeCollectionKeeper,
			app.distributionKeeper, app.accountKeeper, app.depositKeeper), period, 0),
		simulation.PeriodicInvariant(staking.NonNegativePowerInvariant(app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(staking.DelegatorSharesInvariant(app.stakingKeeper), period, 0),
	}
}

func getRandomAuthGenesis(r *rand.Rand) auth.GenesisState {
	return auth.GenesisState{
		Params: auth.Params{
			MaxMemoCharacters:      uint64(simulation.RandIntBetween(r, 100, 200)),
			TxSigLimit:             uint64(r.Intn(7) + 1),
			TxSizeCostPerByte:      uint64(simulation.RandIntBetween(r, 5, 15)),
			SigVerifyCostED25519:   uint64(simulation.RandIntBetween(r, 500, 1000)),
			SigVerifyCostSecp256k1: uint64(simulation.RandIntBetween(r, 500, 1000)),
		},
	}
}

func getRandomBankGenesis(r *rand.Rand) bank.GenesisState {
	return bank.NewGenesisState(r.Int63n(2) == 0)
}

func getRandomGovGenesis(r *rand.Rand) gov.GenesisState {
	vp := time.Duration(r.Intn(2*172800)) * time.Second

	return gov.GenesisState{
		StartingProposalID: uint64(r.Intn(100)),
		DepositParams: gov.DepositParams{
			MinDeposit:       sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(r.Intn(1e3)))},
			MaxDepositPeriod: vp,
		},
		VotingParams: gov.VotingParams{
			VotingPeriod: vp,
		},
		TallyParams: gov.TallyParams{
			Quorum:    sdk.NewDecWithPrec(334, 3),
			Threshold: sdk.NewDecWithPrec(5, 1),
			Veto:      sdk.NewDecWithPrec(334, 3),
		},
	}
}

func getRandomStakingGenesis(r *rand.Rand, amount, initiallyBonded int64,
	accounts []simulation.Account, accountsLen int64) staking.GenesisState {

	stakingGenesis := staking.GenesisState{
		Pool: staking.InitialPool(),
		Params: staking.Params{
			UnbondingTime: time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second,
			MaxValidators: uint16(r.Intn(250) + 1),
			BondDenom:     sdk.DefaultBondDenom,
		},
	}

	var validators []staking.Validator
	var delegations []staking.Delegation
	valAddrs := make([]sdk.ValAddress, initiallyBonded)
	for i := 0; i < int(initiallyBonded); i++ {
		valAddr := sdk.ValAddress(accounts[i].Address)
		valAddrs[i] = valAddr

		validator := staking.NewValidator(valAddr, accounts[i].PubKey, staking.Description{})
		validator.Tokens = sdk.NewInt(amount)
		validator.DelegatorShares = sdk.NewDec(amount)
		delegation := staking.Delegation{DelegatorAddress: accounts[i].Address, ValidatorAddress: valAddr, Shares: sdk.NewDec(amount)}
		validators = append(validators, validator)
		delegations = append(delegations, delegation)
	}

	stakingGenesis.Pool.NotBondedTokens = sdk.NewInt((amount * accountsLen) + (initiallyBonded * amount))
	stakingGenesis.Validators = validators
	stakingGenesis.Delegations = delegations

	return stakingGenesis
}

func getRandomSlashingGenesis(r *rand.Rand, stakingGenesis staking.GenesisState) slashing.GenesisState {
	return slashing.GenesisState{
		Params: slashing.Params{
			MaxEvidenceAge:          stakingGenesis.Params.UnbondingTime,
			SignedBlocksWindow:      int64(simulation.RandIntBetween(r, 10, 1000)),
			MinSignedPerWindow:      sdk.NewDecWithPrec(int64(r.Intn(10)), 1),
			DowntimeJailDuration:    time.Duration(simulation.RandIntBetween(r, 60, 60*60*24)) * time.Second,
			SlashFractionDoubleSign: sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(50) + 1))),
			SlashFractionDowntime:   sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(200) + 1))),
		},
	}
}

func getRandomMintGenesis(r *rand.Rand) mint.GenesisState {
	return mint.GenesisState{
		Minter: mint.InitialMinter(
			sdk.NewDecWithPrec(int64(r.Intn(99)), 2)),
		Params: mint.NewParams(
			sdk.DefaultBondDenom,
			sdk.NewDecWithPrec(int64(r.Intn(99)), 2),
			sdk.NewDecWithPrec(20, 2),
			sdk.NewDecWithPrec(7, 2),
			sdk.NewDecWithPrec(67, 2),
			uint64(60*60*8766/5)),
	}
}

func getRandomDistributionGenesis(r *rand.Rand) distribution.GenesisState {
	return distribution.GenesisState{
		FeePool:             distribution.InitialFeePool(),
		CommunityTax:        sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
		BaseProposerReward:  sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
		BonusProposerReward: sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
	}
}

func getRandomDepositGenesis(r *rand.Rand, accounts []simulation.Account) deposit.GenesisState {
	return deposit.GenesisState(vpnSim.GetRandomDeposits(r, accounts))
}

func getRandomVPNGenesis(r *rand.Rand, accounts []simulation.Account) vpn.GenesisState {
	return vpn.GenesisState{
		Nodes:         vpnSim.GetRandomNodes(r, accounts),
		Subscriptions: vpnSim.GetRandomSubscriptions(r, accounts),
		Sessions:      vpnSim.GetRandomSessions(r),
		Params: vpn.Params{
			FreeNodesCount:          uint64(r.Intn(50)),
			Deposit:                 sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(r.Intn(1000)))),
			NodeInactiveInterval:    int64(r.Intn(10)),
			SessionInactiveInterval: int64(r.Intn(10)),
		},
	}
}

func setAccounts(r *rand.Rand, genesisAccounts []GenesisAccount, accounts []simulation.Account, amount int64,
	initiallyBonded int64, timeStamp time.Time) ([]GenesisAccount, []simulation.Account) {

	for i, account := range accounts {
		coins := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(amount))}
		baseAccount := auth.NewBaseAccountWithAddress(account.Address)
		_ = baseAccount.SetCoins(coins)

		var genesisAccount GenesisAccount
		if int64(i) > initiallyBonded && r.Intn(100) < 50 {
			var (
				vestingAccount auth.VestingAccount
				endTime        int64
			)

			startTime := timeStamp.Unix()
			if r.Intn(100) < 50 {
				endTime = int64(simulation.RandIntBetween(r, int(startTime), int(startTime+(60*60*24*30))))
			} else {
				endTime = int64(simulation.RandIntBetween(r, int(startTime), int(startTime+(60*60*12))))
			}

			if startTime == endTime {
				endTime += 1
			}

			if r.Intn(100) < 50 {
				vestingAccount = auth.NewContinuousVestingAccount(&baseAccount, startTime, endTime)
			} else {
				vestingAccount = auth.NewDelayedVestingAccount(&baseAccount, endTime)
			}

			genesisAccount = NewGenesisAccount(vestingAccount)
		} else {
			genesisAccount = NewGenesisAccount(&baseAccount)
		}

		genesisAccounts = append(genesisAccounts, genesisAccount)
	}

	return genesisAccounts, accounts
}

func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
}

func getSimulation(tb testing.TB, app *HubApp) (testing.TB, *baseapp.BaseApp, simulation.AppStateFn, int64,
	simulation.WeightedOperations, sdk.Invariants, int, int, bool, bool) {

	weightedOps := setWeightedOperations(app)
	invariants := setInvariants(app)

	return tb, app.BaseApp, getRandomGenesisState, seed,
		weightedOps, invariants, blocksLen, blockSize, commit, false
}

func TestFullHubSimulation(t *testing.T) {
	var logger = log.NewNopLogger()
	var db tmDB.DB
	dir, _ := ioutil.TempDir("", "sentinel-hub-simulation.db")
	db, _ = sdk.NewLevelDB("Sentinel-Hub", dir)

	defer func() {
		db.Close()
		_ = os.RemoveAll(dir)
	}()

	app := NewHubApp(logger, db, nil, true, 0, fauxMerkleModeOpt)
	require.Equal(t, appName, app.Name())

	tb, _app, appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean := getSimulation(t, app)
	_, err := simulation.SimulateFromSeed(tb, _app, appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean)
	require.Nil(t, err)
}

func TestHubSimulationImportExport(t *testing.T) {
	var logger = log.NewNopLogger()
	var db1 tmDB.DB
	dir1, _ := ioutil.TempDir("", "sentinel-hub-simulation-import-export.db1")
	db1, _ = sdk.NewLevelDB("Sentinel-Hub-1", dir1)

	defer func() {
		db1.Close()
		_ = os.RemoveAll(dir1)
	}()

	app1 := NewHubApp(logger, db1, nil, true, 0, fauxMerkleModeOpt)
	require.Equal(t, appName, app1.Name())

	tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean := getSimulation(t, app1)
	_, err := simulation.SimulateFromSeed(tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean)
	require.Nil(t, err)

	fmt.Printf("Exporting genesis...\n")
	appState, _, err := app1.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err)

	fmt.Printf("Importing genesis...\n")
	dir2, _ := ioutil.TempDir("", "sentinel-hub-simulation-import-export.db2")
	db2, _ := sdk.NewLevelDB("Sentine-Hub-2", dir1)

	defer func() {
		db2.Close()
		_ = os.RemoveAll(dir2)
	}()

	app2 := NewHubApp(log.NewNopLogger(), db2, nil, true, 0, fauxMerkleModeOpt)
	require.Equal(t, appName, app2.Name())

	var genesisState GenesisState
	err = app1.cdc.UnmarshalJSON(appState, &genesisState)
	if err != nil {
		panic(err)
	}

	ctx2 := app2.NewContext(true, abci.Header{})
	app2.initFromGenesisState(ctx2, genesisState)

	fmt.Printf("Comparing stores...\n")
	ctx1 := app1.NewContext(true, abci.Header{})

	type StoreKeysPrefixes struct {
		key1     sdk.StoreKey
		key2     sdk.StoreKey
		Prefixes [][]byte
	}
	storeKeysPrefixes := []StoreKeysPrefixes{
		{app1.keyMain, app2.keyMain, [][]byte{}},
		{app1.keyAccount, app2.keyAccount, [][]byte{}},
		{app1.keyStaking, app2.keyStaking, [][]byte{staking.UnbondingQueueKey,
			staking.RedelegationQueueKey, staking.ValidatorQueueKey}},
		{app1.keySlashing, app2.keySlashing, [][]byte{}},
		{app1.keyMint, app2.keyMint, [][]byte{}},
		{app1.keyDistribution, app2.keyDistribution, [][]byte{}},
		{app1.keyFeeCollection, app2.keyFeeCollection, [][]byte{}},
		{app1.keyParams, app2.keyParams, [][]byte{}},
		{app1.keyGov, app2.keyGov, [][]byte{}},
		{app1.keyDeposit, app2.keyDeposit, [][]byte{deposit.DepositKeyPrefix}},
		{app1.keyVPNNode, app2.keyVPNNode, [][]byte{vpn.NodeKeyPrefix,
			vpn.NodeIDByAddressKeyPrefix, vpn.NodesCountOfAddressKeyPrefix}},
		{app1.keyVPNSubscription, app2.keyVPNSubscription, [][]byte{vpn.SubscriptionKeyPrefix,
			vpn.SubscriptionIDByAddressKeyPrefix, vpn.SubscriptionsCountOfAddressKeyPrefix,
			vpn.SubscriptionsCountOfNodeKeyPrefix, vpn.SubscriptionIDByNodeIDKeyPrefix}},
		{app1.keyVPNSession, app2.keyVPNSession, [][]byte{vpn.SessionKeyPrefix,
			vpn.SessionsCountOfSubscriptionKeyPrefix, vpn.SessionIDBySubscriptionIDKeyPrefix}},
	}

	for _, storeKeysPrefix := range storeKeysPrefixes {
		storeKey1 := storeKeysPrefix.key1
		storeKey2 := storeKeysPrefix.key2
		prefixes := storeKeysPrefix.Prefixes
		store1 := ctx1.KVStore(storeKey1)
		store2 := ctx2.KVStore(storeKey2)
		kv1, kv2, count, equal := sdk.DiffKVStores(store1, store2, prefixes)
		fmt.Println(kv1.Value, kv2.Value, count, equal)
		fmt.Printf("Compared %d key/value pairs between %s and %s\n", count, storeKey1, storeKey2)
		require.True(t, equal,
			"unequal stores: %s / %s:\nstore key1 %X => %X\nstore key2 %X => %X",
			storeKey1, storeKey2, kv1.Key, kv1.Value, kv2.Key, kv2.Value,
		)
	}
}

func TestHubSimulationAfterImport(t *testing.T) {
	var logger = log.NewNopLogger()
	var db1 tmDB.DB
	dir1, _ := ioutil.TempDir("", "sentinel-hub-simulation-after-import.db1")
	db1, _ = sdk.NewLevelDB("Sentinel-Hub-1", dir1)

	defer func() {
		db1.Close()
		_ = os.RemoveAll(dir1)
	}()

	app1 := NewHubApp(logger, db1, nil, true, 0, fauxMerkleModeOpt)
	require.Equal(t, appName, app1.Name())

	tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean := getSimulation(t, app1)
	_, err := simulation.SimulateFromSeed(tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean)
	require.Nil(t, err)

	fmt.Printf("Exporting genesis...\n")
	appState, _, err := app1.ExportAppStateAndValidators(true, []string{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Importing genesis...\n")
	dir2, _ := ioutil.TempDir("", "sentinel-hub-simulation-after-import.db2")
	db2, _ := sdk.NewLevelDB("Sentinel-Hub-2", dir1)

	defer func() {
		db2.Close()
		_ = os.RemoveAll(dir2)
	}()

	app2 := NewHubApp(log.NewNopLogger(), db2, nil, true, 0, fauxMerkleModeOpt)
	require.Equal(t, appName, app2.Name())

	app2.InitChain(abci.RequestInitChain{
		AppStateBytes: appState,
	})

	tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean = getSimulation(t, app2)
	_, err = simulation.SimulateFromSeed(tb, _app, _appState, seed, weightedOps, invariants, blocksLen, blockSize, commit, lean)
	require.Nil(t, err)
}

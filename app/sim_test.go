package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

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
	
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/staking/keeper"
	"github.com/sentinel-official/hub/x/vpn"
	vpnSim "github.com/sentinel-official/hub/x/vpn/simulation"
)

var (
	seed      int64
	numBlocks int
	blockSize int
	commit    bool
	period    int
)

func init() {
	flag.Int64Var(&seed, "seed", 42, "simulation random seed")
	flag.IntVar(&numBlocks, "num_blocks", 500, "number of blocks")
	flag.IntVar(&blockSize, "block_size", 200, "operations per block")
	flag.BoolVar(&commit, "commit", false, "have the simulation commit")
	flag.IntVar(&period, "period", 1, "run slow invariants only once every period assertions")
}

func getSimulateFromSeedInput(tb testing.TB, app *HubApp) (
	testing.TB, *baseapp.BaseApp, simulation.AppStateFn, int64,
	simulation.WeightedOperations, sdk.Invariants, int, int, bool, bool) {

	return tb, app.BaseApp, appStateFn, seed,
		testAndRunTxs(app), invariants(app), numBlocks, blockSize, commit, false
}

func appStateRandomizedFn(r *rand.Rand, accs []simulation.Account, genesisTimestamp time.Time) (
	json.RawMessage, []simulation.Account, string) {

	var genesisAccounts []GenesisAccount
	amount := int64(r.Intn(1000000000000))
	numInitiallyBonded := int64(r.Intn(250))
	numAccs := int64(len(accs))
	if numInitiallyBonded > numAccs {
		numInitiallyBonded = numAccs
	}
	fmt.Printf("Selected randomly generated parameters for simulated genesis:\n"+
		"\t{amount of stake per account: %v, initially bonded validators: %v}\n",
		amount, numInitiallyBonded)

	for i, acc := range accs {
		coins := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(amount))}
		bacc := auth.NewBaseAccountWithAddress(acc.Address)
		_ = bacc.SetCoins(coins)

		var gacc GenesisAccount
		if int64(i) > numInitiallyBonded && r.Intn(100) < 50 {
			var (
				vacc    auth.VestingAccount
				endTime int64
			)

			startTime := genesisTimestamp.Unix()

			if r.Intn(100) < 50 {
				endTime = int64(simulation.RandIntBetween(r, int(startTime), int(startTime+(60*60*24*30))))
			} else {
				endTime = int64(simulation.RandIntBetween(r, int(startTime), int(startTime+(60*60*12))))
			}

			if startTime == endTime {
				endTime += 1
			}

			if r.Intn(100) < 50 {
				vacc = auth.NewContinuousVestingAccount(&bacc, startTime, endTime)
			} else {
				vacc = auth.NewDelayedVestingAccount(&bacc, endTime)
			}
			gacc = NewGenesisAccount(vacc)
		} else {
			gacc = NewGenesisAccount(&bacc)
		}

		genesisAccounts = append(genesisAccounts, gacc)
	}

	authGenesis := auth.GenesisState{
		Params: auth.Params{
			MaxMemoCharacters:      uint64(simulation.RandIntBetween(r, 100, 200)),
			TxSigLimit:             uint64(r.Intn(7) + 1),
			TxSizeCostPerByte:      uint64(simulation.RandIntBetween(r, 5, 15)),
			SigVerifyCostED25519:   uint64(simulation.RandIntBetween(r, 500, 1000)),
			SigVerifyCostSecp256k1: uint64(simulation.RandIntBetween(r, 500, 1000)),
		},
	}

	bankGenesis := bank.NewGenesisState(r.Int63n(2) == 0)

	vp := time.Duration(r.Intn(2*172800)) * time.Second
	govGenesis := gov.GenesisState{
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

	stakingGenesis := staking.GenesisState{
		Pool: staking.InitialPool(),
		Params: staking.Params{
			UnbondingTime: time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second,
			MaxValidators: uint16(r.Intn(250) + 1),
			BondDenom:     sdk.DefaultBondDenom,
		},
	}

	slashingGenesis := slashing.GenesisState{
		Params: slashing.Params{
			MaxEvidenceAge:          stakingGenesis.Params.UnbondingTime,
			SignedBlocksWindow:      int64(simulation.RandIntBetween(r, 10, 1000)),
			MinSignedPerWindow:      sdk.NewDecWithPrec(int64(r.Intn(10)), 1),
			DowntimeJailDuration:    time.Duration(simulation.RandIntBetween(r, 60, 60*60*24)) * time.Second,
			SlashFractionDoubleSign: sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(50) + 1))),
			SlashFractionDowntime:   sdk.NewDec(1).Quo(sdk.NewDec(int64(r.Intn(200) + 1))),
		},
	}

	mintGenesis := mint.GenesisState{
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

	var validators []staking.Validator
	var delegations []staking.Delegation
	valAddrs := make([]sdk.ValAddress, numInitiallyBonded)
	for i := 0; i < int(numInitiallyBonded); i++ {
		valAddr := sdk.ValAddress(accs[i].Address)
		valAddrs[i] = valAddr

		validator := staking.NewValidator(valAddr, accs[i].PubKey, staking.Description{})
		validator.Tokens = sdk.NewInt(amount)
		validator.DelegatorShares = sdk.NewDec(amount)
		delegation := staking.Delegation{DelegatorAddress: accs[i].Address, ValidatorAddress: valAddr, Shares: sdk.NewDec(amount)}
		validators = append(validators, validator)
		delegations = append(delegations, delegation)
	}

	stakingGenesis.Pool.NotBondedTokens = sdk.NewInt((amount * numAccs) + (numInitiallyBonded * amount))
	stakingGenesis.Validators = validators
	stakingGenesis.Delegations = delegations

	distributionGenesis := distribution.GenesisState{
		FeePool:             distribution.InitialFeePool(),
		CommunityTax:        sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
		BaseProposerReward:  sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
		BonusProposerReward: sdk.NewDecWithPrec(1, 2).Add(sdk.NewDecWithPrec(int64(r.Intn(30)), 2)),
	}

	depositGenesis := deposit.GenesisState(vpnSim.GetRandomDeposits(r, accs))

	vpnGenesis := vpn.GenesisState{
		Nodes:         vpnSim.GetRandomNodes(r, accs),
		Subscriptions: vpnSim.GetRandomSubscriptions(r, accs),
		Sessions:      vpnSim.GetRandomSessions(r, accs),
		Params: vpn.Params{
			FreeNodesCount:          uint64(r.Intn(50)),
			Deposit:                 sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(r.Intn(1000)))),
			NodeInactiveInterval:    int64(r.Intn(10)),
			SessionInactiveInterval: int64(r.Intn(10)),
		},
	}

	genesis := GenesisState{
		Accounts:     genesisAccounts,
		Auth:         authGenesis,
		Bank:         bankGenesis,
		Staking:      stakingGenesis,
		Mint:         mintGenesis,
		Distribution: distributionGenesis,
		Slashing:     slashingGenesis,
		Gov:          govGenesis,

		Deposit: depositGenesis,
		VPN:     vpnGenesis,
	}

	appState, err := MakeCodec().MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return appState, accs, "hub"
}

func appStateFn(r *rand.Rand, accs []simulation.Account, genesisTimestamp time.Time) (json.RawMessage, []simulation.Account, string) {

	return appStateRandomizedFn(r, accs, genesisTimestamp)
}

func testAndRunTxs(app *HubApp) []simulation.WeightedOperation {

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

func invariants(app *HubApp) []sdk.Invariant {

	return []sdk.Invariant{
		simulation.PeriodicInvariant(bank.NonnegativeBalanceInvariant(app.accountKeeper), period, 0),
		simulation.PeriodicInvariant(distribution.AllInvariants(app.distributionKeeper, app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(keeper.SupplyInvariants(app.stakingKeeper, app.feeCollectionKeeper,
			app.distributionKeeper, app.accountKeeper, app.depositKeeper), period, 0),
		simulation.PeriodicInvariant(staking.NonNegativePowerInvariant(app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(staking.DelegatorSharesInvariant(app.stakingKeeper), period, 0),
	}
}

func fauxMerkleModeOpt(bapp *baseapp.BaseApp) {
	bapp.SetFauxMerkleMode()
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

	_, err := simulation.SimulateFromSeed(getSimulateFromSeedInput(t, app))
	require.Nil(t, err)
}

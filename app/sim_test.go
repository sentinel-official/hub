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
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/deposit"
	_staking "github.com/sentinel-official/hub/x/staking"
	"github.com/sentinel-official/hub/x/vpn"
	vpnSim "github.com/sentinel-official/hub/x/vpn/simulation"
)

var (
	seed      int64
	numBlocks int
	blockSize int
	enabled   bool
	commit    bool
	verbose   bool
	period    int
)

func init() {
	flag.Int64Var(&seed, "seed", 42, "simulation random seed")
	flag.IntVar(&numBlocks, "num_blocks", 500, "number of blocks")
	flag.IntVar(&blockSize, "block_size", 200, "operations per block")
	flag.BoolVar(&enabled, "enable", false, "enable the simulation")
	flag.BoolVar(&verbose, "verbose", false, "verbose log output")
	flag.BoolVar(&commit, "commit", true, "have the simulation commit")
	flag.IntVar(&period, "period", 1, "run slow invariants only once every period assertions")
}

func getGenesisAccounts(r *rand.Rand, accounts []simulation.Account, amount, bondedAccountsLen int64,
	timeStamp time.Time) []GenesisAccount {

	genesisAccounts := make([]GenesisAccount, 0, len(accounts))
	for i, account := range accounts {
		coins := sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(amount))}
		baseAccount := auth.NewBaseAccountWithAddress(account.Address)
		if err := baseAccount.SetCoins(coins); err != nil {
			panic(err)
		}

		var genesisAccount GenesisAccount
		if int64(i) > bondedAccountsLen && r.Intn(100) < 50 {
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
				endTime++
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

	return genesisAccounts
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

func getRandomStakingGenesis(r *rand.Rand) staking.GenesisState {
	return staking.GenesisState{
		Pool: staking.InitialPool(),
		Params: staking.Params{
			UnbondingTime: time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second,
			MaxValidators: uint16(r.Intn(250) + 1),
			BondDenom:     sdk.DefaultBondDenom,
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

func getRandomGovGenesis(r *rand.Rand) gov.GenesisState {
	p := time.Duration(r.Intn(2*172800)) * time.Second

	return gov.GenesisState{
		StartingProposalID: uint64(r.Intn(100)),
		DepositParams: gov.DepositParams{
			MinDeposit:       sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(r.Intn(1e3)))},
			MaxDepositPeriod: p,
		},
		VotingParams: gov.VotingParams{
			VotingPeriod: p,
		},
		TallyParams: gov.TallyParams{
			Quorum:    sdk.NewDecWithPrec(334, 3),
			Threshold: sdk.NewDecWithPrec(5, 1),
			Veto:      sdk.NewDecWithPrec(334, 3),
		},
	}
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

func getRandomDepositGenesis() deposit.GenesisState {
	return deposit.GenesisState{}
}

func getRandomVPNGenesis() vpn.GenesisState {
	r := rand.New(rand.NewSource(seed))
	return vpn.GenesisState{
		Params: vpn.Params{
			FreeNodesCount:          uint64(r.Intn(50)),
			Deposit:                 sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(int64(r.Intn(1000)))),
			NodeInactiveInterval:    int64(r.Intn(10)),
			SessionInactiveInterval: int64(r.Intn(10)),
		},
	}
}

func getRandomGenesisState(r *rand.Rand, accounts []simulation.Account,
	genesisTimestamp time.Time) (json.RawMessage, []simulation.Account, string) {

	amount := int64(r.Intn(1000000000000))
	accountsLen := int64(len(accounts))

	bondedAccountsLen := int64(r.Intn(250))
	if bondedAccountsLen > accountsLen {
		bondedAccountsLen = accountsLen
	}

	fmt.Printf("Selected randomly generated parameters for simulated genesis:\n"+
		"\t{amount of stake per account: %v, initially bonded validators: %v}\n",
		amount, bondedAccountsLen)

	genesisAccounts := getGenesisAccounts(r, accounts, amount, bondedAccountsLen, genesisTimestamp)

	authGenesis := getRandomAuthGenesis(r)
	fmt.Printf("Selected randomly generated auth parameters:\n\t%+v\n", authGenesis)

	bankGenesis := getRandomBankGenesis(r)
	fmt.Printf("Selected randomly generated bank parameters:\n\t%+v\n", bankGenesis)

	govGenesis := getRandomGovGenesis(r)
	fmt.Printf("Selected randomly generated governance parameters:\n\t%+v\n", govGenesis)

	stakingGenesis := getRandomStakingGenesis(r)
	fmt.Printf("Selected randomly generated staking parameters:\n\t%+v\n", stakingGenesis)

	slashingGenesis := getRandomSlashingGenesis(r, stakingGenesis)
	fmt.Printf("Selected randomly generated slashing parameters:\n\t%+v\n", slashingGenesis)

	mintGenesis := getRandomMintGenesis(r)
	fmt.Printf("Selected randomly generated minting parameters:\n\t%+v\n", mintGenesis)

	validators := make([]staking.Validator, 0, bondedAccountsLen)
	delegations := make([]staking.Delegation, 0, bondedAccountsLen)
	addresses := make([]sdk.ValAddress, bondedAccountsLen)
	for i := 0; i < int(bondedAccountsLen); i++ {
		address := sdk.ValAddress(accounts[i].Address)
		addresses[i] = address

		validator := staking.NewValidator(address, accounts[i].PubKey, staking.Description{})
		validator.Tokens = sdk.NewInt(amount)
		validator.DelegatorShares = sdk.NewDec(amount)
		delegation := staking.Delegation{
			DelegatorAddress: accounts[i].Address,
			ValidatorAddress: address,
			Shares:           sdk.NewDec(amount),
		}
		validators = append(validators, validator)
		delegations = append(delegations, delegation)
	}

	stakingGenesis.Pool.NotBondedTokens = sdk.NewInt((amount * int64(len(accounts))) + (bondedAccountsLen * amount))
	stakingGenesis.Validators = validators
	stakingGenesis.Delegations = delegations

	distributionGenesis := getRandomDistributionGenesis(r)
	fmt.Printf("Selected randomly generated distribution parameters:\n\t%+v\n", distributionGenesis)

	depositGenesis := getRandomDepositGenesis()
	fmt.Printf("Selected randomly generated deposit parameters:\n\t%+v\n", depositGenesis)

	vpnGenesis := getRandomVPNGenesis()
	fmt.Printf("Selected randomly generated vpn parameters:\n\t%+v\n", vpnGenesis)

	genesis := GenesisState{
		Accounts:     genesisAccounts,
		Auth:         authGenesis,
		Bank:         bankGenesis,
		Staking:      stakingGenesis,
		Mint:         mintGenesis,
		Distribution: distributionGenesis,
		Gov:          govGenesis,
		Slashing:     slashingGenesis,
		Deposit:      depositGenesis,
		VPN:          vpnGenesis,
	}

	appState, err := MakeCodec().MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return appState, accounts, "simulation"
}

func getWeightedOperations(app *HubApp) []simulation.WeightedOperation {
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
		{25, vpnSim.SimulateMsgRegisterNode(app.vpnKeeper)},
		{50, vpnSim.SimulateMsgUpdateNodeInfo(app.vpnKeeper)},
		{75, vpnSim.SimulateMsgUpdateNodeStatus(app.vpnKeeper)},
		{100, vpnSim.SimulateMsgStartSubscription(app.vpnKeeper)},
		{100, vpnSim.SimulateMsgUpdateSessionInfo(app.vpnKeeper)},
		{50, vpnSim.SimulateMsgEndSubscription(app.vpnKeeper)},
	}
}

func getInvariants(app *HubApp) []sdk.Invariant {
	return []sdk.Invariant{
		simulation.PeriodicInvariant(bank.NonnegativeBalanceInvariant(app.accountKeeper), period, 0),
		simulation.PeriodicInvariant(distribution.AllInvariants(app.distributionKeeper, app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(_staking.SupplyInvariants(app.stakingKeeper, app.feeCollectionKeeper,
			app.distributionKeeper, app.accountKeeper, app.depositKeeper), period, 0),
		simulation.PeriodicInvariant(staking.NonNegativePowerInvariant(app.stakingKeeper), period, 0),
		simulation.PeriodicInvariant(staking.DelegatorSharesInvariant(app.stakingKeeper), period, 0),
	}
}

func fauxMerkleModeOpt(baseApp *baseapp.BaseApp) {
	baseApp.SetFauxMerkleMode()
}

func TestFullHubSimulation(t *testing.T) {
	if !enabled {
		t.Skip("skipping hub simulation")
	}

	dir, err := ioutil.TempDir("", "sentinel_hub_simulation_")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = os.RemoveAll(dir); err != nil {
			panic(err)
		}
	}()

	db, err := sdk.NewLevelDB("sentinel_hub", dir)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var logger log.Logger
	if verbose {
		logger = log.TestingLogger()
	} else {
		logger = log.NewNopLogger()
	}

	app := NewHubApp(logger, db, nil, true, 0, fauxMerkleModeOpt)
	_, err = simulation.SimulateFromSeed(t, app.BaseApp, getRandomGenesisState, seed,
		getWeightedOperations(app), getInvariants(app), numBlocks, blockSize, commit, false)
	require.Nil(t, err)
}

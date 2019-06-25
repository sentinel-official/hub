package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	tm "github.com/tendermint/tendermint/types"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn"
)

type GenesisAccount struct {
	Address          sdk.AccAddress `json:"address"`
	Coins            sdk.Coins      `json:"coins"`
	Sequence         uint64         `json:"sequence_number"`
	AccountNumber    uint64         `json:"account_number"`
	OriginalVesting  sdk.Coins      `json:"original_vesting"`
	DelegatedFree    sdk.Coins      `json:"delegated_free"`
	DelegatedVesting sdk.Coins      `json:"delegated_vesting"`
	StartTime        int64          `json:"start_time"`
	EndTime          int64          `json:"end_time"`
}

func NewGenesisAccountFromBaseAccount(acc *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		Address:       acc.Address,
		Coins:         acc.Coins,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}
}

func NewGenesisAccount(acc auth.Account) GenesisAccount {
	genesisAccount := GenesisAccount{
		Address:       acc.GetAddress(),
		Coins:         acc.GetCoins(),
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
	}

	vestingAccount, ok := acc.(auth.VestingAccount)
	if ok {
		genesisAccount.OriginalVesting = vestingAccount.GetOriginalVesting()
		genesisAccount.DelegatedFree = vestingAccount.GetDelegatedFree()
		genesisAccount.DelegatedVesting = vestingAccount.GetDelegatedVesting()
		genesisAccount.StartTime = vestingAccount.GetStartTime()
		genesisAccount.EndTime = vestingAccount.GetEndTime()
	}

	return genesisAccount
}

func (ga *GenesisAccount) ToAccount() auth.Account {
	baseAccount := &auth.BaseAccount{
		Address:       ga.Address,
		Coins:         ga.Coins.Sort(),
		AccountNumber: ga.AccountNumber,
		Sequence:      ga.Sequence,
	}

	if !ga.OriginalVesting.IsZero() {
		baseVestingAccount := &auth.BaseVestingAccount{
			BaseAccount:      baseAccount,
			OriginalVesting:  ga.OriginalVesting,
			DelegatedFree:    ga.DelegatedFree,
			DelegatedVesting: ga.DelegatedVesting,
			EndTime:          ga.EndTime,
		}

		if ga.StartTime != 0 && ga.EndTime != 0 {
			return &auth.ContinuousVestingAccount{
				BaseVestingAccount: baseVestingAccount,
				StartTime:          ga.StartTime,
			}
		} else if ga.EndTime != 0 {
			return &auth.DelayedVestingAccount{
				BaseVestingAccount: baseVestingAccount,
			}
		} else {
			panic(fmt.Sprintf("invalid genesis vesting account: %+v", ga))
		}
	}

	return baseAccount
}

type GenesisState struct {
	Accounts     []GenesisAccount          `json:"accounts"`
	Auth         auth.GenesisState         `json:"auth"`
	Bank         bank.GenesisState         `json:"bank"`
	Staking      staking.GenesisState      `json:"staking"`
	Mint         mint.GenesisState         `json:"mint"`
	Distribution distribution.GenesisState `json:"distribution"`
	Gov          gov.GenesisState          `json:"gov"`
	Crisis       crisis.GenesisState       `json:"crisis"`
	Slashing     slashing.GenesisState     `json:"slashing"`
	Deposit      deposit.GenesisState      `json:"deposit"`
	VPN          vpn.GenesisState          `json:"vpn"`
	GenTxs       []json.RawMessage         `json:"gen_txs"`
}

func NewGenesisState(accounts []GenesisAccount, _auth auth.GenesisState, _bank bank.GenesisState,
	_staking staking.GenesisState, _mint mint.GenesisState, _distribution distribution.GenesisState,
	_gov gov.GenesisState, _crisis crisis.GenesisState, _slashing slashing.GenesisState,
	_deposit deposit.GenesisState, _vpn vpn.GenesisState) GenesisState {

	return GenesisState{
		Accounts:     accounts,
		Auth:         _auth,
		Bank:         _bank,
		Staking:      _staking,
		Mint:         _mint,
		Distribution: _distribution,
		Gov:          _gov,
		Crisis:       _crisis,
		Slashing:     _slashing,
		Deposit:      _deposit,
		VPN:          _vpn,
	}
}

func (gs GenesisState) Sanitize() {
	sort.Slice(gs.Accounts, func(i, j int) bool {
		return gs.Accounts[i].AccountNumber < gs.Accounts[j].AccountNumber
	})

	for _, acc := range gs.Accounts {
		acc.Coins = acc.Coins.Sort()
	}

	sort.Slice(gs.Deposit, func(i, j int) bool {
		return gs.Deposit[i].Address.String() < gs.Deposit[j].Address.String()
	})

	for _, dep := range gs.Deposit {
		dep.Coins = dep.Coins.Sort()
	}
}

func validateGenesisStateAccounts(accounts []GenesisAccount) error {
	addrMap := make(map[string]bool, len(accounts))
	for _, acc := range accounts {
		addrStr := acc.Address.String()

		if _, ok := addrMap[addrStr]; ok {
			return fmt.Errorf("duplicate account found in genesis state; address: %s", addrStr)
		}

		if !acc.OriginalVesting.IsZero() {
			if acc.EndTime == 0 {
				return fmt.Errorf("missing end time for vesting account; address: %s", addrStr)
			}

			if acc.StartTime >= acc.EndTime {
				return fmt.Errorf(
					"vesting start time must before end time; address: %s, start: %s, end: %s",
					addrStr,
					time.Unix(acc.StartTime, 0).UTC().Format(time.RFC3339),
					time.Unix(acc.EndTime, 0).UTC().Format(time.RFC3339),
				)
			}
		}

		addrMap[addrStr] = true
	}

	return nil
}

func NewGenesisStateFromGenesisDoc(cdc *codec.Codec, genesisDoc tm.GenesisDoc,
	genTxs []json.RawMessage) (state GenesisState, err error) {

	if err = cdc.UnmarshalJSON(genesisDoc.AppState, &state); err != nil {
		return state, err
	}

	if len(genTxs) == 0 {
		return state, errors.New("there must be at least one genesis tx")
	}

	_staking := state.Staking
	for i, genTx := range genTxs {
		var tx auth.StdTx
		if err := cdc.UnmarshalJSON(genTx, &tx); err != nil {
			return state, err
		}

		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return state, errors.New(
				"must provide genesis StdTx with exactly 1 CreateValidator message")
		}

		if _, ok := msgs[0].(staking.MsgCreateValidator); !ok {
			return state, fmt.Errorf(
				"genesis transaction %v does not contain a MsgCreateValidator", i)
		}
	}

	for _, acc := range state.Accounts {
		for _, coin := range acc.Coins {
			if coin.Denom == state.Staking.Params.BondDenom {
				_staking.Pool.NotBondedTokens = _staking.Pool.NotBondedTokens.Add(coin.Amount)
			}
		}
	}

	state.Staking = _staking
	state.GenTxs = genTxs

	return state, nil
}

func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		Accounts:     nil,
		Auth:         auth.DefaultGenesisState(),
		Bank:         bank.DefaultGenesisState(),
		Staking:      staking.DefaultGenesisState(),
		Mint:         mint.DefaultGenesisState(),
		Distribution: distribution.DefaultGenesisState(),
		Gov:          gov.DefaultGenesisState(),
		Crisis:       crisis.DefaultGenesisState(),
		Slashing:     slashing.DefaultGenesisState(),
		Deposit:      deposit.DefaultGenesisState(),
		VPN:          vpn.DefaultGenesisState(),
		GenTxs:       nil,
	}
}

// nolint:gocyclo
func ValidateGenesisState(state GenesisState) error {
	if err := validateGenesisStateAccounts(state.Accounts); err != nil {
		return err
	}
	if err := auth.ValidateGenesis(state.Auth); err != nil {
		return err
	}
	if err := bank.ValidateGenesis(state.Bank); err != nil {
		return err
	}
	if err := staking.ValidateGenesis(state.Staking); err != nil {
		return err
	}
	if err := mint.ValidateGenesis(state.Mint); err != nil {
		return err
	}
	if err := distribution.ValidateGenesis(state.Distribution); err != nil {
		return err
	}
	if err := gov.ValidateGenesis(state.Gov); err != nil {
		return err
	}
	if err := crisis.ValidateGenesis(state.Crisis); err != nil {
		return err
	}
	if err := slashing.ValidateGenesis(state.Slashing); err != nil {
		return err
	}
	if err := deposit.ValidateGenesis(state.Deposit); err != nil {
		return err
	}

	return vpn.ValidateGenesis(state.VPN)
}

// nolint:gocyclo
func CollectStdTxs(cdc *codec.Codec, moniker string, genTxsDir string,
	genDoc tm.GenesisDoc) (appGenTxs []auth.StdTx, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	var state GenesisState
	if err = cdc.UnmarshalJSON(genDoc.AppState, &state); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]GenesisAccount, len(state.Accounts))
	for i := 0; i < len(state.Accounts); i++ {
		acc := state.Accounts[i]
		addrMap[acc.Address.String()] = acc
	}

	var addressesIPs []string

	for _, fo := range fos {
		filename := filepath.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (filepath.Ext(filename) != ".json") {
			continue
		}

		var jsonRawTx []byte
		if jsonRawTx, err = ioutil.ReadFile(filename); err != nil {
			return appGenTxs, persistentPeers, err
		}

		var genStdTx auth.StdTx
		if err = cdc.UnmarshalJSON(jsonRawTx, &genStdTx); err != nil {
			return appGenTxs, persistentPeers, err
		}

		appGenTxs = append(appGenTxs, genStdTx)

		nodeAddrIP := genStdTx.GetMemo()
		if len(nodeAddrIP) == 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"couldn't find node's address and IP in %s", fo.Name())
		}

		msgs := genStdTx.GetMsgs()
		if len(msgs) != 1 {
			return appGenTxs, persistentPeers, errors.New(
				"each genesis transaction must provide a single genesis message")
		}

		msg, ok := msgs[0].(staking.MsgCreateValidator)
		if !ok {
			return
		}

		delAddr := msg.DelegatorAddress.String()
		valAddr := sdk.AccAddress(msg.ValidatorAddress).String()

		delAcc, delOk := addrMap[delAddr]
		_, valOk := addrMap[valAddr]

		var accountsNotInGeneses []string
		if !delOk {
			accountsNotInGeneses = append(accountsNotInGeneses, delAddr)
		}
		if !valOk {
			accountsNotInGeneses = append(accountsNotInGeneses, valAddr)
		}
		if len(accountsNotInGeneses) != 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"account(s) %v not in genesis.json: %+v", strings.Join(accountsNotInGeneses, " "), addrMap)
		}

		if delAcc.Coins.AmountOf(msg.Value.Denom).LT(msg.Value.Amount) {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"insufficient fund for delegation %v: %v < %v",
				delAcc.Address, delAcc.Coins.AmountOf(msg.Value.Denom), msg.Value.Amount,
			)
		}

		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}

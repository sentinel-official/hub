package hub

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
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	tm "github.com/tendermint/tendermint/types"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

type GenesisState struct {
	Accounts         []GenesisAccount          `json:"accounts"`
	AuthData         auth.GenesisState         `json:"auth_data"`
	BankData         bank.GenesisState         `json:"bank_data"`
	StakingData      staking.GenesisState      `json:"staking_data"`
	MintData         mint.GenesisState         `json:"mint_data"`
	DistributionData distribution.GenesisState `json:"distribution_data"`
	GovData          gov.GenesisState          `json:"gov_data"`
	CrisisData       crisis.GenesisState       `json:"crisis_data"`
	SlashingData     slashing.GenesisState     `json:"slashing_data"`
	DepositData      deposit.GenesisState      `json:"deposit_data"`
	VPNData          vpn.GenesisState          `json:"vpn_data"`
	GenTxs           []json.RawMessage         `json:"gen_txs"`
}

func NewGenesisState(accounts []GenesisAccount, authData auth.GenesisState, bankData bank.GenesisState,
	stakingData staking.GenesisState, mintData mint.GenesisState, distributionData distribution.GenesisState,
	govData gov.GenesisState, crisisData crisis.GenesisState, slashingData slashing.GenesisState,
	depositData deposit.GenesisState, vpnData vpn.GenesisState) GenesisState {

	return GenesisState{
		Accounts:         accounts,
		AuthData:         authData,
		BankData:         bankData,
		StakingData:      stakingData,
		MintData:         mintData,
		DistributionData: distributionData,
		GovData:          govData,
		CrisisData:       crisisData,
		SlashingData:     slashingData,
		DepositData:      depositData,
		VPNData:          vpnData,
	}
}

func (gs GenesisState) Sanitize() {
	sort.Slice(gs.Accounts, func(i, j int) bool {
		return gs.Accounts[i].AccountNumber < gs.Accounts[j].AccountNumber
	})

	for _, acc := range gs.Accounts {
		acc.Coins = acc.Coins.Sort()
	}

	sort.Slice(gs.DepositData, func(i, j int) bool {
		return gs.DepositData[i].Address.String() < gs.DepositData[j].Address.String()
	})

	for _, dep := range gs.DepositData {
		dep.Coins = dep.Coins.Sort()
	}

	sort.Slice(gs.VPNData.Nodes, func(i, j int) bool {
		return gs.VPNData.Nodes[i].ID < gs.VPNData.Nodes[j].ID
	})

	sort.Slice(gs.VPNData.Sessions, func(i, j int) bool {
		return gs.VPNData.Sessions[i].ID < gs.VPNData.Sessions[j].ID
	})
}

type GenesisAccount struct {
	Address          csdk.AccAddress `json:"address"`
	Coins            csdk.Coins      `json:"coins"`
	Sequence         uint64          `json:"sequence_number"`
	AccountNumber    uint64          `json:"account_number"`
	OriginalVesting  csdk.Coins      `json:"original_vesting"`
	DelegatedFree    csdk.Coins      `json:"delegated_free"`
	DelegatedVesting csdk.Coins      `json:"delegated_vesting"`
	StartTime        int64           `json:"start_time"`
	EndTime          int64           `json:"end_time"`
}

func NewGenesisAccount(acc *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		Address:       acc.Address,
		Coins:         acc.Coins,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}
}

func NewGenesisAccountI(acc auth.Account) GenesisAccount {
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
		baseVestingAcc := &auth.BaseVestingAccount{
			BaseAccount:      baseAccount,
			OriginalVesting:  ga.OriginalVesting,
			DelegatedFree:    ga.DelegatedFree,
			DelegatedVesting: ga.DelegatedVesting,
			EndTime:          ga.EndTime,
		}

		if ga.StartTime != 0 && ga.EndTime != 0 {
			return &auth.ContinuousVestingAccount{
				BaseVestingAccount: baseVestingAcc,
				StartTime:          ga.StartTime,
			}
		} else if ga.EndTime != 0 {
			return &auth.DelayedVestingAccount{
				BaseVestingAccount: baseVestingAcc,
			}
		} else {
			panic(fmt.Sprintf("invalid genesis vesting account: %+v", ga))
		}
	}

	return baseAccount
}

func GenGenesisState(cdc *codec.Codec, genesisDoc tm.GenesisDoc,
	appGenTxs []json.RawMessage) (genesisState GenesisState, err error) {

	if err = cdc.UnmarshalJSON(genesisDoc.AppState, &genesisState); err != nil {
		return genesisState, err
	}

	if len(appGenTxs) == 0 {
		return genesisState, errors.New("there must be at least one genesis tx")
	}

	stakingData := genesisState.StakingData
	for i, genTx := range appGenTxs {
		var tx auth.StdTx
		if err := cdc.UnmarshalJSON(genTx, &tx); err != nil {
			return genesisState, err
		}

		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return genesisState, errors.New(
				"must provide genesis StdTx with exactly 1 CreateValidator message")
		}

		if _, ok := msgs[0].(staking.MsgCreateValidator); !ok {
			return genesisState, fmt.Errorf(
				"genesis transaction %v does not contain a MsgCreateValidator", i)
		}
	}

	for _, acc := range genesisState.Accounts {
		for _, coin := range acc.Coins {
			if coin.Denom == genesisState.StakingData.Params.BondDenom {
				stakingData.Pool.NotBondedTokens = stakingData.Pool.NotBondedTokens.Add(coin.Amount)
			}
		}
	}

	genesisState.StakingData = stakingData
	genesisState.GenTxs = appGenTxs

	return genesisState, nil
}

func NewDefaultGenesisState() GenesisState {
	state := GenesisState{
		Accounts:         nil,
		AuthData:         auth.DefaultGenesisState(),
		BankData:         bank.DefaultGenesisState(),
		StakingData:      staking.DefaultGenesisState(),
		MintData:         mint.DefaultGenesisState(),
		DistributionData: distribution.DefaultGenesisState(),
		GovData:          gov.DefaultGenesisState(),
		CrisisData:       crisis.DefaultGenesisState(),
		SlashingData:     slashing.DefaultGenesisState(),
		DepositData:      deposit.DefaultGenesisState(),
		VPNData:          vpn.DefaultGenesisState(),
		GenTxs:           nil,
	}

	return state
}

func ValidateGenesisState(genesisState GenesisState) error {
	if err := validateGenesisStateAccounts(genesisState.Accounts); err != nil {
		return err
	}
	if len(genesisState.GenTxs) > 0 {
		return nil
	}
	if err := auth.ValidateGenesis(genesisState.AuthData); err != nil {
		return err
	}
	if err := bank.ValidateGenesis(genesisState.BankData); err != nil {
		return err
	}
	if err := staking.ValidateGenesis(genesisState.StakingData); err != nil {
		return err
	}
	if err := slashing.ValidateGenesis(genesisState.SlashingData); err != nil {
		return err
	}
	if err := distribution.ValidateGenesis(genesisState.DistributionData); err != nil {
		return err
	}
	if err := gov.ValidateGenesis(genesisState.GovData); err != nil {
		return err
	}
	if err := mint.ValidateGenesis(genesisState.MintData); err != nil {
		return err
	}
	if err := crisis.ValidateGenesis(genesisState.CrisisData); err != nil {
		return err
	}
	if err := deposit.ValidateGenesis(genesisState.DepositData); err != nil {
		return err
	}

	return vpn.ValidateGenesis(genesisState.VPNData)
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

func GenGenesisStateJSON(cdc *codec.Codec, genDoc tm.GenesisDoc,
	appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	genesisState, err := GenGenesisState(cdc, genDoc, appGenTxs)
	if err != nil {
		return nil, err
	}

	return codec.MarshalJSONIndent(cdc, genesisState)
}

func CollectStdTxs(cdc *codec.Codec, moniker string, genTxsDir string,
	genDoc tm.GenesisDoc) (appGenTxs []auth.StdTx, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	var appState GenesisState
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]GenesisAccount, len(appState.Accounts))
	for i := 0; i < len(appState.Accounts); i++ {
		acc := appState.Accounts[i]
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

		msg := msgs[0].(staking.MsgCreateValidator)
		delAddr := msg.DelegatorAddress.String()
		valAddr := csdk.AccAddress(msg.ValidatorAddress).String()

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

package vpn

import (
	"encoding/json"
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

type GenesisState struct {
	Accounts     []GenesisAccount          `json:"accounts"`
	Auth         auth.GenesisState         `json:"auth"`
	Staking      staking.GenesisState      `json:"staking"`
	Slashing     slashing.GenesisState     `json:"slashing"`
	Distribution distribution.GenesisState `json:"distribution"`
	Gov          gov.GenesisState          `json:"gov"`
	Mint         mint.GenesisState         `json:"mint"`
	GenTxs       []json.RawMessage         `json:"gen_txs"`
}

func NewGenesisState(accounts []GenesisAccount, authData auth.GenesisState,
	stakingData staking.GenesisState, slashingData slashing.GenesisState,
	distrData distribution.GenesisState, govData gov.GenesisState, mintData mint.GenesisState) GenesisState {

	return GenesisState{
		Accounts:     accounts,
		Auth:         authData,
		Staking:      stakingData,
		Slashing:     slashingData,
		Distribution: distrData,
		Gov:          govData,
		Mint:         mintData,
	}
}

type GenesisAccount struct {
	Address       csdkTypes.AccAddress `json:"address"`
	Coins         csdkTypes.Coins      `json:"coins"`
	Sequence      uint64               `json:"sequence_number"`
	AccountNumber uint64               `json:"account_number"`

	OriginalVesting  csdkTypes.Coins `json:"original_vesting"`
	DelegatedFree    csdkTypes.Coins `json:"delegated_free"`
	DelegatedVesting csdkTypes.Coins `json:"delegated_vesting"`
	StartTime        int64           `json:"start_time"`
	EndTime          int64           `json:"end_time"`
}

func NewGenesisAccountI(acc auth.Account) GenesisAccount {
	gacc := GenesisAccount{
		Address:       acc.GetAddress(),
		Coins:         acc.GetCoins(),
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
	}

	vacc, ok := acc.(auth.VestingAccount)
	if ok {
		gacc.OriginalVesting = vacc.GetOriginalVesting()
		gacc.DelegatedFree = vacc.GetDelegatedFree()
		gacc.DelegatedVesting = vacc.GetDelegatedVesting()
		gacc.StartTime = vacc.GetStartTime()
		gacc.EndTime = vacc.GetEndTime()
	}

	return gacc
}

func (ga *GenesisAccount) ToAccount() auth.Account {
	bacc := &auth.BaseAccount{
		Address:       ga.Address,
		Coins:         ga.Coins.Sort(),
		AccountNumber: ga.AccountNumber,
		Sequence:      ga.Sequence,
	}

	if !ga.OriginalVesting.IsZero() {
		baseVestingAcc := &auth.BaseVestingAccount{
			BaseAccount:      bacc,
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

	return bacc
}

func HubValidateGenesisState(genesisState GenesisState) error {
	if err := validateGenesisStateAccounts(genesisState.Accounts); err != nil {
		return err
	}
	if err := auth.ValidateGenesis(genesisState.Auth); err != nil {
		return err
	}
	if err := staking.ValidateGenesis(genesisState.Staking); err != nil {
		return err
	}
	if err := slashing.ValidateGenesis(genesisState.Slashing); err != nil {
		return err
	}
	if err := distribution.ValidateGenesis(genesisState.Distribution); err != nil {
		return err
	}
	if err := gov.ValidateGenesis(genesisState.Gov); err != nil {
		return err
	}
	if err := mint.ValidateGenesis(genesisState.Mint); err != nil {
		return err
	}
	if len(genesisState.GenTxs) > 0 {
		return nil
	}

	return nil
}

func validateGenesisStateAccounts(accs []GenesisAccount) error {
	addrMap := make(map[string]bool, len(accs))
	for i := 0; i < len(accs); i++ {
		acc := accs[i]
		strAddr := string(acc.Address)
		if _, ok := addrMap[strAddr]; ok {
			return fmt.Errorf("Duplicate account in genesis state: Address %v", acc.Address)
		}
		addrMap[strAddr] = true
	}
	return nil
}

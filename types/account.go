package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

var _ auth.Account = (*AppAccount)(nil)

type AppAccount struct {
	auth.BaseAccount

	Name string `json:"name"`
}

func (acc AppAccount) GetName() string {
	return acc.Name
}

func (acc *AppAccount) SetName(name string) {
	acc.Name = name
}

func NewAppAccount(name string, baseAcct auth.BaseAccount) *AppAccount {
	return &AppAccount{
		BaseAccount: baseAcct,
		Name:        name,
	}
}

func GetAccountDecoder(cdc *codec.Codec) auth.AccountDecoder {
	return func(accBytes []byte) (auth.Account, error) {
		if len(accBytes) == 0 {
			return nil, ccsdkTypes.ErrTxDecode("accBytes are empty")
		}

		acct := new(AppAccount)
		err := cdc.UnmarshalBinaryBare(accBytes, &acct)
		if err != nil {
			panic(err)
		}

		return acct, err
	}
}

type GenesisState struct {
	Accounts []*GenesisAccount `json:"accounts"`
}

type GenesisAccount struct {
	Name    string                `json:"name"`
	Address ccsdkTypes.AccAddress `json:"address"`
	Coins   ccsdkTypes.Coins      `json:"coins"`
}

func NewGenesisAccount(aa *AppAccount) *GenesisAccount {
	return &GenesisAccount{
		Name:    aa.Name,
		Address: aa.Address,
		Coins:   aa.Coins.Sort(),
	}
}

func (ga *GenesisAccount) ToAppAccount() (acc *AppAccount, err error) {
	return &AppAccount{
		Name: ga.Name,
		BaseAccount: auth.BaseAccount{
			Address: ga.Address,
			Coins:   ga.Coins.Sort(),
		},
	}, nil
}

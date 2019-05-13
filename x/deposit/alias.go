package deposit

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

type (
	Deposit = types.Deposit
	GenesisState = types.GenesisState
	Keeper = keeper.Keeper
)

var (
	DepositKey          = types.DepositKey
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	NewKeeper           = keeper.NewKeeper
)

const (
	StoreKey = types.StoreKey
)

package deposit

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/querier"
	"github.com/ironman0x7b2/sentinel-sdk/x/deposit/types"
)

const (
	Codespace             = types.Codespace
	StoreKey              = types.StoreKey
	QuerierRoute          = types.QuerierRoute
	QueryDepositOfAddress = querier.QueryDepositOfAddress
	QueryAllDeposits      = querier.QueryAllDeposits
)

type (
	Deposit                    = types.Deposit
	GenesisState               = types.GenesisState
	Keeper                     = keeper.Keeper
	QueryDepositOfAddressPrams = querier.QueryDepositOfAddressPrams
)

// nolint: gochecknoglobals
var (
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	DepositKeyPrefix               = types.DepositKeyPrefix
	DepositKey                     = types.DepositKey
	NewKeeper                      = keeper.NewKeeper
	NewQueryDepositOfAddressParams = querier.NewQueryDepositOfAddressParams
	NewQuerier                     = querier.NewQuerier
)

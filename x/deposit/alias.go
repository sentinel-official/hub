package deposit

import (
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/querier"
	"github.com/sentinel-official/hub/x/deposit/types"
)

const (
	Codespace             = types.Codespace
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	RouterKey             = types.RouterKey
	QuerierRoute          = types.QuerierRoute
	QueryDepositOfAddress = types.QueryDepositOfAddress
	QueryAllDeposits      = types.QueryAllDeposits
)

var (
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorInvalidQueryType          = types.ErrorInvalidQueryType
	ErrorInsufficientDepositFunds  = types.ErrorInsufficientDepositFunds
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	DepositKey                     = types.DepositKey
	NewQueryDepositOfAddressParams = types.NewQueryDepositOfAddressParams
	NewKeeper                      = keeper.NewKeeper
	NewQuerier                     = querier.NewQuerier

	ModuleCdc        = types.ModuleCdc
	DepositKeyPrefix = types.DepositKeyPrefix
)

type (
	Deposit                    = types.Deposit
	GenesisState               = types.GenesisState
	QueryDepositOfAddressPrams = types.QueryDepositOfAddressPrams
	Keeper                     = keeper.Keeper
)

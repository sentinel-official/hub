package app

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	ibcicahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	customminttypes "github.com/sentinel-official/hub/x/mint/types"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"
	vpntypes "github.com/sentinel-official/hub/x/vpn/types"
)

type StoreKeys struct {
	kv        map[string]*sdk.KVStoreKey
	memory    map[string]*sdk.MemoryStoreKey
	transient map[string]*sdk.TransientStoreKey
}

func NewStoreKeys() StoreKeys {
	var (
		kv = sdk.NewKVStoreKeys(
			// Cosmos SDK keys
			authtypes.StoreKey,
			authzkeeper.StoreKey,
			banktypes.StoreKey,
			capabilitytypes.StoreKey,
			distributiontypes.StoreKey,
			evidencetypes.StoreKey,
			feegrant.StoreKey,
			govtypes.StoreKey,
			minttypes.StoreKey,
			paramstypes.StoreKey,
			slashingtypes.StoreKey,
			stakingtypes.StoreKey,
			upgradetypes.StoreKey,

			// Cosmos IBC keys
			ibcfeetypes.StoreKey,
			ibchost.StoreKey,
			ibcicacontrollertypes.StoreKey,
			ibcicahosttypes.StoreKey,
			ibctransfertypes.StoreKey,

			// Sentinel Hub keys
			customminttypes.StoreKey,
			swaptypes.StoreKey,
			vpntypes.StoreKey,

			// Other keys
			wasmtypes.StoreKey,
		)
		memory = sdk.NewMemoryStoreKeys(
			capabilitytypes.MemStoreKey,
		)
		transient = sdk.NewTransientStoreKeys(
			paramstypes.TStoreKey,
		)
	)

	return StoreKeys{
		kv:        kv,
		memory:    memory,
		transient: transient,
	}
}

func (sk *StoreKeys) KVKeys() map[string]*sdk.KVStoreKey               { return sk.kv }
func (sk *StoreKeys) MemoryKeys() map[string]*sdk.MemoryStoreKey       { return sk.memory }
func (sk *StoreKeys) TransientKeys() map[string]*sdk.TransientStoreKey { return sk.transient }

func (sk *StoreKeys) KV(v string) *sdk.KVStoreKey               { return sk.kv[v] }
func (sk *StoreKeys) Memory(v string) *sdk.MemoryStoreKey       { return sk.memory[v] }
func (sk *StoreKeys) Transient(v string) *sdk.TransientStoreKey { return sk.transient[v] }

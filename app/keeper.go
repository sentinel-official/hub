package app

import (
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcicacontroller "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller"
	ibcicacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	ibcicahost "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host"
	ibcicahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	ibcicahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	ibcfee "github.com/cosmos/ibc-go/v4/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"

	custommintkeeper "github.com/sentinel-official/hub/x/mint/keeper"
	customminttypes "github.com/sentinel-official/hub/x/mint/types"
	swapkeeper "github.com/sentinel-official/hub/x/swap/keeper"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"
	vpnkeeper "github.com/sentinel-official/hub/x/vpn/keeper"
	vpntypes "github.com/sentinel-official/hub/x/vpn/types"
)

type Keepers struct {
	// Cosmos SDK keepers
	AccountKeeper      authkeeper.AccountKeeper
	AuthzKeeper        authzkeeper.Keeper
	BankKeeper         bankkeeper.Keeper
	CapabilityKeeper   *capabilitykeeper.Keeper
	CrisisKeeper       crisiskeeper.Keeper
	DistributionKeeper distributionkeeper.Keeper
	EvidenceKeeper     evidencekeeper.Keeper
	FeeGrantKeeper     feegrantkeeper.Keeper
	GovKeeper          govkeeper.Keeper
	MintKeeper         mintkeeper.Keeper
	ParamsKeeper       paramskeeper.Keeper
	SlashingKeeper     slashingkeeper.Keeper
	StakingKeeper      stakingkeeper.Keeper
	UpgradeKeeper      upgradekeeper.Keeper

	// Cosmos IBC keepers
	IBCKeeper              *ibckeeper.Keeper
	IBCFeeKeeper           ibcfeekeeper.Keeper
	IBCICAControllerKeeper ibcicacontrollerkeeper.Keeper
	IBCICAHostKeeper       ibcicahostkeeper.Keeper
	IBCTransferKeeper      ibctransferkeeper.Keeper

	// Sentinel Hub keepers
	CustomMintKeeper custommintkeeper.Keeper
	SwapKeeper       swapkeeper.Keeper
	VPNKeeper        vpnkeeper.Keeper

	// Other keepers
	WasmKeeper wasmkeeper.Keeper

	// Cosmos IBC scoped keepers
	ScopedIBCKeeper              capabilitykeeper.ScopedKeeper
	ScopedIBCFeeKeeper           capabilitykeeper.ScopedKeeper
	ScopedIBCICAControllerKeeper capabilitykeeper.ScopedKeeper
	ScopedIBCICAHostKeeper       capabilitykeeper.ScopedKeeper
	ScopedIBCTransferKeeper      capabilitykeeper.ScopedKeeper

	// Other scoped keepers
	ScopedWasmKeeper capabilitykeeper.ScopedKeeper
}

func (k *Keepers) Subspace(v string) paramstypes.Subspace {
	subspace, _ := k.ParamsKeeper.GetSubspace(v)
	return subspace
}

func (k *Keepers) SetParamSubspaces(app *baseapp.BaseApp) {
	// Tendermint subspaces
	app.SetParamStore(
		k.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	// Cosmos SDK subspaces
	k.ParamsKeeper.Subspace(authtypes.ModuleName)
	k.ParamsKeeper.Subspace(banktypes.ModuleName)
	k.ParamsKeeper.Subspace(crisistypes.ModuleName)
	k.ParamsKeeper.Subspace(distributiontypes.ModuleName)
	k.ParamsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	k.ParamsKeeper.Subspace(minttypes.ModuleName)
	k.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	k.ParamsKeeper.Subspace(stakingtypes.ModuleName)

	// Cosmos IBC subspaces
	k.ParamsKeeper.Subspace(ibchost.ModuleName)
	k.ParamsKeeper.Subspace(ibcicacontrollertypes.SubModuleName)
	k.ParamsKeeper.Subspace(ibcicahosttypes.SubModuleName)
	k.ParamsKeeper.Subspace(ibctransfertypes.ModuleName)

	// Sentinel Hub subspaces
	k.ParamsKeeper.Subspace(swaptypes.ModuleName)

	// Other subspaces
	k.ParamsKeeper.Subspace(wasmtypes.ModuleName)
}

func NewKeepers(
	amino *codec.LegacyAmino,
	app *baseapp.BaseApp,
	blockedAddrs map[string]bool,
	cdc codec.Codec,
	homeDir string,
	invCheckPeriod uint,
	keys StoreKeys,
	mAccPerms map[string][]string,
	skipUpgradeHeights map[int64]bool,
	wasmConfig wasmtypes.WasmConfig,
	wasmOpts []wasmkeeper.Option,
	wasmProposalTypes []wasmtypes.ProposalType,
) (k Keepers) {
	// Cosmos SDK keepers
	k.ParamsKeeper = paramskeeper.NewKeeper(
		cdc, amino, keys.KV(paramstypes.StoreKey), keys.Transient(paramstypes.TStoreKey),
	)
	k.SetParamSubspaces(app)

	k.AccountKeeper = authkeeper.NewAccountKeeper(
		cdc, keys.KV(authtypes.StoreKey), k.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, mAccPerms,
	)
	k.BankKeeper = bankkeeper.NewBaseKeeper(
		cdc, keys.KV(banktypes.StoreKey), k.AccountKeeper, k.Subspace(banktypes.ModuleName), blockedAddrs,
	)
	k.CapabilityKeeper = capabilitykeeper.NewKeeper(
		cdc, keys.KV(capabilitytypes.StoreKey), keys.Memory(capabilitytypes.MemStoreKey),
	)
	k.AuthzKeeper = authzkeeper.NewKeeper(keys.KV(authzkeeper.StoreKey), cdc, app.MsgServiceRouter())
	k.FeeGrantKeeper = feegrantkeeper.NewKeeper(cdc, keys.KV(feegrant.StoreKey), k.AccountKeeper)

	stakingKeeper := stakingkeeper.NewKeeper(
		cdc, keys.KV(stakingtypes.StoreKey), k.AccountKeeper, k.BankKeeper, k.Subspace(stakingtypes.ModuleName),
	)

	k.DistributionKeeper = distributionkeeper.NewKeeper(
		cdc, keys.KV(distributiontypes.StoreKey), k.Subspace(distributiontypes.ModuleName),
		k.AccountKeeper, k.BankKeeper, &stakingKeeper, authtypes.FeeCollectorName, blockedAddrs,
	)
	k.MintKeeper = mintkeeper.NewKeeper(
		cdc, keys.KV(minttypes.StoreKey), k.Subspace(minttypes.ModuleName),
		&stakingKeeper, k.AccountKeeper, k.BankKeeper, authtypes.FeeCollectorName,
	)
	k.SlashingKeeper = slashingkeeper.NewKeeper(
		cdc, keys.KV(slashingtypes.StoreKey), &stakingKeeper, k.Subspace(slashingtypes.ModuleName),
	)

	k.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(k.DistributionKeeper.Hooks(), k.SlashingKeeper.Hooks()),
	)

	k.EvidenceKeeper = *evidencekeeper.NewKeeper(
		cdc, keys.KV(evidencetypes.StoreKey), &k.StakingKeeper, k.SlashingKeeper,
	)
	evidenceRouter := evidencetypes.NewRouter()
	k.EvidenceKeeper.SetRouter(evidenceRouter)

	k.CrisisKeeper = crisiskeeper.NewKeeper(
		k.Subspace(crisistypes.ModuleName), invCheckPeriod, k.BankKeeper, authtypes.FeeCollectorName,
	)
	k.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys.KV(upgradetypes.StoreKey), cdc, homeDir, app)

	// Cosmos IBC keepers
	k.ScopedIBCKeeper = k.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	k.ScopedIBCFeeKeeper = k.CapabilityKeeper.ScopeToModule(ibcfeetypes.ModuleName)
	k.ScopedIBCICAControllerKeeper = k.CapabilityKeeper.ScopeToModule(ibcicacontrollertypes.SubModuleName)
	k.ScopedIBCICAHostKeeper = k.CapabilityKeeper.ScopeToModule(ibcicahosttypes.SubModuleName)
	k.ScopedIBCTransferKeeper = k.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	k.IBCKeeper = ibckeeper.NewKeeper(
		cdc, keys.KV(ibchost.StoreKey), k.Subspace(ibchost.ModuleName),
		k.StakingKeeper, k.UpgradeKeeper, k.ScopedIBCKeeper,
	)
	k.IBCFeeKeeper = ibcfeekeeper.NewKeeper(
		cdc, keys.KV(ibcfeetypes.StoreKey), k.Subspace(ibcfeetypes.ModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper, k.AccountKeeper, k.BankKeeper,
	)
	k.IBCICAControllerKeeper = ibcicacontrollerkeeper.NewKeeper(
		cdc, keys.KV(ibcicacontrollertypes.StoreKey), k.Subspace(ibcicacontrollertypes.SubModuleName),
		k.IBCFeeKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper,
		k.ScopedIBCICAControllerKeeper, app.MsgServiceRouter(),
	)
	k.IBCICAHostKeeper = ibcicahostkeeper.NewKeeper(
		cdc, keys.KV(ibcicahosttypes.StoreKey), k.Subspace(ibcicahosttypes.SubModuleName),
		k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper, k.AccountKeeper,
		k.ScopedIBCICAHostKeeper, app.MsgServiceRouter(),
	)
	k.IBCTransferKeeper = ibctransferkeeper.NewKeeper(
		cdc, keys.KV(ibctransfertypes.StoreKey), k.Subspace(ibctransfertypes.ModuleName),
		k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper, &k.IBCKeeper.PortKeeper,
		k.AccountKeeper, k.BankKeeper, k.ScopedIBCTransferKeeper,
	)

	// Sentinel Hub keepers
	k.CustomMintKeeper = custommintkeeper.NewKeeper(cdc, keys.KV(customminttypes.StoreKey), k.MintKeeper)
	k.SwapKeeper = swapkeeper.NewKeeper(
		cdc, keys.KV(swaptypes.StoreKey), k.Subspace(swaptypes.ModuleName), k.AccountKeeper, k.BankKeeper,
	)
	k.VPNKeeper = vpnkeeper.NewKeeper(
		cdc, keys.KV(vpntypes.StoreKey), k.ParamsKeeper, k.AccountKeeper,
		k.BankKeeper, k.DistributionKeeper, authtypes.FeeCollectorName,
	)

	// Other keepers
	k.ScopedWasmKeeper = k.CapabilityKeeper.ScopeToModule(wasmtypes.ModuleName)

	var (
		wasmCapabilities = "iterator,staking,stargate,cosmwasm_1_1,cosmwasm_1_2"
		wasmDir          = filepath.Join(homeDir, "data")
	)

	k.WasmKeeper = wasmkeeper.NewKeeper(
		cdc, keys.KV(wasmtypes.StoreKey), k.Subspace(wasmtypes.ModuleName), k.AccountKeeper,
		k.BankKeeper, k.StakingKeeper, k.DistributionKeeper, k.IBCKeeper.ChannelKeeper,
		&k.IBCKeeper.PortKeeper, k.ScopedWasmKeeper, k.IBCTransferKeeper, app.MsgServiceRouter(),
		app.GRPCQueryRouter(), wasmDir, wasmConfig, wasmCapabilities, wasmOpts...,
	)

	// Cosmos SDK Governance router
	govRouter := govtypes.NewRouter().
		AddRoute(distributiontypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(k.DistributionKeeper)).
		AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(k.ParamsKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(k.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(k.IBCKeeper.ClientKeeper))
	if len(wasmProposalTypes) != 0 {
		govRouter.AddRoute(wasmtypes.RouterKey, wasmkeeper.NewWasmProposalHandler(k.WasmKeeper, wasmProposalTypes))
	}

	k.GovKeeper = govkeeper.NewKeeper(
		cdc, keys.KV(govtypes.StoreKey), k.Subspace(govtypes.ModuleName),
		k.AccountKeeper, k.BankKeeper, &k.StakingKeeper, govRouter,
	)

	// Cosmos IBC port router
	var (
		ibcICAControllerIBCModule ibcporttypes.IBCModule
		ibcICAHostIBCModule       = ibcicahost.NewIBCModule(k.IBCICAHostKeeper)
		ibcTransferIBCModule      = ibctransfer.NewIBCModule(k.IBCTransferKeeper)
		wasmIBCModule             ibcporttypes.IBCModule
	)

	ibcICAControllerIBCModule = ibcicacontroller.NewIBCMiddleware(ibcICAControllerIBCModule, k.IBCICAControllerKeeper)
	ibcICAControllerIBCModule = ibcfee.NewIBCMiddleware(ibcICAControllerIBCModule, k.IBCFeeKeeper)

	wasmIBCModule = wasm.NewIBCHandler(k.WasmKeeper, k.IBCKeeper.ChannelKeeper, k.IBCKeeper.ChannelKeeper)
	wasmIBCModule = ibcfee.NewIBCMiddleware(wasmIBCModule, k.IBCFeeKeeper)

	ibcPortRouter := ibcporttypes.NewRouter().
		AddRoute(ibcicacontrollertypes.SubModuleName, ibcICAControllerIBCModule).
		AddRoute(ibcicahosttypes.SubModuleName, ibcICAHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, ibcTransferIBCModule).
		AddRoute(wasmtypes.ModuleName, wasmIBCModule)
	k.IBCKeeper.SetRouter(ibcPortRouter)

	k.CapabilityKeeper.Seal()
	return k
}

package vpn

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	depositkeeper "github.com/sentinel-official/hub/x/deposit/keeper"
	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodekeeper "github.com/sentinel-official/hub/x/node/keeper"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plankeeper "github.com/sentinel-official/hub/x/plan/keeper"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providerkeeper "github.com/sentinel-official/hub/x/provider/keeper"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessionkeeper "github.com/sentinel-official/hub/x/session/keeper"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptionkeeper "github.com/sentinel-official/hub/x/subscription/keeper"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
	"github.com/sentinel-official/hub/x/vpn/client/cli"
	"github.com/sentinel-official/hub/x/vpn/expected"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/simulation"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleGenesis    = AppModule{}
	_ module.AppModule           = AppModule{}
	_ module.BeginBlockAppModule = AppModule{}
	_ module.EndBlockAppModule   = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return cdc.MustMarshalJSON(state)
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = deposittypes.RegisterQueryServiceHandlerClient(context.Background(), mux, deposittypes.NewQueryServiceClient(ctx))
	_ = nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, nodetypes.NewQueryServiceClient(ctx))
	_ = plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, plantypes.NewQueryServiceClient(ctx))
	_ = providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, providertypes.NewQueryServiceClient(ctx))
	_ = sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, sessiontypes.NewQueryServiceClient(ctx))
	_ = subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, subscriptiontypes.NewQueryServiceClient(ctx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

type AppModule struct {
	AppModuleBasic
	cdc      codec.Codec
	txConfig client.TxConfig
	ak       expected.AccountKeeper
	bk       expected.BankKeeper
	k        keeper.Keeper
}

func NewAppModule(cdc codec.Codec, txConfig client.TxConfig, ak expected.AccountKeeper, bk expected.BankKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:      cdc,
		txConfig: txConfig,
		ak:       ak,
		bk:       bk,
		k:        k,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	cdc.MustUnmarshalJSON(message, &state)
	InitGenesis(ctx, a.k, &state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	state := ExportGenesis(ctx, a.k)
	return cdc.MustMarshalJSON(state)
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route { return sdk.Route{} }

func (a AppModule) QuerierRoute() string { return "" }

func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

func (a AppModule) RegisterServices(configurator module.Configurator) {
	nodetypes.RegisterMsgServiceServer(configurator.MsgServer(), nodekeeper.NewMsgServiceServer(a.k.Node))
	plantypes.RegisterMsgServiceServer(configurator.MsgServer(), plankeeper.NewMsgServiceServer(a.k.Plan))
	providertypes.RegisterMsgServiceServer(configurator.MsgServer(), providerkeeper.NewMsgServiceServer(a.k.Provider))
	sessiontypes.RegisterMsgServiceServer(configurator.MsgServer(), sessionkeeper.NewMsgServiceServer(a.k.Session))
	subscriptiontypes.RegisterMsgServiceServer(configurator.MsgServer(), subscriptionkeeper.NewMsgServiceServer(a.k.Subscription))

	deposittypes.RegisterQueryServiceServer(configurator.QueryServer(), depositkeeper.NewQueryServiceServer(a.k.Deposit))
	nodetypes.RegisterQueryServiceServer(configurator.QueryServer(), nodekeeper.NewQueryServiceServer(a.k.Node))
	plantypes.RegisterQueryServiceServer(configurator.QueryServer(), plankeeper.NewQueryServiceServer(a.k.Plan))
	providertypes.RegisterQueryServiceServer(configurator.QueryServer(), providerkeeper.NewQueryServiceServer(a.k.Provider))
	sessiontypes.RegisterQueryServiceServer(configurator.QueryServer(), sessionkeeper.NewQueryServiceServer(a.k.Session))
	subscriptiontypes.RegisterQueryServiceServer(configurator.QueryServer(), subscriptionkeeper.NewQueryServiceServer(a.k.Subscription))

	m := keeper.NewMigrator(a.k)
	if err := configurator.RegisterMigration(types.ModuleName, 2, m.Migrate2to3); err != nil {
		panic("failed to migrate x/vpn from version 2 to 3: " + err.Error())
	}
}

func (a AppModule) ConsensusVersion() uint64 { return 3 }

func (a AppModule) BeginBlock(_ sdk.Context, _ abcitypes.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return EndBlock(ctx, a.k)
}

func (a AppModule) GenerateGenesisState(state *module.SimulationState) {
	simulation.RandomizedGenesisState(state)
}

func (a AppModule) ProposalContents(_ module.SimulationState) []simulationtypes.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(r *rand.Rand) []simulationtypes.ParamChange {
	return simulation.RandomizedParams(r)
}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(state module.SimulationState) []simulationtypes.WeightedOperation {
	return simulation.WeightedOperations(a.cdc, a.txConfig, state.AppParams, a.ak, a.bk, a.k)
}

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
	sdksimulation "github.com/cosmos/cosmos-sdk/types/simulation"
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
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(amino)
}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := cdc.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = deposittypes.RegisterQueryServiceHandlerClient(context.Background(), mux, deposittypes.NewQueryServiceClient(ctx))
	_ = providertypes.RegisterQueryServiceHandlerClient(context.Background(), mux, providertypes.NewQueryServiceClient(ctx))
	_ = nodetypes.RegisterQueryServiceHandlerClient(context.Background(), mux, nodetypes.NewQueryServiceClient(ctx))
	_ = plantypes.RegisterQueryServiceHandlerClient(context.Background(), mux, plantypes.NewQueryServiceClient(ctx))
	_ = subscriptiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, subscriptiontypes.NewQueryServiceClient(ctx))
	_ = sessiontypes.RegisterQueryServiceHandlerClient(context.Background(), mux, sessiontypes.NewQueryServiceClient(ctx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

type AppModule struct {
	AppModuleBasic
	cdc codec.Marshaler
	ak expected.AccountKeeper
	k  keeper.Keeper
}

func NewAppModule(cdc codec.Marshaler, ak expected.AccountKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		cdc:            cdc,
		ak:             ak,
		k:              k,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	cdc.MustUnmarshalJSON(message, &state)
	InitGenesis(ctx, a.k, &state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	return cdc.MustMarshalJSON(ExportGenesis(ctx, a.k))
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(a.k))
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

func (a AppModule) RegisterServices(configurator module.Configurator) {
	providertypes.RegisterMsgServiceServer(configurator.MsgServer(), providerkeeper.NewMsgServiceServer(a.k.Provider))
	nodetypes.RegisterMsgServiceServer(configurator.MsgServer(), nodekeeper.NewMsgServiceServer(a.k.Node))
	plantypes.RegisterMsgServiceServer(configurator.MsgServer(), plankeeper.NewMsgServiceServer(a.k.Plan))
	subscriptiontypes.RegisterMsgServiceServer(configurator.MsgServer(), subscriptionkeeper.NewMsgServiceServer(a.k.Subscription))
	sessiontypes.RegisterMsgServiceServer(configurator.MsgServer(), sessionkeeper.NewMsgServiceServer(a.k.Session))

	deposittypes.RegisterQueryServiceServer(configurator.QueryServer(), depositkeeper.NewQueryServiceServer(a.k.Deposit))
	providertypes.RegisterQueryServiceServer(configurator.QueryServer(), providerkeeper.NewQueryServiceServer(a.k.Provider))
	nodetypes.RegisterQueryServiceServer(configurator.QueryServer(), nodekeeper.NewQueryServiceServer(a.k.Node))
	plantypes.RegisterQueryServiceServer(configurator.QueryServer(), plankeeper.NewQueryServiceServer(a.k.Plan))
	subscriptiontypes.RegisterQueryServiceServer(configurator.QueryServer(), subscriptionkeeper.NewQueryServiceServer(a.k.Subscription))
	sessiontypes.RegisterQueryServiceServer(configurator.QueryServer(), sessionkeeper.NewQueryServiceServer(a.k.Session))
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abcitypes.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return EndBlock(ctx, a.k)
}

// App Simulation Methods

func (a AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenesisState(simState)
}

func (a AppModule) ProposalContents(_ module.SimulationState) []sdksimulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(r *rand.Rand) []sdksimulation.ParamChange {
	return simulation.RandomizedParams(r)
}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(simState module.SimulationState) []sdksimulation.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, a.cdc, a.k)
}

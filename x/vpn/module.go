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
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/expected"
	"github.com/sentinel-official/hub/x/vpn/keeper"
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

func (a AppModuleBasic) RegisterRESTRoutes(ctx client.Context, router *mux.Router) {
	panic("implement me")
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	deposit.RegisterQueryServiceHandlerClient(context.Background(), mux, deposit.NewQueryServiceClient(ctx))
	provider.RegisterQueryServiceHandlerClient(context.Background(), mux, provider.NewQueryServiceClient(ctx))
	node.RegisterQueryServiceHandlerClient(context.Background(), mux, node.NewQueryServiceClient(ctx))
	plan.RegisterQueryServiceHandlerClient(context.Background(), mux, plan.NewQueryServiceClient(ctx))
	subscription.RegisterQueryServiceHandlerClient(context.Background(), mux, subscription.NewQueryServiceClient(ctx))
	session.RegisterQueryServiceHandlerClient(context.Background(), mux, session.NewQueryServiceClient(ctx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	panic("implement me")
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	panic("implement me")
}

type AppModule struct {
	AppModuleBasic
	ak expected.AccountKeeper
	k  keeper.Keeper
}

func NewAppModule(ak expected.AccountKeeper, k keeper.Keeper) AppModule {
	return AppModule{
		ak: ak,
		k:  k,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, message json.RawMessage) []abci.ValidatorUpdate {
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
	provider.RegisterMsgServiceServer(configurator.MsgServer(), provider.NewMsgServiceServer(a.k.Provider))
	node.RegisterMsgServiceServer(configurator.MsgServer(), node.NewMsgServiceServer(a.k.Node))
	plan.RegisterMsgServiceServer(configurator.MsgServer(), plan.NewMsgServiceServer(a.k.Plan))
	subscription.RegisterMsgServiceServer(configurator.MsgServer(), subscription.NewMsgServiceServer(a.k.Subscription))
	session.RegisterMsgServiceServer(configurator.MsgServer(), session.NewMsgServiceServer(a.k.Session))
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlock(ctx, a.k)
}

func (a AppModule) GenerateGenesisState(input *module.SimulationState) {
	panic("implement me")
}

func (a AppModule) ProposalContents(simState module.SimulationState) []simulation.WeightedProposalContent {
	panic("implement me")
}

func (a AppModule) RandomizedParams(r *rand.Rand) []simulation.ParamChange {
	panic("implement me")
}

func (a AppModule) RegisterStoreDecoder(registry sdk.StoreDecoderRegistry) {
	panic("implement me")
}

func (a AppModule) WeightedOperations(simState module.SimulationState) []simulation.WeightedOperation {
	panic("implement me")
}

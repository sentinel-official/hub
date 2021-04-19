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
	"github.com/sentinel-official/hub/x/vpn/client/cli"
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

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = deposit.RegisterQueryServiceHandlerClient(context.Background(), mux, deposit.NewQueryServiceClient(ctx))
	_ = provider.RegisterQueryServiceHandlerClient(context.Background(), mux, provider.NewQueryServiceClient(ctx))
	_ = node.RegisterQueryServiceHandlerClient(context.Background(), mux, node.NewQueryServiceClient(ctx))
	_ = plan.RegisterQueryServiceHandlerClient(context.Background(), mux, plan.NewQueryServiceClient(ctx))
	_ = subscription.RegisterQueryServiceHandlerClient(context.Background(), mux, subscription.NewQueryServiceClient(ctx))
	_ = session.RegisterQueryServiceHandlerClient(context.Background(), mux, session.NewQueryServiceClient(ctx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
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

	deposit.RegisterQueryServiceServer(configurator.QueryServer(), deposit.NewQueryServiceServer(a.k.Deposit))
	provider.RegisterQueryServiceServer(configurator.QueryServer(), provider.NewQueryServiceServer(a.k.Provider))
	node.RegisterQueryServiceServer(configurator.QueryServer(), node.NewQueryServiceServer(a.k.Node))
	plan.RegisterQueryServiceServer(configurator.QueryServer(), plan.NewQueryServiceServer(a.k.Plan))
	subscription.RegisterQueryServiceServer(configurator.QueryServer(), subscription.NewQueryServiceServer(a.k.Subscription))
	session.RegisterQueryServiceServer(configurator.QueryServer(), session.NewQueryServiceServer(a.k.Session))
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (a AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlock(ctx, a.k)
}

func (a AppModule) GenerateGenesisState(_ *module.SimulationState) {}

func (a AppModule) ProposalContents(_ module.SimulationState) []simulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(_ *rand.Rand) []simulation.ParamChange {
	return nil
}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(_ module.SimulationState) []simulation.WeightedOperation {
	return nil
}

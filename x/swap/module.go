package swap

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

	"github.com/sentinel-official/hub/x/swap/client/cli"
	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/simulation"
	"github.com/sentinel-official/hub/x/swap/types"
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
	_ = types.RegisterQueryServiceHandlerClient(context.Background(), mux, types.NewQueryServiceClient(ctx))
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
	k   keeper.Keeper
}

func NewAppModule(cdc codec.Marshaler, k keeper.Keeper) AppModule {
	return AppModule{
		cdc: cdc,
		k:   k,
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

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
}

func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(a.k))
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

func (a AppModule) RegisterServices(configurator module.Configurator) {
	types.RegisterMsgServiceServer(configurator.MsgServer(), keeper.NewMsgServiceServer(a.k))
	types.RegisterQueryServiceServer(configurator.QueryServer(), keeper.NewQueryServiceServer(a.k))
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abcitypes.RequestBeginBlock) {}

func (a AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

// AppSimulaion Methods

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenesisState(simState)
}

func (a AppModule) ProposalContents(_ module.SimulationState) []sdksimulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(r *rand.Rand) []sdksimulation.ParamChange {
	return simulation.ParamChanges(r)
}

func (a AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simulation.NewDecodeStore(a.cdc)
}

func (a AppModule) WeightedOperations(simState module.SimulationState) []sdksimulation.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, a.k)
}

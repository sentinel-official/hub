package mint

import (
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

	"github.com/sentinel-official/hub/x/mint/keeper"
	"github.com/sentinel-official/hub/x/mint/types"
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

func (a AppModuleBasic) RegisterInterfaces(_ codectypes.InterfaceRegistry) {}

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

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

func (a AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

func (a AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

type AppModule struct {
	AppModuleBasic
	cdc codec.Codec
	k   keeper.Keeper
}

func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		cdc: cdc,
		k:   k,
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

func (a AppModule) RegisterServices(_ module.Configurator) {}

func (a AppModule) ConsensusVersion() uint64 { return 1 }

func (a AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	BeginBlock(ctx, a.k)
}

func (a AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (a AppModule) GenerateGenesisState(_ *module.SimulationState) {}

func (a AppModule) ProposalContents(_ module.SimulationState) []simulationtypes.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(_ *rand.Rand) []simulationtypes.ParamChange { return nil }

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(_ module.SimulationState) []simulationtypes.WeightedOperation {
	return nil
}

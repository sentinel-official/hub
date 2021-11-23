package mint

import (
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

	"github.com/sentinel-official/hub/x/mint/keeper"
	"github.com/sentinel-official/hub/x/mint/types"
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

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(_ codectypes.InterfaceRegistry) {}

func (a AppModuleBasic) DefaultGenesis(appCodec codec.JSONCodec) json.RawMessage {
	return appCodec.MustMarshalJSON(types.DefaultGenesisState())
}

func (a AppModuleBasic) ValidateGenesis(appCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := appCodec.UnmarshalJSON(message, &state); err != nil {
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
	appCodec codec.Codec
	k        keeper.Keeper
}

func NewAppModule(appCodec codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		appCodec: appCodec,
		k:        k,
	}
}

func (a AppModule) InitGenesis(ctx sdk.Context, appCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	appCodec.MustUnmarshalJSON(message, &state)
	InitGenesis(ctx, a.k, &state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, appCodec codec.JSONCodec) json.RawMessage {
	return appCodec.MustMarshalJSON(ExportGenesis(ctx, a.k))
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a AppModule) Route() sdk.Route {
	return sdk.Route{}
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier { return nil }

func (a AppModule) RegisterServices(_ module.Configurator) {}

func (a AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	BeginBlock(ctx, a.k)
}

func (a AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (a AppModule) GenerateGenesisState(_ *module.SimulationState) {}

func (a AppModule) ProposalContents(_ module.SimulationState) []sdksimulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(_ *rand.Rand) []sdksimulation.ParamChange { return nil }

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(_ module.SimulationState) []sdksimulation.WeightedOperation {
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }

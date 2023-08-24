package swap

import (
	"context"
	"encoding/json"

	abcitypes "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sdksimtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/sentinel-official/hub/x/swap/client/cli"
	"github.com/sentinel-official/hub/x/swap/keeper"
	"github.com/sentinel-official/hub/x/swap/types"
)

var (
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleGenesis    = AppModule{}
	_ module.BeginBlockAppModule = AppModule{}
	_ module.EndBlockAppModule   = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string { return types.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(ctx client.Context, mux *runtime.ServeMux) {
	_ = types.RegisterQueryServiceHandlerClient(context.Background(), mux, types.NewQueryServiceClient(ctx))
}

func (a AppModuleBasic) GetTxCmd() *cobra.Command { return cli.GetTxCmd() }

func (a AppModuleBasic) GetQueryCmd() *cobra.Command { return cli.GetQueryCmd() }

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

func (a AppModule) DefaultGenesis(jsonCodec codec.JSONCodec) json.RawMessage {
	state := types.DefaultGenesisState()
	return jsonCodec.MustMarshalJSON(state)
}

func (a AppModule) ValidateGenesis(jsonCodec codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var state types.GenesisState
	if err := jsonCodec.UnmarshalJSON(message, &state); err != nil {
		return err
	}

	return state.Validate()
}

func (a AppModule) InitGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec, message json.RawMessage) []abcitypes.ValidatorUpdate {
	var state types.GenesisState
	jsonCodec.MustUnmarshalJSON(message, &state)
	InitGenesis(ctx, a.k, &state)

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, jsonCodec codec.JSONCodec) json.RawMessage {
	state := ExportGenesis(ctx, a.k)
	return jsonCodec.MustMarshalJSON(state)
}

func (a AppModule) BeginBlock(_ sdk.Context, _ abcitypes.RequestBeginBlock) {}

func (a AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (AppModule) GenerateGenesisState(_ *module.SimulationState) {}

func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

func (a AppModule) WeightedOperations(_ module.SimulationState) []sdksimtypes.WeightedOperation {
	return nil
}

func (a AppModule) ConsensusVersion() uint64 { return 1 }

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (a AppModule) RegisterServices(configurator module.Configurator) {
	types.RegisterMsgServiceServer(configurator.MsgServer(), keeper.NewMsgServiceServer(a.k))
	types.RegisterQueryServiceServer(configurator.QueryServer(), keeper.NewQueryServiceServer(a.k))
}

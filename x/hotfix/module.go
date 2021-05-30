package hotfix

import (
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	"github.com/sentinel-official/hub/x/hotfix/expected"
	"github.com/sentinel-official/hub/x/hotfix/types"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string                                         { return types.ModuleName }
func (a AppModuleBasic) RegisterLegacyAminoCodec(_ *codec.LegacyAmino)        {}
func (a AppModuleBasic) RegisterInterfaces(_ codectypes.InterfaceRegistry)    {}
func (a AppModuleBasic) DefaultGenesis(_ codec.JSONMarshaler) json.RawMessage { return nil }

func (a AppModuleBasic) ValidateGenesis(_ codec.JSONMarshaler, _ client.TxEncodingConfig, _ json.RawMessage) error {
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router)              {}
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}
func (a AppModuleBasic) GetTxCmd() *cobra.Command                                        { return nil }
func (a AppModuleBasic) GetQueryCmd() *cobra.Command                                     { return nil }

type AppModule struct {
	AppModuleBasic
	ak expected.AccountKeeper
	bk expected.BankKeeper
	sk stakingkeeper.Keeper
}

func NewAppModule(ak expected.AccountKeeper, bk expected.BankKeeper, sk stakingkeeper.Keeper) AppModule {
	return AppModule{
		ak: ak,
		bk: bk,
		sk: sk,
	}
}

func (a AppModule) InitGenesis(_ sdk.Context, _ codec.JSONMarshaler, _ json.RawMessage) []abcitypes.ValidatorUpdate {
	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage { return nil }
func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry)                             {}
func (a AppModule) Route() sdk.Route                                                       { return sdk.NewRoute(types.RouterKey, nil) }
func (a AppModule) QuerierRoute() string                                                   { return types.QuerierRoute }
func (a AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier                  { return nil }
func (a AppModule) RegisterServices(_ module.Configurator)                                 {}

func (a AppModule) BeginBlock(ctx sdk.Context, _ abcitypes.RequestBeginBlock) {
	BeginBlock(ctx, a.ak, a.bk, a.sk)
}

func (a AppModule) EndBlock(_ sdk.Context, _ abcitypes.RequestEndBlock) []abcitypes.ValidatorUpdate {
	return nil
}

func (a AppModule) GenerateGenesisState(_ *module.SimulationState) {}

func (a AppModule) ProposalContents(_ module.SimulationState) []simulation.WeightedProposalContent {
	return nil
}

func (a AppModule) RandomizedParams(_ *rand.Rand) []simulation.ParamChange { return nil }
func (a AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry)        {}

func (a AppModule) WeightedOperations(_ module.SimulationState) []simulation.WeightedOperation {
	return nil
}

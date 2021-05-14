package simulation

import (
	"math/rand"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocdc "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub/x/swap/keeper"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestSimulateMsgSwapRequest(t *testing.T) {
}

func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	// ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	cryptocdc.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)
	k := keeper.NewKeeper(cdc, , app.GetSubspace(swaptypes.ModuleName), app.BankKeeper)

	SimulateMsgSwapRequest(app.BankKeeper, app.AccountKeeper, app.Keeper, app.AppCodec)
}

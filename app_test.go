package hub

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func TestAppExport(t *testing.T) {
	var (
		db  = tmdb.NewMemDB()
		app = NewApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
			db, nil, true, map[int64]bool{}, 0)
		genesis = ModuleBasics.DefaultGenesis()
	)

	state, err := codec.MarshalJSONIndent(app.cdc, genesis)
	require.NoError(t, err)

	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: state,
		},
	)
	app.Commit()

	app = NewApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)),
		db, nil, true, map[int64]bool{}, 0)
	_, _, err = app.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

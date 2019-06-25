package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	_amino "github.com/tendermint/go-amino"
	tmConfig "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"

	"github.com/sentinel-official/sentinel-hub/app"
)

func ExportGenesisFile(genFile, chainID string, validators []types.GenesisValidator, state json.RawMessage) error {
	genDoc := types.GenesisDoc{
		ChainID:    chainID,
		Validators: validators,
		AppState:   state,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

func ExportGenesisFileWithTime(genFile, chainID string, validators []types.GenesisValidator,
	state json.RawMessage, genTime time.Time) error {

	genDoc := types.GenesisDoc{
		GenesisTime: genTime,
		ChainID:     chainID,
		Validators:  validators,
		AppState:    state,
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return err
	}

	return genDoc.SaveAs(genFile)
}

func InitializeNodeValidatorFiles(config *tmConfig.Config) (nodeID string, valPubKey crypto.PubKey, err error) {
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		return nodeID, valPubKey, err
	}

	nodeID = string(nodeKey.ID())
	server.UpgradeOldPrivValFile(config)

	pvKeyFile := config.PrivValidatorKeyFile()
	if err := common.EnsureDir(filepath.Dir(pvKeyFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	pvStateFile := config.PrivValidatorStateFile()
	if err := common.EnsureDir(filepath.Dir(pvStateFile), 0777); err != nil {
		return nodeID, valPubKey, nil
	}

	valPubKey = privval.LoadOrGenFilePV(pvKeyFile, pvStateFile).GetPubKey()
	return nodeID, valPubKey, nil
}

func loadGenesisDoc(cdc *_amino.Codec, genFile string) (genDoc types.GenesisDoc, err error) {
	genContents, err := ioutil.ReadFile(genFile)
	if err != nil {
		return genDoc, err
	}

	if err = cdc.UnmarshalJSON(genContents, &genDoc); err != nil {
		return genDoc, err
	}

	return genDoc, err
}

func initializeEmptyGenesis(cdc *codec.Codec, genFile string, overwrite bool) (state json.RawMessage, err error) {
	if !overwrite && common.FileExists(genFile) {
		return nil, fmt.Errorf("genesis.json file already exists: %v", genFile)
	}

	return codec.MarshalJSONIndent(cdc, app.NewDefaultGenesisState())
}

package keeper_test

import (
	"fmt"

	hub "github.com/sentinel-official/hub/types"
	genesis "github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

func (suite *KeeperTestSuite) TestInitGenesis() {

	testProvider := types.Provider{
		Address: hub.ProvAddress("test-address"),
		Name:    "test-name",
	}

	testGenesis := []struct {
		desc            string
		genesisState    types.GenesisState
		expGenesisState types.GenesisState
	}{
		{
			desc:            "Default Genesis State",
			genesisState:    types.DefaultGenesisState(),
			expGenesisState: types.DefaultGenesisState(),
		},
		{
			desc:            "Non Default Genesis State",
			genesisState:    types.NewGenesisState(types.Providers{testProvider}),
			expGenesisState: types.NewGenesisState(types.Providers{testProvider}),
		},
	}

	for _, test := range testGenesis {
		test := test
		suite.Run(test.desc, func() {
			suite.SetupTest()

			genesis.InitGenesis(suite.ctx, suite.keeper, test.genesisState)
			actualGenesisState := suite.keeper.GetProviders(suite.ctx, 0, 0)

			suite.Require().Equal(test.expGenesisState, actualGenesisState)

		})
	}
}

func (suite *KeeperTestSuite) TestExportGenesis() {
	testProvider := types.Provider{
		Address: hub.ProvAddress("test-address"),
		Name:    "test-name",
	}

	testGenesis := []struct {
		desc            string
		testProviders   types.Providers
		expGenesisState types.GenesisState
	}{
		{
			desc:            "Default Genesis State",
			testProviders:   types.Providers{},
			expGenesisState: types.DefaultGenesisState(),
		},
		{
			desc:            "Non Default Genesis State",
			testProviders:   types.Providers{testProvider},
			expGenesisState: types.NewGenesisState(types.Providers{testProvider}),
		},
	}

	for _, test := range testGenesis {
		test := test
		suite.Run(test.desc, func() {
			suite.SetupTest()

			for _, data := range test.testProviders {
				suite.keeper.SetProvider(suite.ctx, data)
			}
			actualGenesisState := genesis.ExportGenesis(suite.ctx, suite.keeper)
			suite.Require().Equal(test.expGenesisState, actualGenesisState)
		})
	}
}

func (suite *KeeperTestSuite) TestValidateGenesis() {

	testProvider := types.Provider{
		Address: hub.ProvAddress("test-address"),
		Name:    "test-name",
	}

	testGenesis := []struct {
		desc         string
		genesisState types.GenesisState
		expError     bool
		err          error
	}{
		{
			desc: "Invalid Genesis State with duplicate provider",
			genesisState: types.NewGenesisState(
				types.Providers{testProvider,
					types.Provider{
						Address: hub.ProvAddress("test-address"),
						Name:    "test-name-1",
					},
				}),
			expError: true,
			err:      fmt.Errorf("found duplicate provider address %s", hub.ProvAddress("test-address").String()),
		},
		{
			desc:         "Default Genesis State",
			genesisState: types.DefaultGenesisState(),
			expError:     false,
		},
		{
			desc:         "Non Default Genesis State",
			genesisState: types.NewGenesisState(types.Providers{testProvider}),
			expError:     false,
		},
	}

	for _, test := range testGenesis {
		test := test
		suite.Run(test.desc, func() {
			suite.SetupTest()

			err := genesis.ValidateGenesis(test.genesisState)

			if test.expError {
				suite.Require().Error(err)
				suite.Require().Equal(test.err, err)
			} else {
				suite.Require().NoError(err)
			}

		})
	}
}

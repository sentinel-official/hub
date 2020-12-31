package keeper_test

import (
	"fmt"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/provider/types"
)

func (suite *KeeperTestSuite) TestSetProvider() {

	testProvider := types.Provider{
		Address: hub.ProvAddress("test-address"),
		Name:    "test-name",
	}

	suite.keeper.SetProvider(suite.ctx, testProvider)
	actualProvider, found := suite.keeper.GetProvider(suite.ctx, hub.ProvAddress("test-address"))
	suite.Equal(true, found)
	suite.Equal(testProvider, actualProvider)
}

func (suite *KeeperTestSuite) TestHasProvider() {

	testProvider := types.Provider{
		Address: hub.ProvAddress("test-address"),
		Name:    "test-name",
	}

	suite.keeper.SetProvider(suite.ctx, testProvider)
	found := suite.keeper.HasProvider(suite.ctx, hub.ProvAddress("test-address"))
	suite.Equal(true, found)

	//Provider doesn't exist
	found = suite.keeper.HasProvider(suite.ctx, hub.ProvAddress("test-address-2"))
	suite.Equal(false, found)

}

func (suite *KeeperTestSuite) TestGetProvider() {

	testProviders := types.Providers{
		types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
		types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
		types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
		types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
		types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
	}

	for _, provider := range testProviders {

		provider := provider
		suite.Run(fmt.Sprintf("Retrieve Provider %s", provider.Address.String()), func() {
			suite.SetupTest()

			suite.keeper.SetProvider(suite.ctx, provider)

			retrievedProvider, found := suite.keeper.GetProvider(suite.ctx, provider.Address)

			suite.Require().Equal(true, found)
			suite.Require().Equal(retrievedProvider, provider)
		})
	}
}

func (suite *KeeperTestSuite) TestGetProviders() {

	tests := []struct {
		desc          string
		testProviders types.Providers
		expProviders  types.Providers
		limit         int
		skip          int
	}{
		{
			"return complete list of providers",
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
				types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
				types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			0,
			0,
		},

		{
			"return empty list ",
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
				types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			nil,
			0,
			5,
		},

		{
			"return last 3 providers",
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
				types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			0,
			2,
		},

		{
			"return 3rd provider in the list",
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-1"), Name: "test-name-1"},
				types.Provider{Address: hub.ProvAddress("test-address-2"), Name: "test-name-2"},
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
				types.Provider{Address: hub.ProvAddress("test-address-4"), Name: "test-name-4"},
				types.Provider{Address: hub.ProvAddress("test-address-5"), Name: "test-name-5"},
			},
			types.Providers{
				types.Provider{Address: hub.ProvAddress("test-address-3"), Name: "test-name-3"},
			},
			1,
			2,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.desc, func() {
			suite.SetupTest()

			for _, provider := range test.testProviders {
				suite.keeper.SetProvider(suite.ctx, provider)
			}

			actualProviders := suite.keeper.GetProviders(suite.ctx, test.skip, test.limit)

			suite.Equal(test.expProviders, actualProviders)

		})
	}
}

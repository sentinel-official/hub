package types_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestNode_UpdateInfo(t *testing.T) {
	tests := []struct {
		name string
		info types.Node
		want types.Node
	}{
		{
			"node_moniker is empty",
			types.Node{Moniker: ""},
			types.Node{},
		}, {
			"node_moniker length Valid",
			types.Node{Moniker: "moniker"},
			types.Node{Moniker: "moniker"},
		}, {
			"prices_per_gb is nil",
			types.Node{PricesPerGB: nil},
			types.Node{},
		}, {
			"prices_per_gb is empty",
			types.Node{PricesPerGB: sdk.Coins{}},
			types.Node{},
		}, {
			"prices_per_gb is negative",
			types.Node{PricesPerGB: sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}},
			types.Node{},
		}, {
			"prices_per_gb is zero",
			types.Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 0)}},
			types.Node{},
		}, {
			"prices_per_gb is positive",
			types.Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}},
			types.Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}},
		}, {
			"net_speed is empty",
			types.Node{InternetSpeed: hub.Bandwidth{}},
			types.Node{},
		}, {
			"net_speed is negative",
			types.Node{InternetSpeed: hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000))},
			types.Node{},
		}, {
			"net_speed is zero",
			types.Node{InternetSpeed: hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0))},
			types.Node{},
		}, {
			"net_speed is positive",
			types.Node{InternetSpeed: hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000))},
			types.Node{InternetSpeed: hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000))},
		}, {
			"encryption is empty",
			types.Node{Encryption: ""},
			types.Node{},
		}, {
			"encryption is valid",
			types.Node{Encryption: "encryption"},
			types.Node{Encryption: "encryption"},
		}, {
			"node_type is empty",
			types.Node{Type: ""},
			types.Node{},
		}, {
			"node_type is valid",
			types.Node{Type: "node_type"},
			types.Node{Type: "node_type"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node := types.Node{}.UpdateInfo(tc.info)
			if !reflect.DeepEqual(node, tc.want) {
				t.Errorf("\ngot = %vwant = %v", node, tc.want)
			}
		})
	}
}

func TestNode_FindPricePerGB(t *testing.T) {
	var node types.Node
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = types.Node{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = types.Node{PricesPerGB: sdk.Coins{}}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = types.Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.NewInt64Coin("stake", 100))
}

func TestNode_DepositToBandwidth(t *testing.T) {
	node := types.Node{
		PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)},
		Deposit:     sdk.NewInt64Coin("stake", 100),
	}

	_, err := node.DepositToBandwidth(sdk.Coin{})
	require.NotNil(t, err)

	bandwidth, err := node.DepositToBandwidth(sdk.NewInt64Coin("stake", 0))
	require.Nil(t, err)
	reflect.DeepEqual(hub.NewBandwidth(sdk.NewInt(0), sdk.NewInt(0)), bandwidth)

	bandwidth, err = node.DepositToBandwidth(sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.Nil(t, err)
	reflect.DeepEqual(hub.NewBandwidth(sdk.NewInt(-500000000), sdk.NewInt(-500000000)), bandwidth)

	bandwidth, err = node.DepositToBandwidth(sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)
	reflect.DeepEqual(hub.NewBandwidth(sdk.NewInt(500000000), sdk.NewInt(500000000)), bandwidth)
}

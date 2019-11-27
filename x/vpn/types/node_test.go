package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
)

func TestNode_UpdateInfo(t *testing.T) {
	tests := []struct {
		name string
		info Node
		want Node
	}{
		{
			"node_moniker is empty",
			Node{Moniker: ""},
			Node{},
		}, {
			"node_moniker length Valid",
			Node{Moniker: "moniker"},
			Node{Moniker: "moniker"},
		}, {
			"prices_per_gb is nil",
			Node{PricesPerGB: nil},
			Node{},
		}, {
			"prices_per_gb is empty",
			Node{PricesPerGB: sdk.Coins{}},
			Node{},
		}, {
			"prices_per_gb is negative",
			Node{PricesPerGB: sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}},
			Node{},
		}, {
			"prices_per_gb is zero",
			Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 0)}},
			Node{},
		}, {
			"prices_per_gb is positive",
			Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}},
			Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}},
		}, {
			"net_speed is empty",
			Node{InternetSpeed: hub.Bandwidth{}},
			Node{},
		}, {
			"net_speed is negative",
			Node{InternetSpeed: TestBandwidthNeg},
			Node{},
		}, {
			"net_speed is zero",
			Node{InternetSpeed: TestBandwidthZero},
			Node{},
		}, {
			"net_speed is positive",
			Node{InternetSpeed: TestBandwidthPos1},
			Node{InternetSpeed: TestBandwidthPos1},
		}, {
			"encryption is empty",
			Node{Encryption: ""},
			Node{},
		}, {
			"encryption is valid",
			Node{Encryption: "encryption"},
			Node{Encryption: "encryption"},
		}, {
			"node_type is empty",
			Node{Type: ""},
			Node{},
		}, {
			"node_type is valid",
			Node{Type: "node_type"},
			Node{Type: "node_type"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node := Node{}.UpdateInfo(tc.info)
			if !reflect.DeepEqual(node, tc.want) {
				t.Errorf("\ngot = %vwant = %v", node, tc.want)
			}
		})
	}
}

func TestNode_FindPricePerGB(t *testing.T) {
	var node Node
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = Node{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = Node{PricesPerGB: sdk.Coins{}}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.Coin{})

	node = Node{PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)}}
	require.Equal(t, node.FindPricePerGB("stake"), sdk.NewInt64Coin("stake", 100))
}

func TestNode_DepositToBandwidth(t *testing.T) {
	node := Node{
		PricesPerGB: sdk.Coins{sdk.NewInt64Coin("stake", 100)},
		Deposit:     sdk.NewInt64Coin("stake", 100),
	}

	_, err := node.DepositToBandwidth(sdk.Coin{})
	require.NotNil(t, err)

	bandwidth, err := node.DepositToBandwidth(sdk.NewInt64Coin("stake", 0))
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthZero, bandwidth)

	bandwidth, err = node.DepositToBandwidth(sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthNeg, bandwidth)

	bandwidth, err = node.DepositToBandwidth(sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthPos1, bandwidth)
}

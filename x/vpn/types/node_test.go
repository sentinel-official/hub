package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/sentinel-hub/types"
)

func TestNode_UpdateInfo(t *testing.T) {
	tests := []struct {
		name string
		info Node
		want Node
	}{
		{
			"node_moniker is empty",
			Node{Moniker: TestMonikerLenZero},
			Node{},
		}, {
			"node_moniker length Valid",
			Node{Moniker: TestMonikerValid},
			Node{Moniker: TestMonikerValid},
		}, {
			"prices_per_gb is nil",
			Node{PricesPerGB: nil},
			Node{},
		}, {
			"prices_per_gb is empty",
			Node{PricesPerGB: TestCoinsEmpty},
			Node{},
		}, {
			"prices_per_gb is negative",
			Node{PricesPerGB: TestCoinsNeg},
			Node{},
		}, {
			"prices_per_gb is zero",
			Node{PricesPerGB: TestCoinsZero},
			Node{},
		}, {
			"prices_per_gb is positive",
			Node{PricesPerGB: TestCoinsPos},
			Node{PricesPerGB: TestCoinsPos},
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
			Node{Encryption: TestEncryption},
			Node{Encryption: TestEncryption},
		}, {
			"node_type is empty",
			Node{Type: ""},
			Node{},
		}, {
			"node_type is valid",
			Node{Type: TestNodeType},
			Node{Type: TestNodeType},
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
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinEmpty)

	node = Node{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinEmpty)

	node = Node{PricesPerGB: TestCoinsEmpty}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinEmpty)

	node = Node{PricesPerGB: TestCoinsPos}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinPos)
}

func TestNode_DepositToBandwidth(t *testing.T) {
	node := Node{
		PricesPerGB: TestCoinsPos,
		Deposit:     TestCoinPos,
	}

	_, err := node.DepositToBandwidth(TestCoinEmpty)
	require.NotNil(t, err)

	bandwidth, err := node.DepositToBandwidth(TestCoinZero)
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthZero, bandwidth)

	bandwidth, err = node.DepositToBandwidth(TestCoinNeg)
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthNeg, bandwidth)

	bandwidth, err = node.DepositToBandwidth(TestCoinPos)
	require.Nil(t, err)
	reflect.DeepEqual(TestBandwidthPos1, bandwidth)
}

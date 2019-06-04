package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/types"
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
			"prices_per_gb is invalid",
			Node{PricesPerGB: TestCoinsInvalid},
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
			Node{InternetSpeed: types.Bandwidth{}},
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
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinNil)

	node = Node{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinNil)

	node = Node{PricesPerGB: TestCoinsEmpty}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinNil)

	node = Node{PricesPerGB: TestCoinsPos}
	require.Equal(t, node.FindPricePerGB("stake"), TestCoinPos)
}

func TestNode_DepositToBandwidth(t *testing.T) {
	node := Node{
		PricesPerGB: TestCoinsPos,
		Deposit:     TestCoinPos}

	_, err := node.DepositToBandwidth(TestCoinNil)
	require.NotNil(t, err)

	bandwidth, err := node.DepositToBandwidth(TestCoinZero)
	require.Nil(t, err)

	bandwidth, err = node.DepositToBandwidth(TestCoinNeg)
	require.Nil(t, err)

	bandwidth, err = node.DepositToBandwidth(TestCoinPos)
	require.Nil(t, err)
	value := node.Deposit.Amount.Mul(types.MB500).Quo(TestCoinPos.Amount)
	require.Equal(t, types.Bandwidth{value, value}, bandwidth)
}

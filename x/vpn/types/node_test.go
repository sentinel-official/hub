package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/types"
)

func TestNode_UpdateDetails(t *testing.T) {
	tests := []struct {
		name    string
		details Node
		want    Node
	}{
		{
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
			Node{NetSpeed: types.Bandwidth{}},
			Node{},
		}, {
			"net_speed is negative",
			Node{NetSpeed: types.NewBandwidth(TestUploadNeg, TestDownloadNeg)},
			Node{},
		}, {
			"net_speed is zero",
			Node{NetSpeed: types.NewBandwidth(TestUploadZero, TestDownloadZero)},
			Node{},
		}, {
			"net_speed is positive",
			Node{NetSpeed: types.NewBandwidth(TestUploadPos, TestDownloadPos)},
			Node{NetSpeed: types.NewBandwidth(TestUploadPos, TestDownloadPos)},
		}, {
			"api_port is zero",
			Node{APIPort: 0},
			Node{},
		}, {
			"api_port is positive",
			Node{APIPort: 8000},
			Node{APIPort: 8000},
		}, {
			"encryption is empty",
			Node{EncryptionMethod: ""},
			Node{},
		}, {
			"encryption is valid",
			Node{EncryptionMethod: TestEncryptionMethod},
			Node{EncryptionMethod: TestEncryptionMethod},
		}, {
			"node_type is empty",
			Node{Type: ""},
			Node{},
		}, {
			"node_type is valid",
			Node{Type: TestNodeType},
			Node{Type: TestNodeType},
		}, {
			"version is empty",
			Node{Version: ""},
			Node{},
		}, {
			"version is valid",
			Node{Version: TestVersion},
			Node{Version: TestVersion},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node := Node{}
			if node.UpdateDetails(tc.details); !reflect.DeepEqual(node, tc.want) {
				t.Errorf("\ngot = %vwant = %v", node, tc.want)
			}
		})
	}
}

func TestNode_FindPricePerGB(t *testing.T) {
	var node Node
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = Node{PricesPerGB: nil}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = Node{PricesPerGB: TestCoinsEmpty}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinNil)
	node = Node{PricesPerGB: TestCoinsPos}
	require.Equal(t, node.FindPricePerGB("sent"), TestCoinPos)
}

func TestNode_CalculateBandwidth(t *testing.T) {
	node := Node{PricesPerGB: TestCoinsPos}

	b, err := node.CalculateBandwidth(TestCoinNil)
	require.Equal(t, err, ErrorInvalidPriceDenom())
	require.Equal(t, b, types.Bandwidth{})

	b, err = node.CalculateBandwidth(TestCoinZero)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthZero)

	b, err = node.CalculateBandwidth(TestCoinNeg)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthNeg)

	b, err = node.CalculateBandwidth(TestCoinPos)
	require.Nil(t, err)
	require.Equal(t, b, TestBandwidthPos)
}

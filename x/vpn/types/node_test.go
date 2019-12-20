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

func TestNode_IsValid(t *testing.T) {
	var node Node
	require.NotNil(t, node.IsValid())
	
	node.ID = []byte("")
	require.NotNil(t, node.IsValid())
	
	node.ID = nil
	require.NotNil(t, node.IsValid())
	
	node.ID = hub.NewNodeID(0)
	require.NotNil(t, node.IsValid())
	
	node.Owner = []byte("")
	require.NotNil(t, node.IsValid())
	
	node.Owner = nil
	require.NotNil(t, node.IsValid())
	
	node.Owner = TestAddress1
	require.NotNil(t, node.IsValid())
	
	node.Deposit = sdk.NewInt64Coin("stake", 100)
	require.NotNil(t, node.IsValid())
	
	node.Type = ""
	require.NotNil(t, node.IsValid())
	
	node.Type = "typ"
	require.NotNil(t, node.IsValid())
	
	node.Type = "TypeTypeTypeTypeTypeTypeType"
	require.NotNil(t, node.IsValid())
	
	node.Type = "Type"
	node.Version = ""
	require.NotNil(t, node.IsValid())
	
	node.Version = "Ver"
	require.NotNil(t, node.IsValid())
	
	node.Version = "VersionVersionVersionVersionVersion"
	require.NotNil(t, node.IsValid())
	
	node.Version = "Version"
	node.Moniker = ""
	require.NotNil(t, node.IsValid())
	
	node.Moniker = "Mon"
	require.NotNil(t, node.IsValid())
	
	node.Moniker = "MonikerMonikerMonikerMonikerMoniker"
	require.NotNil(t, node.IsValid())
	
	node.Moniker = "Moniker"
	node.PricesPerGB = nil
	require.NotNil(t, node.IsValid())
	
	node.PricesPerGB = sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(0)))
	require.NotNil(t, node.IsValid())
	
	node.PricesPerGB = sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10)))
	node.InternetSpeed = TestBandwidthNeg
	require.NotNil(t, node.IsValid())
	
	node.InternetSpeed = TestBandwidthZero
	require.NotNil(t, node.IsValid())
	
	node.InternetSpeed = TestBandwidthPos1
	require.NotNil(t, node.IsValid())
	
	node.Encryption = ""
	require.NotNil(t, node.IsValid())
	
	node.Encryption = "Enc"
	require.NotNil(t, node.IsValid())
	
	node.Encryption = "EncryptionEncryptionEncryptionEncryption"
	require.NotNil(t, node.IsValid())
	
	node.Encryption = "Encryption"
	node.Status = ""
	require.NotNil(t, node.IsValid())
	
	node.Status = StatusDeRegistered
	require.Nil(t, node.IsValid())
	
	node.Status = StatusRegistered
	require.Nil(t, node.IsValid())
	
}

func TestIsFreeClient(t *testing.T) {
	require.False(t, IsFreeClient(nil, nil))
	require.False(t, IsFreeClient([]sdk.AccAddress{TestAddress1}, nil))
	require.False(t, IsFreeClient([]sdk.AccAddress{TestAddress1}, TestAddress2))
	require.True(t, IsFreeClient([]sdk.AccAddress{TestAddress1}, TestAddress1))
}

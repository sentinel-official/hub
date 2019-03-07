package types

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewBandwidthFromInt64(t *testing.T) {
	b1 := NewBandwidthFromInt64(100, 100)
	require.False(t, b1.Equal(Bandwidth{csdkTypes.NewInt(100), csdkTypes.NewInt(10)}))
	require.False(t, b1.Equal(Bandwidth{csdkTypes.NewInt(10), csdkTypes.NewInt(100)}))
	require.True(t, b1.Equal(Bandwidth{csdkTypes.NewInt(100), csdkTypes.NewInt(100)}))
	require.True(t, b1.IsPositive())
	require.False(t, b1.IsZero())
	require.False(t, b1.IsNegative())
	require.False(t, b1.IsNil())

	b2 := NewBandwidthFromInt64(-100, -100)
	require.False(t, b2.Equal(Bandwidth{csdkTypes.NewInt(-100), csdkTypes.NewInt(-10)}))
	require.False(t, b2.Equal(Bandwidth{csdkTypes.NewInt(-10), csdkTypes.NewInt(-100)}))
	require.True(t, b2.Equal(Bandwidth{csdkTypes.NewInt(-100), csdkTypes.NewInt(-100)}))
	require.False(t, b2.IsPositive())
	require.False(t, b2.IsZero())
	require.True(t, b2.IsNegative())
	require.False(t, b2.IsNil())

	require.False(t, b1.LTE(b2))
	require.False(t, b1.LT(b2))
	require.True(t, b2.LTE(b1))
	require.True(t, b2.LT(b1))
}

func TestNewBandwidthSignData(t *testing.T) {
	b := NewBandwidthSignData(testID, testBandwidth, testAddress1, testAddress2)
	require.Equal(t, b.ID, testID)
	require.Equal(t, b.Bandwidth, testBandwidth)
	require.Equal(t, b.NodeOwner, testAddress1)
	require.Equal(t, b.Client, testAddress2)
	require.NotNil(t, b.GetBytes())
}

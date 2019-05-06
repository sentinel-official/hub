package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionDetails_Amount(t *testing.T) {
	session := TestSessionValid
	require.Equal(t, session.Amount(), TestCoinZero)

	session.BandwidthInfo.Consumed = TestBandwidthPos1
	require.Equal(t, session.Amount(), TestCoinPos)

	session.BandwidthInfo.Consumed = TestBandwidthPos2
	session.LockedAmount = TestCoinPos.Add(TestCoinPos)
	require.Equal(t, session.Amount(), TestCoinPos.Add(TestCoinPos))
	require.NotEqual(t, session.Amount(), TestCoinPos)
	require.NotEqual(t, session.Amount(), TestCoinZero)
}

func TestSessionDetails_SetNewSessionBandwidth(t *testing.T) {
	session := TestSessionValid

	err := session.SetNewSessionBandwidth(TestBandwidthNeg, TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthZero, TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos1, TestNodeOwnerSignBandWidthZero, TestClientSignBandWidthPos1, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos1, TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthNeg, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos1, TestClientSignBandWidthPos1, TestNodeOwnerSignBandWidthPos1, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos1, TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0)
	require.Nil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos2, TestNodeOwnerSignBandWidthPos2, TestClientSignBandWidthPos2, 0)
	require.Nil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos1, TestNodeOwnerSignBandWidthPos1, TestClientSignBandWidthPos1, 0)
	require.NotNil(t, err)
	err = session.SetNewSessionBandwidth(TestBandwidthPos2, TestNodeOwnerSignBandWidthPos2, TestClientSignBandWidthPos2, 0)
	require.Nil(t, err)
}

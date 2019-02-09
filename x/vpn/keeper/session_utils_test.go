package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestVerifyAndUpdateSessionBandwidth(t *testing.T) {
	ctx, _, _, accountKeeper, _ := TestCreateInput()
	account := accountKeeper.NewAccountWithAddress(ctx, types.TestAddress1)
	require.Nil(t, account.SetPubKey(types.TestPubkey1))
	require.Nil(t, account.SetCoins(types.TestCoinsPos))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, account.GetCoins())

	account = accountKeeper.NewAccountWithAddress(ctx, types.TestAddress2)
	require.Nil(t, account.SetPubKey(types.TestPubkey2))
	accountKeeper.SetAccount(ctx, account)

	account = accountKeeper.GetAccount(ctx, types.TestAddress2)
	require.Equal(t, types.TestPubkey2, account.GetPubKey())

	sign := types.NewBandwidthSign(types.TestSessionIDValid, types.TestBandwidthZero, types.TestAddress1, types.TestAddress2)
	signBytes, err := sign.GetBytes()
	require.Nil(t, err)
	nodeOwnerSign, err1 := types.TestPrivKey1.Sign(signBytes)
	require.Nil(t, err1)
	clientSign, err1 := types.TestPrivKey2.Sign(signBytes)
	require.Nil(t, err1)

	require.Equal(t, types.ErrorInvalidBandwidth(),
		VerifyAndUpdateSessionBandwidth(ctx, accountKeeper, &TestSessionValid, sign, clientSign, nodeOwnerSign))

	sign = types.NewBandwidthSign(types.TestSessionIDValid,
		sdkTypes.NewBandwidthFromInt64(types.TestUploadPos.Int64()+1, types.TestDownloadPos.Int64()+1),
		types.TestAddress1, types.TestAddress2)
	signBytes, err = sign.GetBytes()
	require.Nil(t, err)
	nodeOwnerSign, err1 = types.TestPrivKey1.Sign(signBytes)
	require.Nil(t, err1)
	clientSign, err1 = types.TestPrivKey2.Sign(signBytes)
	require.Nil(t, err1)

	require.Equal(t, types.ErrorInvalidBandwidth(),
		VerifyAndUpdateSessionBandwidth(ctx, accountKeeper, &TestSessionValid, sign, clientSign, nodeOwnerSign))

	sign = types.NewBandwidthSign(types.TestSessionIDValid, types.TestBandwidthPos, types.TestAddress1, types.TestAddress2)
	signBytes, err = sign.GetBytes()
	require.Nil(t, err)
	nodeOwnerSign, err1 = types.TestPrivKey1.Sign(signBytes)
	require.Nil(t, err1)
	clientSign, err1 = types.TestPrivKey2.Sign(signBytes)
	require.Nil(t, err1)

	require.Equal(t, types.ErrorInvalidSign(),
		VerifyAndUpdateSessionBandwidth(ctx, accountKeeper, &TestSessionValid, sign, []byte("invalid"), nodeOwnerSign))
	require.Equal(t, types.ErrorInvalidSign(),
		VerifyAndUpdateSessionBandwidth(ctx, accountKeeper, &TestSessionValid, sign, clientSign, []byte("invalid")))
	require.Nil(t, VerifyAndUpdateSessionBandwidth(ctx, accountKeeper, &TestSessionValid, sign, clientSign, nodeOwnerSign))
	require.Equal(t, sign.Bandwidth, TestSessionValid.Bandwidth.Consumed)
	require.Equal(t, clientSign, TestSessionValid.Bandwidth.ClientSign)
	require.Equal(t, nodeOwnerSign, TestSessionValid.Bandwidth.NodeOwnerSign)
	require.Equal(t, types.StatusActive, TestSessionValid.Status)
	require.Equal(t, int64(0), TestSessionValid.StatusAtHeight)
}

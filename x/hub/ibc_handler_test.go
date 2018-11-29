package hub

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestMsgLockCoinsHandler(t *testing.T) {
	cdc := codec.New()
	cdc.RegisterInterface((*types.Interface)(nil), nil)
	cdc.RegisterConcrete(MsgLockerStatus{}, "", nil)

	multiStore, authKey, hubKey, ibcKey := setupMultiStore()

	ctx := csdkTypes.NewContext(multiStore, abci.Header{}, false, log.NewNopLogger())
	auth.RegisterBaseAccount(cdc)

	accountMapper := auth.NewAccountKeeper(cdc, authKey, auth.ProtoBaseAccount)

	account1 := auth.NewBaseAccountWithAddress(csdkTypes.AccAddress(pk1.Address()))
	account1.SetCoins(coins1)
	account1.SetPubKey(pk1)
	accountMapper.SetAccount(ctx, &account1)

	account2 := auth.NewBaseAccountWithAddress(csdkTypes.AccAddress(pk2.Address()))
	account2.SetCoins(coins1)
	account2.SetPubKey(pk2)
	accountMapper.SetAccount(ctx, &account2)

	bankKeeper := bank.NewBaseKeeper(accountMapper)
	keeper := NewBaseKeeper(cdc, hubKey, bankKeeper)
	ibcKeeper := ibc.NewKeeper(ibcKey, cdc)
	handler := NewIBCHubHandler(ibcKeeper, keeper)

	msgLockCoins1 := TestNewMsgIBCTransactionForLockCoins1()
	msgLockCoinsRes1 := handler(ctx, msgLockCoins1)
	require.EqualValues(t, csdkTypes.ToABCICode(codeSpaceHub, types.ErrCodeIBCPacketMsgVerificationFailed), msgLockCoinsRes1.Code)

	msgLockCoins2 := TestNewMsgIBCTransactionForLockCoins2()
	msgLockCoinsRes2 := handler(ctx, msgLockCoins2)
	require.True(t, msgLockCoinsRes2.IsOK(), "expected coins to lock but %v got", msgLockCoinsRes2)

	msgReleaseCoins := TestNewMsgIBCTransactionForReleaseCoins()
	msgReleaseCoinsRes := handler(ctx, msgReleaseCoins)
	require.True(t, msgReleaseCoinsRes.IsOK(), "expected coins to release but %v got", msgReleaseCoinsRes)

	getAccount2 := accountMapper.GetAccount(ctx, csdkTypes.AccAddress(pk2.Address()))
	require.Equal(t, getAccount2.GetCoins(), coins1)

}

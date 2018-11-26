package hub

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

//TestLockCoins: lock the coins respective address based on the lockerId
//and release the coins
func TestLockCoins(t *testing.T) {
	cdc := codec.New()

	multiStore, authkey, hubKey := setupMultiStore()
	ctx := csdkTypes.NewContext(multiStore, abci.Header{}, false, log.NewNopLogger())
	auth.RegisterBaseAccount(cdc)

	accountMapper := auth.NewAccountKeeper(cdc, authkey, auth.ProtoBaseAccount)
	bankKeeper := bank.NewBaseKeeper(accountMapper)
	keeper := NewBaseKeeper(cdc, hubKey, bankKeeper)

	account1 := auth.NewBaseAccountWithAddress(addr1)
	account1.SetCoins(coins1)
	accountMapper.SetAccount(ctx, &account1)

	account2 := auth.NewBaseAccountWithAddress(addr2)
	accountMapper.SetAccount(ctx, &account2)

	account3 := auth.NewBaseAccountWithAddress(addr3)
	accountMapper.SetAccount(ctx, &account3)

	err := keeper.LockCoins(ctx, lockerId, addr1, coins2)
	require.Nil(t, err)

	err = keeper.LockCoins(ctx, "", addr1, csdkTypes.Coins{coin3})
	require.NotNil(t, err)

	getAccount1 := accountMapper.GetAccount(ctx, addr1)
	require.Equal(t, getAccount1.GetCoins(), csdkTypes.Coins{coin1, coin2})

	locker := keeper.GetLocker(ctx, lockerId)
	require.Equal(t, locker.Coins, coins2)

	err = keeper.ReleaseCoins(ctx, lockerId)
	require.Nil(t, err)

	//check releasecoins called multiple times
	//err=keeper.ReleaseCoins(ctx,lockerId)
	//require.Nil(t,err)
	//
	//err = keeper.ReleaseCoins(ctx, lockerId)
	//require.Nil(t,err)

	keeper.SetLocker(ctx, lockerId3, emptyLocker1)
	err = keeper.ReleaseCoins(ctx, lockerId3)
	require.NotNil(t, err)

	getAccount2 := accountMapper.GetAccount(ctx, addr1)
	require.Equal(t, getAccount2.GetCoins(), account1.GetCoins())

	err = keeper.ReleaseCoinsToMany(ctx, lockerId, []csdkTypes.AccAddress{addr1, addr2, addr3}, []csdkTypes.Coins{csdkTypes.Coins{
		csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))}})

	require.Nil(t, err)

	err = keeper.ReleaseCoinsToMany(ctx, lockerId, []csdkTypes.AccAddress{addr1, {}, addr3}, []csdkTypes.Coins{csdkTypes.Coins{
		csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(0))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(-100))}})

	require.NotNil(t, err)

	getAccount3 := accountMapper.GetAccount(ctx, addr3)
	require.Equal(t, getAccount3.GetCoins(), csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))})

	respose := keeper.GetLocker(ctx, "")
	require.Nil(t, respose)
}

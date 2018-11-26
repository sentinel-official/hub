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

	if err := keeper.LockCoins(ctx, lockerId, addr1, coins2); err != nil {
		panic(err)
	}

	getAccount1 := accountMapper.GetAccount(ctx, addr1)
	require.Equal(t, getAccount1.GetCoins(), csdkTypes.Coins{coin1, coin2})

	locker := keeper.GetLocker(ctx, lockerId)
	require.Equal(t, locker.Coins, coins2)

	if err := keeper.ReleaseCoins(ctx, lockerId); err != nil {
		panic(err)
	}

	getAccountStatus := accountMapper.GetAccount(ctx, addr1)
	require.Equal(t, getAccountStatus.GetCoins(), account1.GetCoins())

	if err := keeper.ReleaseCoinsToMany(ctx, lockerId, []csdkTypes.AccAddress{addr1, addr2, addr3}, []csdkTypes.Coins{csdkTypes.Coins{
		csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))},
		csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))}}); err != nil {
		panic(err)
	}

	getAccount3 := accountMapper.GetAccount(ctx, addr3)
	require.Equal(t, getAccount3.GetCoins(), csdkTypes.Coins{csdkTypes.NewCoin("sent", csdkTypes.NewInt(10))})

}

package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_AddDeposit(t *testing.T) {
	ctx, _, keeper, bankKeeper := TestCreateInput()

	deposit, found := keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err := keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinEmpty)
	require.NotNil(t, err)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinNeg)
	require.NotNil(t, err)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinZero)
	require.NotNil(t, err)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinEmpty)
	require.NotNil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinNeg)
	require.NotNil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinZero)
	require.NotNil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos.Add(types.TestCoinPos))
	require.NotNil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)
}

func TestKeeper_SubtractDeposit(t *testing.T) {
	ctx, _, keeper, bankKeeper := TestCreateInput()

	coins := bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsEmpty, coins)
	deposit, found := keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err := keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinEmpty)
	require.NotNil(t, err)
	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinNeg)
	require.NotNil(t, err)
	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinZero)
	require.NotNil(t, err)
	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinNeg)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsNil, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinZero)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsNil, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos.Add(types.TestCoinPos))
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsNil, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsPos, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)
}

func TestKeeper_SendDeposit(t *testing.T) {
	ctx, _, keeper, bankKeeper := TestCreateInput()

	coins := bankKeeper.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, types.TestCoinsEmpty, coins)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsEmpty, coins)
	deposit, found := keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err := keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinEmpty)
	require.NotNil(t, err)
	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinNeg)
	require.NotNil(t, err)
	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinZero)
	require.NotNil(t, err)
	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinNeg)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsEmpty, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinZero)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsEmpty, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinPos.Add(types.TestCoinPos))
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsEmpty, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinPos)
	require.Nil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinPos)
	require.NotNil(t, err)
	coins = bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsNil, deposit.Coins)
}

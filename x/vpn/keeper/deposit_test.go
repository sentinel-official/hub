package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_AddDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)

	_, err := keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)
	deposit, found := keeper.depositKeeper.GetDeposit(ctx, TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsPos, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	_, err = keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsPos.Add(TestCoinsPos), deposit.Coins)
}

func TestKeeper_SubtractDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	_, err := keeper.SubtractDeposit(ctx, TestAddress1, TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)

	_, err = keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.SubtractDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)
}

func TestKeeper_SendDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos.Add(TestCoinsPos))
	_, err := keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.AddDeposit(ctx, TestAddress1, TestCoinPos)
	require.Nil(t, err)

	deposit, found := keeper.depositKeeper.GetDeposit(ctx, TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsPos.Add(TestCoinsPos), deposit.Coins)

	_, found = keeper.depositKeeper.GetDeposit(ctx, TestAddress2)
	require.Equal(t, false, found)

	_, err = keeper.SendDeposit(ctx, TestAddress1, TestAddress2, TestCoinPos)
	require.Nil(t, err)

	coins := bankKeeper.GetCoins(ctx, TestAddress2)
	require.Equal(t, TestCoinsPos, coins)
}

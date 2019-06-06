package keeper

import (
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_AddDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)

	_, err := keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)
	deposit, found := keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos, deposit.Coins)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)
	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)
	deposit, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)
}

func TestKeeper_SubtractDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	_, err := keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.NotNil(t, err)

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.SubtractDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)
}

func TestKeeper_SendDeposit(t *testing.T) {
	ctx, _, _, keeper, _, bankKeeper := TestCreateInput()

	bankKeeper.AddCoins(ctx, types.TestAddress1, types.TestCoinsPos.Add(types.TestCoinsPos))
	_, err := keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)

	_, err = keeper.AddDeposit(ctx, types.TestAddress1, types.TestCoinPos)
	require.Nil(t, err)

	deposit, found := keeper.depositKeeper.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.TestCoinsPos.Add(types.TestCoinsPos), deposit.Coins)

	_, found = keeper.depositKeeper.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)

	_, err = keeper.SendDeposit(ctx, types.TestAddress1, types.TestAddress2, types.TestCoinPos)
	require.Nil(t, err)

	coins := bankKeeper.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, types.TestCoinsPos, coins)
}

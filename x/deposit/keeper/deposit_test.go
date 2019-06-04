package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKeeper_SetDeposit(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, TestDepositZero)
	deposit, found := depositKeeper.GetDeposit(ctx, TestDepositZero.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestDepositZero, deposit)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	deposit, found = depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestDepositPos, deposit)

	coins = bankKeeper.GetCoins(ctx, TestDepositPos.Address)
	require.Equal(t, TestDepositPos.Coins, coins)
}

func TestKeeper_GetDeposit(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	_, found := depositKeeper.GetDeposit(ctx, TestAddress1)
	require.Equal(t, false, found)

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	deposit, found := depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestDepositPos, deposit)
}

func TestKeeper_GetAllDeposits(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	deposits := depositKeeper.GetAllDeposits(ctx)
	require.Equal(t, TestDepositsNil, deposits)

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	deposits = depositKeeper.GetAllDeposits(ctx)
	require.Equal(t, TestDepositsPos, deposits)
}

func TestKeeper_Add(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	deposit, found := depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestDepositPos, deposit)

	_, err = depositKeeper.Add(ctx, TestDepositPos.Address, TestCoinsPos)
	require.Nil(t, err)

	deposit, found = depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsPos.Add(TestCoinsPos), deposit.Coins)
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	deposit, found := depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestDepositPos, deposit)

	_, err = depositKeeper.Subtract(ctx, TestAddress1, coins)
	deposit, found = depositKeeper.GetDeposit(ctx, TestDepositPos.Address)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsNil, deposit.Coins)
}

func TestKeeper_Send(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	coins = bankKeeper.GetCoins(ctx, TestAddress2)
	require.Equal(t, TestCoinsEmpty, coins)

	depositKeeper.SetDeposit(ctx, TestDepositPos)
	_, err = depositKeeper.Send(ctx, TestAddress1, TestAddress2, TestCoinsPos)
	require.Nil(t, err)

	deposit, found := depositKeeper.GetDeposit(ctx, TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, TestCoinsNil, deposit.Coins)

	coins = bankKeeper.GetCoins(ctx, TestAddress2)
	require.Equal(t, TestCoinsPos, coins)
}

func TestKeeper_Receive(t *testing.T) {
	ctx, _, depositKeeper, _, bankKeeper := TestCreateInput()

	coins, _, err := bankKeeper.AddCoins(ctx, TestAddress1, TestCoinsPos)
	require.Nil(t, err)
	require.Equal(t, TestCoinsPos, coins)

	_, err = depositKeeper.Receive(ctx, TestAddress1, TestAddress2, TestCoinsPos)
	require.Nil(t, err)

	deposit, found := depositKeeper.GetDeposit(ctx, TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, TestAddress2, deposit.Address)
	require.Equal(t, TestCoinsPos, deposit.Coins)

}

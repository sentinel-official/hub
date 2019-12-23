package keeper

import (
	"testing"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	
	"github.com/sentinel-official/hub/x/deposit/types"
)

func TestKeeper_SetDeposit(t *testing.T) {
	ctx, dk, _ := CreateTestInput(t, false)
	
	deposit, found := dk.GetDeposit(ctx, types.Deposit{}.Address)
	require.Equal(t, false, found)
	require.Equal(t, types.Deposit{}, deposit)
	dk.SetDeposit(ctx, types.Deposit{})
	deposit, found = dk.GetDeposit(ctx, types.Deposit{}.Address)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{}, deposit)
	
	dk.SetDeposit(ctx, types.Deposit{Coins: sdk.Coins(nil)})
	deposit, found = dk.GetDeposit(ctx, types.Deposit{Coins: sdk.Coins(nil)}.Address)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{Coins: sdk.Coins(nil)}, deposit)
	
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.Deposit{}, deposit)
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}}})
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}}}, deposit)
	
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 0)}})
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 0)}}, deposit)
	
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	deposit, found = dk.GetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}.Address)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
}

func TestKeeper_GetDeposit(t *testing.T) {
	TestKeeper_SetDeposit(t)
}

func TestKeeper_GetAllDeposits(t *testing.T) {
	ctx, dk, _ := CreateTestInput(t, false)
	
	deposits := dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 0)
	require.Equal(t, []types.Deposit(nil), deposits)
	
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	deposits = dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 1)
	require.Equal(t, []types.Deposit{{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}}, deposits)
	
	depositPos2 := types.Deposit{Address: types.TestAddress2, Coins: sdk.Coins{sdk.NewInt64Coin("stake", 10)}}
	dk.SetDeposit(ctx, depositPos2)
	deposits = dk.GetAllDeposits(ctx)
	require.Len(t, deposits, 2)
}

func TestKeeper_Add(t *testing.T) {
	ctx, dk, bk := CreateTestInput(t, false)
	
	err := dk.Add(ctx, types.TestAddress1, sdk.Coins{})
	require.Nil(t, err)
	deposit, found := dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins(nil)}, deposit)
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins(nil))
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins(nil)}, deposit)
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins(nil)}, deposit)
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins(nil)}, deposit)
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins(nil)}, deposit)
	
	coins, err := bk.AddCoins(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, bk.GetCoins(ctx, types.TestAddress1))
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	require.Equal(t, sdk.Coins(nil), bk.GetCoins(ctx, types.TestAddress1))
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	require.Equal(t, sdk.Coins(nil), bk.GetCoins(ctx, types.TestAddress1))
	
	coinsPos2 := sdk.Coins{sdk.NewInt64Coin("stake", 10)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	depositPos2 := types.Deposit{Address: types.TestAddress1, Coins: coinsPos2}
	
	coins, err = bk.AddCoins(ctx, types.TestAddress1, coinsPos2)
	require.Nil(t, err)
	require.Equal(t, coinsPos2, bk.GetCoins(ctx, types.TestAddress1))
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos2, deposit)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, bk.GetCoins(ctx, types.TestAddress1))
	
	coinsPos3 := coinsPos2.Add(sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	depositPos3 := types.Deposit{Address: types.TestAddress1, Coins: coinsPos3}
	
	err = dk.Add(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, depositPos3, deposit)
	require.Equal(t, sdk.Coins(nil), bk.GetCoins(ctx, types.TestAddress1))
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, dk, bk := CreateTestInput(t, false)
	
	deposit, found := dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.Deposit{}, deposit)
	
	err := dk.Subtract(ctx, types.TestAddress1, sdk.Coins{})
	require.NotNil(t, err)
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins(nil))
	require.NotNil(t, err)
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.Deposit{}, deposit)
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins(nil))
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.Subtract(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, deposit.Coins)
	coins := bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins(nil), coins)
}

func TestKeeper_SendFromDepositToAccount(t *testing.T) {
	ctx, dk, bk := CreateTestInput(t, false)
	
	deposit, found := dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, false, found)
	require.Equal(t, types.Deposit{}, deposit)
	
	err := dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{})
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins(nil))
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{})
	require.Nil(t, err)
	coins := bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins(nil))
	require.Nil(t, err)
	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	err = dk.SendFromDepositToAccount(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, types.TestAddress2)
	require.Equal(t, sdk.Coins(nil), coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, deposit.Coins)
}

func TestKeeper_ReceiveFromAccountToDeposit(t *testing.T) {
	ctx, dk, bk := CreateTestInput(t, false)
	
	deposit, found := dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, false, found)
	coins := bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{}, coins)
	
	err := dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{})
	require.Nil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins(nil))
	require.Nil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
	coins, err = bk.AddCoins(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins(nil))
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-10)}})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 0)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins(nil), coins)
	
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins(nil), coins)
	
	coins, err = bk.AddCoins(ctx, types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 10)}))
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 10)}), coins)
	err = dk.ReceiveFromAccountToDeposit(ctx, types.TestAddress1, types.TestAddress2, sdk.Coins{sdk.NewInt64Coin("stake", 10)})
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, types.TestAddress2)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}.Add(sdk.Coins{sdk.NewInt64Coin("stake", 10)}), deposit.Coins)
	coins = bk.GetCoins(ctx, types.TestAddress1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 10)}, coins)
}

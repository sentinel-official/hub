package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/hub/x/deposit/types"
)

func TestKeeper_SetDeposit(t *testing.T) {
	ctx, keeper, _ := CreateTestInput(t, false)

	deposit, found := keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.False(t, found)
	require.Equal(t, types.Deposit{}, deposit)

	keeper.SetDeposit(ctx, types.Deposit{})
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress{})
	require.True(t, found)
	require.Equal(t, deposit, types.Deposit{})

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1")})
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1")}, deposit)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}})
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}}, deposit)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}})
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}}, deposit)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}})
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)
}

func TestKeeper_GetDeposit(t *testing.T) {
	TestKeeper_SetDeposit(t)
}

func TestKeeper_GetDeposits(t *testing.T) {
	ctx, keeper, _ := CreateTestInput(t, false)

	deposits := keeper.GetDeposits(ctx)
	require.Len(t, deposits, 0)
	require.Equal(t, types.Deposits(nil), deposits)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}})
	deposits = keeper.GetDeposits(ctx)
	require.Len(t, deposits, 1)
	require.Equal(t, types.Deposits{types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}}}, deposits)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}})
	deposits = keeper.GetDeposits(ctx)
	require.Len(t, deposits, 1)
	require.Equal(t, types.Deposits{types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}}}, deposits)

	keeper.SetDeposit(ctx, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}})
	deposits = keeper.GetDeposits(ctx)
	require.Len(t, deposits, 2)
	require.Equal(t, types.Deposits{types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}}, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}}}, deposits)
}

func TestKeeper_Add(t *testing.T) {
	ctx, keeper, bankKeeper := CreateTestInput(t, false)

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress{}, sdk.Coins{}))
	deposit, found := keeper.GetDeposit(ctx, sdk.AccAddress{})
	require.True(t, found)
	require.Equal(t, types.Deposit{Coins: sdk.Coins(nil)}, deposit)

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)

	require.NotNil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)

	require.NotNil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)

	require.NotNil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)

	coins, err := bankKeeper.AddCoins(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, coins)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.NotNil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	coins, err = bankKeeper.AddCoins(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(20)}})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(20)}}, coins)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(20)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(20)}}}, deposit)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))
}

func TestKeeper_Subtract(t *testing.T) {
	ctx, keeper, bankKeeper := CreateTestInput(t, false)

	deposit, found := keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.False(t, found)
	require.Equal(t, types.Deposit{}, deposit)

	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{}))
	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.False(t, found)
	require.Equal(t, types.Deposit{}, deposit)

	coins, err := bankKeeper.AddCoins(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, coins)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.Nil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.NotNil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.Nil(t, keeper.Subtract(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))
}

func TestKeeper_SendCoinsFromDepositToAccount(t *testing.T) {
	ctx, keeper, bankKeeper := CreateTestInput(t, false)

	deposit, found := keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.False(t, found)
	require.Equal(t, types.Deposit{}, deposit)

	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{}))
	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))

	coins, err := bankKeeper.AddCoins(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, coins)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.Nil(t, keeper.Add(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.Nil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{}))
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-2")))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-2")))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.NotNil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-2")))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)

	require.Nil(t, keeper.SendCoinsFromDepositToAccount(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-2")))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-1"), Coins: sdk.Coins(nil)}, deposit)
}

func TestKeeper_SendCoinsFromAccountToDeposit(t *testing.T) {
	ctx, keeper, bankKeeper := CreateTestInput(t, false)

	deposit, found := keeper.GetDeposit(ctx, sdk.AccAddress("address-1"))
	require.False(t, found)
	require.Equal(t, types.Deposit{}, deposit)

	require.Nil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{}))
	require.NotNil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	require.NotNil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	require.NotNil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))

	coins, err := bankKeeper.AddCoins(ctx, sdk.AccAddress("address-1"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, coins)

	require.Nil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-2"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins(nil)}, deposit)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.NotNil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(-10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-2"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins(nil)}, deposit)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.NotNil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-2"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins(nil)}, deposit)
	require.Equal(t, sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}, bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))

	require.Nil(t, keeper.SendCoinsFromAccountToDeposit(ctx, sdk.AccAddress("address-1"), sdk.AccAddress("address-2"), sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}))
	deposit, found = keeper.GetDeposit(ctx, sdk.AccAddress("address-2"))
	require.True(t, found)
	require.Equal(t, types.Deposit{Address: sdk.AccAddress("address-2"), Coins: sdk.Coins{sdk.Coin{Denom: "one", Amount: sdk.NewInt(10)}}}, deposit)
	require.Equal(t, sdk.Coins(nil), bankKeeper.GetCoins(ctx, sdk.AccAddress("address-1")))
}

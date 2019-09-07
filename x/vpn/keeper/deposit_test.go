package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_AddDeposit(t *testing.T) {
	ctx, k, dk, bk := CreateTestInput(t, false)

	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err := k.AddDeposit(ctx, address1, sdk.Coin{})
	require.NotNil(t, err)
	err = k.AddDeposit(ctx, address1, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)

	coins, err := bk.AddCoins(ctx, address1, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)

	err = k.AddDeposit(ctx, address1, sdk.Coin{})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.AddDeposit(ctx, address1, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100).Add(sdk.NewInt64Coin("stake", 100)))
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)
}

func TestKeeper_SubtractDeposit(t *testing.T) {
	ctx, k, dk, bk := CreateTestInput(t, false)

	coins := bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err := k.SubtractDeposit(ctx, address1, sdk.Coin{})
	require.NotNil(t, err)
	err = k.SubtractDeposit(ctx, address1, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)

	coins, err = bk.AddCoins(ctx, address1, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)

	err = k.SubtractDeposit(ctx, address1, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100).Add(sdk.NewInt64Coin("stake", 100)))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins(nil), coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.SubtractDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
}

func TestKeeper_SendDeposit(t *testing.T) {
	ctx, k, dk, bk := CreateTestInput(t, false)

	coins := bk.GetCoins(ctx, address1)
	require.Equal(t, sdk.Coins{}, coins)

	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found := dk.GetDeposit(ctx, address1)
	require.Equal(t, false, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err := k.SendDeposit(ctx, address1, address2, sdk.Coin{})
	require.NotNil(t, err)
	err = k.SendDeposit(ctx, address1, address2, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)

	coins, err = bk.AddCoins(ctx, address1, sdk.Coins{sdk.NewInt64Coin("stake", 100)})
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	err = k.AddDeposit(ctx, address1, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)

	err = k.SendDeposit(ctx, address1, address2, sdk.Coin{Denom: "stake", Amount: sdk.NewInt(-100)})
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 0))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 100).Add(sdk.NewInt64Coin("stake", 100)))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, deposit.Coins)

	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 100))
	require.Nil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)

	err = k.SendDeposit(ctx, address1, address2, sdk.NewInt64Coin("stake", 100))
	require.NotNil(t, err)
	coins = bk.GetCoins(ctx, address2)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin("stake", 100)}, coins)
	deposit, found = dk.GetDeposit(ctx, address1)
	require.Equal(t, true, found)
	require.Equal(t, sdk.Coins(nil), deposit.Coins)
}

package querier

import (
	"fmt"
	"testing"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	
	"github.com/sentinel-official/hub/x/deposit/keeper"
	"github.com/sentinel-official/hub/x/deposit/types"
)

func Test_queryDepositOfAddress(t *testing.T) {
	ctx, dk, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	req := abci.RequestQuery{
		Path: fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDepositOfAddress),
		Data: []byte{},
	}
	
	res, err := queryDepositOfAddress(ctx, req, dk)
	require.NotNil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams([]byte("")))
	
	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	var deposit types.Deposit
	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
	require.NotEqual(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	
	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams([]byte("")))
	
	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.Len(t, res, 0)
	
	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
	require.NotEqual(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(types.TestAddress1))
	require.Nil(t, err)
	
	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)
	
	cdc.MustUnmarshalJSON(res, &deposit)
	require.Nil(t, err)
	require.Equal(t, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}, deposit)
	
	req.Data = cdc.MustMarshalJSON(types.NewQueryDepositOfAddressParams(types.TestAddress2))
	require.Nil(t, err)
	
	res, err = queryDepositOfAddress(ctx, req, dk)
	require.Nil(t, err)
	require.Equal(t, []byte(nil), res)
	require.NotNil(t, cdc.UnmarshalJSON(res, &deposit))
}

func Test_queryAllDeposits(t *testing.T) {
	ctx, dk, _ := keeper.CreateTestInput(t, false)
	cdc := keeper.MakeTestCodec()
	
	res, err := queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.Equal(t, []byte("null"), res)
	
	var deposits []types.Deposit
	cdc.MustUnmarshalJSON(res, &deposits)
	require.NotEqual(t, []types.Deposit{{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}}, deposits)
	
	dk.SetDeposit(ctx, types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}})
	
	res, err = queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)
	
	cdc.MustUnmarshalJSON(res, &deposits)
	require.Equal(t, []types.Deposit{{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}}, deposits)
	
	deposit := types.Deposit{types.TestAddress1, sdk.Coins{sdk.NewInt64Coin("stake", 10)}}
	deposit.Address = types.TestAddress2
	dk.SetDeposit(ctx, deposit)
	
	res, err = queryAllDeposits(ctx, dk)
	require.Nil(t, err)
	require.NotEqual(t, []byte(nil), res)
	
	cdc.MustUnmarshalJSON(res, &deposits)
	require.Nil(t, err)
	require.Len(t, deposits, 2)
}

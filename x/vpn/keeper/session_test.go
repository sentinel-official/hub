package keeper

import (
	"testing"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

func TestKeeper_AddSession(t *testing.T) {
	keeper, ctx := CreateTestInput()

	tags, err := keeper.AddSession(ctx, &ParamsOfSessionDetails()[0])
	require.Nil(t, err)
	require.Equal(t, csdkTypes.MakeTag("session_id", []byte("cosmos1vdkxjetww3qkgerjv4ehxtf30dsznq/0")), tags[0])
}

func TestKeeper_SessionDetails(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, sessionDetails := range ParamsOfSessionDetails() {
		err := keeper.SetSessionDetails(ctx, &sessionDetails)
		require.Nil(t, err)
	}

	res, err := keeper.GetSessionDetails(ctx, types.NewSessionID("new-session-id-2"))
	require.Nil(t, err)
	require.Equal(t, types.NewNodeID("new-node-id-1"), res.NodeID)

	res, err = keeper.GetSessionDetails(ctx, "id/0")
	require.Nil(t, res)
}

func TestKeeper_SessionsCount(t *testing.T) {
	keeper, ctx := CreateTestInput()

	res, err := keeper.GetSessionsCount(ctx, types.ClientAddress1)
	require.Equal(t, uint64(0), res)
	require.Nil(t, err)

	err = keeper.SetSessionsCount(ctx, types.ClientAddress2, uint64(10))
	require.Nil(t, err)

	res, err = keeper.GetSessionsCount(ctx, types.ClientAddress2)
	require.Equal(t, uint64(10), res)
}

func TestKeeper_ActiveSessionIDsAtHeight(t *testing.T) {
	keeper, ctx := CreateTestInput()

	for _, sessionDetails := range ParamsOfSessionDetails() {
		err := keeper.AddActiveSessionIDsAtHeight(ctx, sessionDetails.StartedAtHeight, sessionDetails.ID)
		require.Nil(t, err)
	}

	res, err := keeper.GetActiveSessionIDsAtHeight(ctx, 10)
	require.Nil(t, err)
	require.Equal(t, 2, len(res))

	err = keeper.RemoveActiveSessionIDsAtHeight(ctx, 10, ParamsOfSessionDetails()[0].ID)
	require.Nil(t, err)

	res, err = keeper.GetActiveSessionIDsAtHeight(ctx, 10)
	require.Nil(t, err)
	require.Equal(t, types.NewSessionID("new-session-id-2"), res[0])
	require.Equal(t, 1, len(res))

	err = keeper.AddActiveSessionIDsAtHeight(ctx, 10, types.NewSessionID("new-session-id-2"))
	require.Equal(t, nil, err)
}

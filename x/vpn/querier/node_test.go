package querier

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/types"

	vk "github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	vpnTypes "github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	store = "custom"
	acc   = "vpn"
)

func TestNewQueryNodeParams(t *testing.T) {
	keeper, ctx := vk.CreateTestInput()
	cdc := vk.MakeCdc()
	qurier := NewQuerier(keeper, cdc)

	for _, nodeDetails := range vk.ParamsOfNodeDetails() {
		err := keeper.SetNodeDetails(ctx, &nodeDetails)
		require.Nil(t, err)
	}

	query := types.RequestQuery{
		Path: strings.Join([]string{store, acc, QueryNode}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryNodeParams(vpnTypes.NewNodeID("new-node-id-1"))),
	}

	res, err := qurier(ctx, []string{QueryNode}, query)
	require.Nil(t, err)

	var nodeDetails vpnTypes.NodeDetails
	err1 := cdc.UnmarshalJSON(res, &nodeDetails)
	require.Equal(t, nil, err1)

	query.Data = []byte{}

	res, err = qurier(ctx, []string{QueryNode}, query)
	require.NotNil(t, err)
}

func TestNewQueryNodesOfOwnerParams(t *testing.T) {
	keeper, ctx := vk.CreateTestInput()
	cdc := vk.MakeCdc()
	querier := NewQuerier(keeper, cdc)

	for _, nodeDetails := range vk.ParamsOfNodeDetails() {
		tags, err := keeper.AddNode(ctx, &nodeDetails)
		require.Nil(t, err)
		require.NotNil(t, tags)
	}

	query := types.RequestQuery{
		Path: strings.Join([]string{store, acc, QueryNodesOfOwner}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryNodesOfOwnerParams(vpnTypes.NodeAddress1)),
	}

	res, err := querier(ctx, []string{QueryNodesOfOwner}, query)
	require.Nil(t, err)
	t.Log(res)

	var nodes []vpnTypes.NodeDetails
	err1 := cdc.UnmarshalJSON(res, &nodes)
	require.Nil(t, err1)
}

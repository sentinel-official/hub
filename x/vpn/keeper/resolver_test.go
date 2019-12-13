package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_Resolver(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	resolver, found := k.GetResolver(ctx, types.TestAddress2)
	require.False(t, found)
	require.Equal(t, types.Resolver{}, resolver)

	resolvers := k.GetAllResolvers(ctx)
	require.Equal(t, 0, len(resolvers))

	resolver = types.TestResolver

	resolver1 := resolver
	resolver1.Owner = types.TestAddress2

	k.SetResolver(ctx, resolver)
	k.SetResolver(ctx, resolver1)

	res, found := k.GetResolver(ctx, resolver.Owner)
	require.True(t, found)
	require.Equal(t, res, resolver)

	resolvers = k.GetAllResolvers(ctx)
	require.Equal(t, 2, len(resolvers))

	_, found = k.GetNodeOfResolver(ctx, resolver.Owner, hub.NewNodeID(1))
	require.False(t, found)

	k.SetResolverOfNode(ctx, hub.NewNodeID(1), resolver.Owner)
	k.SetNodeOfResolver(ctx, resolver.Owner, hub.NewNodeID(1))

	k.SetResolverOfNode(ctx, hub.NewNodeID(1), resolver1.Owner)
	k.SetNodeOfResolver(ctx, resolver1.Owner, hub.NewNodeID(1))

	addresses := k.GetResolversOfNode(ctx, hub.NewNodeID(1))
	require.Equal(t, 2, len(addresses))

	k.SetNodeOfResolver(ctx, resolver1.Owner, hub.NewNodeID(2))
	nodes := k.GetNodesOfResolver(ctx, resolver1.Owner)
	require.Equal(t, 2, len(nodes))

	k.GetNodeOfResolver(ctx, resolver.Owner, hub.NewNodeID(1))
	nodes = k.GetNodesOfResolver(ctx, resolver.Owner)
	require.Equal(t, 1, len(nodes))

	_, found = k.GetResolverOfNode(ctx, hub.NewNodeID(2), types.TestAddress1)
	require.False(t, found)

	k.SetResolverOfNode(ctx, hub.NewNodeID(2), types.TestAddress1)

	address, found := k.GetResolverOfNode(ctx, hub.NewNodeID(2), types.TestAddress1)
	require.False(t, found)
	require.Equal(t, address, types.TestAddress1)

	k.RemoveVPNNodeOnResolver(ctx, hub.NewNodeID(2), resolver.Owner)
}

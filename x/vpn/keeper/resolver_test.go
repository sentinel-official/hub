package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sentinel-official/hub/x/vpn/types"
)

func TestKeeper_Resolver(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	resolvers := k.GetAllResolvers(ctx)
	require.Equal(t, 0, len(resolvers))

	resolver := types.TestResolver

	resolver1 := resolver
	resolver1.Owner = types.TestAddress2

	k.SetResolver(ctx, resolver)
	k.SetResolver(ctx, resolver1)

	resolvers = k.GetAllResolvers(ctx)
	require.Equal(t, 2, len(resolvers))
}

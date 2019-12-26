package common

import (
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/client/context"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func QueryResolvers(ctx context.CLIContext, resolverID string) (types.Resolvers, error) {
	var res []byte
	var err error
	
	if resolverID != "" {
		
		_id, err := hub.NewResolverIDFromString(resolverID)
		if err != nil {
			return nil, err
		}
		bytes, err := ctx.Codec.MarshalJSON(_id)
		if err != nil {
			return nil, err
		}
		res, _, err = ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers), bytes)
		if err != nil {
			return nil, err
		}
	} else {
		res, _, err = ctx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryResolvers), nil)
		if err != nil {
			return nil, err
		}
	}
	
	var resolvers types.Resolvers
	if res == nil {
		return nil, fmt.Errorf("resolvers doesnot exist")
	}
	
	if err := ctx.Codec.UnmarshalJSON(res, &resolvers); err != nil {
		return nil, err
	}
	
	return resolvers, nil
}

// nolint:dupl
package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QuerySession                = "session"
	QuerySessionOfSubscription  = "sessionOfSubscription"
	QuerySessionsOfSubscription = "sessionsOfSubscription"
	QueryAllSessions            = "allSessions"
)

type QuerySessionParams struct {
	ID sdk.ID
}

func NewQuerySessionParams(id sdk.ID) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

// nolint:dupl
func querySession(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

	var params QuerySessionParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	session, found := k.GetSession(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(session)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySessionOfSubscriptionPrams struct {
	ID    sdk.ID
	Index uint64
}

func NewQuerySessionOfSubscriptionPrams(id sdk.ID, index uint64) QuerySessionOfSubscriptionPrams {
	return QuerySessionOfSubscriptionPrams{
		ID:    id,
		Index: index,
	}
}

func querySessionOfSubscription(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

	var params QuerySessionOfSubscriptionPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	id, found := k.GetSessionIDBySubscriptionID(ctx, params.ID, params.Index)
	if !found {
		return nil, nil
	}

	session, found := k.GetSession(ctx, id)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(session)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySessionsOfSubscriptionPrams struct {
	ID sdk.ID
}

func NewQuerySessionsOfSubscriptionPrams(id sdk.ID) QuerySessionsOfSubscriptionPrams {
	return QuerySessionsOfSubscriptionPrams{
		ID: id,
	}
}

func querySessionsOfSubscription(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

	var params QuerySessionsOfSubscriptionPrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	sessions := k.GetSessionsOfSubscription(ctx, params.ID)

	res, resErr := cdc.MarshalJSON(sessions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllSessions(ctx csdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdk.Error) {
	sessions := k.GetAllSessions(ctx)

	res, resErr := cdc.MarshalJSON(sessions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

// nolint:dupl
package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/keeper"
	"github.com/sentinel-official/hub/x/vpn/types"
)

const (
	QuerySession                = "session"
	QuerySessionOfSubscription  = "sessionOfSubscription"
	QuerySessionsOfSubscription = "sessionsOfSubscription"
	QueryAllSessions            = "allSessions"
)

type QuerySessionParams struct {
	ID hub.ID
}

func NewQuerySessionParams(id hub.ID) QuerySessionParams {
	return QuerySessionParams{
		ID: id,
	}
}

// nolint:dupl
func querySession(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

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
	ID    hub.ID
	Index uint64
}

func NewQuerySessionOfSubscriptionPrams(id hub.ID, index uint64) QuerySessionOfSubscriptionPrams {
	return QuerySessionOfSubscriptionPrams{
		ID:    id,
		Index: index,
	}
}

func querySessionOfSubscription(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

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
	ID hub.ID
}

func NewQuerySessionsOfSubscriptionPrams(id hub.ID) QuerySessionsOfSubscriptionPrams {
	return QuerySessionsOfSubscriptionPrams{
		ID: id,
	}
}

func querySessionsOfSubscription(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

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

func queryAllSessions(ctx sdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, sdk.Error) {
	sessions := k.GetAllSessions(ctx)

	res, resErr := cdc.MarshalJSON(sessions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/sentinel-hub/types"
	"github.com/sentinel-official/sentinel-hub/x/vpn/keeper"
	"github.com/sentinel-official/sentinel-hub/x/vpn/types"
)

const (
	QuerySubscription                = "subscription"
	QuerySubscriptionsOfNode         = "subscriptionsOfNode"
	QuerySubscriptionsOfAddress      = "subscriptionsOfAddress"
	QueryAllSubscriptions            = "allSubscriptions"
	QuerySessionsCountOfSubscription = "sessionsCountOfSubscription"
)

type QuerySubscriptionParams struct {
	ID hub.ID
}

func NewQuerySubscriptionParams(id hub.ID) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

// nolint:dupl
func querySubscription(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

	var params QuerySubscriptionParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscription, found := k.GetSubscription(ctx, params.ID)
	if !found {
		return nil, nil
	}

	res, resErr := cdc.MarshalJSON(subscription)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySubscriptionsOfNodePrams struct {
	ID hub.ID
}

func NewQuerySubscriptionsOfNodePrams(id hub.ID) QuerySubscriptionsOfNodePrams {
	return QuerySubscriptionsOfNodePrams{
		ID: id,
	}
}

func querySubscriptionsOfNode(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

	var params QuerySubscriptionsOfNodePrams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsOfNode(ctx, params.ID)

	res, resErr := cdc.MarshalJSON(subscriptions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySubscriptionsOfAddressParams struct {
	Address sdk.AccAddress
}

func NewQuerySubscriptionsOfAddressParams(address sdk.AccAddress) QuerySubscriptionsOfAddressParams {
	return QuerySubscriptionsOfAddressParams{
		Address: address,
	}
}

func querySubscriptionsOfAddress(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

	var params QuerySubscriptionsOfAddressParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	subscriptions := k.GetSubscriptionsOfAddress(ctx, params.Address)

	res, resErr := cdc.MarshalJSON(subscriptions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

func queryAllSubscriptions(ctx sdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, sdk.Error) {
	subscriptions := k.GetAllSubscriptions(ctx)

	res, resErr := cdc.MarshalJSON(subscriptions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySessionsCountOfSubscriptionParams struct {
	ID hub.ID
}

func NewQuerySessionsCountOfSubscriptionParams(id hub.ID) QuerySessionsCountOfSubscriptionParams {
	return QuerySessionsCountOfSubscriptionParams{
		ID: id,
	}
}

func querySessionsCountOfSubscription(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, sdk.Error) {

	var params QuerySessionsCountOfSubscriptionParams
	if err := cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, types.ErrorUnmarshal()
	}

	count := k.GetSessionsCountOfSubscription(ctx, params.ID)

	res, resErr := cdc.MarshalJSON(count)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

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
	QuerySubscription                = "subscription"
	QuerySubscriptionsOfNode         = "subscriptionsOfNode"
	QuerySubscriptionsOfAddress      = "subscriptionsOfAddress"
	QueryAllSubscriptions            = "allSubscriptions"
	QuerySessionsCountOfSubscription = "sessionsCountOfSubscription"
)

type QuerySubscriptionParams struct {
	ID sdk.ID
}

func NewQuerySubscriptionParams(id sdk.ID) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

// nolint:dupl
func querySubscription(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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
	ID sdk.ID
}

func NewQuerySubscriptionsOfNodePrams(id sdk.ID) QuerySubscriptionsOfNodePrams {
	return QuerySubscriptionsOfNodePrams{
		ID: id,
	}
}

func querySubscriptionsOfNode(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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
	Address csdk.AccAddress
}

func NewQuerySubscriptionsOfAddressParams(address csdk.AccAddress) QuerySubscriptionsOfAddressParams {
	return QuerySubscriptionsOfAddressParams{
		Address: address,
	}
}

func querySubscriptionsOfAddress(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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

func queryAllSubscriptions(ctx csdk.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdk.Error) {
	subscriptions := k.GetAllSubscriptions(ctx)

	res, resErr := cdc.MarshalJSON(subscriptions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

type QuerySessionsCountOfSubscriptionParams struct {
	ID sdk.ID
}

func NewQuerySessionsCountOfSubscriptionParams(id sdk.ID) QuerySessionsCountOfSubscriptionParams {
	return QuerySessionsCountOfSubscriptionParams{
		ID: id,
	}
}

func querySessionsCountOfSubscription(ctx csdk.Context, cdc *codec.Codec, req abci.RequestQuery,
	k keeper.Keeper) ([]byte, csdk.Error) {

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

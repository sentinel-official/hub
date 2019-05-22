package querier

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/keeper"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/types"
)

const (
	QuerySubscription           = "subscription"
	QuerySubscriptionsOfNode    = "subscriptionsOfNode"
	QuerySubscriptionsOfAddress = "subscriptionsOfAddress"
	QueryAllSubscriptions       = "allSubscriptions"
)

type QuerySubscriptionParams struct {
	ID sdkTypes.ID
}

func NewQuerySubscriptionParams(id sdkTypes.ID) QuerySubscriptionParams {
	return QuerySubscriptionParams{
		ID: id,
	}
}

func querySubscription(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

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
	ID sdkTypes.ID
}

func NewQuerySubscriptionsOfNodePrams(id sdkTypes.ID) QuerySubscriptionsOfNodePrams {
	return QuerySubscriptionsOfNodePrams{
		ID: id,
	}
}

func querySubscriptionsOfNode(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

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
	Address csdkTypes.AccAddress
}

func NewQuerySubscriptionsOfAddressParams(address csdkTypes.AccAddress) QuerySubscriptionsOfAddressParams {
	return QuerySubscriptionsOfAddressParams{
		Address: address,
	}
}

func querySubscriptionsOfAddress(ctx csdkTypes.Context, cdc *codec.Codec, req abciTypes.RequestQuery,
	k keeper.Keeper) ([]byte, csdkTypes.Error) {

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

func queryAllSubscriptions(ctx csdkTypes.Context, cdc *codec.Codec, k keeper.Keeper) ([]byte, csdkTypes.Error) {
	subscriptions := k.GetAllSubscriptions(ctx)

	res, resErr := cdc.MarshalJSON(subscriptions)
	if resErr != nil {
		return nil, types.ErrorMarshal()
	}

	return res, nil
}

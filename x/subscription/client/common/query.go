package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func QuerySubscription(ctx context.CLIContext, id uint64) (*types.Subscription, error) {
	params := types.NewQuerySubscriptionParams(id)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscription)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscription found")
	}

	var subscription types.Subscription
	if err := ctx.Codec.UnmarshalJSON(res, &subscription); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func QuerySubscriptions(ctx context.CLIContext, skip, limit int) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsParams(skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptions)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForNode(ctx context.CLIContext, address hub.NodeAddress, skip, limit int) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForNodeParams(address, skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForNode)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForPlan(ctx context.CLIContext, id uint64, skip, limit int) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForPlanParams(id, skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForPlan)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QuerySubscriptionsForAddress(ctx context.CLIContext, address sdk.AccAddress, status hub.Status, skip, limit int) (types.Subscriptions, error) {
	params := types.NewQuerySubscriptionsForAddressParams(address, status, skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QuerySubscriptionsForAddress)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no subscriptions found")
	}

	var subscriptions types.Subscriptions
	if err := ctx.Codec.UnmarshalJSON(res, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func QueryQuota(ctx context.CLIContext, id uint64, address sdk.AccAddress) (*types.Quota, error) {
	params := types.NewQueryQuotaParams(id, address)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryQuota)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no quota found")
	}

	var quota types.Quota
	if err := ctx.Codec.UnmarshalJSON(res, &quota); err != nil {
		return nil, err
	}

	return &quota, nil
}

func QueryQuotas(ctx context.CLIContext, id uint64, skip, limit int) (types.Quotas, error) {
	params := types.NewQueryQuotasParams(id, skip, limit)
	bytes, err := ctx.Codec.MarshalJSON(params)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("custom/%s/%s/%s", types.StoreKey, types.QuerierRoute, types.QueryQuotas)
	res, _, err := ctx.QueryWithData(path, bytes)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, fmt.Errorf("no quotas found")
	}

	var quotas types.Quotas
	if err := ctx.Codec.UnmarshalJSON(res, &quotas); err != nil {
		return nil, err
	}

	return quotas, nil
}

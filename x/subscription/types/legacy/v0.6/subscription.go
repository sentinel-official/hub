package v0_6

import (
	hubtypes "github.com/sentinel-official/hub/types/legacy/v0.6"
	"github.com/sentinel-official/hub/x/subscription/types"
	legacy "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.5"
)

func MigrateSubscription(item legacy.Subscription) types.Subscription {
	return types.Subscription{
		Id:       item.ID,
		Owner:    item.Owner.String(),
		Node:     item.Node.String(),
		Price:    item.Price,
		Deposit:  item.Deposit,
		Plan:     item.Plan,
		Denom:    "",
		Expiry:   item.Expiry,
		Free:     item.Free,
		Status:   hubtypes.MigrateStatus(item.Status),
		StatusAt: item.StatusAt,
	}
}

func MigrateSubscriptions(items legacy.Subscriptions) types.Subscriptions {
	var subscriptions types.Subscriptions
	for _, item := range items {
		subscriptions = append(subscriptions, MigrateSubscription(item))
	}

	return subscriptions
}

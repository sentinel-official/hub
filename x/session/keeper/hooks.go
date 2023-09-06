package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v1/types"
	"github.com/sentinel-official/hub/v1/x/session/types"
)

func (k *Keeper) SubscriptionInactivePendingHook(ctx sdk.Context, id uint64) error {
	// Get the status change delay from the store.
	statusChangeDelay := k.StatusChangeDelay(ctx)

	// Iterate through sessions associated with the subscription.
	k.IterateSessionsForSubscription(ctx, id, func(_ int, item types.Session) (stop bool) {
		// Skip non-active sessions.
		if !item.Status.Equal(hubtypes.StatusActive) {
			return false
		}

		// Delete the session's entry from the InactiveAt index before updating the InactiveAt value.
		k.DeleteSessionForInactiveAt(ctx, item.InactiveAt, item.ID)

		// Calculate the new InactiveAt value by adding the status change delay to the current block time.
		item.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)

		// Set the session status to 'InactivePending' to mark it for an upcoming status update.
		item.Status = hubtypes.StatusInactivePending

		// Record the time of the status update in 'StatusAt' field.
		item.StatusAt = ctx.BlockTime()

		// Update the session entry in the store with the new status and status update time.
		k.SetSession(ctx, item)

		// Update the session entry in the InactiveAt index with the new InactiveAt value.
		k.SetSessionForInactiveAt(ctx, item.InactiveAt, item.ID)

		// Emit an event to notify that the session status has been updated.
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				Status:         hubtypes.StatusInactivePending,
				Address:        item.Address,
				NodeAddress:    item.NodeAddress,
				ID:             item.ID,
				PlanID:         0,
				SubscriptionID: item.SubscriptionID,
			},
		)

		return false
	})

	return nil
}
